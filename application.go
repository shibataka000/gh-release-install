// Package main provides a service to find and install GitHub release assets.
package main

import (
	"context"

	"github.com/shibataka000/gh-release-install/github2"
)

// ApplicationService provides a service to find and install GitHub release assets.
type ApplicationService struct {
	asset      github2.IAssetRepository
	execBinary github2.IExecBinaryRepository
}

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset github2.IAssetRepository, execBinary github2.IExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// Find a GitHub release asset in given release which matches given patterns and returns it and an executable binary in it.
func (app *ApplicationService) Find(ctx context.Context, tag string, patterns map[string]string) (github2.FindResult, error) {
	release := github2.NewRelease(tag)

	ps, err := github2.ParsePatternMap(patterns)
	if err != nil {
		return github2.FindResult{}, err
	}

	assets, err := app.asset.List(ctx, release)
	if err != nil {
		return github2.FindResult{}, err
	}

	asset, pattern, err := github2.FindAssetAndPattern(assets, ps)
	if err != nil {
		return github2.FindResult{}, err
	}

	execBinary, err := pattern.Execute(asset)
	if err != nil {
		return github2.FindResult{}, err
	}

	return github2.NewFindResult(asset, execBinary), nil
}

// Install downloads a GitHub release asset, extracts an executable binary from it, and writes it into given directory.
func (app *ApplicationService) Install(ctx context.Context, result github2.FindResult, dir string) error {
	assetContent, err := app.asset.Download(ctx, result.Asset)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.Extract(result.ExecBinary)
	if err != nil {
		return err
	}

	return app.execBinary.Write(result.ExecBinary, execBinaryContent, dir)
}
