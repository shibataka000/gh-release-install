package external

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"text/template"

	"github.com/cheggaaa/pb/v3"
	"github.com/shibataka000/gh-release-install/github"
)

// newAssetFromString returns a new [github.com/shibataka000/gh-release-install/github.Asset] object.
func newAssetFromString(downloadURL string) (github.Asset, error) {
	// ID should be 0. See [github.com/shibataka000/gh-release-install/github.Asset]'s comments for more details.
	return github.NewAssetFromString(0, downloadURL)
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

// execute applies a [AssetTemplate] to [github.com/shibataka000/gh-release-install/github.Release] and returns [github.com/shibataka000/gh-release-install/github.Asset].
func (a AssetTemplate) execute(release github.Release) (github.Asset, error) {
	var buf bytes.Buffer
	data := map[string]string{
		"Tag":    release.Tag,
		"SemVer": release.SemVer(),
	}
	if err := a.downloadURL.Execute(&buf, data); err != nil {
		return github.Asset{}, err
	}
	return newAssetFromString(buf.String())
}

// AssetTemplateList is a list of [AssetTemplate].
type AssetTemplateList []AssetTemplate

// execute applies a each [AssetTemplate] element in [AssetTemplateList] to [github.com/shibataka000/gh-release-install/github.Release] and returns a list of [github.com/shibataka000/gh-release-install/github.Asset].
func (as AssetTemplateList) execute(release github.Release) ([]github.Asset, error) {
	assets := []github.Asset{}
	for _, t := range as {
		a, err := t.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}
	return assets, nil
}

// AssetTemplateMap is a map which contains [AssetTemplateList] for each [github.com/shibataka000/gh-release-install/github.Repository].
type AssetTemplateMap map[github.Repository]AssetTemplateList

// get a value.
func (m AssetTemplateMap) get(repo github.Repository) (AssetTemplateList, error) {
	val, ok := m[repo]
	if !ok {
		return nil, ErrNoAssetTemplatesFound
	}
	return val, nil
}

// has returns true if map contains given key.
func (m AssetTemplateMap) has(repo github.Repository) bool {
	_, ok := m[repo]
	return ok
}

// AssetRepository is a repository for [github.com/shibataka000/gh-release-install/github.Asset] and [github.com/shibataka000/gh-release-install/github.AssetContent].
type AssetRepository struct {
	templates AssetTemplateMap
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository() *AssetRepository {
	return &AssetRepository{
		templates: defaultAssetTemplates,
	}
}

// Has returns true if [AssetRepository] contains asset templates about given repository.
func (r *AssetRepository) Has(repo github.Repository) bool {
	return r.templates.has(repo)
}

// List GitHub release assets and returns them.
func (r *AssetRepository) List(ctx context.Context, repo github.Repository, release github.Release) ([]github.Asset, error) {
	templates, err := r.templates.get(repo)
	if err != nil {
		return nil, err
	}
	return templates.execute(release)
}

// Download a GitHub release asset content and returns it.
// Progress bar is written into w.
func (r *AssetRepository) Download(ctx context.Context, _ github.Repository, asset github.Asset, w io.Writer) (github.AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(w).NewProxyReader(resp.Body)

	return io.ReadAll(pr)
}
