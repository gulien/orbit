package config

type (
	// OrbitConfig struct represents a YAML file defining commands.
	OrbitConfig struct {
		// Commands represents the list of commands.
		Commands []*OrbitCommand `yaml:"commands"`
	}

	// OrbitCommand struct represents a section in the YAML file defining a command.
	OrbitCommand struct {
		// Use is the name of the command.
		Use string `yaml:"use"`
		// Short describes what the command does.
		Short string `yaml:"short,omitempty"`
		// Run is the stacks of commands to execute.
		Run []string `yaml:"run"`
	}
)
