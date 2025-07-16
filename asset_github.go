package main

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/cheggaaa/pb/v3"
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/google/go-github/v67/github"
)

// GitHubAssetRepository is a repository for [Asset] and [AssetContent].
type GitHubAssetRepository struct {
	client      *github.Client
	repo        Repository
	progressBar io.Writer // written progress bar into when downloading a GitHub release asset.
}

// newGitHubAssetRepository returns a new [GitHubAssetRepository] object.
func newGitHubAssetRepository(repo Repository, progressBar io.Writer) *GitHubAssetRepository {
	token, _ := auth.TokenForHost(repo.Host)
	return &GitHubAssetRepository{
		client:      github.NewClient(http.DefaultClient).WithAuthToken(token),
		repo:        repo,
		progressBar: progressBar,
	}
}

// List lists GitHub release assets in a given GitHub release and returns them.
func (r *GitHubAssetRepository) List(ctx context.Context, release Release) ([]Asset, error) {
	assets := []Asset{}

	repositoryRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, r.repo.Owner, r.repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	for page := 1; page != 0; {
		releaseAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, r.repo.Owner, r.repo.Name, repositoryRelease.GetID(), &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, releaseAsset := range releaseAssets {
			downloadURL, err := url.Parse(releaseAsset.GetBrowserDownloadURL())
			if err != nil {
				return nil, err
			}
			assets = append(assets, Asset{
				ID:          releaseAsset.GetID(),
				DownloadURL: downloadURL,
			})
		}
		page = resp.NextPage
	}

	return assets, nil
}

// Download downloads a GitHub release asset content and returns it.
func (r *GitHubAssetRepository) Download(ctx context.Context, asset Asset) (AssetContent, error) {
	releaseAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, r.repo.Owner, r.repo.Name, asset.ID)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, r.repo.Owner, r.repo.Name, asset.ID, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close() // nolint:errcheck

	total := int64(releaseAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(r.progressBar).NewProxyReader(rc)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
