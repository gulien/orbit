// Package commands implements all commands of the application.
package commands

import "github.com/spf13/cobra"

var (
	// ValuesFiles is the path or a map of paths of YAML files listing values used in a data-driven template.
	ValuesFiles string

	// EnvFiles is the path or a map of paths of .env files listing values used in a data-driven template.
	EnvFiles string

	// RootCmd is the instance of the root of all commands.
	RootCmd = &cobra.Command{
		Use:           "orbit",
		Short:         "A simple tool for running commands and generating files from templates",
		SilenceErrors: true,
	}
)
