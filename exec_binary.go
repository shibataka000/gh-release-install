package main

// ExecBinary represents a executable binary in a GitHub release asset.
type ExecBinary struct {
	name string
}

// ExecBinaryContent represents an executable binary content in a GitHub release asset content.
type ExecBinaryContent []byte

// ExecBinaryRepository is an interface about repository for [ExecBinary] and [ExecBinaryContent].
type ExecBinaryRepository interface {
	write(meta ExecBinary, content ExecBinaryContent) error
}

// newExecBinaryRepository returns a new [ExecBinaryRepository] object.
func newExecBinaryRepository(dir string) ExecBinaryRepository {
	return newFSExecBinaryRepository(dir)
}
