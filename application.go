// package main provides a service to find and install GitHub release assets.
package main

import (
	"context"
)

// ApplicationService provides a service to find and install GitHub release assets.
type ApplicationService struct {
	asset      AssetRepository
	execBinary *ExecBinaryRepository
}

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset AssetRepository, execBinary *ExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// AssetRepository is an interface about repository for [Asset] and [AssetContent].
type AssetRepository interface {
	list(ctx context.Context, release Release) ([]Asset, error)
	download(ctx context.Context, asset Asset) (AssetContent, error)
}

// Find a GitHub release asset in given release which matches given patterns and returns it and an executable binary in it.
func (app *ApplicationService) Find(ctx context.Context, tag string, patterns map[string]string) (Asset, ExecBinary, error) {
	release := Release{
		tag: tag,
	}

	ps, err := parsePatternMap(patterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	assets, err := app.asset.list(ctx, release)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	asset, pattern, err := findAssetAndPattern(assets, ps)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	execBinary, err := pattern.execute(asset)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	return asset, execBinary, nil
}

// Install downloads a GitHub release asset, extracts an executable binary from it, and writes it into given directory.
func (app *ApplicationService) Install(ctx context.Context, asset Asset, execBinary ExecBinary, dir string) error {
	assetContent, err := app.asset.download(ctx, asset)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.extract(execBinary)
	if err != nil {
		return err
	}

	return app.execBinary.write(execBinary, execBinaryContent, dir)
}
