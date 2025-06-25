package pattern

import "errors"

var (
	// ErrNoAssetsMatchPattern means that no assets match the pattern.
	ErrNoAssetsMatchPattern = errors.New("no assets match the pattern")
)
