package github

import (
	"context"
	"io"
)

// ApplicationService.
type ApplicationService struct {
	repository    *RepositoryRepository
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
// See [newRepositoryFromFullName], [newRelease], [newPatternArrayFromStringMap] for details about each arguments.
func (app *ApplicationService) Find(ctx context.Context, repoFullName string, tag string, patterns map[string]string) (FindResult, error) {
	repo, err := app.findRepository(ctx, repoFullName)
	if err != nil {
		return FindResult{}, err
	}

	release := newRelease(tag)

	asset, pattern, err := app.findAssetAndPattern(ctx, repo, release, patterns)
	if err != nil {
		return FindResult{}, err
	}

	execBinary, err := app.findExecBinary(asset, pattern)
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

	execBinaryContent, err := app.extract(assetContent, result.ExecBinary)
	if err != nil {
		return err
	}

	return app.install(result.ExecBinary, execBinaryContent, dir)
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

func (app *ApplicationService) findRepository(ctx context.Context, query string) (Repository, error) {
	parsed, err := parseRepositoryName(query)
	if err != nil {
		repo, err := app.repository.get(ctx, parsed.owner, parsed.name)
		if err == nil {
			return repo, nil
		}
	}
	return app.repository.search(ctx, query)
}

// findAssetAndPattern finds a GitHub release asset in given GitHub release which matches any of given patterns and returns them.
func (app *ApplicationService) findAssetAndPattern(ctx context.Context, repo Repository, release Release, patterns map[string]string) (Asset, Pattern, error) {
	assets, err := app.listAssets(ctx, repo, release)
	if err != nil {
		return Asset{}, Pattern{}, err
	}

	ps, err := newPatternArrayFromStringMap(patterns)
	if err != nil {
		return Asset{}, Pattern{}, err
	}

	return find(assets, ps)
}

// findExecBinary makes [ExecBinary] object from [Asset] and [Pattern] and returns it.
func (app *ApplicationService) findExecBinary(asset Asset, pattern Pattern) (ExecBinary, error) {
	return pattern.execute(asset)
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

// extract [ExecBinaryContent] from [AssetContent] and returns it.
func (app *ApplicationService) extract(asset AssetContent, execBinary ExecBinary) (ExecBinaryContent, error) {
	return asset.extract(execBinary)
}

// install [ExecBinaryContent] into given directory.
func (app *ApplicationService) install(meta ExecBinary, content ExecBinaryContent, dir string) error {
	return app.execBinary.write(meta, content, dir)
}
