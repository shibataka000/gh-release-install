package external

import "github.com/shibataka000/gh-release-install/github"

var defaultAssetTemplates = map[github.Repository]AssetTemplateList{
	github.NewRepository("gravitational", "teleport"): {},
	github.NewRepository("hashicorp", "terraform"): {
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_darwin_amd64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_darwin_arm64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_freebsd_386.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_freebsd_amd64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_freebsd_arm.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_linux_386.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_linux_amd64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_linux_arm.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_linux_arm64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_openbsd_386.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_openbsd_amd64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_solaris_amd64.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_windows_386.zip"),
		mustNewAssetTemplateFromString("https://releases.hashicorp.com/terraform/1.9.1/terraform_1.9.1_windows_amd64.zip"),
	},
	github.NewRepository("helm", "helm"): {
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-darwin-amd64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-darwin-arm64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-386.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-amd64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-arm.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-arm64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-ppc64le.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-riscv64.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-linux-s390x.tar.gz"),
		mustNewAssetTemplateFromString("https://get.helm.sh/helm-v3.15.2-windows-amd64.zip"),
	},
	github.NewRepository("kubernetes", "kubernetes"): {
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/v1.30.2/bin/darwin/amd64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/v1.30.2/bin/darwin/arm64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/v1.30.2/bin/linux/amd64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/v1.30.2/bin/linux/arm64/kubectl"),
		mustNewAssetTemplateFromString("https://dl.k8s.io/release/v1.30.2/bin/windows/amd64/kubectl.exe"),
	},
}
