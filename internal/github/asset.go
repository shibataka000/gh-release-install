package github

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"slices"

	"github.com/cheggaaa/pb/v3"
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/go-github/v67/github"
	"github.com/ulikunitz/xz"
)

// Asset represents a GitHub release asset.
type Asset struct {
	ID          int64
	DownloadURL *url.URL
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extracts [ExecBinaryContent] from [AssetContent] and returns it.
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
	binaryMIMEs := []string{"application/octet-stream", "application/x-executable", "application/x-sharedlib"}
	mime := mimetype.Detect(b)
	return slices.Contains(binaryMIMEs, mime.String())
}

// newReaderToExtract returns an [io.Reader] to unarchive/decompress given bytes.
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
		r, err := newTarReader(br, execBinary.name)
		return r, nil, err
	case "application/zip":
		r, err := newZipReader(br, br.Size(), execBinary.name)
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("MIME type of asset content was unexpected: %s", mime.String())
	}
}

// newTarReader returns a [io.Reader] to read file which is given name in tarball.
func newTarReader(r io.Reader, name string) (io.Reader, error) {
	for tr := tar.NewReader(r); ; {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if filepath.Base(header.Name) == name && header.Typeflag == tar.TypeReg {
			return tr, nil
		}
	}
}

// newZipReader returns a [io.ReadCloser] to read file which is given name in zip file.
// Closing [io.ReadCloser] is caller's responsibility.
func newZipReader(r io.ReaderAt, size int64, name string) (io.ReadCloser, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if filepath.Base(f.Name) == name && !f.FileInfo().IsDir() {
			return f.Open()
		}
	}
	return nil, io.EOF
}

// AssetRepository is a repository for [Asset] and [AssetContent].
type AssetRepository struct {
	client      *github.Client
	repo        Repository
	progressBar io.Writer // written progress bar into when downloading a GitHub release asset.
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository(repo Repository, progressBar io.Writer) *AssetRepository {
	token, _ := auth.TokenForHost(repo.Host)
	return &AssetRepository{
		client:      github.NewClient(http.DefaultClient).WithAuthToken(token),
		repo:        repo,
		progressBar: progressBar,
	}
}

// List lists GitHub release assets in a given GitHub release and returns them.
func (r *AssetRepository) List(ctx context.Context, release Release) ([]Asset, error) {
	assets := []Asset{}

	repositoryRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, r.repo.Owner, r.repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	for page := 1; page != 0; {
		releaseAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, r.repo.Owner, r.repo.Name, repositoryRelease.GetID(), &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, releaseAsset := range releaseAssets {
			downloadURL, err := url.Parse(releaseAsset.GetBrowserDownloadURL())
			if err != nil {
				return nil, err
			}
			assets = append(assets, Asset{
				ID:          releaseAsset.GetID(),
				DownloadURL: downloadURL,
			})
		}
		page = resp.NextPage
	}

	return assets, nil
}

// Download downloads a GitHub release asset content and returns it.
func (r *AssetRepository) Download(ctx context.Context, asset Asset) (AssetContent, error) {
	releaseAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, r.repo.Owner, r.repo.Name, asset.ID)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, r.repo.Owner, r.repo.Name, asset.ID, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close() // nolint:errcheck

	total := int64(releaseAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(r.progressBar).NewProxyReader(rc)
	defer pr.Close() // nolint:errcheck

	return io.ReadAll(pr)
}
