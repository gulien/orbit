package commands

import (
	"fmt"

	OrbitVersion "github.com/gulien/orbit/version"

	"github.com/spf13/cobra"
)

var (
	// versionCmd is the instance of version command.
	versionCmd = &cobra.Command{
		Use:           "version",
		Short:         "Prints the version number of Orbit",
		Long:          "Prints the version number of Orbit.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(OrbitVersion.Current)
		},
	}
)

// init initializes a versionCmd instance and adds it to the RootCmd.
func init() {
	RootCmd.AddCommand(versionCmd)
}
