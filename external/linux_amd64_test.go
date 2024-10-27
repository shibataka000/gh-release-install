package external

import (
	"context"
	"io"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/stretchr/testify/require"
)

func TestDefaultCorePatterns(t *testing.T) {
	for k := range DefaultCorePatterns {
		name := k
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			re, err := regexp.Compile(k)
			require.NoError(err)
			prefix, _ := re.LiteralPrefix()
			require.Equal("", prefix)
		})
	}
}

func TestDefaultExtPatterns(t *testing.T) {
	for k := range DefaultExtPatterns {
		name := k
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			re, err := regexp.Compile(k)
			require.NoError(err)
			prefix, _ := re.LiteralPrefix()
			require.Greater(len(prefix), 0)
		})
	}
}

func TestApplicationServiceForLinuxAmd64(t *testing.T) {
	tests := []struct {
		repoFullName string
		tag          string
		asset        github.Asset
		execBinary   github.ExecBinary
		test         *exec.Cmd
	}{
		{
			repoFullName: "gravitational/teleport",
			tag:          "v16.4.6",
			asset:        must(github.NewAssetFromString(0, "https://cdn.teleport.dev/teleport-v16.4.6-linux-amd64-bin.tar.gz")),
			execBinary:   github.NewExecBinary("tsh"),
			test:         exec.Command("./tsh", "version"),
		},
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
