package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/cheggaaa/pb"
	"github.com/shibataka000/go-get-release/internal/github"
	"github.com/ulikunitz/xz"
)

func install(name, token, goos, goarch, dir string, showPrompt bool) error {
	repo, release, asset, err := findAsset(name, token, goos, goarch)
	if err != nil {
		return err
	}

	if showPrompt {
		fmt.Printf("repo:\t%s/%s\ntag:\t%s\nasset:\t%s\n\n", repo.Owner(), repo.Name(), release.Tag(), asset.Name())
		if !prompter.YN("Are you sure to install release binary from above repository?", true) {
			return nil
		}
		fmt.Println()
	}

	tempDir, err := ioutil.TempDir("", "go-get-release-")
	if err != nil {
		return err
	}

	downloadFilePath := filepath.Join(tempDir, asset.Name())
	err = downloadFile(asset.DownloadURL(), downloadFilePath, showPrompt)
	if err != nil {
		return err
	}

	binaryName, err := asset.BinaryName()
	if err != nil {
		return err
	}

	var downloadBinaryPath string
	if asset.IsArchived() {
		err = extract(downloadFilePath, tempDir, binaryName)
		if err != nil {
			return err
		}
		downloadBinaryPath, err = findFile(tempDir, binaryName)
		if err != nil {
			return err
		}
	} else {
		downloadBinaryPath = downloadFilePath
	}

	installBinaryPath := filepath.Join(dir, binaryName)
	err = os.Rename(downloadBinaryPath, installBinaryPath)
	if err != nil {
		return err
	}
	err = os.Chmod(installBinaryPath, 0775)
	if err != nil {
		return err
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func findAsset(name, token, goos, goarch string) (github.Repository, github.Release, github.Asset, error) {
	owner, repoName, tag, err := parse(name)
	if err != nil {
		return nil, nil, nil, err
	}

	client, err := github.NewClient(token)
	if err != nil {
		return nil, nil, nil, err
	}

	var repo github.Repository
	if owner != "" {
		repo, err = client.Repository(owner, repoName)
	} else {
		repo, err = client.FindRepository(repoName)
	}
	if err != nil {
		return nil, nil, nil, err
	}

	var release github.Release
	if tag != "" {
		release, err = repo.Release(tag)
	} else {
		release, err = repo.LatestRelease()
	}
	if err != nil {
		return nil, nil, nil, err
	}

	asset, err := release.AssetByPlatform(goos, goarch)
	if err != nil {
		return nil, nil, nil, err
	}

	return repo, release, asset, nil
}

func parse(name string) (string, string, string, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^=]+))?`)
	if re.MatchString(name) {
		match := re.FindStringSubmatch(name)
		return match[2], match[3], match[5], nil
	}
	return "", "", "", fmt.Errorf("parsing package name failed: %s\npackage name should be \"owner/repo=tag\" format", name)
}

func downloadFile(url, filePath string, showProgress bool) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	r := resp.Body
	if showProgress {
		bar := pb.Full.Start64(resp.ContentLength)
		r = bar.NewProxyReader(resp.Body)
	}
	_, err = io.Copy(out, r)
	return err
}

func extract(srcFile, dstDir, dstFileName string) error {
	switch {
	case strings.HasSuffix(srcFile, ".zip"):
		return extractZip(srcFile, dstDir)
	case strings.HasSuffix(srcFile, ".tar.gz"):
		return extractTarGz(srcFile, dstDir)
	case strings.HasSuffix(srcFile, ".tgz"):
		return extractTarGz(srcFile, dstDir)
	case strings.HasSuffix(srcFile, ".tar.xz"):
		return extractTarXz(srcFile, dstDir)
	case strings.HasSuffix(srcFile, ".gz"):
		return extractGz(srcFile, filepath.Join(dstDir, dstFileName))
	default:
		return fmt.Errorf("unexpected archive type: %s", srcFile)
	}
}

func extractZip(srcFile, dstDir string) error {
	r, err := zip.OpenReader(srcFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		if f.FileInfo().IsDir() {
			path := filepath.Join(dstDir, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize)
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				return err
			}

			path := filepath.Join(dstDir, f.Name)
			err := ioutil.WriteFile(path, buf, f.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func extractTar(in io.Reader, dstDir string) error {
	tarReader := tar.NewReader(in)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		path := filepath.Join(dstDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return err
				}
			}
			outFile, err := os.Create(path)
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		default:
			return fmt.Errorf("fail to extract tarball: %s %v", header.Name, header.Typeflag)
		}
	}

	return nil
}

func extractTarGz(srcFile, dstDir string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	uncompressedStream, err := gzip.NewReader(in)
	if err != nil {
		return err
	}

	return extractTar(uncompressedStream, dstDir)
}

func extractTarXz(srcFile, dstDir string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	uncompressedStream, err := xz.NewReader(in)
	if err != nil {
		return err
	}

	return extractTar(uncompressedStream, dstDir)
}

func extractGz(srcFile, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	gr, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gr.Close()

	if _, err := io.Copy(out, gr); err != nil {
		return err
	}
	return nil
}

func findFile(dirPath, fileName string) (string, error) {
	filePath := ""
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			filePath = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if filePath == "" {
		return "", fmt.Errorf("%s is not found in %s", fileName, dirPath)
	}
	return filePath, nil
}
