/*
Package generator implements a solution to parse data-driven templates and generate output.

A data-driven template is executed by applying it the data structure provided by the application context.
*/
package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gulien/orbit/context"
)

// OrbitGenerator provides a set of functions which help to execute a data-driven template.
type OrbitGenerator struct {
	// context is an instance of OrbitContext.
	context *context.OrbitContext
}

// NewOrbitGenerator instantiates a new instance of OrbitGenerator.
func NewOrbitGenerator(context *context.OrbitContext) *OrbitGenerator {
	return &OrbitGenerator{
		context: context,
	}
}

/*
Parse executes a data-driven template by applying it the data structure provided by the application context.

Returns the resulting bytes.
*/
func (g *OrbitGenerator) Parse() (bytes.Buffer, error) {
	var data bytes.Buffer

	tmpl, err := template.ParseFiles(g.context.TemplateFilePath)
	if err != nil {
		return data, fmt.Errorf("unable to parse the template file \"%s\":\n%s", g.context.TemplateFilePath, err)
	}

	if err := tmpl.Execute(&data, g.context); err != nil {
		return data, fmt.Errorf("unable to execute the template file \"%s\":\n%s", g.context.TemplateFilePath, err)
	}

	return data, nil
}

/*
WriteOutputFile writes bytes into a file.

If the file does not exist, this function will create it.

This function should be called after Parse function.
*/
func (g *OrbitGenerator) WriteOutputFile(outputPath string, data bytes.Buffer) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create the output file \"%s\":\n%s", outputPath, err)
	}

	_, err = file.Write(data.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write into the output file \"%s\":\n%s", outputPath, err)
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
