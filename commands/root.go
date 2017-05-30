package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "orbit",
	Short: "A simple tool for running commands and generating templates",
}
