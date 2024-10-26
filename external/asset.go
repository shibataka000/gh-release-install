package external

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"text/template"

	"github.com/cheggaaa/pb/v3"
	"github.com/shibataka000/gh-release-install/github"
)

func newAsset(downloadURL *url.URL) github.Asset {
	return github.Asset{
		DownloadURL: downloadURL,
	}
}

type AssetTemplate struct {
	downloadURL *template.Template
}

func newAssetTemplate(downloadURL *template.Template) AssetTemplate {
	return AssetTemplate{
		downloadURL: downloadURL,
	}
}

func mustNewAssetTemplateFromString(downloadURL string) AssetTemplate {
	return newAssetTemplate(template.Must(template.New("DownloadURL").Parse(downloadURL)))
}

func (a AssetTemplate) execute(release github.Release) (github.Asset, error) {
	var buf bytes.Buffer
	if err := a.downloadURL.Execute(&buf, release); err != nil {
		return github.Asset{}, err
	}
	downloadURL, err := url.Parse(buf.String())
	if err != nil {
		return github.Asset{}, err
	}
	return newAsset(downloadURL), nil
}

type AssetTemplateList []AssetTemplate

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

type AssetRepository struct {
	templates AssetTemplateMap
}

func NewAssetRepository() *AssetRepository {
	return &AssetRepository{
		templates: defaultAssetTemplates,
	}
}

func (r *AssetRepository) Has(repo github.Repository) bool {
	return r.templates.has(repo)
}

func (r *AssetRepository) List(ctx context.Context, repo github.Repository, release github.Release) ([]github.Asset, error) {
	templates, err := r.templates.get(repo)
	if err != nil {
		return nil, err
	}
	return templates.execute(release)
}

func (r *AssetRepository) Download(ctx context.Context, _ github.Repository, asset github.Asset, w io.Writer) (github.AssetContent, error) {
	resp, err := http.Get(asset.DownloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := pb.Full.Start64(resp.ContentLength).SetWriter(w).NewProxyReader(resp.Body)

	return io.ReadAll(pr)
}

type AssetTemplateMap map[string]AssetTemplateList

func (m AssetTemplateMap) get(repo github.Repository) (AssetTemplateList, error) {
	key := m.key(repo)
	val, ok := m[key]
	if !ok {
		return nil, errors.New("")
	}
	return val, nil
}

func (m AssetTemplateMap) has(repo github.Repository) bool {
	key := m.key(repo)
	_, ok := m[key]
	return ok
}

func (m AssetTemplateMap) key(repo github.Repository) string {
	return fmt.Sprintf("%s/%s", repo.Owner, repo.Name)
}
