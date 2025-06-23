package github_test

import (
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/stretchr/testify/require"
)

func TestReleaseSemVer(t *testing.T) {
	tests := []struct {
		name    string
		release github.Release
		semver  string
	}{
		{
			name: "v1.0.0",
			release: github.Release{
				Tag: "v1.0.0",
			},
			semver: "1.0.0",
		},
		{
			name: "1.0.0",
			release: github.Release{
				Tag: "1.0.0",
			},
			semver: "1.0.0",
		},
		{
			name: "x.y.z",
			release: github.Release{
				Tag: "x.y.z",
			},
			semver: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.semver, tt.release.SemVer())
		})
	}
}
