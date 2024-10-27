package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReleaseSemVer(t *testing.T) {
	tests := []struct {
		name    string
		release Release
		semver  string
	}{
		{
			name:    "v1.0.0",
			release: NewRelease("v1.0.0"),
			semver:  "1.0.0",
		},
		{
			name:    "1.0.0",
			release: NewRelease("1.0.0"),
			semver:  "1.0.0",
		},
		{
			name:    "x.y.z",
			release: NewRelease("x.y.z"),
			semver:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.semver, tt.release.SemVer())
		})
	}
}
