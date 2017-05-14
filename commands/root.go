package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "orbit",
	Short: "A swiss knife for managing the Docker environment of your projects.",
	Long: `Orbit is able to:

- Downloads Orbit's templates from github.com
- Generates your containers' configuration files on the fly
- Runs commands defined in the orbit.yml file`,
}
