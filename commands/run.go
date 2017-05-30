package commands

import (
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	// TODO describe
	configFile string
)

// TODO describe
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs one or more stack of commands defined in a configuration file",
	Run:   run,
}

// TODO describe
func init() {
	runCmd.Flags().StringVarP(&configFile, "config", "c", "orbit.yml", "specify an alternate configuration file")
	runCmd.Flags().StringVarP(&ValuesFile, "values", "v", "", "specify a YAML file listing values used in the configuration file")
	runCmd.Flags().StringVarP(&EnvFile, "env_file", "e", "", "specify a .env file listing values used in the configuration file")
	RootCmd.AddCommand(runCmd)
}

// TODO describe
func run(cmd *cobra.Command, args []string) {
	// if no args, bye!
	if len(args) == 0 {
		jww.ERROR.Println("No command to run")
		os.Exit(1)
	}
}
