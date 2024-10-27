package external

import (
	"context"
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/stretchr/testify/require"
)

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

func TestAssetTemplateListExecute(t *testing.T) {
	tests := []struct {
		name      string
		templates AssetTemplateList
		release   github.Release
		assets    []github.Asset
	}{
		{
			name:      "hashicorp/terraform",
			templates: defaultAssetTemplates[github.NewRepository("hashicorp", "terraform")],
			release:   github.NewRelease("v1.9.0"),
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
			name:      "helm/helm",
			templates: defaultAssetTemplates[github.NewRepository("helm", "helm")],
			release:   github.NewRelease("v3.16.2"),
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
			assets, err := tt.templates.execute(tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestAssetTemplateMapGet(t *testing.T) {
	tests := []struct {
		name      string
		templates AssetTemplateMap
		repo      github.Repository
		value     AssetTemplateList
	}{
		{
			name:      "hashicorp/terraform",
			templates: defaultAssetTemplates,
			repo:      github.NewRepository("hashicorp", "terraform"),
			value: AssetTemplateList{
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_arm64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_386.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_amd64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_arm.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_386.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_arm.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_arm64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_openbsd_386.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_openbsd_amd64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_solaris_amd64.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_386.zip"),
				mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			v, err := tt.templates.get(tt.repo)
			require.NoError(err)
			require.Equal(tt.value, v)
		})
	}
}

func TestAssetTemplateMapHas(t *testing.T) {
	tests := []struct {
		name      string
		templates AssetTemplateMap
		repo      github.Repository
		has       bool
	}{
		{
			name:      "Exists",
			templates: defaultAssetTemplates,
			repo:      github.NewRepository("hashicorp", "terraform"),
			has:       true,
		},
		{
			name:      "NotExist",
			templates: defaultAssetTemplates,
			repo:      github.NewRepository("", ""),
			has:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.has, tt.templates.has(tt.repo))
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

func TestAssetRepositoryHas(t *testing.T) {
	tests := []struct {
		name string
		repo github.Repository
		has  bool
	}{
		{
			name: "hashicorp/terraform",
			repo: github.NewRepository("hashicorp", "terraform"),
			has:  true,
		},
		{
			name: "NotExist",
			repo: github.NewRepository("", ""),
			has:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			repository := NewAssetRepository()
			require.Equal(tt.has, repository.Has(tt.repo))
		})
	}
}
