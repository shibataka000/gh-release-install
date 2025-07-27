package main

import (
	"context"
	"io"
)

// newAssetRepository returns a new [GitHubAssetRepository] object or [ExternalAssetRepository] object based on given repository name.
func newAssetRepository(repo string, progressBar io.Writer) (AssetRepository, error) {
	r, err := parseRepository(repo)
	if err != nil {
		return nil, err
	}
	if templates, ok := externalAssetTemplates[r]; ok {
		return newExternalAssetRepository(templates, progressBar), nil
	}
	return newGitHubAssetRepository(r, progressBar), nil
}

// AssetRepository is an interface about repository for [Asset] and [AssetContent].
type AssetRepository interface {
	list(ctx context.Context, release Release) ([]Asset, error)
	download(ctx context.Context, asset Asset) (AssetContent, error)
}

// ExecBinaryRepository is an interface about repository for [ExecBinary] and [ExecBinaryContent].
type ExecBinaryRepository interface {
	write(meta ExecBinary, content ExecBinaryContent) error
}

// newExecBinaryRepository returns a new [ExecBinaryRepository] object.
func newExecBinaryRepository(dir string) ExecBinaryRepository {
	return newFSExecBinaryRepository(dir)
}
