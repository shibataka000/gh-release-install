package github

import (
	"context"
	"io"
)

// ApplicationService.
type ApplicationService struct {
	asset      *AssetRepository
	execBinary *ExecBinaryRepository
}

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset *AssetRepository, execBinary *ExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// Find a GitHub release asset and an executable binary and returns them.
// See [newRepositoryFromFullName], [newRelease], [newPatternArrayFromStringMap] for details about each arguments.
func (app *ApplicationService) Find(ctx context.Context, repoFullName string, tag string, patterns map[string]string) (Asset, ExecBinary, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	release := newRelease(tag)

	ps, err := newPatternArrayFromStringMap(patterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	assets, err := app.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	asset, pattern, err := find(assets, ps)
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
// Progress bar is written into w when downloading a GitHub release asset.
func (app *ApplicationService) Install(ctx context.Context, repoFullName string, asset Asset, execBinary ExecBinary, dir string, w io.Writer) error {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return err
	}

	assetContent, err := app.asset.download(ctx, repo, asset, w)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.extract(execBinary)
	if err != nil {
		return err
	}

	return app.execBinary.write(execBinary, execBinaryContent, dir)
}
