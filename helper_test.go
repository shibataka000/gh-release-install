package main

import (
	"os/exec"
	"testing"
)

// clone [exec.Cmd] and return it.
func clone(t *testing.T, cmd *exec.Cmd) *exec.Cmd {
	t.Helper()
	newCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	newCmd.Dir = cmd.Dir
	return newCmd
}
