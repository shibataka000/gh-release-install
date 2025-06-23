package github

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/cheggaaa/pb/v3"
	"github.com/cli/go-gh/v2/pkg/auth"
	gogithub "github.com/google/go-github/v67/github"
	"github.com/shibataka000/gh-release-install/github"
)

// AssetRepository is a repository for [Asset] and [AssetContent].
type AssetRepository struct {
	client      *gogithub.Client
	repo        github.Repository
	progressBar io.Writer // This is written progress bar into when downloading a GitHub release asset.
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository(repo github.Repository, progressBar io.Writer) *AssetRepository {
	token, _ := auth.TokenForHost(repo.Host)
	return &AssetRepository{
		client:      gogithub.NewClient(http.DefaultClient).WithAuthToken(token),
		repo:        repo,
		progressBar: progressBar,
	}
}

// list GitHub release assets in a given GitHub release and returns them.
func (r *AssetRepository) List(ctx context.Context, release github.Release) ([]github.Asset, error) {
	assets := []github.Asset{}

	repositoryRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, r.repo.Owner, r.repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	for page := 1; page != 0; {
		releaseAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, r.repo.Owner, r.repo.Name, repositoryRelease.GetID(), &gogithub.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, releaseAsset := range releaseAssets {
			asset, err := parse(releaseAsset.GetID(), releaseAsset.GetBrowserDownloadURL())
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
		page = resp.NextPage
	}

	return assets, nil
}

// Download a GitHub release asset content and returns it.
func (r *AssetRepository) Download(ctx context.Context, asset github.Asset) (github.AssetContent, error) {
	githubAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, r.repo.Owner, r.repo.Name, asset.ID)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, r.repo.Owner, r.repo.Name, asset.ID, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close() // nolint:errcheck

	total := int64(githubAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(r.progressBar).NewProxyReader(rc)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}

// parse returns a new [Asset] object.
func parse(id int64, downloadURL string) (github.Asset, error) {
	parsed, err := url.Parse(downloadURL)
	if err != nil {
		return github.Asset{}, err
	}
	return github.Asset{
		ID:          id,
		DownloadURL: parsed,
	}, nil
}
