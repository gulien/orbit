/*
Package commands implements all commands of the application:

- generate for generating a file from a data-driven template.
- run for executing one or more Orbit commands.
- version for printing the Orbit version
*/
package commands

import "github.com/spf13/cobra"

var (
	// ValuesFiles is the path or a map of paths of YAML files listing values used in data-driven templates.
	ValuesFiles string

	// EnvFiles is the path or a map of paths of .env files listing values used in data-driver templates.
	EnvFiles string

	// RootCmd is the instance of the root of all Orbit's commands.
	RootCmd = &cobra.Command{
		Use:           "orbit",
		Short:         "A simple tool for running commands and generating files from templates",
		SilenceErrors: true,
	}
)
