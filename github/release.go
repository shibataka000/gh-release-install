package github

import (
	"strings"

	"golang.org/x/mod/semver"
)

// Release represents a GitHub release.
type Release struct {
	Tag string
}

// NewRelease returns a new [Release] object.
func NewRelease(tag string) Release {
	return Release{
		Tag: tag,
	}
}

// SemVer returns semantic version of this release.
// For example, if release tag is "v1.2.3", this returns "1.2.3".
// If release tag can't be converted into semantic version, this returns empty string.
func (r Release) SemVer() string {
	v := strings.TrimLeft(r.Tag, "v")
	if semver.IsValid("v" + v) {
		return v
	}
	return ""
}
