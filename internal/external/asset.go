package external

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"slices"
	"text/template"

	"github.com/cheggaaa/pb/v3"
	"github.com/shibataka000/gh-release-install/internal/github"
)

// AssetTemplates are templates of known release asset hosted on server other than GitHub.
var AssetTemplates = map[github.Repository][]AssetTemplate{
	// https://github.com/gravitational/teleport
	{
		Host:  "github.com",
		Owner: "gravitational",
		Name:  "teleport",
	}: {
		must(parseAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-amd64-bin.tar.gz")),
	},
	// https://github.com/hashicorp/terraform
	{
		Host:  "github.com",
		Owner: "hashicorp",
		Name:  "terraform",
	}: {
		must(parseAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
	},
	// https://github.com/helm/helm
	{
		Host:  "github.com",
		Owner: "helm",
		Name:  "helm",
	}: {
		must(parseAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz")),
	},
	// https://github.com/kubernetes/kubernetes
	{
		Host:  "github.com",
		Owner: "kubernetes",
		Name:  "kubernetes",
	}: {
		must(parseAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/linux/amd64/kubectl")),
	},
}

// AssetTemplate is a template of [Asset] hosted on server other than GitHub.
type AssetTemplate struct {
	downloadURL *template.Template
}

// parseAssetTemplate returns a new [AssetTemplate] object.
func parseAssetTemplate(downloadURL string) (AssetTemplate, error) {
	tmpl, err := template.New("DownloadURL").Parse(downloadURL)
	if err != nil {
		return AssetTemplate{}, err
	}
	return AssetTemplate{
		downloadURL: tmpl,
	}, nil
}

// execute applies an [AssetTemplate] to [Release] object and returns [Asset] object.
func (a AssetTemplate) execute(release github.Release) (github.Asset, error) {
	var buf bytes.Buffer
	data := map[string]string{
		"Tag":    release.Tag,
		"SemVer": release.SemVer(),
	}
	if err := a.downloadURL.Execute(&buf, data); err != nil {
		return github.Asset{}, err
	}
	downloadURL, err := url.Parse(buf.String())
	return github.Asset{
		ID:          0, // fake ID of [Asset] hosted on server other than GitHub.
		DownloadURL: downloadURL,
	}, err
}

// AssetRepository is a repository for [Asset] and [AssetContent] hosted on server other than GitHub.
type AssetRepository struct {
	templates   []AssetTemplate
	progressBar io.Writer // written progress bar into when downloading a GitHub release asset.
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository(templates []AssetTemplate, progressBar io.Writer) *AssetRepository {
	return &AssetRepository{
		templates:   slices.Clone(templates),
		progressBar: progressBar,
	}
}

// List lists GitHub release assets in a given GitHub release and returns them.
func (r *AssetRepository) List(_ context.Context, release github.Release) ([]github.Asset, error) {
	assets := []github.Asset{}
	for _, tmpl := range r.templates {
		asset, err := tmpl.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

// Download downloads a GitHub release asset content and returns it.
func (r *AssetRepository) Download(_ context.Context, asset github.Asset) (github.AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint:errcheck

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(r.progressBar).NewProxyReader(resp.Body)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}

// must is a helper that wraps a call to a function returning (E, error) and panics if the error is non-nil.
// This is intended for use in variable initializations.
func must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}
