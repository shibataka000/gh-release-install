package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPatternMatch(t *testing.T) {
	tests := []struct {
		name    string
		pattern Pattern
		asset   Asset
		match   bool
	}{
		{
			name:    "Match",
			pattern: must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/gh_.+_linux_amd64\.tar\.gz`, "gh")),
			asset:   must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			match:   true,
		},
		{
			name:    "NotMatch",
			pattern: must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/gh_.+_linux_amd64\.tar\.gz`, "gh")),
			asset:   must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz")),
			match:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.match, tt.pattern.match(tt.asset))
		})
	}
}

func TestPatternPriority(t *testing.T) {
	tests := []struct {
		name     string
		pattern  Pattern
		priority int
	}{
		{
			name:     "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz",
			pattern:  must(newPatternFromString(`https://github\.com/cli/cli/releases/download/v2\.52\.0/gh_2\.52\.0_linux_amd64\.tar\.gz`, "gh")),
			priority: len("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
		},
		{
			name:     "https://github.com/cli/cli/releases/download/.+/gh_.+_linux_amd64.tar.gz",
			pattern:  must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/gh_.+_linux_amd64\.tar\.gz`, "gh")),
			priority: len("https://github.com/cli/cli/releases/download/"),
		},
		{
			name:     "https://github.com/.+/.+/releases/download/.+/.+",
			pattern:  must(newPatternFromString(`https://github\.com/.+/.+/releases/download/.+/.+`, "gh")),
			priority: len("https://github.com/"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			priority := tt.pattern.priority()
			require.Equal(tt.priority, priority)
		})
	}
}

func TestPatternExecute(t *testing.T) {
	tests := []struct {
		name       string
		pattern    Pattern
		asset      Asset
		execBinary ExecBinary
	}{
		{
			name:       "CapturingGroup",
			pattern:    must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/(\w+)_[\d\.]+_linux_amd64\.tar\.gz`, `{{index . "1"}}`)),
			asset:      must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			execBinary: newExecBinary("gh"),
		},
		{
			name:       "NamedCapturingGroup",
			pattern:    must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/(?P<name>\w+)_[\d\.]+_linux_amd64\.tar\.gz`, "{{.name}}")),
			asset:      must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			execBinary: newExecBinary("gh"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			execBinary, err := tt.pattern.execute(tt.asset)
			require.NoError(err)
			require.Equal(tt.execBinary, execBinary)
		})
	}
}

func TestFindAssetAndPattern(t *testing.T) {
	tests := []struct {
		name     string
		assets   []Asset
		patterns []Pattern
		asset    Asset
		pattern  Pattern
	}{
		{
			name: "FindMatchingAssetAndPattern",
			assets: []Asset{
				must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
				must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz")),
			},
			patterns: []Pattern{
				must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/gh_.+_linux_amd64\.tar\.gz`, "gh")),
				must(newPatternFromString(`https://github\.com/istio/istio/releases/download/.+/istioctl-.+-linux-amd64\.tar\.gz`, "istioctl")),
			},
			asset:   must(newAssetFromString(0, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			pattern: must(newPatternFromString(`https://github\.com/cli/cli/releases/download/.+/gh_.+_linux_amd64\.tar\.gz`, "gh")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset, pattern, err := findAssetAndPattern(tt.assets, tt.patterns)
			require.NoError(err)
			require.Equal(tt.asset, asset)
			require.Equal(tt.pattern, pattern)
		})
	}
}
