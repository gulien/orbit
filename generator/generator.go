/*
Package generator implements a solution to parse data-driven templates and generate output.

A data-driven template is executed by applying it the data structure provided by the application context.
*/
package generator

import (
	"bytes"
	"os"
	"text/template"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/logger"
)

// OrbitGenerator provides a set of functions which help to execute a data-driven template.
type OrbitGenerator struct {
	// context is an instance of OrbitContext.
	context *context.OrbitContext
}

// NewOrbitGenerator creates an instance of OrbitGenerator.
func NewOrbitGenerator(context *context.OrbitContext) *OrbitGenerator {
	g := &OrbitGenerator{
		context: context,
	}

	logger.Debugf("generator has been instantiated with context %s", g.context)

	return g
}

/*
Parse executes a data-driven template by applying it the data structure provided by the application context.

Returns the resulting bytes.
*/
func (g *OrbitGenerator) Parse() (bytes.Buffer, error) {
	var data bytes.Buffer

	tmpl, err := template.ParseFiles(g.context.TemplateFilePath)
	if err != nil {
		return data, errors.NewOrbitErrorf("unable to parse the template file %s. Details:\n%s", g.context.TemplateFilePath, err)
	}

	if err := tmpl.Execute(&data, g.context); err != nil {
		return data, errors.NewOrbitErrorf("unable to execute the template file %s. Details:\n%s", g.context.TemplateFilePath, err)
	}

	logger.Debugf("template file %s has been parsed and the following data have been retrieved:\n%s", g.context.TemplateFilePath, data.String())

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
		return errors.NewOrbitErrorf("unable to create the output file %s. Details:\n%s", outputPath, err)
	}

	_, err = file.Write(data.Bytes())
	if err != nil {
		return errors.NewOrbitErrorf("unable to write into the output file %s. Details:\n%s", outputPath, err)
	}

	err = file.Close()
	if err != nil {
		return err
	}

	logger.Debugf("the template file %s has been executed to the output file %s", g.context.TemplateFilePath, outputPath)

	return nil
}
