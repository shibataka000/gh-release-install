package main

// ExecBinary represents a executable binary in a GitHub release asset.
type ExecBinary struct {
	Name string
}

// ExecBinaryContent represents an executable binary content in a GitHub release asset content.
type ExecBinaryContent []byte

// ExecBinaryRepository is an interface about repository for [ExecBinary] and [ExecBinaryContent].
type ExecBinaryRepository interface {
	Write(meta ExecBinary, content ExecBinaryContent) error
}

// NewExecBinaryRepository returns a new [ExecBinaryRepository] object.
func NewExecBinaryRepository(dir string) ExecBinaryRepository {
	return newFSExecBinaryRepository(dir)
}
