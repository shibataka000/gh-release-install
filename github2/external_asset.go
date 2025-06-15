package github2

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

const externalAssetID = 0

func newExternalAsset(downloadURL *url.URL) Asset {
	return newAsset(externalAssetID, downloadURL)
}

func parseExternalAsset(downloadURL string) (Asset, error) {
	url, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newExternalAsset(url), nil
}

// ExternalAssetTemplate is a template for Asset hosted on server other than GitHub.
type ExternalAssetTemplate struct {
	downloadURL *template.Template
}

func newExternalAssetTemplate(downloadURL *template.Template) ExternalAssetTemplate {
	return ExternalAssetTemplate{
		downloadURL: downloadURL,
	}
}

func parseExternalAssetTemplate(downloadURL string) (ExternalAssetTemplate, error) {
	tmpl, err := template.New("DownloadURL").Parse(downloadURL)
	if err != nil {
		return ExternalAssetTemplate{}, err
	}
	return newExternalAssetTemplate(tmpl), nil
}

func (a ExternalAssetTemplate) execute(release Release) (Asset, error) {
	var buf bytes.Buffer
	data := map[string]string{
		"Tag":    release.tag,
		"SemVer": release.semVer(),
	}
	if err := a.downloadURL.Execute(&buf, data); err != nil {
		return Asset{}, err
	}
	return parseExternalAsset(buf.String())
}

// ExternalAssetRepository is a repository for Asset and AssetContent hosted on server other than GitHub.
type ExternalAssetRepository struct {
	templates []ExternalAssetTemplate
	stdout    io.Writer
}

func newExternalAssetRepository(templates []ExternalAssetTemplate, stdout io.Writer) *ExternalAssetRepository {
	return &ExternalAssetRepository{
		templates: slices.Clone(templates),
		stdout:    stdout,
	}
}

// List lists assets in a given release and returns them.
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

// Download downloads an asset content and returns it.
func (r *ExternalAssetRepository) Download(_ context.Context, asset Asset) (AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	total := resp.ContentLength
	pr := pb.Full.Start64(total).SetWriter(r.stdout).NewProxyReader(resp.Body)
	defer func() { _ = pr.Close() }()

	return io.ReadAll(pr)
}
