package commands

import (
	"fmt"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/generator"

	"github.com/spf13/cobra"
)

var (
	// templateFilePath is the path of the template.
	templateFilePath string

	// outputFilePath is the path of the resulting file from the template.
	outputFilePath string

	// generateCmd is the instance of Orbit's generate command.
	generateCmd = &cobra.Command{
		Use:           "generate",
		Short:         "Generates a file according to a template",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          generate,
	}
)

// init initializes a generateCmd instance with some flags and adds it to the RootCmd.
func init() {
	generateCmd.Flags().StringVarP(&templateFilePath, "template", "t", "", "specify the template")
	generateCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "specify the output file which will be generated from the template")
	generateCmd.Flags().StringVarP(&ValuesFiles, "values", "v", "", "specify a YAML file or a map of YAML files listing values used in the template")
	generateCmd.Flags().StringVarP(&EnvFiles, "env", "e", "", "specify a .env file or a map of .env files listing values used in the template")
	RootCmd.AddCommand(generateCmd)
}

// generate transforms a template to a resulting file.
// If no output file is given, prints the result to Stdout.
func generate(cmd *cobra.Command, args []string) error {
	// first, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(templateFilePath, ValuesFiles, EnvFiles)
	if err != nil {
		return err
	}

	// then retrieves the data from the template file.
	gen := generator.NewOrbitGenerator(ctx)
	data, err := gen.Parse()
	if err != nil {
		return err
	}

	// if an output file has been given, writes the result into it.
	if outputFilePath != "" {
		if err := gen.WriteOutputFile(outputFilePath, data); err != nil {
			return err
		}
	} else {
		// ok, no output file given, let's print the result to Stdout.
		fmt.Println(string(data.Bytes()))
	}

	return nil
}
