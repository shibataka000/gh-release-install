package external

import (
	"text/template"

	"github.com/shibataka000/gh-release-install/github"
)

func newAssetTemplate(text string) AssetTemplate {
	return AssetTemplate{downloadURL: template.Must(template.New("").Parse(text))}
}

// defaultExternalAssetTemplates are templates of known release asset hosted on server other than GitHub.
var defaultExternalAssetTemplates = map[github.Repository][]AssetTemplate{
	{Host: "github.com", Owner: "gravitational", Name: "teleport"}: {
		// Linux
		newAssetTemplate("https://cdn.teleport.dev/teleport-{{.SemVer}}-1.arm64.rpm"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-{{.SemVer}}-1.arm.rpm"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-{{.SemVer}}-1.i386.rpm"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-{{.SemVer}}-1.x86_64.rpm"),
		newAssetTemplate("https://cdn.teleport.dev/teleport_{{.SemVer}}_amd64.deb"),
		newAssetTemplate("https://cdn.teleport.dev/teleport_{{.SemVer}}_arm64.deb"),
		newAssetTemplate("https://cdn.teleport.dev/teleport_{{.SemVer}}_arm.deb"),
		newAssetTemplate("https://cdn.teleport.dev/teleport_{{.SemVer}}_i386.deb"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-datadog-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-datadog-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-discord-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-discord-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-email-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-email-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-jira-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-jira-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-mattermost-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-mattermost-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-msteams-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-msteams-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-pagerduty-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-pagerduty-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-slack-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-access-slack-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-event-handler-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-event-handler-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-386-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-amd64-centos7-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-arm-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/terraform-provider-teleport-v{{.SemVer}}-linux-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/terraform-provider-teleport-v{{.SemVer}}-linux-arm64-bin.tar.gz"),
		// macOS
		newAssetTemplate("https://cdn.teleport.dev/teleport-{{.SemVer}}.pkg"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-event-handler-v{{.SemVer}}-darwin-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-event-handler-v{{.SemVer}}-darwin-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-darwin-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-darwin-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-darwin-universal-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/terraform-provider-teleport-v{{.SemVer}}-darwin-amd64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/terraform-provider-teleport-v{{.SemVer}}-darwin-arm64-bin.tar.gz"),
		newAssetTemplate("https://cdn.teleport.dev/terraform-provider-teleport-v{{.SemVer}}-darwin-universal-bin.tar.gz"),
		// Windows
		newAssetTemplate("https://cdn.teleport.dev/Teleport%20Connect%20Setup-{{.SemVer}}.exe"),
	},
	{Host: "github.com", Owner: "hashicorp", Name: "terraform"}: {
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_arm64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_386.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_amd64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_arm.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_386.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_arm.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_arm64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_openbsd_386.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_openbsd_amd64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_solaris_amd64.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_386.zip"),
		newAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip"),
	},
	{Host: "github.com", Owner: "helm", Name: "helm"}: {
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-darwin-amd64.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-darwin-arm64.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-386.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-arm.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-arm64.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-ppc64le.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-riscv64.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-s390x.tar.gz"),
		newAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-windows-amd64.zip"),
	},
	{Host: "github.com", Owner: "kubernetes", Name: "kubernetes"}: {
		newAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/darwin/amd64/kubectl"),
		newAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/darwin/arm64/kubectl"),
		newAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/linux/amd64/kubectl"),
		newAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/linux/arm64/kubectl"),
		newAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/windows/amd64/kubectl.exe"),
	},
}
