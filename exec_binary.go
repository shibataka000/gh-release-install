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

// NewExecBinaryRepository returns a new [ExecBinaryRepository] object.
func NewExecBinaryRepository(dir string) ExecBinaryRepository {
	return newFSExecBinaryRepository(dir)
}

// FSExecBinaryRepository is a repository for [ExecBinary] and [ExecBinaryContent].
type FSExecBinaryRepository struct {
	dir string
}

// newFSExecBinaryRepository returns a new [FSExecBinaryRepository] object.
func newFSExecBinaryRepository(dir string) *FSExecBinaryRepository {
	return &FSExecBinaryRepository{
		dir: dir,
	}
}

// Write [ExecBinaryContent] into file in given repository.
func (r *FSExecBinaryRepository) Write(meta ExecBinary, content ExecBinaryContent) error {
	path := filepath.Join(r.dir, meta.Name)
	return os.WriteFile(path, content, 0755)
}
