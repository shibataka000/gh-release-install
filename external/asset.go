package external

import (
	"bytes"
	"context"
	"io"
	"maps"
	"net/http"
	"text/template"

	"github.com/cheggaaa/pb/v3"
	"github.com/shibataka000/gh-release-install/github"
)

// Exists returns true if given repository's assets are hosted on server outside from GitHub.
func Exists(repoFullName string) (bool, error) {
	repo, err := github.NewRepositoryFromFullName(repoFullName)
	if err != nil {
		return false, err
	}
	_, ok := defaultAssetTemplates[repo]
	return ok, nil
}

// AssetTemplate is a template of [github.com/shibataka000/gh-release-install/github.Asset].
type AssetTemplate struct {
	downloadURL *template.Template
}

// newAssetTemplate returns a new [AssetTemplate] object.
func newAssetTemplate(downloadURL *template.Template) AssetTemplate {
	return AssetTemplate{
		downloadURL: downloadURL,
	}
}

// mustNewAssetTemplateFromString returns a new [AssetTemplate] object.
// This is like [newAssetTemplate] but panics if the downloadURL cannot be parsed.
func mustNewAssetTemplateFromString(downloadURL string) AssetTemplate {
	return newAssetTemplate(template.Must(template.New("DownloadURL").Parse(downloadURL)))
}

// execute applies a [AssetTemplate] to [github.com/shibataka000/gh-release-install/github.Release] object and returns [github.com/shibataka000/gh-release-install/github.Asset] object.
func (a AssetTemplate) execute(release github.Release) (github.Asset, error) {
	var buf bytes.Buffer
	data := map[string]string{
		"Tag":    release.Tag,
		"SemVer": release.SemVer(),
	}
	if err := a.downloadURL.Execute(&buf, data); err != nil {
		return github.Asset{}, err
	}
	// ID should be 0. See [github.com/shibataka000/gh-release-install/github.Asset]'s comments for more details.
	return github.NewAssetFromString(0, buf.String())
}

// AssetRepository is a repository for [github.com/shibataka000/gh-release-install/github.Asset] and [github.com/shibataka000/gh-release-install/github.AssetContent].
type AssetRepository struct {
	templates map[github.Repository][]AssetTemplate
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository() *AssetRepository {
	return &AssetRepository{
		templates: maps.Clone(defaultAssetTemplates),
	}
}

// List GitHub release assets and returns them.
func (r *AssetRepository) List(_ context.Context, repo github.Repository, release github.Release) ([]github.Asset, error) {
	templates, ok := r.templates[repo]
	if !ok {
		return nil, ErrAssetTemplateNotFound
	}

	assets := []github.Asset{}
	for _, tmpl := range templates {
		asset, err := tmpl.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

// Download a GitHub release asset content and returns it.
// Progress bar is written into w.
func (r *AssetRepository) Download(_ context.Context, _ github.Repository, asset github.Asset, w io.Writer) (github.AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(w).NewProxyReader(resp.Body)

	return io.ReadAll(pr)
}
