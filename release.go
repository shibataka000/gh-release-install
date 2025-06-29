package main

import (
	"strings"

	"golang.org/x/mod/semver"
)

// Release represents a GitHub release.
type Release struct {
	tag string
}

// semVer returns semantic version of this release.
// For example, if release tag is "v1.2.3", this returns "1.2.3".
// If release tag can't be converted into semantic version, this returns empty string.
func (r Release) semVer() string {
	v := strings.TrimLeft(r.tag, "v")
	if semver.IsValid("v" + v) {
		return v
	}
	return ""
}
