package external

import (
	"context"
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name         string
		repoFullName string
		exists       bool
	}{
		{
			name:         "hashicorp/terraform",
			repoFullName: "hashicorp/terraform",
			exists:       true,
		},
		{
			name:         "NotExist",
			repoFullName: "/",
			exists:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			exists, err := Exists(tt.repoFullName)
			require.NoError(err)
			require.Equal(tt.exists, exists)
		})
	}
}

func TestAssetTemplateExecute(t *testing.T) {
	tests := []struct {
		name     string
		template AssetTemplate
		release  github.Release
		asset    github.Asset
	}{
		{
			name:     "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip",
			template: mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"),
			release:  github.NewRelease("v1.9.0"),
			asset:    must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
		},
		{
			name:     "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz",
			template: mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz"),
			release:  github.NewRelease("v3.16.2"),
			asset:    must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.template.execute(tt.release)
			require.NoError(err)
			require.Equal(tt.asset, asset)
		})
	}
}

func TestAssetRepositoryList(t *testing.T) {
	tests := []struct {
		name    string
		repo    github.Repository
		release github.Release
		assets  []github.Asset
	}{
		{
			name:    "hashicorp/terraform",
			repo:    github.NewRepository("hashicorp", "terraform"),
			release: github.NewRelease("v1.9.0"),
			assets: []github.Asset{
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_darwin_amd64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_darwin_arm64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_386.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_amd64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_arm.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_386.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_arm.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_arm64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_openbsd_386.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_openbsd_amd64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_solaris_amd64.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_windows_386.zip")),
				must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_windows_amd64.zip")),
			},
		},
		{
			name:    "helm/helm",
			repo:    github.NewRepository("helm", "helm"),
			release: github.NewRelease("v3.16.2"),
			assets: []github.Asset{
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-darwin-amd64.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-darwin-arm64.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-386.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-arm.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-arm64.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-ppc64le.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-riscv64.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-s390x.tar.gz")),
				must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-windows-amd64.zip")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository()
			assets, err := repository.List(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestAssetRepositoryDownload(_ *testing.T) {
	// todo: implement this.
}
