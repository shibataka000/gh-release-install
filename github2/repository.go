package github2

import (
	"github.com/cli/go-gh/v2/pkg/repository"
)

var defaultHost = "github.com"

// Repository represents a GitHub repository.
type Repository struct {
	host  string
	owner string
	name  string
}

func newRepository(owner string, name string) Repository {
	return newRepositoryWithHost(defaultHost, owner, name)
}

func newRepositoryWithHost(host string, owner string, name string) Repository {
	return Repository{
		host:  host,
		owner: owner,
		name:  name,
	}
}

func parseRepository(s string) (Repository, error) {
	repo, err := repository.Parse(s)
	if err != nil {
		return Repository{}, err
	}
	return newRepositoryWithHost(repo.Host, repo.Owner, repo.Name), nil
}
