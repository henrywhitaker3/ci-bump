package main

import (
	"os"

	"github.com/henrywhitaker3/ci-bump/cmd/root"
)

var (
	version = "unknown"
)

func main() {
	if err := root.NewCommand(version).Execute(); err != nil {
		os.Exit(1)
	}
}
