package commands

import (
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	// configFile is the path of a file describing commands.
	configFile string
)

// runCmd is the instance of orbit's run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs one or more stack of commands defined in a configuration file",
	Run:   run,
}

// init function initializes runCmd instance with some flags and adds it to the RootCmd.
func init() {
	runCmd.Flags().StringVarP(&configFile, "config", "c", "orbit.yml", "specify an alternate configuration file")
	runCmd.Flags().StringVarP(&ValuesFile, "values", "v", "", "specify a YAML file listing values used in the configuration file")
	runCmd.Flags().StringVarP(&EnvFile, "env_file", "e", "", "specify a .env file listing values used in the configuration file")
	RootCmd.AddCommand(runCmd)
}

// run function executes one or more stacks of commands defined in the configuration file.
func run(cmd *cobra.Command, args []string) {
	// if no args, bye!
	if len(args) == 0 {
		jww.ERROR.Println("No command to run")
		os.Exit(1)
	}
}
