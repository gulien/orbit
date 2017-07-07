package commands

import (
	"fmt"

	orbitVersion "github.com/gulien/orbit/version"

	"github.com/spf13/cobra"
)

var (
	// version is the current version of Orbit.
	//version = "1.0.0-alpha1"

	// versionCmd is the instance of version command.
	versionCmd = &cobra.Command{
		Use:           "version",
		Short:         "Prints the version number of Orbit",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(orbitVersion.GetVersion())
		},
	}
)

// init adds versionCmd to the RootCmd.
func init() {
	RootCmd.AddCommand(versionCmd)
}
