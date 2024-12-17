package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/prompter"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/shibataka000/gh-release-install/github"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		repo     string
		tag      string
		patterns map[string]string
		dir      string
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()

			app, err := newApplicationService(repo)
			if err != nil {
				return err
			}

			result, err := app.Find(ctx, tag, patterns)
			if err != nil {
				return err
			}

			prompt := fmt.Sprintf("Do you want to install %s from %s ?", result.ExecBinary.Name, result.Asset.DownloadURL.String())
			confirm, err := prompter.New(os.Stdin, os.Stdout, os.Stderr).Confirm(prompt, true)
			if !confirm || err != nil {
				return err
			}

			return app.Install(ctx, result, dir)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	command.Flags().StringVarP(&repo, "repo", "R", currentRepository(), "GitHub repository name. This should be [HOST/]OWNER/REPO format.")
	command.Flags().StringVar(&tag, "tag", "", "GitHub release tag.")
	command.Flags().StringToStringVar(&patterns, "pattern", github.DefaultPatterns, "Map whose key should be regular expressions of GitHub release asset download URL to download and value should be templates of executable binary name to install.")
	command.Flags().StringVarP(&dir, "dir", "D", ".", "Directory where executable binary will be installed into.")

	if err := command.MarkFlagRequired("tag"); err != nil {
		panic(err)
	}

	return command
}

// newApplicationService returns a new [github.com/shibataka000/gh-release-install/github.ApplicationService] object.
func newApplicationService(repo string) (*github.ApplicationService, error) {
	asset, err := github.NewAssetRepository(repo, os.Stdout)
	if err != nil {
		return nil, err
	}
	execBinary := github.NewExecBinaryRepository()
	return github.NewApplicationService(asset, execBinary), nil
}

// currentRepository returns the GitHub repository the current directory is tracking, or empty string if not found.
func currentRepository() string {
	repo, err := repository.Current()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", repo.Host, repo.Owner, repo.Name)
}
