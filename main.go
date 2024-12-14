package main

import (
	"log"
	"os"

	"github.com/shibataka000/gh-release-install/cmd"
)

func main() {
	log.SetFlags(0)
	if err := cmd.NewCommand().Execute(); err != nil {
		log.SetOutput(os.Stderr)
		log.Fatal("Error: ", err)
	}
}
