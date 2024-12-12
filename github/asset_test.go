package github

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssetContentExtract(_ *testing.T) {
	// todo: implement this.
}

func TestIsExecBinaryContent(_ *testing.T) {
	// todo: implement this.
}

func TestNewReaderToExtract(_ *testing.T) {
	// todo: implement this.
}

func TestAssetRepositoryList(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  []Asset
	}{
		{
			name:    "cli/cli",
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.52.0"),
			assets: []Asset{
				must(newAssetFromString(175682878, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_checksums.txt")),
				must(newAssetFromString(175682881, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.deb")),
				must(newAssetFromString(175682882, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.rpm")),
				must(newAssetFromString(175682880, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.tar.gz")),
				must(newAssetFromString(175682879, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.deb")),
				must(newAssetFromString(175682883, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.rpm")),
				must(newAssetFromString(175682889, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
				must(newAssetFromString(175682892, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.deb")),
				must(newAssetFromString(175682891, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.rpm")),
				must(newAssetFromString(175682895, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz")),
				must(newAssetFromString(175682896, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.deb")),
				must(newAssetFromString(175682899, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.rpm")),
				must(newAssetFromString(175682905, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.tar.gz")),
				must(newAssetFromString(175682903, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_amd64.zip")),
				must(newAssetFromString(175682902, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_arm64.zip")),
				must(newAssetFromString(175682904, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_universal.pkg")),
				must(newAssetFromString(175682911, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.msi")),
				must(newAssetFromString(175682913, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.zip")),
				must(newAssetFromString(175682914, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.msi")),
				must(newAssetFromString(175682915, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.zip")),
				must(newAssetFromString(175682917, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_arm64.zip")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(githubTokenForTest)
			assets, err := repository.list(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestAssetRepositoryDownload(_ *testing.T) {
	// todo: implement this.
}
