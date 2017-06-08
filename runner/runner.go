package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/notifier"
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

		// Run is the stack of commands to execute.
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
func NewOrbitRunner(config *OrbitRunnerConfig, context *context.OrbitContext) *OrbitRunner {
	return &OrbitRunner{
		config:  config,
		context: context,
	}
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

// exec executes the stack of commands from the given Orbit command.
func (r *OrbitRunner) exec(cmd *OrbitCommand) error {
	notifier.Info("starting Orbit Command \"%s\"", cmd.Use)

	for _, c := range cmd.Run {
		notifier.Info("executing \"%s\"", c)
		parts := strings.Fields(c)

		// parts[0] contains the name of the current command.
		// parts[1:] contains the arguments of the current command.
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
