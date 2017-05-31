package commands

import (
	"github.com/spf13/cobra"
)

var (
	// EnvFile is the path of a .env file listing values used in templates.
	EnvFile string
	// ValuesFile is the path of a YAML file listing values used in templates.
	ValuesFile string
)

// RootCmd is the instance of the root of all orbit's commands.
var RootCmd = &cobra.Command{
	Use:   "orbit",
	Short: "A simple tool for running commands and generating templates",
}
