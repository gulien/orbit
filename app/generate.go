package app

import (
	"github.com/gulien/orbit/app/context"
	"github.com/gulien/orbit/app/generator"

	"github.com/spf13/cobra"
)

var (
	// outputFilePath is the path of the resulting file.
	outputFilePath string
	// templateDelimiters is the optional (overriding) pair of template delimiters.
	templateDelimiters []string

	// generateCmd is the instance of generate command.
	generateCmd = &cobra.Command{
		Use:           "generate",
		Short:         "Generates a file according to a data-driven template",
		Long:          "Generates a file according to a data-driven template.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          generate,
	}
)

// init initializes a generateCmd instance with some flags and adds it to the RootCmd.
func init() {
	generateCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "specify the output file which will be generated from a data-driven template")
	generateCmd.Flags().StringSliceVar(&templateDelimiters, "delimiters", make([]string, 2), "optionally specify template delimiters")
	RootCmd.AddCommand(generateCmd)
}

/*
generate transforms a data-driven template to a resulting file.

If no output file is given, prints the result to Stdout.
*/
func generate(cmd *cobra.Command, args []string) error {
	// first, let's instantiate our Orbit context.
	ctx, err := context.NewOrbitContext(templateFilePath, payload, templates, templateDelimiters)
	if err != nil {
		return err
	}

	// then retrieves the data from the template file.
	g := generator.NewOrbitGenerator(ctx)
	data, err := g.Execute()
	if err != nil {
		return err
	}

	return g.Flush(outputFilePath, data)
}
