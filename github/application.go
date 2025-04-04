// Package github provides a service to find and install GitHub release assets.
package github

import (
	"context"
)

// ApplicationService provides a service to find and install GitHub release assets.
type ApplicationService struct {
	asset      IAssetRepository
	execBinary IExecBinaryRepository
}

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset IAssetRepository, execBinary IExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// Find a GitHub release asset in given release which matches given patterns and returns it and an executable binary in it.
func (app *ApplicationService) Find(ctx context.Context, tag string, patterns map[string]string) (FindResult, error) {
	release := newRelease(tag)

	ps, err := parsePatternMap(patterns)
	if err != nil {
		return FindResult{}, err
	}

	assets, err := app.asset.list(ctx, release)
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

	return newFindResult(asset, execBinary), nil
}

// Install downloads a GitHub release asset, extracts an executable binary from it, and writes it into given directory.
func (app *ApplicationService) Install(ctx context.Context, result FindResult, dir string) error {
	assetContent, err := app.asset.download(ctx, result.Asset)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.extract(result.ExecBinary)
	if err != nil {
		return err
	}

	return app.execBinary.write(result.ExecBinary, execBinaryContent, dir)
}

// FindResult represents the result of [ApplicationService.Find].
type FindResult struct {
	Asset      Asset
	ExecBinary ExecBinary
}

// newFindResult returns a new [FindResult] object.
func newFindResult(asset Asset, execBinary ExecBinary) FindResult {
	return FindResult{
		Asset:      asset,
		ExecBinary: execBinary,
	}
}
