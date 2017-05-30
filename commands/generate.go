package commands

import (
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	// TODO describe
	templateFile string
	// TODO describe
	outputFile   string
)

// TODO describe
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a file according to a template",
	Run:   generate,
}

// TODO describe
func init() {
	generateCmd.Flags().StringVarP(&templateFile, "template", "t", "", "Template to read")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to generate from template")
	generateCmd.Flags().StringVarP(&ValuesFile, "values", "v", "", "File (YAML) defining values used in the template")
	generateCmd.Flags().StringVarP(&EnvFile, "env_file", "e", "", "File storing env values used in the template")
	RootCmd.AddCommand(generateCmd)
}

// TODO describe
func generate(cmd *cobra.Command, args []string) {
	jww.ERROR.Println("Nothing to generate")
	os.Exit(1)
}
