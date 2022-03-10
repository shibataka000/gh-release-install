package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var version string

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:                 "go-get-release",
		Usage:                "install golang release binary",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "github-token",
				Value:   "",
				EnvVars: []string{"GITHUB_TOKEN"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Install golang release binary",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "goos",
						Value:   "linux",
						EnvVars: []string{"GOOS"},
					},
					&cli.StringFlag{
						Name:    "goarch",
						Value:   "amd64",
						EnvVars: []string{"GOARCH"},
					},
					&cli.StringFlag{
						Name:  "install-dir",
						Value: filepath.Join(os.Getenv("GOPATH"), "bin"),
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 0 {
						return fmt.Errorf("no repository is specified")
					}
					return install(c.Args().Get(0), c.String("github-token"), c.String("goos"), c.String("goarch"), c.String("install-dir"), true)
				},
			},
			{
				Name:  "search",
				Usage: "Search GitHub repository",
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 0 {
						return fmt.Errorf("no repository is specified")
					}
					return search(c.Args().Get(0), c.String("github-token"))
				},
			},
			{
				Name:  "tags",
				Usage: "Show tags of GitHub repository",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "max",
						Value: 20,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 0 {
						return fmt.Errorf("no repository is specified")
					}
					return tags(c.Args().Get(0), c.String("github-token"), c.Int("max"))
				},
			},
			{
				Name:  "completion",
				Usage: "Outputs shell completion for the given shell (bash only)",
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 0 {
						return fmt.Errorf("no shell is specified")
					}
					return completion(c.Args().Get(0))
				},
			},
			{
				Name:  "version",
				Usage: "Show client version",
				Action: func(c *cli.Context) error {
					fmt.Println(version)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
