// Package commands implements all commands of the application.
package commands

import (
	"github.com/gulien/orbit/logger"

	"github.com/spf13/cobra"
)

var (
	// templateFilePath is the path of a data-driven template.
	templateFilePath string

	// payload represents a map of raw data, YAML files, TOML files, JSON files, and .env files.
	// Value format: key,path;key,path;key,data...
	payload string

	// debug enables logging if true.
	debug bool

	// RootCmd is the instance of the root of all commands.
	RootCmd = &cobra.Command{
		Use:           "orbit",
		Short:         "A task runner and a simple tool for generating files from data-driven templates",
		Long:          "A task runner and a simple tool for generating files from data-driven templates.",
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
	RootCmd.PersistentFlags().StringVarP(&payload, "payload", "p", "", "specify a map of raw data, YAML files, TOML files, JSON files, and .env files")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "display a detailed output")
}
