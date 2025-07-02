// Package main provides the main entry point for the gh-release-install command line tool.
package main

import (
	"log"
)

func main() {
	log.SetFlags(0)
	if err := NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
