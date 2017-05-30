package commands

import (
	"github.com/spf13/cobra"
)

var (
	// TODO describe
	EnvFile    string
	// TODO describe
	ValuesFile string
)

// TODO describe
var RootCmd = &cobra.Command{
	Use:   "orbit",
	Short: "A simple tool for running commands and generating templates",
}
