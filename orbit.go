/*
Package main is the root package of the application.

Orbit started with the need to find a cross-platform alternative of "make"
and "sed -i" commands. As it does not aim to be as powerful as these two
commands, Orbit offers an elegant solution for running commands and generating
files from templates, whatever the platform you're using.

For more information, go to https://github.com/gulien/orbit.
*/
package main

import (
	"os"

	"github.com/gulien/orbit/commands"
	"github.com/gulien/orbit/notifier"
)

// main is the root function of the application.
func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		notifier.Error(err)
		os.Exit(1)
	}
}
