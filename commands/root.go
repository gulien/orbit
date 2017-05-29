package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "orbit",
	Short: "A simple tool for generating templates and running commands",
}
