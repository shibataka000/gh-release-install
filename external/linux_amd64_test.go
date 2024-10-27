package external

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/stretchr/testify/require"
)

func TestApplicationServiceForLinuxAmd64(t *testing.T) {
	tests := []struct {
		repoFullName string
		tag          string
		asset        github.Asset
		execBinary   github.ExecBinary
		test         *exec.Cmd
	}{
		{
			repoFullName: "aquasecurity/trivy",
			tag:          "v0.53.0",
			asset:        must(github.NewAssetFromString(176981043, "https://github.com/aquasecurity/trivy/releases/download/v0.53.0/trivy_0.53.0_Linux-64bit.tar.gz")),
			execBinary:   github.NewExecBinary("trivy"),
			test:         exec.Command("./trivy", "version"),
		},
		{
			repoFullName: "argoproj/argo-cd",
			tag:          "v2.9.18",
			asset:        must(github.NewAssetFromString(177293568, "https://github.com/argoproj/argo-cd/releases/download/v2.9.18/argocd-linux-amd64")),
			execBinary:   github.NewExecBinary("argocd"),
			test:         exec.Command("./argocd", "version", "--client"),
		},
		{
			repoFullName: "argoproj/argo-rollouts",
			tag:          "v1.7.1",
			asset:        must(github.NewAssetFromString(175717897, "https://github.com/argoproj/argo-rollouts/releases/download/v1.7.1/kubectl-argo-rollouts-linux-amd64")),
			execBinary:   github.NewExecBinary("kubectl-argo-rollouts"),
			test:         exec.Command("./kubectl-argo-rollouts", "version"),
		},
		{
			repoFullName: "argoproj/argo-workflows",
			tag:          "v3.5.8",
			asset:        must(github.NewAssetFromString(174415137, "https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-amd64.gz")),
			execBinary:   github.NewExecBinary("argo"),
			test:         exec.Command("./argo", "version"),
		},
		{
			repoFullName: "buildpacks/pack",
			tag:          "v0.34.2",
			asset:        must(github.NewAssetFromString(172104571, "https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux.tgz")),
			execBinary:   github.NewExecBinary("pack"),
			test:         exec.Command("./pack", "version"),
		},
		{
			repoFullName: "cli/cli",
			tag:          "v2.52.0",
			asset:        must(github.NewAssetFromString(175682889, "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			execBinary:   github.NewExecBinary("gh"),
			test:         exec.Command("./gh", "version"),
		},
		{
			repoFullName: "getsops/sops",
			tag:          "v3.9.0",
			asset:        must(github.NewAssetFromString(176438234, "https://github.com/getsops/sops/releases/download/v3.9.0/sops-v3.9.0.linux.amd64")),
			execBinary:   github.NewExecBinary("sops"),
			test:         exec.Command("./sops", "--version"),
		},
		{
			repoFullName: "goodwithtech/dockle",
			tag:          "v0.4.14",
			asset:        must(github.NewAssetFromString(149683239, "https://github.com/goodwithtech/dockle/releases/download/v0.4.14/dockle_0.4.14_Linux-64bit.tar.gz")),
			execBinary:   github.NewExecBinary("dockle"),
			test:         exec.Command("./dockle", "--version"),
		},
		{
			repoFullName: "istio/istio",
			tag:          "1.22.2",
			asset:        must(github.NewAssetFromString(176364493, "https://github.com/istio/istio/releases/download/1.22.2/istioctl-1.22.2-linux-amd64.tar.gz")),
			execBinary:   github.NewExecBinary("istioctl"),
			test:         exec.Command("./istioctl", "version"),
		},
		{
			repoFullName: "mikefarah/yq",
			tag:          "v4.44.2",
			asset:        must(github.NewAssetFromString(174040565, "https://github.com/mikefarah/yq/releases/download/v4.44.2/yq_linux_amd64")),
			execBinary:   github.NewExecBinary("yq"),
			test:         exec.Command("./yq", "version"),
		},
		{
			repoFullName: "open-policy-agent/conftest",
			tag:          "v0.53.0",
			asset:        must(github.NewAssetFromString(172540735, "https://github.com/open-policy-agent/conftest/releases/download/v0.53.0/conftest_0.53.0_Linux_x86_64.tar.gz")),
			execBinary:   github.NewExecBinary("conftest"),
			test:         exec.Command("./conftest", "--version"),
		},
		{
			repoFullName: "open-policy-agent/gatekeeper",
			tag:          "v3.16.3",
			asset:        must(github.NewAssetFromString(169950399, "https://github.com/open-policy-agent/gatekeeper/releases/download/v3.16.3/gator-v3.16.3-linux-amd64.tar.gz")),
			execBinary:   github.NewExecBinary("gator"),
			test:         exec.Command("./gator", "version"),
		},
		{
			repoFullName: "open-policy-agent/opa",
			tag:          "v0.66.0",
			asset:        must(github.NewAssetFromString(176292835, "https://github.com/open-policy-agent/opa/releases/download/v0.66.0/opa_linux_amd64")),
			execBinary:   github.NewExecBinary("opa"),
			test:         exec.Command("./opa", "version"),
		},
		{
			repoFullName: "protocolbuffers/protobuf",
			tag:          "v27.2",
			asset:        must(github.NewAssetFromString(175919234, "https://github.com/protocolbuffers/protobuf/releases/download/v27.2/protoc-27.2-linux-x86_64.zip")),
			execBinary:   github.NewExecBinary("protoc"),
			test:         exec.Command("./protoc", "--version"),
		},
		{
			repoFullName: "snyk/cli",
			tag:          "v1.1292.1",
			asset:        must(github.NewAssetFromString(176276540, "https://github.com/snyk/cli/releases/download/v1.1292.1/snyk-linux")),
			execBinary:   github.NewExecBinary("snyk"),
			test:         exec.Command("./snyk", "version"),
		},
		{
			repoFullName: "starship/starship",
			tag:          "v1.19.0",
			asset:        must(github.NewAssetFromString(168103285, "https://github.com/starship/starship/releases/download/v1.19.0/starship-x86_64-unknown-linux-gnu.tar.gz")),
			execBinary:   github.NewExecBinary("starship"),
			test:         exec.Command("./starship", "--version"),
		},
		{
			repoFullName: "viaduct-ai/kustomize-sops",
			tag:          "v4.3.2",
			asset:        must(github.NewAssetFromString(176582858, "https://github.com/viaduct-ai/kustomize-sops/releases/download/v4.3.2/ksops_4.3.2_Linux_x86_64.tar.gz")),
			execBinary:   github.NewExecBinary("ksops"),
			test:         exec.Command("test", "-f", "./ksops"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.repoFullName, func(t *testing.T) {
			require := require.New(t)

			dir, err := os.MkdirTemp("", "")
			require.NoError(err)
			defer os.RemoveAll(dir)
			tt.test.Dir = dir

			before := cloneCommand(t, tt.test)
			require.Error(before.Run(), "executable binary was already installed")

			app := github.NewApplicationService(
				NewAssetRepository(),
				github.NewExecBinaryRepository(),
			)

			ctx := context.Background()

			asset, execBinary, err := app.Find(ctx, tt.repoFullName, tt.tag, DefaultPatterns)
			require.NoError(err)
			require.Equal(tt.asset, asset)
			require.Equal(tt.execBinary, execBinary)

			err = app.Install(ctx, tt.repoFullName, asset, execBinary, dir, io.Discard)
			require.NoError(err)

			after := cloneCommand(t, tt.test)
			require.NoError(after.Run())
		})
	}
}

// cloneCommand clones [exec.Cmd] and return it.
func cloneCommand(t *testing.T, cmd *exec.Cmd) *exec.Cmd {
	t.Helper()
	newCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	newCmd.Dir = cmd.Dir
	return newCmd
}
