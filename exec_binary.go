package main

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
