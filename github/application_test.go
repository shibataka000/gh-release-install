package github

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationServiceFind(_ *testing.T) {
	// todo: implement this.
}

func TestApplicationServiceInstall(_ *testing.T) {
	// todo: implement this.
}

func TestApplicationServiceListAssets(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  []Asset
	}{
		{
			name:    "helm/helm",
			repo:    newRepository("helm", "helm"),
			release: newRelease("v3.16.2"),
			assets: []Asset{
				must(newAssetFromString(197985305, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-darwin-amd64.tar.gz.asc")),
				must(newAssetFromString(197985309, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-darwin-amd64.tar.gz.sha256.asc")),
				must(newAssetFromString(197985311, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-darwin-amd64.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985312, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-darwin-arm64.tar.gz.asc")),
				must(newAssetFromString(197985314, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-darwin-arm64.tar.gz.sha256.asc")),
				must(newAssetFromString(197985317, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-darwin-arm64.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985318, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-386.tar.gz.asc")),
				must(newAssetFromString(197985320, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-386.tar.gz.sha256.asc")),
				must(newAssetFromString(197985322, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-386.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985323, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-amd64.tar.gz.asc")),
				must(newAssetFromString(197985328, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-amd64.tar.gz.sha256.asc")),
				must(newAssetFromString(197985331, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-amd64.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985332, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-arm.tar.gz.asc")),
				must(newAssetFromString(197985333, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-arm.tar.gz.sha256.asc")),
				must(newAssetFromString(197985335, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-arm.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985337, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-arm64.tar.gz.asc")),
				must(newAssetFromString(197985338, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-arm64.tar.gz.sha256.asc")),
				must(newAssetFromString(197985339, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-arm64.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985344, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-ppc64le.tar.gz.asc")),
				must(newAssetFromString(197985345, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-ppc64le.tar.gz.sha256.asc")),
				must(newAssetFromString(197985346, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-ppc64le.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985348, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-riscv64.tar.gz.asc")),
				must(newAssetFromString(197985349, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-riscv64.tar.gz.sha256.asc")),
				must(newAssetFromString(197985350, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-riscv64.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985354, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-s390x.tar.gz.asc")),
				must(newAssetFromString(197985355, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-s390x.tar.gz.sha256.asc")),
				must(newAssetFromString(197985357, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-linux-s390x.tar.gz.sha256sum.asc")),
				must(newAssetFromString(197985360, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-windows-amd64.zip.asc")),
				must(newAssetFromString(197985362, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-windows-amd64.zip.sha256.asc")),
				must(newAssetFromString(197985364, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-windows-amd64.zip.sha256sum.asc")),
				must(newAssetFromString(197985367, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-windows-arm64.zip.asc")),
				must(newAssetFromString(197985370, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-windows-arm64.zip.sha256.asc")),
				must(newAssetFromString(197985372, "https://github.com/helm/helm/releases/download/v3.16.2/helm-v3.16.2-windows-arm64.zip.sha256sum.asc")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-darwin-amd64.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-darwin-arm64.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-386.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-amd64.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-arm.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-arm64.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-ppc64le.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-riscv64.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-linux-s390x.tar.gz")),
				must(newExternalAssetFromString("https://get.helm.sh/helm-v3.16.2-windows-amd64.zip")),
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			app := NewApplicationService(
				NewAssetRepository(githubTokenForTest),
				NewExternalAssetRepository(DefaultExternalAssetTemplates),
				NewExecBinaryRepository(),
			)
			assets, err := app.listAssets(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestApplicationServiceDonwload(_ *testing.T) {
	// todo: implement this.
}
