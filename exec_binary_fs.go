package main

import (
	"os"
	"path/filepath"
)

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
