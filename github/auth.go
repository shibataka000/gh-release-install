package github

import (
	"net/http"

	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/google/go-github/v67/github"
)

// newGitHubClient returns a new [github.com/google/go-github/v67/github.Client] object.
func newGitHubClient(host string) *github.Client {
	token, _ := auth.TokenForHost(host)
	return github.NewClient(http.DefaultClient).WithAuthToken(token)
}
