package github

import (
	"os"
	"path/filepath"
)

// ExecBinary represents a executable binary in a GitHub release asset.
type ExecBinary struct {
	Name string
}

// ExecBinaryContent represents an executable binary content in a GitHub release asset content.
type ExecBinaryContent []byte

// FSExecBinaryRepository is a repository for [ExecBinary] and [ExecBinaryContent].
type FSExecBinaryRepository struct {
	dir string
}

// newExecBinaryRepository returns a new [FSExecBinaryRepository] object.
func newExecBinaryRepository(dir string) *FSExecBinaryRepository {
	return &FSExecBinaryRepository{
		dir: dir,
	}
}

// write writes [ExecBinaryContent] into a file in given repository.
func (r *FSExecBinaryRepository) write(meta ExecBinary, content ExecBinaryContent) error {
	path := filepath.Join(r.dir, meta.Name)
	return os.WriteFile(path, content, 0755)
}
