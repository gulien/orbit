// Package commands implements all commands of the application.
package commands

import (
	"github.com/gulien/orbit/logger"

	"github.com/spf13/cobra"
)

var (
	// ValuesFiles is the path or a map of paths of YAML files listing values used in a data-driven template.
	ValuesFiles string

	// EnvFiles is the path or a map of paths of .env files listing values used in a data-driven template.
	EnvFiles string

	// RawData are a map of values used in a data-driven template.
	RawData string

	// debug enables logging if true.
	debug bool

	// RootCmd is the instance of the root of all commands.
	RootCmd = &cobra.Command{
		Use:           "orbit",
		Short:         "A simple tool for running commands and generating files from templates",
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				logger.Mute()
			}
		},
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&ValuesFiles, "values", "v", "", "specify a YAML file or a map of YAML files listing values used in the template")
	RootCmd.PersistentFlags().StringVarP(&EnvFiles, "env", "e", "", "specify a .env file or a map of .env files listing values used in the template")
	RootCmd.PersistentFlags().StringVarP(&RawData, "raw", "r", "", "specify a map of values used in the template")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "display a detailed output")
}
