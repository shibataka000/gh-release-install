package github

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"

	"github.com/cheggaaa/pb/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/go-github/v67/github"
	"github.com/ulikunitz/xz"
)

// Asset represents a GitHub release asset.
type Asset struct {
	id          int64
	DownloadURL *url.URL
}

// newAsset returns a new [Asset] object.
func newAsset(id int64, downloadURL *url.URL) Asset {
	return Asset{
		id:          id,
		DownloadURL: downloadURL,
	}
}

// parseAsset returns a new [Asset] object.
func parseAsset(id int64, downloadURL string) (Asset, error) {
	parsed, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newAsset(id, parsed), nil
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extract [ExecBinaryContent] from [AssetContent] and returns it.
func (a AssetContent) extract(execBinary ExecBinary) (ExecBinaryContent, error) {
	b := []byte(a)

	for !isExecBinaryContent(b) {
		r, c, err := newReaderToExtract(b, execBinary)
		if err != nil {
			return nil, err
		}
		b, err = io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		if c != nil {
			if err := c.Close(); err != nil {
				return nil, err
			}
		}
	}

	return ExecBinaryContent(b), nil
}

// isExecBinaryContent returns true if MIME type of given bytes means executable binary content.
func isExecBinaryContent(b []byte) bool {
	expect := []string{"application/octet-stream", "application/x-executable", "application/x-sharedlib"}
	mime := mimetype.Detect(b)
	return slices.Contains(expect, mime.String())
}

// newReaderToExtract returns [io.Reader] to unarchive/decompress given bytes.
// Closing [io.ReadCloser] is caller's responsibility if it is not nil.
func newReaderToExtract(b []byte, execBinary ExecBinary) (io.Reader, io.Closer, error) {
	br := bytes.NewReader(b)
	mime := mimetype.Detect(b)

	switch mime.String() {
	case "application/gzip":
		r, err := gzip.NewReader(br)
		return r, nil, err
	case "application/x-xz":
		r, err := xz.NewReader(br)
		return r, nil, err
	case "application/x-tar":
		r, err := newFileReaderInTar(br, execBinary.Name)
		return r, nil, err
	case "application/zip":
		r, err := newFileReaderInZip(br, br.Size(), execBinary.Name)
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("%w: %s", ErrUnexpectedMIMEType, mime.String())
	}
}

// IAssetRepository is an interface about repository for [Asset] and [AssetContent].
type IAssetRepository interface {
	list(ctx context.Context, release Release) ([]Asset, error)
	download(ctx context.Context, asset Asset) (AssetContent, error)
}

// NewAssetRepository returns a new [AssetRepository] object or [ExternalAssetRepository] object based on given repository name.
func NewAssetRepository(repo string, stdout io.Writer) (IAssetRepository, error) {
	r, err := parseRepository(repo)
	if err != nil {
		return nil, err
	}
	if templates, ok := defaultExternalAssetTemplates[r]; ok {
		return newExternalAssetRepository(templates, stdout), nil
	}
	return newAssetRepository(r, stdout), nil
}

// AssetRepository is a repository for [Asset] and [AssetContent].
type AssetRepository struct {
	client *github.Client
	repo   Repository

	// stdout written progress bar into when downloading a GitHub release asset.
	stdout io.Writer
}

// newAssetRepository returns a new [AssetRepository] object.
func newAssetRepository(repo Repository, stdout io.Writer) *AssetRepository {
	return &AssetRepository{
		client: newGitHubClient(repo.host),
		repo:   repo,
		stdout: stdout,
	}
}

// list GitHub release assets in a given GitHub release and returns them.
func (r *AssetRepository) list(ctx context.Context, release Release) ([]Asset, error) {
	assets := []Asset{}

	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, r.repo.owner, r.repo.name, release.tag)
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
func (r *AssetRepository) download(ctx context.Context, asset Asset) (AssetContent, error) {
	githubAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, r.repo.owner, r.repo.name, asset.id)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, r.repo.owner, r.repo.name, asset.id, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close() // nolint:errcheck

	total := int64(githubAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(r.stdout).NewProxyReader(rc)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
