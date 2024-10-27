package external

import "errors"

var (
	// ErrAssetTemplateNotFound is returned when no asset templates was found for given repository.
	ErrAssetTemplateNotFound = errors.New("no asset templates was found for given repository")
)
