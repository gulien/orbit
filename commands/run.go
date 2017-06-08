package commands

import (
	"errors"
	"fmt"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/generator"
	"github.com/gulien/orbit/runner"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	// configFilePath is the path of the file describing commands.
	configFilePath string

	// runCmd is the instance of Orbit's runner command.
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

	// then retrieves the data from the configuration file.
	gen := generator.NewOrbitGenerator(ctx)
	data, err := gen.Parse()
	if err != nil {
		return err
	}

	// then handles the data as YAML.
	var config = &runner.OrbitRunnerConfig{}
	if err := yaml.Unmarshal(data.Bytes(), &config); err != nil {
		return fmt.Errorf("configuration file %s is not a valid YAML file:\n%s", configFilePath, err)
	}

	r := runner.NewOrbitRunner(config, ctx)
	if err := r.Exec(args[:]...); err != nil {
		return err
	}

	return nil
}
