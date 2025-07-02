package main

import (
	"context"
)

// ApplicationService provides a service to find and install GitHub release assets.
type ApplicationService struct {
	asset      AssetRepository
	execBinary *ExecBinaryRepository
}

// FindResult represents the result of [ApplicationService.Find].
type FindResult struct {
	Asset      Asset
	ExecBinary ExecBinary
}

// AssetRepository is an interface about repository for [Asset] and [AssetContent].
type AssetRepository interface {
	list(ctx context.Context, release Release) ([]Asset, error)
	download(ctx context.Context, asset Asset) (AssetContent, error)
}

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset AssetRepository, execBinary *ExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// Find a GitHub release asset in given release which matches given patterns and returns it and an executable binary in it.
func (app *ApplicationService) Find(ctx context.Context, tag string, patterns map[string]string) (FindResult, error) {
	release := Release{
		Tag: tag,
	}

	ps, err := parsePatterns(patterns)
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

	return FindResult{
		Asset:      asset,
		ExecBinary: execBinary,
	}, nil
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
