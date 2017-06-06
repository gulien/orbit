package commands

import (
	"fmt"
	"os"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/notifier"

	"github.com/gulien/orbit/generator"
	"github.com/spf13/cobra"
)

var (
	// templateFilePath is the path of the template.
	templateFilePath string
	// outputFilePath is the path of the resulting file from the template.
	outputFilePath string
)

// generateCmd is the instance of Orbit's generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a file according to a template",
	Run:   generate,
}

// init function initializes a generateCmd instance with some flags and adds it to the RootCmd.
func init() {
	generateCmd.Flags().StringVarP(&templateFilePath, "template", "t", "", "specify the template")
	generateCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "specify the output file which will be generated from the template")
	generateCmd.Flags().StringVarP(&ValuesFiles, "values", "v", "", "specify a YAML file or a map of YAML files listing values used in the template")
	generateCmd.Flags().StringVarP(&EnvFiles, "env", "e", "", "specify a .env file or a map of .env files listing values used in the template")
	RootCmd.AddCommand(generateCmd)
}

// generate function transforms a template to a resulting file.
// if no output specified, prints the result to stdout.
func generate(cmd *cobra.Command, args []string) {
	// first, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(templateFilePath, ValuesFiles, EnvFiles)
	if err != nil {
		notifier.Error(err)
	}

	// then retrieves the data from the template file.
	gen := generator.NewOrbitGenerator(ctx)
	data, err := gen.Parse()
	if err != nil {
		notifier.Error(err)
	}

	// if an output file path has been specified, writes the result into it.
	if outputFilePath != "" {
		if err := gen.WriteOutputFile(outputFilePath, data); err != nil {
			notifier.Error(err)
		}
	} else {
		// ok, no output specified, let's print the result to stdout.
		fmt.Println(data)
	}

	// everything good!
	os.Exit(0)
}
