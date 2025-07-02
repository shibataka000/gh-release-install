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

// FSExecBinaryRepository is a repository for [ExecBinary] and [ExecBinaryContent].
type FSExecBinaryRepository struct {
	dir string
}

// newExecBinaryRepository returns a new [ExecBinaryRepository] object.
func newExecBinaryRepository(dir string) *FSExecBinaryRepository {
	return &FSExecBinaryRepository{
		dir: dir,
	}
}

// newFSExecBinaryRepository returns a new [FSExecBinaryRepository] object.
func newFSExecBinaryRepository(dir string) *FSExecBinaryRepository {
	return &FSExecBinaryRepository{
		dir: dir,
	}
}

// write [ExecBinaryContent] into file in given directory.
func (r *FSExecBinaryRepository) Write(meta ExecBinary, content ExecBinaryContent) error {
	path := filepath.Join(r.dir, meta.name)
	return os.WriteFile(path, content, 0755)
}
