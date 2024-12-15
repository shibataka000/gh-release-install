package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/prompter"
	"github.com/shibataka000/gh-release-install/github"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		repoFullName string
		tag          string
		patterns     map[string]string
		dir          string
		token        string

		defaultRepoFullName = os.Getenv("GH_REPO")
		defaultToken        = os.Getenv("GH_TOKEN")
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			app := newApplicationService(token)

			result, err := app.Find(ctx, repoFullName, tag, patterns)
			if err != nil {
				return err
			}

			prompt := fmt.Sprintf("Do you want to install %s from %s ?", result.ExecBinary.Name, result.Asset.DownloadURL.String())
			confirm, err := prompter.New(os.Stdin, os.Stdout, os.Stderr).Confirm(prompt, false)
			if err != nil {
				return err
			}
			if !confirm {
				return nil
			}

			return app.Install(ctx, result, dir, os.Stdout)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	command.Flags().StringVarP(&repoFullName, "repo", "R", defaultRepoFullName, "GitHub repository name. This should be [HOST/]OWNER/REPO format.")
	command.Flags().StringVar(&tag, "tag", "", "GitHub release tag.")
	command.Flags().StringToStringVar(&patterns, "pattern", github.DefaultPatterns, "Map whose key should be regular expressions of GitHub release asset download URL to download and value should be templates of executable binary name to install.")
	command.Flags().StringVarP(&dir, "dir", "D", ".", "Directory where executable binary will be installed into.")
	command.Flags().StringVar(&token, "token", defaultToken, "Authentication token for GitHub API requests.")

	markFlagRequired(command, "repo", defaultRepoFullName)
	markFlagRequired(command, "tag", "")
	markFlagRequired(command, "token", defaultToken)

	return command
}

// markFlagRequired marks flag as required if default value is not set.
func markFlagRequired(command *cobra.Command, name string, defaultValue string) {
	if defaultValue != "" {
		return
	}
	if err := command.MarkFlagRequired(name); err != nil {
		panic(err)
	}
}

// newApplicationService returns a new [github.com/shibataka000/gh-release-install/github.ApplicationService] object.
func newApplicationService(token string) *github.ApplicationService {
	return github.NewApplicationService(
		github.NewAssetRepository(token),
		github.NewExternalAssetRepository(github.DefaultExternalAssetTemplates),
		github.NewExecBinaryRepository(),
	)
}
