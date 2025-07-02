package main

import (
	"os"
	"path/filepath"
)

// ExecBinary represents a executable binary in a GitHub release asset.
type ExecBinary struct {
	name string
}

// ExecBinaryContent represents a executable binary content in a GitHub release asset content.
type ExecBinaryContent []byte

// ExecBinaryRepository is a repository for [ExecBinary] and [ExecBinaryContent].
type ExecBinaryRepository2 struct{}

// newExecBinaryRepository returns a new [ExecBinaryRepository] object.
func newExecBinaryRepository() ExecBinaryRepository {
	return &ExecBinaryRepository2{}
}

// write [ExecBinaryContent] into file in given directory.
func (r *ExecBinaryRepository2) Write(meta ExecBinary, content ExecBinaryContent, dir string) error {
	path := filepath.Join(dir, meta.name)
	return os.WriteFile(path, content, 0755)
}
