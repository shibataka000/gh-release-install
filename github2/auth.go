package github2

import (
	"net/http"

	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/google/go-github/v67/github"
)

func newGitHubClient(host string) *github.Client {
	token, _ := auth.TokenForHost(host)
	return github.NewClient(http.DefaultClient).WithAuthToken(token)
}
