package github

import (
	"github.com/cli/go-gh/v2/pkg/repository"
)

// defaultHost represents a default GitHub host.
var defaultHost = "github.com"

// Repository represents a GitHub repository.
type Repository struct {
	host  string
	owner string
	name  string
}

// newRepositoryWithHost returns a new [Repository] object.
func newRepositoryWithHost(host string, owner string, name string) Repository {
	return Repository{
		host:  host,
		owner: owner,
		name:  name,
	}
}

// newRepository returns a new [Repository] object whose host is 'github.com' .
func newRepository(owner string, name string) Repository {
	return newRepositoryWithHost(defaultHost, owner, name)
}

// parseRepository extracts the repository information from the following string formats: "OWNER/REPO", "HOST/OWNER/REPO", and a full URL.
// If the format does not specify a host, use the config to determine a host.
func parseRepository(s string) (Repository, error) {
	repo, err := repository.Parse(s)
	if err != nil {
		return Repository{}, err
	}
	return newRepositoryWithHost(repo.Host, repo.Owner, repo.Name), nil
}
