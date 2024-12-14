package github

import (
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/google/go-github/v62/github"
)

// githubTokenForTest is authentication token for GitHub API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GH_TOKEN")

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

// newGitHubClient returns a new GitHub API client.
func newGitHubClient(token string) *github.Client {
	return github.NewClient(http.DefaultClient).WithAuthToken(token)
}
