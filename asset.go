package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"slices"

	"github.com/gabriel-vasile/mimetype"
	"github.com/ulikunitz/xz"
)

// Asset represents a GitHub release asset.
type Asset struct {
	id          int64
	downloadURL *url.URL
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extracts [ExecBinaryContent] from [AssetContent] and returns it.
func (a AssetContent) extract(execBinary ExecBinary) (ExecBinaryContent, error) {
	b := []byte(a)
	for !isExecBinaryContent(b) {
		r, err := newReaderToExtract(b, execBinary)
		if err != nil {
			return nil, err
		}
		b, err = io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		if err := r.Close(); err != nil {
			return nil, err
		}
	}
	return ExecBinaryContent(b), nil
}

// isExecBinaryContent returns true if MIME type of given bytes means executable binary content.
func isExecBinaryContent(b []byte) bool {
	binaryMIMEs := []string{"application/octet-stream", "application/x-executable", "application/x-sharedlib"}
	mime := mimetype.Detect(b)
	return slices.Contains(binaryMIMEs, mime.String())
}

// newReaderToExtract returns a [io.ReadCloser] to unarchive/decompress given bytes.
// Closing [io.ReadCloser] is caller's responsibility.
func newReaderToExtract(b []byte, execBinary ExecBinary) (io.ReadCloser, error) {
	br := bytes.NewReader(b)
	mime := mimetype.Detect(b)

	switch mime.String() {
	case "application/gzip":
		return gzip.NewReader(br)
	case "application/x-xz":
		r, err := xz.NewReader(br)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(r), nil
	case "application/x-tar":
		return newTarReader(br, execBinary.name)
	case "application/zip":
		return newZipReader(br, br.Size(), execBinary.name)
	default:
		return nil, fmt.Errorf("MIME type of asset content was unexpected: %s", mime.String())
	}
}

// newTarReader returns a [io.ReadCloser] to read file which is given name in tarball.
// Closing [io.ReadCloser] is caller's responsibility.
func newTarReader(r io.Reader, name string) (io.ReadCloser, error) {
	for tr := tar.NewReader(r); ; {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if filepath.Base(header.Name) == name && header.Typeflag == tar.TypeReg {
			return io.NopCloser(tr), nil
		}
	}
}

// newZipReader returns a [io.ReadCloser] to read file which is given name in zip file.
// Closing [io.ReadCloser] is caller's responsibility.
func newZipReader(r io.ReaderAt, size int64, name string) (io.ReadCloser, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if filepath.Base(f.Name) == name && !f.FileInfo().IsDir() {
			return f.Open()
		}
	}
	return nil, io.EOF
}

// AssetRepository is an interface about repository for [Asset] and [AssetContent].
type AssetRepository interface {
	list(ctx context.Context, release Release) ([]Asset, error)
	download(ctx context.Context, asset Asset) (AssetContent, error)
}

// newAssetRepository returns a new [GitHubAssetRepository] object or [ExternalAssetRepository] object based on given repository name.
func newAssetRepository(repo string, progressBar io.Writer) (AssetRepository, error) {
	r, err := parseRepository(repo)
	if err != nil {
		return nil, err
	}
	if templates, ok := externalAssetTemplates[r]; ok {
		return newExternalAssetRepository(templates, progressBar), nil
	}
	return newGitHubAssetRepository(r, progressBar), nil
}
