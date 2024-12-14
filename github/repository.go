package github

import (
	"fmt"
	"regexp"
)

// repositoryFullNameFormat represents a GitHub repository full name format.
var repositoryFullNameFormat = regexp.MustCompile("(?P<owner>.*)/(?P<name>.*)")

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
// Repository full name should be 'OWNER/REPO' format.
func newRepositoryFromFullName(fullName string) (Repository, error) {
	if !repositoryFullNameFormat.MatchString(fullName) {
		return Repository{}, fmt.Errorf("%w: %s", ErrInvalidRepositoryFullName, fullName)
	}

	submatch := repositoryFullNameFormat.FindStringSubmatch(fullName)

	ownerSubexpIndex := repositoryFullNameFormat.SubexpIndex("owner")
	if ownerSubexpIndex < 0 && ownerSubexpIndex >= len(submatch) {
		return Repository{}, errOutOfIndex
	}
	owner := submatch[ownerSubexpIndex]

	nameSubexpIndex := repositoryFullNameFormat.SubexpIndex("name")
	if nameSubexpIndex < 0 && nameSubexpIndex >= len(submatch) {
		return Repository{}, errOutOfIndex
	}
	name := submatch[nameSubexpIndex]

	return newRepository(owner, name), nil
}
