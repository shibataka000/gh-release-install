package github

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

// externalAssetID is fake ID of [Asset] hosted on server other than GitHub.
const externalAssetID = 0

// newExternalAsset return a new [Asset] object hosted on server other than GitHub.
func newExternalAsset(downloadURL *url.URL) Asset {
	return newAsset(externalAssetID, downloadURL)
}

// parseExternalAsset return a new [Asset] object hosted on server other than GitHub.
func parseExternalAsset(downloadURL string) (Asset, error) {
	url, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newExternalAsset(url), nil
}

// ExternalAssetTemplate is a template of [Asset] hosted on server other than GitHub.
type ExternalAssetTemplate struct {
	downloadURL *template.Template
}

// newExternalAssetTemplate returns a new [ExternalAssetTemplate] object.
func newExternalAssetTemplate(downloadURL *template.Template) ExternalAssetTemplate {
	return ExternalAssetTemplate{
		downloadURL: downloadURL,
	}
}

// parseExternalAssetTemplate returns a new [ExternalAssetTemplate] object.
func parseExternalAssetTemplate(downloadURL string) (ExternalAssetTemplate, error) {
	tmpl, err := template.New("DownloadURL").Parse(downloadURL)
	if err != nil {
		return ExternalAssetTemplate{}, err
	}
	return newExternalAssetTemplate(tmpl), nil
}

// execute applies a [ExternalAssetTemplate] to [Release] object and returns [Asset] object.
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

// ExternalAssetRepository is a repository for [Asset] and [AssetContent] hosted on server other than GitHub.
type ExternalAssetRepository struct {
	templates []ExternalAssetTemplate

	// stdout written progress bar into when downloading a GitHub release asset.
	stdout io.Writer
}

// newExternalAssetRepository returns a new [ExternalAssetRepository] object.
func newExternalAssetRepository(templates []ExternalAssetTemplate, stdout io.Writer) *ExternalAssetRepository {
	return &ExternalAssetRepository{
		templates: slices.Clone(templates),
		stdout:    stdout,
	}
}

// list GitHub release assets in a given GitHub release and returns them.
func (r *ExternalAssetRepository) list(_ context.Context, release Release) ([]Asset, error) {
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

// download a GitHub release asset content and returns it.
func (r *ExternalAssetRepository) download(_ context.Context, asset Asset) (AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint:errcheck

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(r.stdout).NewProxyReader(resp.Body)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
