package github

import (
	"github.com/cli/go-gh/v2/pkg/repository"
)

// Repository represents a GitHub repository.
type Repository struct {
	owner string
	name  string
}

// newRepository returns a new [Repository] object.
func newRepository(owner string, name string) Repository {
	return Repository{
		owner: owner,
		name:  name,
	}
}

// newRepositoryFromFullName returns a new [Repository] object from repository full name.
func newRepositoryFromFullName(fullName string) (Repository, error) {
	repo, err := repository.Parse(fullName)
	if err != nil {
		return Repository{}, err
	}
	return newRepository(repo.Owner, repo.Name), nil
}
