package external

import "errors"

var (
	// ErrNoAssetTemplatesFound is returned when no asset templates was found for given repository.
	ErrNoAssetTemplatesFound = errors.New("no asset templates was found for given repository")
)
