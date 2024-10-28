package github

var (
	// DefaultPatterns are recommended patterns.
	DefaultPatterns = appendMap(DefaultCorePatterns, DefaultExtPatterns)

	// DefaultCorePatterns are recommended patterns for general repository.
	// These are same as [github.com/shibataka000/gh-release-install/external.DefaultCorePatterns].
	// These are for linux/amd64.
	DefaultCorePatterns = map[string]string{
		`(?i)^.+/(?P<name>[^\.]+)([\-\._]v?\d+\.\d+\.\d+)?[\-\._]linux([\-\._](amd64|x86_64|64bit))?(\.tar\.gz|\.tar\.xz|\.zip|\.gz|\.tgz)?$`: "{{.name}}",
	}

	// DefaultExtPatterns are recommended patterns for specific repository.
	// These should start with literals containing host and repository name to avoid conflict with other patterns.
	// These are for linux/amd64.
	DefaultExtPatterns = map[string]string{
		`https://github\.com/istio/istio/releases/download/.+/istioctl-\d+\.\d+\.\d+-linux-amd64\.tar\.gz$`:      "istioctl",
		`https://github\.com/starship/starship/releases/download/.+/starship-x86_64-unknown-linux-gnu\.tar\.gz$`: "starship",
		`https://github\.com/protocolbuffers/protobuf/releases/download/.+/protoc-\d+\.\d+-linux-x86_64\.zip$`:   "protoc",
	}
)
