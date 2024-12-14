package main

import (
	"log"

	"github.com/shibataka000/gh-release-install/cmd"
)

func main() {
	log.SetFlags(0)
	if err := cmd.NewCommand().Execute(); err != nil {
		log.Fatal("Error: ", err)
	}
}
