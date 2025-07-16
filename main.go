// Package main provides the main entry point for the gh-release-install command line tool.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/prompter"
	"github.com/spf13/cobra"
)

func runE(ctx context.Context, repo string, tag string, patterns map[string]string, dir string) error {
	assetRepository, err := NewAssetRepository(repo, os.Stdout)
	if err != nil {
		return err
	}
	app := NewApplicationService(assetRepository, NewExecBinaryRepository(dir))

	asset, execBinary, err := app.Find(ctx, tag, patterns)
	if err != nil {
		return err
	}

	prompt := fmt.Sprintf("Do you want to install %s from %s ?", execBinary.Name, asset.DownloadURL.String())
	confirm, err := prompter.New(os.Stdin, os.Stdout, os.Stderr).Confirm(prompt, true)
	if !confirm || err != nil {
		return err
	}

	return app.Install(ctx, asset, execBinary)
}

func main() {
	var (
		repo     string
		tag      string
		patterns map[string]string
		dir      string
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install an executable binary from a GitHub release asset.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runE(cmd.Context(), repo, tag, patterns, dir)
		},
		SilenceUsage: true,
	}

	currentRepositoryName := ""
	if r, err := currentRepository(); err == nil {
		currentRepositoryName = fmt.Sprintf("%s/%s/%s", r.Host, r.Owner, r.Name)
	}

	command.Flags().StringVarP(&repo, "repo", "R", currentRepositoryName, "GitHub repository name. This should be [HOST/]OWNER/REPO format.")
	command.Flags().StringVar(&tag, "tag", "", "GitHub release tag.")
	command.Flags().StringToStringVar(&patterns, "pattern", DefaultPatterns, "Map whose key should be regular expressions of GitHub release asset download URL to download and value should be templates of executable binary name to install.")
	command.Flags().StringVarP(&dir, "dir", "D", ".", "Directory where executable binary will be installed into.")

	if err := command.MarkFlagRequired("tag"); err != nil {
		panic(err)
	}

	if err := command.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
