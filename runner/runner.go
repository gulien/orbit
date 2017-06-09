/*
Package runner implements a solution to runs one or more commands which have been defined in
a configuration file (by default "orbit.yml").

These commands, also called Orbit commands, lists one ore more external commands which will be executed
one by one.

Thanks to the generator package, the configuration file may be a data-driven template which is executed at runtime.
*/
package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/generator"
	"github.com/gulien/orbit/notifier"

	"gopkg.in/yaml.v2"
)

type (
	// OrbitRunnerConfig represents a YAML configuration file defining Orbit commands.
	OrbitRunnerConfig struct {
		// Commands slice represents the Orbit commands.
		Commands []*OrbitCommand `yaml:"commands"`
	}

	// OrbitCommand represents an Orbit command as defined in the configuration file.
	OrbitCommand struct {
		// Use is the name of the Orbit command.
		Use string `yaml:"use"`

		// Short describes what the Orbit command does.
		Short string `yaml:"short,omitempty"`

		// Run is the stack of external commands to execute.
		Run []string `yaml:"run"`
	}

	// OrbitRunner helps running Orbit commands.
	OrbitRunner struct {
		// config is an instance of OrbitRunnerConfig.
		config *OrbitRunnerConfig

		// context is an instance of OrbitContext.
		context *context.OrbitContext
	}
)

// NewOrbitRunner instantiates a new instance of OrbitRunner.
func NewOrbitRunner(context *context.OrbitContext) (*OrbitRunner, error) {
	// first retrieves the data from the configuration file.
	gen := generator.NewOrbitGenerator(context)
	data, err := gen.Parse()
	if err != nil {
		return nil, err
	}

	// then handles the data as YAML.
	var config = &OrbitRunnerConfig{}
	if err := yaml.Unmarshal(data.Bytes(), &config); err != nil {
		return nil, fmt.Errorf("configuration file \"%s\" is not a valid YAML file:\n%s", context.TemplateFilePath, err)
	}

	return &OrbitRunner{
		config:  config,
		context: context,
	}, nil
}

// Exec executes the given Orbit commands.
func (r *OrbitRunner) Exec(names ...string) error {
	// populates a slice of instances of Orbit Command.
	// if a given name doest not match with any Orbit Command defined in the configuration file, throws an error.
	cmds := make([]*OrbitCommand, len(names))
	for index, name := range names {
		cmds[index] = r.getOrbitCommand(name)
		if cmds[index] == nil {
			return fmt.Errorf("orbit command %s does not exist in configuration file \"%s\"", name, r.context.TemplateFilePath)
		}
	}

	// alright, let's execute each Orbit command.
	for _, cmd := range cmds {
		if err := r.exec(cmd); err != nil {
			return err
		}
	}

	return nil
}

// exec executes the stack of external commands from the given Orbit command.
func (r *OrbitRunner) exec(cmd *OrbitCommand) error {
	notifier.Info("starting Orbit Command \"%s\"", cmd.Use)

	for _, c := range cmd.Run {
		notifier.Info("executing \"%s\"", c)
		parts := strings.Fields(c)

		// parts[0] contains the name of the current external command.
		// parts[1:] contains the arguments of the current external command.
		e := exec.Command(parts[0], parts[1:]...)
		e.Stdout = os.Stdout
		e.Stderr = os.Stderr
		e.Stdin = os.Stdin

		if err := e.Run(); err != nil {
			return err
		}
	}

	return nil
}

// getOrbitCommand returns an instance of OrbitCommand if found or nil.
func (r *OrbitRunner) getOrbitCommand(name string) *OrbitCommand {
	for _, c := range r.config.Commands {
		if name == c.Use {
			return c
		}
	}

	return nil
}
