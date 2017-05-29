package commands

import (
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

func init() {
	RootCmd.AddCommand(RunCmd)
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run one or more commands",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			jww.ERROR.Println("No command to run")
			os.Exit(0)
		}
	},
}
