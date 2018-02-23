/*
Package main is the root package of the application.

Orbit is a cross-platform task runner for executing commands and generating files from templates.

It started with the need to find a cross-platform alternative of "make"
and "sed -i" commands. As it does not aim to be as powerful as these two
commands, Orbit offers an elegant solution for running tasks and generating
files from templates, whatever the platform you're using.

For more information, go to https://github.com/gulien/orbit.
*/
package main

import (
	"os"

	"github.com/gulien/orbit/commands"
	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/logger"
	OrbitVersion "github.com/gulien/orbit/version"
)

/*
version will be set by GoReleaser.

It will be the current Git tag (with v prefix stripped) or
the name of the snapshot if you're using the --snapshot flag.
*/
var version = "master"

// main is the root function of the application.
func main() {
	OrbitVersion.Current = version

	if err := commands.RootCmd.Execute(); err != nil {
		if orbitError, ok := err.(*errors.OrbitError); ok {
			logger.NotifyOrbitError(orbitError)
		} else {
			logger.Error(err)
		}

		os.Exit(1)
	}
}
