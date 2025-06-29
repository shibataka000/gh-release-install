package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/url"
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

// extract [ExecBinaryContent] from [AssetContent] and returns it.
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

// newReaderToExtract returns [io.Reader] to unarchive/decompress given bytes.
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
		r, err := newTarReader(br, execBinary.name)
		return r, nil, err
	case "application/zip":
		r, err := newZipReader(br, br.Size(), execBinary.name)
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("%w: %s", ErrUnexpectedMIMEType, mime.String())
	}
}

// newAssetRepository returns a new [GitHubAssetRepository] object or [ExternalAssetRepository] object based on given repository name.
func newAssetRepository(repo string, progressBar io.Writer) (AssetRepository, error) {
	r, err := parseRepository(repo)
	if err != nil {
		return nil, err
	}
	if templates, ok := defaultExternalAssetTemplates[r]; ok {
		return newExternalAssetRepository(templates, progressBar), nil
	}
	return newGitHubAssetRepository(r, progressBar), nil
}
