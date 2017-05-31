package commands

import (
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	// templateFile is the path of a template.
	templateFile string
	// outputFile is the path of the resulting file from the template.
	outputFile string
)

// generateCmd is the instance of orbit's generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a file according to a template",
	Run:   generate,
}

// init function initializes generateCmd instance with some flags and adds it to the RootCmd.
func init() {
	generateCmd.Flags().StringVarP(&templateFile, "template", "t", "", "specify the template")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "specify the output file which will be generated from the template")
	generateCmd.Flags().StringVarP(&ValuesFile, "values", "v", "", "specify a YAML file listing values used in the template")
	generateCmd.Flags().StringVarP(&EnvFile, "env_file", "e", "", "specify a .env file listing values used in the template")
	RootCmd.AddCommand(generateCmd)
}

// generate function transforms a template to a resulting file.
func generate(cmd *cobra.Command, args []string) {
	jww.ERROR.Println("Nothing to generate")
	os.Exit(1)
}
