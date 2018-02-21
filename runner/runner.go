/*
Package runner implements a solution to executes one or more commands which have been defined
in a configuration file (by default "orbit.yml").

These commands, also called Orbit commands, runs one ore more external commands one by one.

Thanks to the generator package, the configuration file may be a data-driven template which is executed at runtime
(e.g. no file generated).
*/
package runner

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/generator"
	"github.com/gulien/orbit/logger"

	"gopkg.in/yaml.v2"
)

const defaultWindowsShellEnvVariable = "COMSPEC"
const defaultPosixShellEnvVariable = "SHELL"

type (
	// orbitRunnerConfig represents a YAML configuration file defining Orbit commands.
	orbitRunnerConfig struct {
		// Commands array represents the Orbit commands.
		Commands []*orbitCommand `yaml:"commands"`
	}

	// orbitCommand represents an Orbit command as defined in the configuration file.
	orbitCommand struct {
		// Use is the name of the Orbit command.
		Use string `yaml:"use"`

		// Short is the short description of the Orbit Command.
		Short string `yaml:"short,omitempty"`

		// Private allows to hide the Orbit command when
		// printing the available Orbit commands.
		Private bool `yaml:"private,omitempty"`

		// Run is the stack of external commands to run.
		Run []string `yaml:"run"`
	}

	// OrbitRunner helps executing Orbit commands.
	OrbitRunner struct {
		// config is an instance of orbitRunnerConfig.
		config *orbitRunnerConfig

		// context is an instance of OrbitContext.
		context *context.OrbitContext
	}
)

// NewOrbitRunner creates an instance of OrbitRunner.
func NewOrbitRunner(context *context.OrbitContext) (*OrbitRunner, error) {
	// first retrieves the data from the configuration file...
	g := generator.NewOrbitGenerator(context)
	data, err := g.Execute()
	if err != nil {
		return nil, err
	}

	// then populates the orbitRunnerConfig.
	var config = &orbitRunnerConfig{}
	if err := yaml.Unmarshal(data.Bytes(), &config); err != nil {
		return nil, errors.NewOrbitErrorf("configuration file %s is not a valid YAML file. Details:\n%s", context.TemplateFilePath, err)
	}

	r := &OrbitRunner{
		config:  config,
		context: context,
	}

	logger.Debugf("runner has been instantiated with config %s and context %s", r.config, r.context)

	return r, nil
}

// Print prints the available Orbit commands from the configuration file
// to Stdout.
func (r *OrbitRunner) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprint(w, "Configuration file:")
	fmt.Fprintf(w, "\n  %s\t\n", r.context.TemplateFilePath)
	fmt.Fprint(w, "\nAvailable Commands:")

	for _, c := range r.config.Commands {
		if !c.Private {
			fmt.Fprintf(w, "\n  %s\t%s", c.Use, c.Short)
		}
	}

	// clears the writer as it may contain some weird characters.
	fmt.Fprintln(w, "")

	w.Flush()
}

// Exec executes the given Orbit commands.
func (r *OrbitRunner) Exec(names ...string) error {
	// populates an array of instances of Orbit Command.
	// if a given name doest not match with any Orbit Command defined in the configuration file, throws an error.
	cmds := make([]*orbitCommand, len(names))
	for index, name := range names {
		cmds[index] = r.getOrbitCommand(name)
		if cmds[index] == nil {
			return errors.NewOrbitErrorf("Orbit command %s does not exist in configuration file %s", name, r.context.TemplateFilePath)
		}
	}

	// alright, let's run each Orbit command.
	for _, cmd := range cmds {
		if err := r.run(cmd); err != nil {
			return err
		}
	}

	return nil
}

// run runs the stack of external commands from the given Orbit command.
func (r *OrbitRunner) run(cmd *orbitCommand) error {
	logger.Debugf("starting Orbit command %s", cmd.Use)

	for _, c := range cmd.Run {

		var e *exec.Cmd
		if runtime.GOOS == "windows" {
			e = exec.Command(os.Getenv(defaultWindowsShellEnvVariable), "/c", c)
		} else {
			e = exec.Command(os.Getenv(defaultPosixShellEnvVariable), "-c", c)
		}

		logger.Debugf("running external command %s", e.Args)

		e.Stdout = os.Stdout
		e.Stderr = os.Stderr
		e.Stdin = os.Stdin

		if err := e.Run(); err != nil {
			return err
		}
	}

	return nil
}

// getOrbitCommand returns an instance of orbitCommand if found or nil.
func (r *OrbitRunner) getOrbitCommand(name string) *orbitCommand {
	for _, c := range r.config.Commands {
		if name == c.Use {
			return c
		}
	}

	return nil
}
