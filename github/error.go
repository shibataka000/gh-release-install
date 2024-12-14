package github

import "errors"

var (
	// ErrUnexpectedMIME is returned when MIME of release asset content is unexpected.
	ErrUnexpectedMIME = errors.New("unexpected mime type")
	// ErrNoAssetsMatchPatterns is returned when no release assets matched given patterns.
	ErrNoAssetsMatchPatterns = errors.New("no release assets matched given patterns")
	// ErrNotFound means something was not found.
	ErrNotFound = errors.New("not found")
)
