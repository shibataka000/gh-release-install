package github

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationServiceForLinuxAmd64(t *testing.T) {
	tests := []struct {
		repo   string
		tag    string
		result FindResult
		test   *exec.Cmd
	}{
		{
			repo: "aquasecurity/trivy",
			tag:  "v0.53.0",
			result: newFindResult(
				must(parseAsset(176981043, "https://github.com/aquasecurity/trivy/releases/download/v0.53.0/trivy_0.53.0_Linux-64bit.tar.gz")),
				newExecBinary("trivy"),
			),
			test: exec.Command("./trivy", "version"),
		},
		{
			repo: "argoproj/argo-cd",
			tag:  "v2.9.18",
			result: newFindResult(
				must(parseAsset(177293568, "https://github.com/argoproj/argo-cd/releases/download/v2.9.18/argocd-linux-amd64")),
				newExecBinary("argocd"),
			),
			test: exec.Command("./argocd", "version", "--client"),
		},
		{
			repo: "argoproj/argo-rollouts",
			tag:  "v1.7.1",
			result: newFindResult(
				must(parseAsset(175717897, "https://github.com/argoproj/argo-rollouts/releases/download/v1.7.1/kubectl-argo-rollouts-linux-amd64")),
				newExecBinary("kubectl-argo-rollouts"),
			),
			test: exec.Command("./kubectl-argo-rollouts", "version"),
		},
		{
			repo: "argoproj/argo-workflows",
			tag:  "v3.5.8",
			result: newFindResult(
				must(parseAsset(174415137, "https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-amd64.gz")),
				newExecBinary("argo"),
			),
			test: exec.Command("./argo", "version"),
		},
		{
			repo: "buildpacks/pack",
			tag:  "v0.34.2",
			result: newFindResult(
				must(parseAsset(172104571, "https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux.tgz")),
				newExecBinary("pack"),
			),
			test: exec.Command("./pack", "version"),
		},
		{
			repo: "cli/cli",
			tag:  "v2.52.0",
			result: newFindResult(
				must(parseAsset(175682889, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
				newExecBinary("gh"),
			),
			test: exec.Command("./gh", "version"),
		},
		{
			repo: "getsops/sops",
			tag:  "v3.9.0",
			result: newFindResult(
				must(parseAsset(176438234, "https://github.com/getsops/sops/releases/download/v3.9.0/sops-v3.9.0.linux.amd64")),
				newExecBinary("sops"),
			),
			test: exec.Command("./sops", "--version"),
		},
		{
			repo: "goodwithtech/dockle",
			tag:  "v0.4.14",
			result: newFindResult(
				must(parseAsset(149683239, "https://github.com/goodwithtech/dockle/releases/download/v0.4.14/dockle_0.4.14_Linux-64bit.tar.gz")),
				newExecBinary("dockle"),
			),
			test: exec.Command("./dockle", "--version"),
		},
		{
			repo: "gravitational/teleport",
			tag:  "v16.4.6",
			result: newFindResult(
				must(parseAsset(0, "https://cdn.teleport.dev/teleport-v16.4.6-linux-amd64-bin.tar.gz")),
				newExecBinary("tsh"),
			),
			test: exec.Command("./tsh", "version"),
		},
		{
			repo: "hashicorp/terraform",
			tag:  "v1.9.0",
			result: newFindResult(
				must(parseAsset(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
				newExecBinary("terraform"),
			),
			test: exec.Command("./terraform", "version"),
		},
		{
			repo: "helm/helm",
			tag:  "v3.16.2",
			result: newFindResult(
				must(parseAsset(0, "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
				newExecBinary("helm"),
			),
			test: exec.Command("./helm", "version"),
		},
		{
			repo: "istio/istio",
			tag:  "1.22.2",
			result: newFindResult(
				must(parseAsset(176364493, "https://github.com/istio/istio/releases/download/1.22.2/istioctl-1.22.2-linux-amd64.tar.gz")),
				newExecBinary("istioctl"),
			),
			test: exec.Command("./istioctl", "version"),
		},
		{
			repo: "koalaman/shellcheck",
			tag:  "v0.10.0",
			result: newFindResult(
				must(parseAsset(155543215, "https://github.com/koalaman/shellcheck/releases/download/v0.10.0/shellcheck-v0.10.0.linux.x86_64.tar.xz")),
				newExecBinary("shellcheck"),
			),
			test: exec.Command("./shellcheck", "--version"),
		},
		{
			repo: "kubernetes/kubernetes",
			tag:  "v1.31.0",
			result: newFindResult(
				must(parseAsset(0, "https://dl.k8s.io/release/v1.31.0/bin/linux/amd64/kubectl")),
				newExecBinary("kubectl"),
			),
			test: exec.Command("./kubectl", "version", "--client"),
		},
		{
			repo: "mikefarah/yq",
			tag:  "v4.44.2",
			result: newFindResult(
				must(parseAsset(174040565, "https://github.com/mikefarah/yq/releases/download/v4.44.2/yq_linux_amd64")),
				newExecBinary("yq"),
			),
			test: exec.Command("./yq", "version"),
		},
		{
			repo: "open-policy-agent/conftest",
			tag:  "v0.53.0",
			result: newFindResult(
				must(parseAsset(172540735, "https://github.com/open-policy-agent/conftest/releases/download/v0.53.0/conftest_0.53.0_Linux_x86_64.tar.gz")),
				newExecBinary("conftest"),
			),
			test: exec.Command("./conftest", "--version"),
		},
		{
			repo: "open-policy-agent/gatekeeper",
			tag:  "v3.16.3",
			result: newFindResult(
				must(parseAsset(169950399, "https://github.com/open-policy-agent/gatekeeper/releases/download/v3.16.3/gator-v3.16.3-linux-amd64.tar.gz")),
				newExecBinary("gator"),
			),
			test: exec.Command("./gator", "version"),
		},
		{
			repo: "open-policy-agent/opa",
			tag:  "v0.66.0",
			result: newFindResult(
				must(parseAsset(176292835, "https://github.com/open-policy-agent/opa/releases/download/v0.66.0/opa_linux_amd64")),
				newExecBinary("opa"),
			),
			test: exec.Command("./opa", "version"),
		},
		{
			repo: "protocolbuffers/protobuf",
			tag:  "v27.2",
			result: newFindResult(
				must(parseAsset(175919234, "https://github.com/protocolbuffers/protobuf/releases/download/v27.2/protoc-27.2-linux-x86_64.zip")),
				newExecBinary("protoc"),
			),
			test: exec.Command("./protoc", "--version"),
		},
		{
			repo: "snyk/cli",
			tag:  "v1.1292.1",
			result: newFindResult(
				must(parseAsset(176276540, "https://github.com/snyk/cli/releases/download/v1.1292.1/snyk-linux")),
				newExecBinary("snyk"),
			),
			test: exec.Command("./snyk", "version"),
		},
		{
			repo: "starship/starship",
			tag:  "v1.19.0",
			result: newFindResult(
				must(parseAsset(168103285, "https://github.com/starship/starship/releases/download/v1.19.0/starship-x86_64-unknown-linux-gnu.tar.gz")),
				newExecBinary("starship"),
			),
			test: exec.Command("./starship", "--version"),
		},
		{
			repo: "viaduct-ai/kustomize-sops",
			tag:  "v4.3.2",
			result: newFindResult(
				must(parseAsset(176582858, "https://github.com/viaduct-ai/kustomize-sops/releases/download/v4.3.2/ksops_4.3.2_Linux_x86_64.tar.gz")),
				newExecBinary("ksops"),
			),
			test: exec.Command("test", "-f", "./ksops"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.repo, func(t *testing.T) {
			require := require.New(t)

			dir, err := os.MkdirTemp("", "")
			require.NoError(err)
			defer os.RemoveAll(dir)
			tt.test.Dir = dir

			before := cloneCommand(t, tt.test)
			require.Error(before.Run(), "executable binary was already installed")

			assetRepository, err := NewAssetRepository(tt.repo, io.Discard)
			require.NoError(err)
			execBinary := NewExecBinaryRepository()
			app := NewApplicationService(assetRepository, execBinary)

			ctx := context.Background()

			result, err := app.Find(ctx, tt.tag, DefaultPatterns)
			require.NoError(err)
			require.Equal(tt.result, result)

			if !testing.Short() {
				err := app.Install(ctx, result, dir)
				require.NoError(err)

				after := cloneCommand(t, tt.test)
				require.NoError(after.Run())
			}
		})
	}
}
