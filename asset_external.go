package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"slices"
	"text/template"

	"github.com/cheggaaa/pb/v3"
)

// externalAssetTemplates are templates of known release asset hosted on server other than GitHub.
var externalAssetTemplates = map[Repository][]ExternalAssetTemplate{
	{
		Host:  "github.com",
		Owner: "gravitational",
		Name:  "teleport",
	}: {
		must(parseExternalAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-amd64-bin.tar.gz")),
	},
	{
		Host:  "github.com",
		Owner: "hashicorp",
		Name:  "terraform",
	}: {
		must(parseExternalAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
	},
	{
		Host:  "github.com",
		Owner: "helm",
		Name:  "helm",
	}: {
		must(parseExternalAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz")),
	},
	{
		Host:  "github.com",
		Owner: "kubernetes",
		Name:  "kubernetes",
	}: {
		must(parseExternalAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/linux/amd64/kubectl")),
	},
}

// ExternalAssetTemplate is a template of [Asset] hosted on server other than GitHub.
type ExternalAssetTemplate struct {
	downloadURL *template.Template
}

// parseExternalAssetTemplate returns a new [ExternalAssetTemplate] object.
func parseExternalAssetTemplate(downloadURL string) (ExternalAssetTemplate, error) {
	tmpl, err := template.New("DownloadURL").Parse(downloadURL)
	if err != nil {
		return ExternalAssetTemplate{}, err
	}
	return ExternalAssetTemplate{
		downloadURL: tmpl,
	}, nil
}

// execute applies an [ExternalAssetTemplate] to [Release] object and returns [Asset] object.
func (a ExternalAssetTemplate) execute(release Release) (Asset, error) {
	var buf bytes.Buffer
	data := map[string]string{
		"Tag":    release.Tag,
		"SemVer": release.semVer(),
	}
	if err := a.downloadURL.Execute(&buf, data); err != nil {
		return Asset{}, err
	}
	downloadURL, err := url.Parse(buf.String())
	return Asset{
		ID:          0, // fake ID of [Asset] hosted on server other than GitHub.
		DownloadURL: downloadURL,
	}, err
}

// ExternalAssetRepository is a repository for [Asset] and [AssetContent] hosted on server other than GitHub.
type ExternalAssetRepository struct {
	templates   []ExternalAssetTemplate
	progressBar io.Writer // written progress bar into when downloading a GitHub release asset.
}

// newExternalAssetRepository returns a new [ExternalAssetRepository] object.
func newExternalAssetRepository(templates []ExternalAssetTemplate, progressBar io.Writer) *ExternalAssetRepository {
	return &ExternalAssetRepository{
		templates:   slices.Clone(templates),
		progressBar: progressBar,
	}
}

// List lists GitHub release assets in a given GitHub release and returns them.
func (r *ExternalAssetRepository) List(_ context.Context, release Release) ([]Asset, error) {
	assets := []Asset{}
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
func (r *ExternalAssetRepository) Download(_ context.Context, asset Asset) (AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint:errcheck

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(r.progressBar).NewProxyReader(resp.Body)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
