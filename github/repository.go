package github

import (
	"github.com/cli/go-gh/v2/pkg/repository"
)

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
	return Repository{
		Host:  repo.Host,
		Owner: repo.Owner,
		Name:  repo.Name,
	}, nil
}
