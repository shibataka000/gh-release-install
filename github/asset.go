package github

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
	ID          int64
	DownloadURL *url.URL
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extract [ExecBinaryContent] from [AssetContent] and returns it.
func (a AssetContent) Extract(execBinary ExecBinary) (ExecBinaryContent, error) {
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
	expect := []string{"application/octet-stream", "application/x-executable", "application/x-sharedlib"}
	mime := mimetype.Detect(b)
	return slices.Contains(expect, mime.String())
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
		r, err := newFileReaderInTar(br, execBinary.Name)
		return r, nil, err
	case "application/zip":
		r, err := newFileReaderInZip(br, br.Size(), execBinary.Name)
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("%w: %s", ErrUnexpectedMIMEType, mime.String())
	}
}
