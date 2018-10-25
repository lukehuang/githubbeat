package main

import (
	"os"

	"github.com/ressedpanda/githubbeat/cmd"

	_ "github.com/ressedpanda/githubbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
