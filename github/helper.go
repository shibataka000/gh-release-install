package github

import (
	"maps"
	"os"
	"os/exec"
	"testing"
)

// githubTokenForTest is authentication token for GitHub API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GH_TOKEN")

// appendMap merges m1, m2 and return new map.
func appendMap[M ~map[K]V, K comparable, V any](m1 M, m2 M) M {
	m3 := M{}
	maps.Copy(m3, m1)
	maps.Copy(m3, m2)
	return m3
}

// cloneCommand clones [exec.Cmd] and return it.
func cloneCommand(t *testing.T, cmd *exec.Cmd) *exec.Cmd {
	t.Helper()
	newCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	newCmd.Dir = cmd.Dir
	return newCmd
}

// must is a helper that wraps a call to a function returning (E, error) and panics if the error is non-nil.
// This is intended for use in variable initializations.
func must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}
