package main

import (
	"github.com/tmiddlet2666/ghstats/pkg/cmd"
)

var (
	// Version is the cohctl version injected by the Go linker at build time
	Version string
	// Commit is the git commit hash injected by the Go linker at build time
	Commit string
	// Date is the build timestamp injected by the Go linker at build time
	Date string
)

// main is the main entry point to Coherence CLI
func main() {
	cmd.Execute(Version, Date, Commit)
}
