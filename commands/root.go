// Package commands implements all commands of the application.
package commands

import (
	"github.com/gulien/orbit/logger"

	"github.com/spf13/cobra"
)

var (
	// templateFilePath is the path of a data-driven template.
	templateFilePath string

	// payload represents a map of YAML files, TOML files, JSON files, .env files and raw data.
	// Value format: key,path;key,path;key,data...
	payload string

	// debug enables logging if true.
	debug bool

	// RootCmd is the instance of the root of all commands.
	RootCmd = &cobra.Command{
		Use:           "orbit",
		Short:         "A cross-platform task runner for executing commands and generating files from templates",
		Long:          "A cross-platform task runner for executing commands and generating files from templates.",
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				logger.Mute()
			}
		},
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&templateFilePath, "file", "f", "", "specify the path of a data-driven template")
	RootCmd.PersistentFlags().StringVarP(&payload, "payload", "p", "", "specify a map of YAML files, TOML files, JSON files, .env files and raw data")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "display a detailed output")
}
