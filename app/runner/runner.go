/*
Package runner implements a solution to run one or more tasks which have been defined
in a configuration file (by default "orbit.yml").

These tasks executes one ore more commands one by one.

Thanks to the generator package, the configuration file may be a data-driven template which is executed at runtime
(e.g. no file generated).
*/
package runner

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/tabwriter"

	"github.com/gulien/orbit/app/context"
	OrbitError "github.com/gulien/orbit/app/error"
	"github.com/gulien/orbit/app/generator"
	"github.com/gulien/orbit/app/logger"

	"gopkg.in/yaml.v2"
)

const defaultWindowsShellEnvVariable = "COMSPEC"
const defaultPosixShellEnvVariable = "SHELL"

type (
	// orbitRunnerConfig represents a YAML configuration file defining tasks.
	orbitRunnerConfig struct {
		// Tasks array represents the tasks defined in the configuration file.
		Tasks []*orbitTask `yaml:"tasks"`
	}

	// orbitTask represents a task as defined in the configuration file.
	orbitTask struct {
		// Use is the name of the task.
		Use string `yaml:"use"`

		// Short is the short description of the task.
		Short string `yaml:"short,omitempty"`

		// Private allows to hide the task when
		// printing the available tasks.
		Private bool `yaml:"private,omitempty"`

		// Shell allows to choose which binary will
		// be called to run the commands.
		Shell string `yaml:"shell,omitempty"`

		// Run is the stack of commands to execute.
		Run []string `yaml:"run"`
	}

	// OrbitRunner helps executing tasks.
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
		return nil, OrbitError.NewOrbitErrorf("configuration file %s is not a valid YAML file. Details:\n%s", context.TemplateFilePath, err)
	}

	r := &OrbitRunner{
		config:  config,
		context: context,
	}

	logger.Debugf("runner has been instantiated with config %s and context %s", r.config, r.context)

	return r, nil
}

// Print prints the available tasks from the configuration file
// to Stdout.
func (r *OrbitRunner) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprint(w, "Configuration file:")
	fmt.Fprintf(w, "\n  %s\t\n", r.context.TemplateFilePath)
	fmt.Fprint(w, "\nAvailable tasks:")

	for _, task := range r.config.Tasks {
		if !task.Private {
			fmt.Fprintf(w, "\n  %s\t%s", task.Use, task.Short)
		}
	}

	// clears the writer as it may contain some weird characters.
	fmt.Fprintln(w, "")

	w.Flush()
}

// Run runs the given tasks.
func (r *OrbitRunner) Run(names ...string) error {
	// populates an array of instances of orbitTask.
	// if a given name doest not match with any tasks defined in the configuration file, throws an error.
	tasks := make([]*orbitTask, len(names))
	for index, name := range names {
		tasks[index] = r.getTask(name)
		if tasks[index] == nil {
			return OrbitError.NewOrbitErrorf("task %s does not exist in configuration file %s", name, r.context.TemplateFilePath)
		}
	}

	// alright, let's run each task.
	for _, task := range tasks {
		if err := r.run(task); err != nil {
			return err
		}
	}

	return nil
}

// run executes the stack of commands from the given task.
func (r *OrbitRunner) run(task *orbitTask) error {
	logger.Debugf("running task %s", task.Use)

	for _, cmd := range task.Run {

		var e *exec.Cmd
		if task.Shell != "" {
			shellAndParams := strings.Fields(task.Shell)
			shell := shellAndParams[0]
			parameters := append(shellAndParams[1:], cmd)
			e = exec.Command(shell, parameters...)
		} else {
			if runtime.GOOS == "windows" {
				e = exec.Command(os.Getenv(defaultWindowsShellEnvVariable), "/c", cmd)
			} else {
				e = exec.Command(os.Getenv(defaultPosixShellEnvVariable), "-c", cmd)
			}
		}

		logger.Debugf("executing command %s", e.Args)

		e.Stdout = os.Stdout
		e.Stderr = os.Stderr
		e.Stdin = os.Stdin

		if err := e.Run(); err != nil {
			return err
		}
	}

	return nil
}

// getTask returns an instance of orbitTask if found or nil.
func (r *OrbitRunner) getTask(name string) *orbitTask {
	for _, task := range r.config.Tasks {
		if name == task.Use {
			return task
		}
	}

	return nil
}
