package github

import (
	"context"
	"io"
)

// ApplicationService.
type ApplicationService struct {
	asset         *AssetRepository
	externalAsset *ExternalAssetRepository
	execBinary    *ExecBinaryRepository
}

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset *AssetRepository, externalAsset *ExternalAssetRepository, execBinary *ExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:         asset,
		externalAsset: externalAsset,
		execBinary:    execBinary,
	}
}

// Find a GitHub release asset and an executable binary and returns them.
// repoFullName should be 'OWNER/REPO' format.
// patterns should be map whose key should be regular expressions of GitHub release asset download URL to download and value should be templates of executable binary name to install.
func (app *ApplicationService) Find(ctx context.Context, repoFullName string, tag string, patterns map[string]string) (FindResult, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return FindResult{}, err
	}

	release := newRelease(tag)

	ps, err := newPatternArrayFromStringMap(patterns)
	if err != nil {
		return FindResult{}, err
	}

	assets, err := app.listAssets(ctx, repo, release)
	if err != nil {
		return FindResult{}, err
	}

	asset, pattern, err := findAssetAndPattern(assets, ps)
	if err != nil {
		return FindResult{}, err
	}

	execBinary, err := pattern.execute(asset)
	if err != nil {
		return FindResult{}, err
	}

	return newFindResult(repo, release, asset, execBinary), nil
}

// Install downloads a GitHub release asset, extracts an executable binary from it, and writes it into given directory.
// Progress bar is written into w when downloading a GitHub release asset.
func (app *ApplicationService) Install(ctx context.Context, result FindResult, dir string, w io.Writer) error {
	assetContent, err := app.download(ctx, result.repo, result.Asset, w)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.extract(result.ExecBinary)
	if err != nil {
		return err
	}

	return app.execBinary.write(result.ExecBinary, execBinaryContent, dir)
}

// listAssets lists GitHub release assets in given GitHub release and returns them.
func (app *ApplicationService) listAssets(ctx context.Context, repo Repository, release Release) ([]Asset, error) {
	assets, err := app.asset.list(ctx, repo, release)
	if err != nil {
		return nil, err
	}

	externalAssets, err := app.externalAsset.list(repo, release)
	if err != nil {
		return nil, err
	}

	return append(assets, externalAssets...), nil
}

// download a GitHub release asset content and returns it. Progress bar is written into w.
func (app *ApplicationService) download(ctx context.Context, repo Repository, asset Asset, w io.Writer) (AssetContent, error) {
	if asset.isExternal() {
		return app.externalAsset.download(asset, w)
	}
	return app.asset.download(ctx, repo, asset, w)
}

// FindResult represents the result of [ApplicationService.Find].
type FindResult struct {
	repo       Repository
	release    Release
	Asset      Asset
	ExecBinary ExecBinary
}

// newFindResult returns a new [FindResult] object.
func newFindResult(repo Repository, release Release, asset Asset, execBinary ExecBinary) FindResult {
	return FindResult{
		repo:       repo,
		release:    release,
		Asset:      asset,
		ExecBinary: execBinary,
	}
}
