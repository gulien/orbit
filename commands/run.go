package commands

import (
	"errors"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/runner"

	"github.com/spf13/cobra"
)

var (
	// configFilePath is the path of the file describing commands.
	configFilePath string

	// runCmd is the instance of Orbit's run command.
	runCmd = &cobra.Command{
		Use:           "run",
		Short:         "Runs one or more stack of commands defined in a configuration file",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          run,
	}
)

// init initializes a runCmd instance with some flags and adds it to the RootCmd.
func init() {
	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "orbit.yml", "specify an alternate configuration file")
	runCmd.Flags().StringVarP(&ValuesFiles, "values", "v", "", "specify a YAML file or a map of YAML files listing values used in the configuration file")
	runCmd.Flags().StringVarP(&EnvFiles, "env", "e", "", "specify a .env file or a map of .env files listing values used in the configuration file")
	RootCmd.AddCommand(runCmd)
}

// runner executes one or more stacks of commands defined in the configuration file.
func run(cmd *cobra.Command, args []string) error {
	// if no args, bye!
	if len(args) == 0 {
		return errors.New("no command to run")
	}

	// alright, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(configFilePath, ValuesFiles, EnvFiles)
	if err != nil {
		return err
	}

	r, err := runner.NewOrbitRunner(ctx)
	if err != nil {
		return err
	}

	if err := r.Exec(args[:]...); err != nil {
		return err
	}

	return nil
}
