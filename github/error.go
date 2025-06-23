package github

import "errors"

var (
	// ErrUnexpectedMIMEType means that MIME type of asset content was unexpected.
	ErrUnexpectedMIMEType = errors.New("MIME type of asset content was unexpected")
)
