package github

var defaultPatternsForLinuxAmd64 = map[string]string{
	// General patterns.
	`(?i)^https://github\.com/.+/.+/releases/download/.+/(?P<name>[^\.]+)([\-\._]v?\d+\.\d+\.\d+)?[\-\._]linux([\-\._](amd64|x86_64|64bit))?(\.tar\.gz|\.zip|\.gz|\.tgz)?$`: "{{.name}}",

	// Patterns for specific repository. These must start with literals.
	`https://github\.com/istio/istio/releases/download/.+/istioctl-\d+\.\d+\.\d+-linux-amd64\.tar\.gz$`:      "istioctl",
	`https://github\.com/starship/starship/releases/download/.+/starship-x86_64-unknown-linux-gnu\.tar\.gz$`: "starship",
	`https://github\.com/protocolbuffers/protobuf/releases/download/.+/protoc-\d+\.\d+-linux-x86_64\.zip$`:   "protoc",
}

func init() {
	DefaultPatterns = defaultPatternsForLinuxAmd64
}
