package commands

import (
	"fmt"
	"os"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/generator"
	"github.com/gulien/orbit/helpers"
	"github.com/gulien/orbit/runner"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
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
		jww.ERROR.Println("No command to run")
		os.Exit(1)
	}

	// alright, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(configFilePath, ValuesFilePath, EnvFilePath)
	if err != nil {
		jww.ERROR.Println(err)
		os.Exit(1)
	}

	// checks if the config file is a YAML file.
	if !helpers.IsYAML(configFilePath) {
		err := fmt.Errorf("Configuration file %s is not a valid YAML file", configFilePath)
		jww.ERROR.Println(err)
		os.Exit(1)
	}

	// then retrieves the data from the configuration file.
	gen := generator.NewOrbitGenerator(ctx)
	data, err := gen.Parse()
	if err != nil {
		jww.ERROR.Println(err)
		os.Exit(1)
	}

	// then handles the data as YAML.
	var config = &runner.OrbitRunnerConfig{}
	if err := yaml.Unmarshal(data.Bytes(), &config); err != nil {
		jww.ERROR.Printf("Configuration file %s is not a valid YAML file: %s", configFilePath, err)
		os.Exit(1)
	}

	r := runner.NewOrbitRunner(config, ctx)
	if err := r.Exec(args[:]...); err != nil {
		jww.ERROR.Println(err)
		os.Exit(1)
	}

	// everything good!
	os.Exit(0)
}
