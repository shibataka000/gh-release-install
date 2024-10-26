package github

import "errors"

var (
	// ErrInvalidRepositoryFullName is returned when given repository full name is not 'OWNER/REPO' format.
	ErrInvalidRepositoryFullName = errors.New("repository full name was not 'OWNER/REPO' format")
	// ErrUnexpectedMIME is returned when MIME of release asset content is unexpected.
	ErrUnexpectedMIME = errors.New("unexpected mime type")
	// ErrNoAssetsMatchPatterns is returned when no release assets matched given patterns.
	ErrNoAssetsMatchPatterns = errors.New("no release assets matched given patterns")
)
