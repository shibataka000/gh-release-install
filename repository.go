package main

import (
	"github.com/cli/go-gh/v2/pkg/repository"
)

// Repository represents a GitHub repository.
type Repository struct {
	Host  string
	Name  string
	Owner string
}

// parseRepository extracts the repository information from the following string formats: "OWNER/REPO", "HOST/OWNER/REPO", and a full URL.
// If the format does not specify a host, use the config to determine a host.
func parseRepository(s string) (Repository, error) {
	repo, err := repository.Parse(s)
	if err != nil {
		return Repository{}, err
	}
	return Repository(repo), nil
}

// currentRepository returns the GitHub repository the current directory is tracking.
func currentRepository() (Repository, error) {
	repo, err := repository.Current()
	if err != nil {
		return Repository{}, err
	}
	return Repository(repo), nil
}
