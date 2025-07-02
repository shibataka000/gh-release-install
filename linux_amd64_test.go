package main_test

import (
	"context"
	"io"
	"net/url"
	"os"
	"os/exec"
	"testing"

	main "github.com/shibataka000/gh-release-install"
	"github.com/stretchr/testify/require"
)

func TestApplicationServiceForLinuxAmd64(t *testing.T) {
	tests := []struct {
		repo       string
		tag        string
		asset      main.Asset
		execBinary main.ExecBinary
		test       *exec.Cmd
	}{
		{
			repo: "aquasecurity/trivy",
			tag:  "v0.53.0",
			asset: main.Asset{
				ID:          176981043,
				DownloadURL: must(url.Parse("https://github.com/aquasecurity/trivy/releases/download/v0.53.0/trivy_0.53.0_Linux-64bit.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "trivy",
			},
			test: exec.Command("./trivy", "version"),
		},
		{
			repo: "argoproj/argo-cd",
			tag:  "v2.9.18",
			asset: main.Asset{
				ID:          177293568,
				DownloadURL: must(url.Parse("https://github.com/argoproj/argo-cd/releases/download/v2.9.18/argocd-linux-amd64")),
			},
			execBinary: main.ExecBinary{
				Name: "argocd",
			},
			test: exec.Command("./argocd", "version", "--client"),
		},
		{
			repo: "argoproj/argo-rollouts",
			tag:  "v1.7.1",
			asset: main.Asset{
				ID:          175717897,
				DownloadURL: must(url.Parse("https://github.com/argoproj/argo-rollouts/releases/download/v1.7.1/kubectl-argo-rollouts-linux-amd64")),
			},
			execBinary: main.ExecBinary{
				Name: "kubectl-argo-rollouts",
			},
			test: exec.Command("./kubectl-argo-rollouts", "version"),
		},
		{
			repo: "argoproj/argo-workflows",
			tag:  "v3.5.8",
			asset: main.Asset{
				ID:          174415137,
				DownloadURL: must(url.Parse("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-amd64.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "argo",
			},
			test: exec.Command("./argo", "version"),
		},
		{
			repo: "buildpacks/pack",
			tag:  "v0.34.2",
			asset: main.Asset{
				ID:          172104571,
				DownloadURL: must(url.Parse("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux.tgz")),
			},
			execBinary: main.ExecBinary{
				Name: "pack",
			},
			test: exec.Command("./pack", "version"),
		},
		{
			repo: "cli/cli",
			tag:  "v2.52.0",
			asset: main.Asset{
				ID:          175682889,
				DownloadURL: must(url.Parse("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "gh",
			},
			test: exec.Command("./gh", "version"),
		},
		{
			repo: "getsops/sops",
			tag:  "v3.9.0",
			asset: main.Asset{
				ID:          176438234,
				DownloadURL: must(url.Parse("https://github.com/getsops/sops/releases/download/v3.9.0/sops-v3.9.0.linux.amd64")),
			},
			execBinary: main.ExecBinary{
				Name: "sops",
			},
			test: exec.Command("./sops", "--version"),
		},
		{
			repo: "goodwithtech/dockle",
			tag:  "v0.4.14",
			asset: main.Asset{
				ID:          149683239,
				DownloadURL: must(url.Parse("https://github.com/goodwithtech/dockle/releases/download/v0.4.14/dockle_0.4.14_Linux-64bit.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "dockle",
			},
			test: exec.Command("./dockle", "--version"),
		},
		{
			repo: "gravitational/teleport",
			tag:  "v16.4.6",
			asset: main.Asset{
				ID:          0,
				DownloadURL: must(url.Parse("https://cdn.teleport.dev/teleport-v16.4.6-linux-amd64-bin.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "tsh",
			},
			test: exec.Command("./tsh", "version"),
		},
		{
			repo: "hashicorp/terraform",
			tag:  "v1.9.0",
			asset: main.Asset{
				ID:          0,
				DownloadURL: must(url.Parse("https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
			},
			execBinary: main.ExecBinary{
				Name: "terraform",
			},
			test: exec.Command("./terraform", "version"),
		},
		{
			repo: "helm/helm",
			tag:  "v3.16.2",
			asset: main.Asset{
				ID:          0,
				DownloadURL: must(url.Parse("https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "helm",
			},
			test: exec.Command("./helm", "version"),
		},
		{
			repo: "istio/istio",
			tag:  "1.22.2",
			asset: main.Asset{
				ID:          176364493,
				DownloadURL: must(url.Parse("https://github.com/istio/istio/releases/download/1.22.2/istioctl-1.22.2-linux-amd64.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "istioctl",
			},
			test: exec.Command("./istioctl", "version"),
		},
		{
			repo: "koalaman/shellcheck",
			tag:  "v0.10.0",
			asset: main.Asset{
				ID:          155543215,
				DownloadURL: must(url.Parse("https://github.com/koalaman/shellcheck/releases/download/v0.10.0/shellcheck-v0.10.0.linux.x86_64.tar.xz")),
			},
			execBinary: main.ExecBinary{
				Name: "shellcheck",
			},
			test: exec.Command("./shellcheck", "--version"),
		},
		{
			repo: "kubernetes/kubernetes",
			tag:  "v1.31.0",
			asset: main.Asset{
				ID:          0,
				DownloadURL: must(url.Parse("https://dl.k8s.io/release/v1.31.0/bin/linux/amd64/kubectl")),
			},
			execBinary: main.ExecBinary{
				Name: "kubectl",
			},
			test: exec.Command("./kubectl", "version", "--client"),
		},
		{
			repo: "mikefarah/yq",
			tag:  "v4.44.2",
			asset: main.Asset{
				ID:          174040565,
				DownloadURL: must(url.Parse("https://github.com/mikefarah/yq/releases/download/v4.44.2/yq_linux_amd64")),
			},
			execBinary: main.ExecBinary{
				Name: "yq",
			},
			test: exec.Command("./yq", "version"),
		},
		{
			repo: "open-policy-agent/conftest",
			tag:  "v0.53.0",
			asset: main.Asset{
				ID:          172540735,
				DownloadURL: must(url.Parse("https://github.com/open-policy-agent/conftest/releases/download/v0.53.0/conftest_0.53.0_Linux_x86_64.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "conftest",
			},
			test: exec.Command("./conftest", "--version"),
		},
		{
			repo: "open-policy-agent/gatekeeper",
			tag:  "v3.16.3",
			asset: main.Asset{
				ID:          169950399,
				DownloadURL: must(url.Parse("https://github.com/open-policy-agent/gatekeeper/releases/download/v3.16.3/gator-v3.16.3-linux-amd64.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "gator",
			},
			test: exec.Command("./gator", "version"),
		},
		{
			repo: "open-policy-agent/opa",
			tag:  "v0.66.0",
			asset: main.Asset{
				ID:          176292835,
				DownloadURL: must(url.Parse("https://github.com/open-policy-agent/opa/releases/download/v0.66.0/opa_linux_amd64")),
			},
			execBinary: main.ExecBinary{
				Name: "opa",
			},
			test: exec.Command("./opa", "version"),
		},
		{
			repo: "protocolbuffers/protobuf",
			tag:  "v27.2",
			asset: main.Asset{
				ID:          175919234,
				DownloadURL: must(url.Parse("https://github.com/protocolbuffers/protobuf/releases/download/v27.2/protoc-27.2-linux-x86_64.zip")),
			},
			execBinary: main.ExecBinary{
				Name: "protoc",
			},
			test: exec.Command("./protoc", "--version"),
		},
		{
			repo: "snyk/cli",
			tag:  "v1.1292.1",
			asset: main.Asset{
				ID:          176276540,
				DownloadURL: must(url.Parse("https://github.com/snyk/cli/releases/download/v1.1292.1/snyk-linux")),
			},
			execBinary: main.ExecBinary{
				Name: "snyk",
			},
			test: exec.Command("./snyk", "version"),
		},
		{
			repo: "starship/starship",
			tag:  "v1.19.0",
			asset: main.Asset{
				ID:          168103285,
				DownloadURL: must(url.Parse("https://github.com/starship/starship/releases/download/v1.19.0/starship-x86_64-unknown-linux-gnu.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "starship",
			},
			test: exec.Command("./starship", "--version"),
		},
		{
			repo: "viaduct-ai/kustomize-sops",
			tag:  "v4.3.2",
			asset: main.Asset{
				ID:          176582858,
				DownloadURL: must(url.Parse("https://github.com/viaduct-ai/kustomize-sops/releases/download/v4.3.2/ksops_4.3.2_Linux_x86_64.tar.gz")),
			},
			execBinary: main.ExecBinary{
				Name: "ksops",
			},
			test: exec.Command("test", "-f", "./ksops"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.repo, func(t *testing.T) {
			require := require.New(t)

			dir, err := os.MkdirTemp("", "")
			require.NoError(err)
			defer os.RemoveAll(dir) // nolint:errcheck
			tt.test.Dir = dir

			before := clone(t, tt.test)
			require.Error(before.Run(), "executable binary was already installed")

			assetRepository, err := main.NewAssetRepository(tt.repo, io.Discard)
			require.NoError(err)
			app := main.NewApplicationService(assetRepository, main.NewExecBinaryRepository(dir))

			ctx := context.Background()

			asset, execBinary, err := app.Find(ctx, tt.tag, main.DefaultPatterns)
			require.NoError(err)
			require.Equal(tt.asset, asset)
			require.Equal(tt.execBinary, execBinary)

			err = app.Install(ctx, asset, execBinary)
			require.NoError(err)

			after := clone(t, tt.test)
			require.NoError(after.Run())
		})
	}
}
