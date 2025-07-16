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
	ID          int64
	DownloadURL *url.URL
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extracts [ExecBinaryContent] from [AssetContent] and returns it.
func (a AssetContent) extract(execBinary ExecBinary) (ExecBinaryContent, error) {
	b := []byte(a)
	for !isExecBinaryContent(b) {
		r, c, err := newReaderToExtract(b, execBinary)
		if err != nil {
			return nil, err
		}
		b, err = io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		if c != nil {
			if err := c.Close(); err != nil {
				return nil, err
			}
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

// newReaderToExtract returns an [io.Reader] to unarchive/decompress given bytes.
// Closing [io.ReadCloser] is caller's responsibility if it is not nil.
func newReaderToExtract(b []byte, execBinary ExecBinary) (io.Reader, io.Closer, error) {
	br := bytes.NewReader(b)
	mime := mimetype.Detect(b)

	switch mime.String() {
	case "application/gzip":
		r, err := gzip.NewReader(br)
		return r, nil, err
	case "application/x-xz":
		r, err := xz.NewReader(br)
		return r, nil, err
	case "application/x-tar":
		r, err := newTarReader(br, execBinary.Name)
		return r, nil, err
	case "application/zip":
		r, err := newZipReader(br, br.Size(), execBinary.Name)
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("MIME type of asset content was unexpected: %s", mime.String())
	}
}

// newTarReader returns a [io.Reader] to read file which is given name in tarball.
func newTarReader(r io.Reader, name string) (io.Reader, error) {
	for tr := tar.NewReader(r); ; {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if filepath.Base(header.Name) == name && header.Typeflag == tar.TypeReg {
			return tr, nil
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
	List(ctx context.Context, release Release) ([]Asset, error)
	Download(ctx context.Context, asset Asset) (AssetContent, error)
}

// NewAssetRepository returns a new [GitHubAssetRepository] object or [ExternalAssetRepository] object based on given repository name.
func NewAssetRepository(repo string, progressBar io.Writer) (AssetRepository, error) {
	r, err := parseRepository(repo)
	if err != nil {
		return nil, err
	}
	if templates, ok := externalAssetTemplates[r]; ok {
		return newExternalAssetRepository(templates, progressBar), nil
	}
	return newGitHubAssetRepository(r, progressBar), nil
}
