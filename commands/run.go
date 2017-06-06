package commands

import (
	"os"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/generator"
	"github.com/gulien/orbit/helpers"
	"github.com/gulien/orbit/notifier"
	"github.com/gulien/orbit/runner"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	// configFilePath is the path of a file describing commands.
	configFilePath string
)

// runCmd is the instance of Orbit's runner command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs one or more stack of commands defined in a configuration file",
	Run:   run,
}

// init function initializes runCmd instance with some flags and adds it to the RootCmd.
func init() {
	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "orbit.yml", "specify an alternate configuration file")
	runCmd.Flags().StringVarP(&ValuesFilePath, "values", "v", "", "specify a YAML file listing values used in the configuration file")
	runCmd.Flags().StringVarP(&EnvFilePath, "env_file", "e", "", "specify a .env file listing values used in the configuration file")
	RootCmd.AddCommand(runCmd)
}

// runner function executes one or more stacks of commands defined in the configuration file.
func run(cmd *cobra.Command, args []string) {
	// if no args, bye!
	if len(args) == 0 {
		notifier.Error("No command to run")
	}

	// alright, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(configFilePath, ValuesFilePath, EnvFilePath)
	if err != nil {
		notifier.Error(err)
	}

	// checks if the config file is a YAML file.
	if !helpers.IsYAML(configFilePath) {
		notifier.Errorf("Configuration file %s is not a valid YAML file", configFilePath)
	}

	// then retrieves the data from the configuration file.
	gen := generator.NewOrbitGenerator(ctx)
	data, err := gen.Parse()
	if err != nil {
		notifier.Error(err)
	}

	// then handles the data as YAML.
	var config = &runner.OrbitRunnerConfig{}
	if err := yaml.Unmarshal(data.Bytes(), &config); err != nil {
		notifier.Errorf("Configuration file %s is not a valid YAML file:\n%s", configFilePath, err)
	}

	r := runner.NewOrbitRunner(config, ctx)
	if err := r.Exec(args[:]...); err != nil {
		notifier.Error(err)
	}

	// everything good!
	os.Exit(0)
}
