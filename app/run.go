package app

import (
	"github.com/gulien/orbit/app/context"
	"github.com/gulien/orbit/app/runner"

	"github.com/spf13/cobra"
)

// default Orbit configuration file path.
const orbitFilePath = "orbit.yml"

var (
	// runCmd is the instance of run command.
	runCmd = &cobra.Command{
		Use:           "run",
		Short:         "Runs one or more tasks defined in a configuration file",
		Long:          "Runs one or more tasks defined in a configuration file.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          run,
	}
)

// init initializes a runCmd instance and adds it to the RootCmd.
func init() {
	RootCmd.AddCommand(runCmd)
}

// run runs one or more tasks defined in a configuration file.
func run(cmd *cobra.Command, args []string) error {
	// alright, let's instantiate our Orbit context...
	if templateFilePath == "" {
		templateFilePath = orbitFilePath
	}

	ctx, err := context.NewOrbitContext(templateFilePath, payload, templates, nil)
	if err != nil {
		return err
	}

	// then our runner...
	r, err := runner.NewOrbitRunner(ctx)
	if err != nil {
		return err
	}

	// if no args, prints the available tasks to Stdout...
	if len(args) == 0 {
		r.Print()
		return nil
	}

	// ... or runs given tasks.
	return r.Run(args[:]...)
}
