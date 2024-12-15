package github

import (
	"bytes"
	"context"
	"io"
	"maps"
	"net/http"
	"text/template"

	"github.com/cheggaaa/pb/v3"
)

// externalAssetID is fake ID of [Asset] hosted on server other than GitHub.
const externalAssetID = 0

// newExternalAssetFromString return a new [Asset] object hosted on server other than GitHub.
func newExternalAssetFromString(downloadURL string) (Asset, error) {
	return newAssetFromString(externalAssetID, downloadURL)
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

// newExternalAssetTemplateFromString returns a new [ExternalAssetTemplate] object.
func newExternalAssetTemplateFromString(downloadURL string) (ExternalAssetTemplate, error) {
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
	return newExternalAssetFromString(buf.String())
}

// AssetRepository is a repository for [Asset] and [AssetContent] hosted on server other than GitHub.
type ExternalAssetRepository struct {
	templates map[Repository][]ExternalAssetTemplate
}

// NewExternalAssetRepository returns a new [ExternalAssetRepository] object.
func NewExternalAssetRepository(templates map[Repository][]ExternalAssetTemplate) *ExternalAssetRepository {
	return &ExternalAssetRepository{
		templates: maps.Clone(templates),
	}
}

// list GitHub release assets in given GitHub release and returns them.
func (r *ExternalAssetRepository) list(_ context.Context, repo Repository, release Release) ([]Asset, error) {
	templates, ok := r.templates[repo]
	if !ok {
		return []Asset{}, nil
	}

	assets := []Asset{}
	for _, tmpl := range templates {
		asset, err := tmpl.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

// download a GitHub release asset content and returns it. Progress bar is written into w.
func (r *ExternalAssetRepository) download(_ context.Context, _ Repository, asset Asset, w io.Writer) (AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(w).NewProxyReader(resp.Body)

	return io.ReadAll(pr)
}
