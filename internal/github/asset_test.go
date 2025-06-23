package github_test

import (
	"context"
	"io"
	"net/url"
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	internalgithub "github.com/shibataka000/gh-release-install/internal/github"
	"github.com/stretchr/testify/require"
)

func TestAssetRepositoryList(t *testing.T) {
	tests := []struct {
		name    string
		repo    github.Repository
		release github.Release
		assets  []github.Asset
	}{
		{
			name: "cli/cli",
			repo: github.Repository{
				Host:  "github.com",
				Owner: "cli",
				Name:  "cli",
			},
			release: github.Release{
				Tag: "v2.52.0",
			},
			assets: []github.Asset{
				{ID: 175682878, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_checksums.txt"))},
				{ID: 175682881, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.deb"))},
				{ID: 175682882, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.rpm"))},
				{ID: 175682880, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.tar.gz"))},
				{ID: 175682879, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.deb"))},
				{ID: 175682883, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.rpm"))},
				{ID: 175682889, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"))},
				{ID: 175682892, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.deb"))},
				{ID: 175682891, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.rpm"))},
				{ID: 175682895, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz"))},
				{ID: 175682896, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.deb"))},
				{ID: 175682899, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.rpm"))},
				{ID: 175682905, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.tar.gz"))},
				{ID: 175682903, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_amd64.zip"))},
				{ID: 175682902, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_arm64.zip"))},
				{ID: 175682904, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_universal.pkg"))},
				{ID: 175682911, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.msi"))},
				{ID: 175682913, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.zip"))},
				{ID: 175682914, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.msi"))},
				{ID: 175682915, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.zip"))},
				{ID: 175682917, DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_arm64.zip"))},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := internalgithub.NewAssetRepository(tt.repo, io.Discard)
			assets, err := repository.List(ctx, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestAssetRepositoryDownload(_ *testing.T) {
	// todo: implement this.
}

// must is a helper that wraps a call to a function returning (E, error) and panics if the error is non-nil.
// This is intended for use in variable initializations.
func must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}
