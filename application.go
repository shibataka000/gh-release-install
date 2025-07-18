package main

import (
	"context"
)

// ApplicationService provides a service to find and install GitHub release assets.
type ApplicationService struct {
	asset      AssetRepository
	execBinary ExecBinaryRepository
}

// newApplicationService returns a new [ApplicationService] object.
func newApplicationService(asset AssetRepository, execBinary ExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// find finds a GitHub release asset in given release which matches given patterns and returns it and an executable binary in it.
func (app *ApplicationService) find(ctx context.Context, tag string, patterns map[string]string) (Asset, ExecBinary, error) {
	ps, err := parsePatterns(patterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	assets, err := app.asset.list(ctx, Release{
		tag: tag,
	})
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

// install downloads a GitHub release asset, extracts an executable binary from it, and writes it.
func (app *ApplicationService) install(ctx context.Context, asset Asset, execBinary ExecBinary) error {
	assetContent, err := app.asset.download(ctx, asset)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.extract(execBinary)
	if err != nil {
		return err
	}

	return app.execBinary.write(execBinary, execBinaryContent)
}
