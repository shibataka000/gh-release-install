package main

import (
	"archive/tar"
	"archive/zip"
	"io"
	"path/filepath"
)

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
