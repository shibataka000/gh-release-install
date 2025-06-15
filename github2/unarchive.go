package github2

import (
	"archive/tar"
	"archive/zip"
	"io"
	"path/filepath"
)

func newFileReaderInTar(r io.Reader, name string) (io.Reader, error) {
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

func newFileReaderInZip(r io.ReaderAt, size int64, name string) (io.ReadCloser, error) {
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
