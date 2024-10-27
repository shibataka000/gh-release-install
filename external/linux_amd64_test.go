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
			repoFullName: "hashicorp/terraform",
			tag:          "v1.9.0",
			asset:        must(github.NewAssetFromString(0, "https://releases.hashicorp.com/terraform/1.9.0/terraform_1.9.0_linux_amd64.zip")),
			execBinary:   github.NewExecBinary("terraform"),
			test:         exec.Command("./terraform", "version"),
		},
		{
			repoFullName: "helm/helm",
			tag:          "v3.16.2",
			asset:        must(github.NewAssetFromString(0, "https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
			execBinary:   github.NewExecBinary("helm"),
			test:         exec.Command("./helm", "version"),
		},
		{
			repoFullName: "kubernetes/kubernetes",
			tag:          "v1.31.0",
			asset:        must(github.NewAssetFromString(0, "https://dl.k8s.io/release/v1.31.0/bin/linux/amd64/kubectl")),
			execBinary:   github.NewExecBinary("kubectl"),
			test:         exec.Command("./kubectl", "version", "--client"),
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
