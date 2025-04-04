// Package main provides the main entry point for the gh-release-install command line tool.
package main

import (
	"log"

	"github.com/shibataka000/gh-release-install/cmd"
)

func main() {
	log.SetFlags(0)
	if err := cmd.NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
