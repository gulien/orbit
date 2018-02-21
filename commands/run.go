package commands

import (
	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/runner"

	"github.com/spf13/cobra"
)

// default Orbit configuration file path.
const orbitFilePath = "orbit.yml"

var (
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

// init initializes a runCmd instance and adds it to the RootCmd.
func init() {
	RootCmd.AddCommand(runCmd)
}

// run executes one or more stacks of commands defined in a configuration file.
func run(cmd *cobra.Command, args []string) error {
	// alright, let's instantiate our Orbit context...
	if templateFilePath == "" {
		templateFilePath = orbitFilePath
	}

	ctx, err := context.NewOrbitContext(templateFilePath, payload)
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
