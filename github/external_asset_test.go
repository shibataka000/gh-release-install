package github

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExternalAssetTemplateExecute(t *testing.T) {
	tests := []struct {
		name     string
		template ExternalAssetTemplate
		release  Release
		asset    Asset
	}{
		{
			name:     "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip",
			template: must(newExternalAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
			release:  newRelease("v1.9.0"),
			asset:    must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
		},
		{
			name:     "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz",
			template: must(newExternalAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz")),
			release:  newRelease("v3.16.2"),
			asset:    must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
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

func TestExternalAssetRepositoryList(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  []Asset
	}{
		{
			name:    "hashicorp/terraform",
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.9.0"),
			assets: []Asset{
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_darwin_amd64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_darwin_arm64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_386.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_amd64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_freebsd_arm.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_386.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_arm.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_arm64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_openbsd_386.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_openbsd_amd64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_solaris_amd64.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_windows_386.zip")),
				must(newAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_windows_amd64.zip")),
			},
		},
		{
			name:    "helm/helm",
			repo:    newRepository("helm", "helm"),
			release: newRelease("v3.16.2"),
			assets: []Asset{
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-darwin-amd64.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-darwin-arm64.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-386.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-arm.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-arm64.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-ppc64le.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-riscv64.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-s390x.tar.gz")),
				must(newAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-windows-amd64.zip")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			templates, ok := defaultExternalAssetTemplates[tt.repo]
			require.True(ok)
			repository := newExternalAssetRepository(templates, io.Discard)
			assets, err := repository.list(ctx, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestExternalAssetRepositoryDownload(_ *testing.T) {
	// todo: implement this.
}
