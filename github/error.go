package github

import "errors"

var (
	// ErrInvalidRepositoryFullName means that given repository name was not 'OWNER/REPO' format.
	ErrInvalidRepositoryFullName = errors.New("given repository name was not 'OWNER/REPO' format")
	// ErrUnexpectedMIMEType means that MIME type of asset content was unexpected.
	ErrUnexpectedMIMEType = errors.New("MIME type of asset content was unexpected")
	// ErrNoAssetsMatchPattern means that no assets match the pattern.
	ErrNoAssetsMatchPattern = errors.New("no assets match the pattern")
)
