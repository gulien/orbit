package commands

import (
	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/runner"

	"github.com/spf13/cobra"
)

var (
	// configFilePath is the path of the file describing commands.
	configFilePath string

	// runCmd is the instance of run command.
	runCmd = &cobra.Command{
		Use:           "run",
		Short:         "Runs one or more stack of commands defined in a configuration file",
		Long:          "Runs one or more stack of commands defined in a configuration file.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          run,
	}
)

// init initializes a runCmd instance with some flags and adds it to the RootCmd.
func init() {
	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "orbit.yml", "specify an alternate configuration file")
	RootCmd.AddCommand(runCmd)
}

// run executes one or more stacks of commands defined in the configuration file.
func run(cmd *cobra.Command, args []string) error {
	// alright, let's instantiate our Orbit context...
	ctx, err := context.NewOrbitContext(configFilePath, ValuesFiles, EnvFiles, RawData)
	if err != nil {
		return err
	}

	// then our runner...
	r, err := runner.NewOrbitRunner(ctx)
	if err != nil {
		return err
	}

	// if no args, prints the available Orbit commands to Stdout.
	if len(args) == 0 {
		r.Print()
		return nil
	}

	// last but not least, executes Orbit commands.
	return r.Exec(args[:]...)
}
