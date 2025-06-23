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
	"github.com/shibataka000/gh-release-install/github"
)

// AssetTemplate is a template of [Asset] hosted on server other than GitHub.
type AssetTemplate struct {
	downloadURL *template.Template
}

// execute applies a [AssetTemplate] to [Release] object and returns [Asset] object.
func (a AssetTemplate) execute(release github.Release) (github.Asset, error) {
	var buf bytes.Buffer
	data := map[string]string{
		"Tag":    release.Tag,
		"SemVer": release.SemVer(),
	}
	if err := a.downloadURL.Execute(&buf, data); err != nil {
		return github.Asset{}, err
	}
	url, err := url.Parse(buf.String())
	if err != nil {
		return github.Asset{}, err
	}
	return github.Asset{
		ID:          0,
		DownloadURL: url,
	}, nil
}

// AssetRepository is a repository for [Asset] and [AssetContent] hosted on server other than GitHub.
type AssetRepository struct {
	templates   []AssetTemplate
	progressBar io.Writer // This is written progress bar into when downloading a GitHub release asset.
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository(stdout io.Writer) *AssetRepository {
	return &AssetRepository{
		templates:   slices.Clone(defaultExternalAssetTemplates),
		progressBar: stdout,
	}
}

// list GitHub release assets in a given GitHub release and returns them.
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

// download a GitHub release asset content and returns it.
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
