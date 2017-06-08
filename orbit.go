package main

import (
	"os"

	"github.com/gulien/orbit/commands"
	"github.com/gulien/orbit/notifier"
)

// main is the root function of the Orbit application.
func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		notifier.Error(err)
		os.Exit(1)
	}
}
