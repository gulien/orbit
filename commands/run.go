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
	Short: "Run one or more commands defined in a configuration file",
	Run:   run,
}

// TODO describe
func init() {
	runCmd.Flags().StringVarP(&configFile, "config", "c", "orbit.yml", "File or template (YAML) defining commands")
	runCmd.Flags().StringVarP(&ValuesFile, "values", "v", "", "File (YAML) defining values used in the template")
	runCmd.Flags().StringVarP(&EnvFile, "env_file", "e", "", "File storing env values used in the template")
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
