package external

import "github.com/shibataka000/gh-release-install/github"

// defaultAssetTemplates are templates of release asset hosted on server outside from GitHub.
var defaultAssetTemplates = map[github.Repository][]AssetTemplate{
	github.NewRepository("hashicorp", "terraform"): {
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
	github.NewRepository("helm", "helm"): {
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-darwin-amd64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-darwin-arm64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-386.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-arm.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-arm64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-ppc64le.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-riscv64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-linux-s390x.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-{{.Tag}}-windows-amd64.zip"),
	},
	github.NewRepository("kubernetes", "kubernetes"): {
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/{{.Tag}}/bin/darwin/amd64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/{{.Tag}}/bin/darwin/arm64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/{{.Tag}}/bin/linux/amd64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/{{.Tag}}/bin/linux/arm64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/{{.Tag}}/bin/windows/amd64/kubectl.exe"),
	},
}
