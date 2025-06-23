package github

import (
	"github.com/cli/go-gh/v2/pkg/repository"
)

// defaultHost represents a default GitHub host.
var defaultHost = "github.com"

// Repository represents a GitHub repository.
type Repository struct {
	Host  string
	Owner string
	Name  string
}

// ParseRepository extracts the repository information from the following string formats: "OWNER/REPO", "HOST/OWNER/REPO", and a full URL.
// If the format does not specify a host, use the config to determine a host.
func ParseRepository(s string) (Repository, error) {
	repo, err := repository.Parse(s)
	if err != nil {
		return Repository{}, err
	}
	// return newRepositoryWithHost(repo.Host, repo.Owner, repo.Name), nil
	return Repository{
		Host:  repo.Host,
		Owner: repo.Owner,
		Name:  repo.Name,
	}, nil
}
