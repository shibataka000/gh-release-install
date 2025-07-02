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

// ExternalAssetTemplate is a template of [Asset] hosted on server other than GitHub.
type ExternalAssetTemplate struct {
	downloadURL *template.Template
}

// execute applies a [ExternalAssetTemplate] to [Release] object and returns [Asset] object.
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
		id:          0, // fake ID of [Asset] hosted on server other than GitHub.
		downloadURL: downloadURL,
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
	resp, err := http.Get(asset.downloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint:errcheck

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(r.progressBar).NewProxyReader(resp.Body)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
