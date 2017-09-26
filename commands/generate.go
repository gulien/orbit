package commands

import (
	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/generator"

	"github.com/spf13/cobra"
)

var (
	// templateFilePath is the path of the data-driven template.
	templateFilePath string

	// outputFilePath is the path of the resulting file.
	outputFilePath string

	// generateCmd is the instance of generate command.
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
	RootCmd.AddCommand(generateCmd)
}

/*
generate transforms a template to a resulting file.

If no output file is given, prints the result to Stdout.
*/
func generate(cmd *cobra.Command, args []string) error {
	// first, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(templateFilePath, ValuesFiles, EnvFiles, RawData)
	if err != nil {
		return err
	}

	// then retrieves the data from the template file.
	g := generator.NewOrbitGenerator(ctx)
	data, err := g.Parse()
	if err != nil {
		return err
	}

	return g.Output(outputFilePath, data)
}
