package github

import (
	"fmt"
	"strings"
)

// Repository represents a GitHub repository.
type Repository struct {
	Owner string
	Name  string
}

// newRepository returns a new [Repository] object.
func newRepository(owner string, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}

// newRepositoryFromFullName returns a new [Repository] object from repository full name.
// Repository full name should be 'OWNER/REPO' format.
func newRepositoryFromFullName(fullName string) (Repository, error) {
	s := strings.Split(fullName, "/")
	if len(s) != 2 {
		return Repository{}, fmt.Errorf("%w: %s", ErrInvalidRepositoryFullName, fullName)
	}
	return newRepository(s[0], s[1]), nil
}
