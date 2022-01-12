package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-github/github"
)

// Release in GitHub repository
type Release interface {
	Tag() string
	Version() string
	Assets() ([]Asset, error)
	Asset(name string) (Asset, error)
	FindAssetByPlatform(goos string, goarch string) (Asset, error)
}

type release struct {
	client *client
	repo   *repository
	tag    string
	id     int64
}

type TeleportReleasesOutput struct {
	Next       int               `json:"next"`
	Prev       int               `json:"prev"`
	First      int               `json:"first"`
	Last       int               `json:"last"`
	Items      []TeleportRelease `json:"items"`
	Prerelease bool              `json:"prerelease"`
}

type TeleportRelease struct {
	ID          string          `json:"id"`
	Version     string          `json:"version"`
	Description string          `json:"description"`
	PublishedAt string          `json:"publishedAt"`
	Downloads   []TeleportAsset `json:"downloads"`
}

type TeleportAsset struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	DisplaySize string `json:"displaySize"`
	Sha256      string `json:"sha256"`
}

// Tag return tag name of GitHub release
func (r *release) Tag() string {
	return r.tag
}

// Version return version string, which does not have 'v' prefix
func (r *release) Version() string {
	return strings.TrimLeft(r.Tag(), "v")
}

// Assets return assets in release
func (r *release) Assets() ([]Asset, error) {
	switch {
	case r.repo.Owner() == "hashicorp" && r.repo.Name() == "terraform":
		return r.terraformAssets()
	case r.repo.Owner() == "helm" && r.repo.Name() == "helm":
		return r.helmAssets()
	case r.repo.Owner() == "gravitational" && r.repo.Name() == "teleport":
		return r.teleportAssets()
	default:
		return r.assets()
	}
}

// assets return assets in GitHub release
func (r *release) assets() ([]Asset, error) {
	c := r.client
	assets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, r.repo.owner, r.repo.name, r.id, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}

	result := []Asset{}
	for _, a := range assets {
		result = append(result, &asset{
			client:      r.client,
			repo:        r.repo,
			release:     r,
			downloadURL: a.GetBrowserDownloadURL(),
		})
	}
	return result, nil
}

// terraformAssets return hashicorp/terraform's assets
func (r *release) terraformAssets() ([]Asset, error) {
	baseURL, err := url.Parse("https://releases.hashicorp.com")
	if err != nil {
		return nil, err
	}
	baseURL.Path = path.Join(baseURL.Path, "terraform", r.Version())

	doc, err := newGoqueryDocument(baseURL.String())
	if err != nil {
		return nil, err
	}

	assetNames := []string{}
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		assetName := s.Find("a").Text()
		if assetName != "../" {
			assetNames = append(assetNames, assetName)
		}
	})

	results := []Asset{}
	for _, assetName := range assetNames {
		downloadURL, err := url.Parse(baseURL.String())
		if err != nil {
			return nil, err
		}
		downloadURL.Path = path.Join(downloadURL.Path, assetName)
		results = append(results, &asset{
			client:      r.client,
			repo:        r.repo,
			release:     r,
			downloadURL: downloadURL.String(),
		})
	}
	return results, nil
}

// newGoqueryDocument return new goquery.Document object by URL
func newGoqueryDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return goquery.NewDocumentFromReader(res.Body)
}

// helmAssets return helm/helm's assets
func (r *release) helmAssets() ([]Asset, error) {
	baseURL, err := url.Parse("https://get.helm.sh")
	if err != nil {
		return nil, err
	}

	assets, err := r.assets()
	if err != nil {
		return nil, err
	}

	result := []Asset{}
	for _, a := range assets {
		assetName := strings.TrimRight(a.Name(), ".asc")
		if strings.Contains(assetName, "sha256") {
			continue
		}
		downloadURL, err := url.Parse(baseURL.String())
		if err != nil {
			return nil, err
		}
		downloadURL.Path = path.Join(downloadURL.Path, assetName)
		result = append(result, &asset{
			client:      r.client,
			repo:        r.repo,
			release:     r,
			downloadURL: downloadURL.String(),
		})
	}
	return result, nil
}

// teleportAssets return gravitational/teleport's assets
func (r *release) teleportAssets() ([]Asset, error) {
	var teleportRelease TeleportRelease
	found := false

	page := 0
	for {
		req, err := http.NewRequest("GET", "https://dashboard.gravitational.com/webapi/releases-oss", nil)
		if err != nil {
			return nil, err
		}

		params := req.URL.Query()
		params.Add("product", "teleport")
		params.Add("page", fmt.Sprintf("%d", page))
		req.URL.RawQuery = params.Encode()

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		var out TeleportReleasesOutput
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &out)
		if err != nil {
			return nil, err
		}

		for _, tr := range out.Items {
			if tr.Version == r.Tag() {
				teleportRelease = tr
				found = true
				break
			}
		}

		if found || out.Last == page {
			break
		}

		page = out.Next
	}

	if !found {
		return nil, fmt.Errorf("not found")
	}

	result := []Asset{}
	for _, a := range teleportRelease.Downloads {
		result = append(result, &asset{
			client:      r.client,
			repo:        r.repo,
			release:     r,
			downloadURL: a.URL,
		})
	}
	return result, nil
}

// Asset return asset which have specific name
func (r *release) Asset(name string) (Asset, error) {
	assets, err := r.Assets()
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		if a.Name() == name {
			return a, nil
		}
	}
	return nil, fmt.Errorf("no asset found: %s", name)
}

// FindAssetByPlatform return asset which have specific goos and goarch
func (r *release) FindAssetByPlatform(goos, goarch string) (Asset, error) {
	assetName, err := r.renderPredefinedAssetName(goos, goarch)
	if err == nil {
		return r.Asset(assetName)
	}

	assets, err := r.listAssetsFilteredByPlatform(goos, goarch)
	if err != nil {
		return nil, err
	}

	switch len(assets) {
	case 0:
		return nil, fmt.Errorf("no asset found")
	case 1:
		return assets[0], nil
	case 2:
		if assets[0].IsExecBinary() && assets[1].IsArchived() {
			return assets[0], nil
		} else if assets[0].IsArchived() && assets[1].IsExecBinary() {
			return assets[1], nil
		}
		fallthrough
	default:
		assetNames := []string{}
		for _, a := range assets {
			assetNames = append(assetNames, a.Name())
		}
		return nil, fmt.Errorf("too many assets found: %v", strings.Join(assetNames, ", "))
	}
}

// renderPredefinedAssetName return asset name which is predefined in assetNameMap
func (r *release) renderPredefinedAssetName(goos, goarch string) (string, error) {
	key := fmt.Sprintf("%s/%s", r.repo.Owner(), r.repo.Name())
	tmplMap, ok := assetNameMap[key]
	if !ok {
		return "", fmt.Errorf("asset name is not predefined: %s", key)
	}

	key = fmt.Sprintf("%s/%s", goos, goarch)
	tmplStr := ""
	if s, ok := tmplMap["default"]; ok {
		tmplStr = s
	}
	if s, ok := tmplMap[key]; ok {
		tmplStr = s
	}
	if tmplStr == "" {
		return "", fmt.Errorf("unsupported GOOS and GOARCH: %s", key)
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("assetName").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buf, struct {
		Tag     string
		Version string
		Goos    string
		Goarch  string
	}{
		Tag:     r.tag,
		Version: r.Version(),
		Goos:    goos,
		Goarch:  goarch,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// listAssetsFilteredByPlatform return assets which have specific goos and goarch
func (r *release) listAssetsFilteredByPlatform(goos, goarch string) ([]Asset, error) {
	var result []Asset
	assets, err := r.Assets()
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		if !a.ContainReleaseBinary() {
			continue
		}
		assetGoos, err := a.Goos()
		if err != nil {
			continue
		}
		assetGoarch, err := a.Goarch()
		if err != nil {
			continue
		}
		if assetGoos == goos && assetGoarch == goarch {
			result = append(result, a)
		}
	}
	return result, nil
}
