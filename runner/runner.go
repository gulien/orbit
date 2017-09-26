/*
Package runner implements a solution to executes one or more commands which have been defined
in a configuration file (by default "orbit.yml").

These commands, also called Orbit commands, runs one ore more external commands one by one.

Thanks to the generator package, the configuration file may be a data-driven template which is executed at runtime
(e.g. no file generated).
*/
package runner

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/generator"
	"github.com/gulien/orbit/logger"

	"gopkg.in/yaml.v2"
)

type (
	// OrbitRunnerConfig represents a YAML configuration file defining Orbit commands.
	OrbitRunnerConfig struct {
		// Commands array represents the Orbit commands.
		Commands []*OrbitCommand `yaml:"commands"`
	}

	// OrbitCommand represents an Orbit command as defined in the configuration file.
	OrbitCommand struct {
		// Use is the name of the Orbit command.
		Use string `yaml:"use"`

		// Run is the stack of external commands to run.
		Run []string `yaml:"run"`
	}

	// OrbitRunner helps executing Orbit commands.
	OrbitRunner struct {
		// config is an instance of OrbitRunnerConfig.
		config *OrbitRunnerConfig

		// context is an instance of OrbitContext.
		context *context.OrbitContext
	}

	// orbitExternalCommand represents an external command from an Orbit command.
	orbitExternalCommand struct {
		// argc is the binary to call.
		argc string

		// argv are the arguments of the external command.
		argv []string
	}
)

// NewOrbitRunner creates an instance of OrbitRunner.
func NewOrbitRunner(context *context.OrbitContext) (*OrbitRunner, error) {
	// first retrieves the data from the configuration file...
	g := generator.NewOrbitGenerator(context)
	data, err := g.Parse()
	if err != nil {
		return nil, err
	}

	// then populates the OrbitRunnerConfig.
	var config = &OrbitRunnerConfig{}
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

// newOrbitExternalCommand creates an instance of orbitExternalCommand.
func newOrbitExternalCommand(c string) *orbitExternalCommand {
	pattern := regexp.MustCompile("'.+'|\".+\"|`.+`|\\S+")
	parts := pattern.FindAllString(c, -1)

	extCmd := &orbitExternalCommand{
		argc: parts[0],
	}

	for _, argv := range parts[1:] {
		argv = strings.TrimPrefix(argv, "'")
		argv = strings.TrimSuffix(argv, "'")
		argv = strings.TrimPrefix(argv, "\"")
		argv = strings.TrimSuffix(argv, "\"")
		argv = strings.TrimPrefix(argv, "`")
		argv = strings.TrimSuffix(argv, "`")

		extCmd.argv = append(extCmd.argv, argv)
	}

	return extCmd
}

// Exec executes the given Orbit commands.
func (r *OrbitRunner) Exec(names ...string) error {
	// populates an array of instances of Orbit Command.
	// if a given name doest not match with any Orbit Command defined in the configuration file, throws an error.
	cmds := make([]*OrbitCommand, len(names))
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
func (r *OrbitRunner) run(cmd *OrbitCommand) error {
	logger.Debugf("starting Orbit command %s", cmd.Use)

	for _, c := range cmd.Run {
		extCmd := newOrbitExternalCommand(c)
		e := exec.Command(extCmd.argc, extCmd.argv...)

		logger.Debugf("running external command %s with args %s", extCmd.argc, extCmd.argv)

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
