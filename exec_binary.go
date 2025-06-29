package main

import (
	"os"
	"path/filepath"
)

// ExecBinary represents a executable binary in a GitHub release asset.
type ExecBinary struct {
	Name string
}

// ExecBinaryContent represents a executable binary content in a GitHub release asset content.
type ExecBinaryContent []byte

// ExecBinaryRepository is a repository for [ExecBinary] and [ExecBinaryContent].
type ExecBinaryRepository struct{}

// NewExecBinaryRepository returns a new [ExecBinaryRepository] object.
func NewExecBinaryRepository() *ExecBinaryRepository {
	return &ExecBinaryRepository{}
}

// write [ExecBinaryContent] into file in given directory.
func (r *ExecBinaryRepository) Write(meta ExecBinary, content ExecBinaryContent, dir string) error {
	path := filepath.Join(dir, meta.Name)
	return os.WriteFile(path, content, 0755)
}
