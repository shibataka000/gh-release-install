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

// parseAsset returns a new [Asset] object.
func parseAsset(id int64, downloadURL string) (Asset, error) {
	parsed, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newAsset(id, parsed), nil
}

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

// list GitHub release assets in a given GitHub release and returns them.
func (r *GitHubAssetRepository) list(ctx context.Context, release Release) ([]Asset, error) {
	assets := []Asset{}

	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, r.repo.owner, r.repo.name, release.Tag)
	if err != nil {
		return nil, err
	}

	for page := 1; page != 0; {
		githubAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, r.repo.owner, r.repo.name, githubRelease.GetID(), &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, githubAsset := range githubAssets {
			asset, err := parseAsset(githubAsset.GetID(), githubAsset.GetBrowserDownloadURL())
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
		page = resp.NextPage
	}

	return assets, nil
}

// download a GitHub release asset content and returns it.
func (r *GitHubAssetRepository) download(ctx context.Context, asset Asset) (AssetContent, error) {
	githubAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, r.repo.owner, r.repo.name, asset.ID)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, r.repo.owner, r.repo.name, asset.ID, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close() // nolint:errcheck

	total := int64(githubAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(r.progressBar).NewProxyReader(rc)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
