package github

import (
	"fmt"
	"os"
	"testing"
)

func TestAsset(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		asset       string
		downloadURL string
		binaryName  string
		goos        string
		goarch      string
	}{
		{
			description: "shibataka000/go-get-release-test:v0.0.2(linux,amd64)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			asset:       "go-get-release_v0.0.2_linux_amd64",
			downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
			binaryName:  "go-get-release-test",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "shibataka000/go-get-release-test:v0.0.2(windows,amd64)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			asset:       "go-get-release_v0.0.2_windows_amd64.exe",
			downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe",
			binaryName:  "go-get-release-test.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "shibataka000/go-get-release-test:v0.0.2(darwin,amd64)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			asset:       "go-get-release_v0.0.2_darwin_amd64",
			downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64",
			binaryName:  "go-get-release-test",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "docker/compose:1.25.4(linux,amd64)",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			asset:       "docker-compose-Linux-x86_64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
			binaryName:  "docker-compose",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "docker/compose:1.25.4(windows,amd64)",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			asset:       "docker-compose-Windows-x86_64.exe",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
			binaryName:  "docker-compose.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "docker/compose:1.25.4(darwin,amd64)",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			asset:       "docker-compose-Darwin-x86_64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
			binaryName:  "docker-compose",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "docker/machine:v0.16.2(linux,amd64)",
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			asset:       "docker-machine-Linux-x86_64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
			binaryName:  "docker-machine",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "docker/machine:v0.16.2(windows,amd64)",
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			asset:       "docker-machine-Windows-x86_64.exe",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
			binaryName:  "docker-machine.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "docker/machine:v0.16.2(darwin,amd64)",
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			asset:       "docker-machine-Darwin-x86_64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
			binaryName:  "docker-machine",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "helm/helm(linux,amd64)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			asset:       "helm-v3.1.0-linux-amd64.tar.gz",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
			binaryName:  "helm",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "helm/helm(windows,amd64)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			asset:       "helm-v3.1.0-windows-amd64.zip",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
			binaryName:  "helm.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "helm/helm(darwin,amd64)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			asset:       "helm-v3.1.0-darwin-amd64.tar.gz",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
			binaryName:  "helm",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "istio/istio:1.6.0(linux,amd64)",
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			asset:       "istioctl-1.6.0-linux-amd64.tar.gz",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-linux-amd64.tar.gz",
			binaryName:  "istioctl",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "istio/istio:1.6.0(windows,amd64)",
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			asset:       "istioctl-1.6.0-win.zip",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-win.zip",
			binaryName:  "istioctl.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "istio/istio:1.6.0(darwin,amd64)",
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			asset:       "istioctl-1.6.0-osx.tar.gz",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-osx.tar.gz",
			binaryName:  "istioctl",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "hashicorp/terraform:v0.12.20(linux,amd64)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			asset:       "terraform_0.12.20_linux_amd64.zip",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
			binaryName:  "terraform",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "hashicorp/terraform:v0.12.20(windows,amd64)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			asset:       "terraform_0.12.20_windows_amd64.zip",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			binaryName:  "terraform.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "hashicorp/terraform:v0.12.20(darwin,amd64)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			asset:       "terraform_0.12.20_darwin_amd64.zip",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
			binaryName:  "terraform",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-cd:v1.4.2(linux,amd64)",
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v1.4.2",
			asset:       "argocd-linux-amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v1.4.2/argocd-linux-amd64",
			binaryName:  "argocd",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-cd:v1.4.2(darwin,amd64)",
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v1.4.2",
			asset:       "argocd-darwin-amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v1.4.2/argocd-darwin-amd64",
			binaryName:  "argocd",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "protocolbuffers/protobuf:v3.11.4(linux,amd64)",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			asset:       "protoc-3.11.4-linux-x86_64.zip",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			binaryName:  "protoc",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "protocolbuffers/protobuf:v3.11.4(windows,amd64)",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			asset:       "protoc-3.11.4-win64.zip",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
			binaryName:  "protoc.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "protocolbuffers/protobuf:v3.11.4(darwin,amd64)",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			asset:       "protoc-3.11.4-osx-x86_64.zip",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
			binaryName:  "protoc",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "mozilla/sops:v3.5.0(linux,amd64)",
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			asset:       "sops-v3.5.0.linux",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
			binaryName:  "sops",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "mozilla/sops:v3.5.0(windows,amd64)",
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			asset:       "sops-v3.5.0.exe",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
			binaryName:  "sops.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "mozilla/sops:v3.5.0(darwin,amd64)",
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			asset:       "sops-v3.5.0.darwin",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
			binaryName:  "sops",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "CircleCI-Public/circleci-cli:v0.1.8764(linux,amd64)",
			owner:       "CircleCI-Public",
			repo:        "circleci-cli",
			tag:         "v0.1.8764",
			asset:       "circleci-cli_0.1.8764_linux_amd64.tar.gz",
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_linux_amd64.tar.gz",
			binaryName:  "circleci",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "CircleCI-Public/circleci-cli:v0.1.8764(windows,amd64)",
			owner:       "CircleCI-Public",
			repo:        "circleci-cli",
			tag:         "v0.1.8764",
			asset:       "circleci-cli_0.1.8764_windows_amd64.zip",
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_windows_amd64.zip",
			binaryName:  "circleci.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "CircleCI-Public/circleci-cli:v0.1.8764(darwin,amd64)",
			owner:       "CircleCI-Public",
			repo:        "circleci-cli",
			tag:         "v0.1.8764",
			asset:       "circleci-cli_0.1.8764_darwin_amd64.tar.gz",
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_darwin_amd64.tar.gz",
			binaryName:  "circleci",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-rollouts:v0.9.0(linux,amd64)",
			owner:       "argoproj",
			repo:        "argo-rollouts",
			tag:         "v0.9.0",
			asset:       "kubectl-argo-rollouts-linux-amd64",
			downloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-amd64",
			binaryName:  "kubectl-argo-rollouts",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-rollouts:v0.9.0(darwin,amd64)",
			owner:       "argoproj",
			repo:        "argo-rollouts",
			tag:         "v0.9.0",
			asset:       "kubectl-argo-rollouts-darwin-amd64",
			downloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-amd64",
			binaryName:  "kubectl-argo-rollouts",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "open-policy-agent/conftest:v0.21.0(linux,amd64)",
			owner:       "open-policy-agent",
			repo:        "conftest",
			tag:         "v0.21.0",
			asset:       "conftest_0.21.0_Linux_x86_64.tar.gz",
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Linux_x86_64.tar.gz",
			binaryName:  "conftest",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "open-policy-agent/conftest:v0.21.0(windows,amd64)",
			owner:       "open-policy-agent",
			repo:        "conftest",
			tag:         "v0.21.0",
			asset:       "conftest_0.21.0_Windows_x86_64.zip",
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Windows_x86_64.zip",
			binaryName:  "conftest.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "open-policy-agent/conftest:v0.21.0(darwin,amd64)",
			owner:       "open-policy-agent",
			repo:        "conftest",
			tag:         "v0.21.0",
			asset:       "conftest_0.21.0_Darwin_x86_64.tar.gz",
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Darwin_x86_64.tar.gz",
			binaryName:  "conftest",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "goodwithtech/dockle:v0.3.1(linux,amd64)",
			owner:       "goodwithtech",
			repo:        "dockle",
			tag:         "v0.3.1",
			asset:       "dockle_0.3.1_Linux-64bit.tar.gz",
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Linux-64bit.tar.gz",
			binaryName:  "dockle",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "goodwithtech/dockle:v0.3.1(windows,amd64)",
			owner:       "goodwithtech",
			repo:        "dockle",
			tag:         "v0.3.1",
			asset:       "dockle_0.3.1_Windows-64bit.zip",
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Windows-64bit.zip",
			binaryName:  "dockle.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "goodwithtech/dockle:v0.3.1(darwin,amd64)",
			owner:       "goodwithtech",
			repo:        "dockle",
			tag:         "v0.3.1",
			asset:       "dockle_0.3.1_macOS-64bit.tar.gz",
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_macOS-64bit.tar.gz",
			binaryName:  "dockle",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "starship/starship:v0.47.1(linux,amd64)",
			owner:       "starship",
			repo:        "starship",
			tag:         "v0.47.0",
			asset:       "starship-x86_64-unknown-linux-gnu.tar.gz",
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-unknown-linux-gnu.tar.gz",
			binaryName:  "starship",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "starship/starship:v0.47.1(windows,amd64)",
			owner:       "starship",
			repo:        "starship",
			tag:         "v0.47.0",
			asset:       "starship-x86_64-pc-windows-msvc.zip",
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-pc-windows-msvc.zip",
			binaryName:  "starship.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "starship/starship:v0.47.1(darwin,amd64)",
			owner:       "starship",
			repo:        "starship",
			tag:         "v0.47.0",
			asset:       "starship-x86_64-apple-darwin.tar.gz",
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-apple-darwin.tar.gz",
			binaryName:  "starship",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "viaduct-ai/kustomize-sops:v2.3.3(linux,amd64)",
			owner:       "viaduct-ai",
			repo:        "kustomize-sops",
			tag:         "v2.3.3",
			asset:       "ksops_2.3.3_Linux_x86_64.tar.gz",
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Linux_x86_64.tar.gz",
			binaryName:  "ksops",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "viaduct-ai/kustomize-sops:v2.3.3(windows,amd64)",
			owner:       "viaduct-ai",
			repo:        "kustomize-sops",
			tag:         "v2.3.3",
			asset:       "ksops_2.3.3_Windows_x86_64.tar.gz",
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Windows_x86_64.tar.gz",
			binaryName:  "ksops.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "viaduct-ai/kustomize-sops:v2.3.3(darwin,amd64)",
			owner:       "viaduct-ai",
			repo:        "kustomize-sops",
			tag:         "v2.3.3",
			asset:       "ksops_2.3.3_Darwin_x86_64.tar.gz",
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Darwin_x86_64.tar.gz",
			binaryName:  "ksops",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "fluxcd/flux2:v0.8.0(linux,amd64)",
			owner:       "fluxcd",
			repo:        "flux2",
			tag:         "v0.8.0",
			asset:       "flux_0.8.0_linux_amd64.tar.gz",
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_linux_amd64.tar.gz",
			binaryName:  "flux",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "fluxcd/flux2:v0.8.0(windows,amd64)",
			owner:       "fluxcd",
			repo:        "flux2",
			tag:         "v0.8.0",
			asset:       "flux_0.8.0_windows_amd64.zip",
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_windows_amd64.zip",
			binaryName:  "flux.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "fluxcd/flux2:v0.8.0(darwin,amd64)",
			owner:       "fluxcd",
			repo:        "flux2",
			tag:         "v0.8.0",
			asset:       "flux_0.8.0_darwin_amd64.tar.gz",
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_darwin_amd64.tar.gz",
			binaryName:  "flux",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "mikefarah/yq:v4.7.1(linux,amd64)",
			owner:       "mikefarah",
			repo:        "yq",
			tag:         "v4.7.1",
			asset:       "yq_linux_amd64",
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_linux_amd64",
			binaryName:  "yq",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "mikefarah/yq:v4.7.1(windows,amd64)",
			owner:       "mikefarah",
			repo:        "yq",
			tag:         "v4.7.1",
			asset:       "yq_windows_amd64.exe",
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_windows_amd64.exe",
			binaryName:  "yq.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "mikefarah/yq:v4.7.1(darwin,amd64)",
			owner:       "mikefarah",
			repo:        "yq",
			tag:         "v4.7.1",
			asset:       "yq_darwin_amd64",
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_darwin_amd64",
			binaryName:  "yq",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "aquasecurity/trivy:v0.17.2(linux,amd64)",
			owner:       "aquasecurity",
			repo:        "trivy",
			tag:         "v0.17.2",
			asset:       "trivy_0.17.2_Linux-64bit.tar.gz",
			downloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_Linux-64bit.tar.gz",
			binaryName:  "trivy",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "aquasecurity/trivy:v0.17.2(darwin,amd64)",
			owner:       "aquasecurity",
			repo:        "trivy",
			tag:         "v0.17.2",
			asset:       "trivy_0.17.2_macOS-64bit.tar.gz",
			downloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_macOS-64bit.tar.gz",
			binaryName:  "trivy",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "aws/amazon-ec2-instance-selector:v2.0.2(linux,amd64)",
			owner:       "aws",
			repo:        "amazon-ec2-instance-selector",
			tag:         "v2.0.2",
			asset:       "ec2-instance-selector-linux-amd64",
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-linux-amd64",
			binaryName:  "ec2-instance-selector",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "aws/amazon-ec2-instance-selector:v2.0.2(windows,amd64)",
			owner:       "aws",
			repo:        "amazon-ec2-instance-selector",
			tag:         "v2.0.2",
			asset:       "ec2-instance-selector-windows-amd64",
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-windows-amd64",
			binaryName:  "ec2-instance-selector.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "aws/amazon-ec2-instance-selector:v2.0.2(darwin,amd64)",
			owner:       "aws",
			repo:        "amazon-ec2-instance-selector",
			tag:         "v2.0.2",
			asset:       "ec2-instance-selector-darwin-amd64",
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-darwin-amd64",
			binaryName:  "ec2-instance-selector",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-workflows:v3.0.7(linux,amd64)",
			owner:       "argoproj",
			repo:        "argo-workflows",
			tag:         "v3.0.7",
			asset:       "argo-linux-amd64.gz",
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-linux-amd64.gz",
			binaryName:  "argo",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-workflows:v3.0.7(windows,amd64)",
			owner:       "argoproj",
			repo:        "argo-workflows",
			tag:         "v3.0.7",
			asset:       "argo-windows-amd64.gz",
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-windows-amd64.gz",
			binaryName:  "argo.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-workflows:v3.0.7(darwin,amd64)",
			owner:       "argoproj",
			repo:        "argo-workflows",
			tag:         "v3.0.7",
			asset:       "argo-darwin-amd64.gz",
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-darwin-amd64.gz",
			binaryName:  "argo",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "open-policy-agent/opa:v0.29.4(linux,amd64)",
			owner:       "open-policy-agent",
			repo:        "opa",
			tag:         "v0.29.4",
			asset:       "opa_linux_amd64",
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_linux_amd64",
			binaryName:  "opa",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "open-policy-agent/opa:v0.29.4(windows,amd64)",
			owner:       "open-policy-agent",
			repo:        "opa",
			tag:         "v0.29.4",
			asset:       "opa_windows_amd64.exe",
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_windows_amd64.exe",
			binaryName:  "opa.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "open-policy-agent/opa:v0.29.4(darwin,amd64)",
			owner:       "open-policy-agent",
			repo:        "opa",
			tag:         "v0.29.4",
			asset:       "opa_darwin_amd64",
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_darwin_amd64",
			binaryName:  "opa",
			goos:        "darwin",
			goarch:      "amd64",
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo, err := c.Repository(tt.owner, tt.repo)
			if err != nil {
				t.Fatal(err)
			}
			release, err := repo.Release(tt.tag)
			if err != nil {
				t.Fatal(err)
			}
			asset, err := release.Asset(tt.goos, tt.goarch)
			if err != nil {
				t.Fatal(err)
			}
			if asset.Name() != tt.asset || asset.DownloadURL() != tt.downloadURL || asset.BinaryName() != tt.binaryName {
				t.Fatalf("Expected is {%s %s %s} but actual is {%s %s %s}", tt.asset, tt.downloadURL, tt.binaryName, asset.Name(), asset.DownloadURL(), asset.BinaryName())
			}
		})
	}
}

func TestIsSpecialAsset(t *testing.T) {
	tests := []struct {
		owner          string
		repo           string
		isSpecialAsset bool
	}{
		{
			owner:          "hashicorp",
			repo:           "terraform",
			isSpecialAsset: true,
		},
		{
			owner:          "shibataka000",
			repo:           "go-get-release",
			isSpecialAsset: false,
		},
	}

	for _, tt := range tests {
		description := fmt.Sprintf("%s/%s", tt.owner, tt.repo)
		t.Run(description, func(t *testing.T) {
			actual := isSpecialAsset(tt.owner, tt.repo)
			if tt.isSpecialAsset != actual {
				t.Fatalf("Expected is %t but actual is %t", tt.isSpecialAsset, actual)
			}
		})
	}
}

func TestGetGoosAndGoarchByAsset(t *testing.T) {
	tests := []struct {
		asset  string
		goos   string
		goarch string
	}{
		{
			asset:  "argo-linux-amd64",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "argo-windows-amd64",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "argo-darwin-amd64",
			goos:   "darwin",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Linux-x86_64",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Windows-x86_64.exe",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Darwin-x86_64",
			goos:   "darwin",
			goarch: "amd64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.asset, func(t *testing.T) {
			goos, err := getGoosByAsset(tt.asset)
			if err != nil {
				t.Fatal(err)
			}
			if goos != tt.goos {
				t.Fatalf("Expected is %v but actual is %v", tt.goos, goos)
			}
			goarch, err := getGoarchByAsset(tt.asset)
			if err != nil {
				t.Fatal(err)
			}
			if goarch != tt.goarch {
				t.Fatalf("Expected is %v but actual is %v", tt.goarch, goarch)
			}
		})
	}
}