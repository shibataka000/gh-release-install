package external_test

import (
	"context"
	"io"
	"net/url"
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/shibataka000/gh-release-install/internal/external"
	"github.com/stretchr/testify/require"
)

func TestExternalAssetRepositoryList(t *testing.T) {
	tests := []struct {
		name    string
		repo    github.Repository
		release github.Release
		assets  []github.Asset
	}{
		{
			name: "hashicorp/terraform",
			repo: github.Repository{
				Host:  "github.com",
				Owner: "hashicorp",
				Name:  "terraform",
			},
			release: github.Release{
				Tag: "v1.9.0",
			},
			assets: []github.Asset{
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_darwin_amd64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_darwin_arm64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_386.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_amd64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_arm.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_386.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_arm.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_arm64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_openbsd_386.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_openbsd_amd64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_solaris_amd64.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_windows_386.zip"))},
				{ID: 0, DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_windows_amd64.zip"))},
			},
		},
		{
			name: "helm/helm",
			repo: github.Repository{
				Host:  "github.com",
				Owner: "helm",
				Name:  "helm",
			},
			release: github.Release{
				Tag: "v3.16.2",
			},
			assets: []github.Asset{
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-darwin-amd64.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-darwin-arm64.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-386.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-arm.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-arm64.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-ppc64le.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-riscv64.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-s390x.tar.gz"))},
				{ID: 0, DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-windows-amd64.zip"))},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := external.NewAssetRepository(tt.repo, io.Discard)
			assets, err := repository.List(ctx, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestExternalAssetRepositoryDownload(_ *testing.T) {
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
