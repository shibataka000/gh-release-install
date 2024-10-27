package external

import (
	"maps"

	"github.com/shibataka000/gh-release-install/github"
)

var (
	// DefaultPatterns are recommended patterns.
	DefaultPatterns = appendMap(DefaultCorePatterns, DefaultExtPatterns)

	// DefaultCorePatterns are recommended patterns for general repository.
	// These are same as [github.com/shibataka000/gh-release-install/github.DefaultCorePatterns].
	// These are for linux/amd64.
	DefaultCorePatterns = maps.Clone(github.DefaultCorePatterns)

	// DefaultExtPatterns are recommended patterns for specific repository.
	// These should start with literals containing host to avoid conflict with other patterns.
	// These are for linux/amd64.
	DefaultExtPatterns = map[string]string{
		`https://dl\.k8s\.io/release/.+/bin/linux/amd64/kubectl`:         "kubectl",
		`https://cdn\.teleport\.dev/teleport-v.+-linux-amd64-bin.tar.gz`: "tsh",
	}
)
