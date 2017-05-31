package commands

import (
	"github.com/spf13/cobra"
)

var (
	// EnvFilePath is the path of a .env file listing values used in templates.
	EnvFilePath string
	// ValuesFilePath is the path of a YAML file listing values used in templates.
	ValuesFilePath string
)

// RootCmd is the instance of the root of all Orbit's commands.
var RootCmd = &cobra.Command{
	Use:   "orbit",
	Short: "A simple tool for running commands and generating files from templates",
}
