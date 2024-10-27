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
	"github.com/google/go-github/v62/github"
	"github.com/ulikunitz/xz"
)

// Asset represents a GitHub release asset.
type Asset struct {
	// ID. If asset is hosted on server outside from GitHub, this should be 0.
	ID int64

	// DownloadURL is an URL to download an asset content.
	DownloadURL *url.URL
}

// NewAsset returns a new [Asset] object.
func NewAsset(id int64, downloadURL *url.URL) Asset {
	return Asset{
		ID:          id,
		DownloadURL: downloadURL,
	}
}

// NewAssetFromString returns a new [Asset] object.
func NewAssetFromString(id int64, downloadURL string) (Asset, error) {
	parsed, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return NewAsset(id, parsed), nil
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extract [ExecBinaryContent] from [AssetContent] and return it.
func (a AssetContent) extract(meta ExecBinary) (ExecBinaryContent, error) {
	b := []byte(a)

	for !isExecBinaryContent(b) {
		r, c, err := newReaderToExtract(b, meta)
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
func newReaderToExtract(b []byte, meta ExecBinary) (io.Reader, io.Closer, error) {
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
		r, err := newFileReaderInTar(br, meta.Name)
		return r, nil, err
	case "application/zip":
		r, err := newFileReaderInZip(br, br.Size(), meta.Name)
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("%w: %s", ErrUnexpectedMIME, mime.String())
	}
}

// IAssetRepository is an interface about repository for [Asset] and [AssetContent].
type IAssetRepository interface {
	List(ctx context.Context, repo Repository, release Release) ([]Asset, error)
	Download(ctx context.Context, repo Repository, asset Asset, w io.Writer) (AssetContent, error)
}

// AssetRepository is a repository for [Asset] and [AssetContent].
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository(token string) *AssetRepository {
	return &AssetRepository{
		client: github.NewClient(http.DefaultClient).WithAuthToken(token),
	}
}

// List GitHub release assets and returns them.
func (r *AssetRepository) List(ctx context.Context, repo Repository, release Release) ([]Asset, error) {
	assets := []Asset{}

	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	for page := 1; page != 0; {
		githubAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, githubRelease.GetID(), &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, githubAsset := range githubAssets {
			asset, err := NewAssetFromString(githubAsset.GetID(), githubAsset.GetBrowserDownloadURL())
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
		page = resp.NextPage
	}

	return assets, nil
}

// Download a GitHub release asset content and returns it.
// Progress bar is written into w.
func (r *AssetRepository) Download(ctx context.Context, repo Repository, asset Asset, w io.Writer) (AssetContent, error) {
	githubAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, repo.Owner, repo.Name, asset.ID)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, repo.Owner, repo.Name, asset.ID, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	total := int64(githubAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(w).NewProxyReader(rc)
	defer pr.Close()

	return io.ReadAll(pr)
}
