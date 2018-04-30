/*
Package generator implements a solution to parse a data-driven template and generate an output from it.

A data-driven template is executed by applying it the data structure provided by the payload from the application context.
*/
package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gulien/orbit/app/context"
	OrbitError "github.com/gulien/orbit/app/error"
	"github.com/gulien/orbit/app/logger"

	"github.com/Masterminds/sprig"
)

type (
	// OrbitGenerator provides a set of functions which helps to execute a data-driven template.
	OrbitGenerator struct {
		// context is an instance of OrbitContext.
		context *context.OrbitContext

		// funcMap contains sprig functions and custom os function.
		funcMap template.FuncMap
	}

	// orbitData is a simple handler of the payload given by the user.
	orbitData struct {
		// Orbit will be filled by the payload from the context.
		// The goal here is to allow the use of the syntax {{ .Orbit }}
		// in a data-driven template.
		Orbit map[string]interface{}
	}
)

// NewOrbitGenerator creates an instance of OrbitGenerator.
func NewOrbitGenerator(context *context.OrbitContext) *OrbitGenerator {
	funcMap := sprig.TxtFuncMap()
	funcMap["os"] = getOS
	funcMap["verbose"] = isVerbose
	funcMap["debug"] = isDebug

	g := &OrbitGenerator{
		context: context,
		funcMap: funcMap,
	}

	logger.Debugf("generator has been instantiated with context %s", g.context)

	return g
}

/*
Execute executes a data-driven template by applying it the data structure provided by the application context.

Returns the resulting bytes.
*/
func (g *OrbitGenerator) Execute() (bytes.Buffer, error) {
	var data bytes.Buffer

	tmpl, err := template.New(filepath.Base(g.context.TemplateFilePath)).Funcs(g.funcMap).ParseFiles(g.context.TemplateFilePath)
	if err != nil {
		return data, OrbitError.NewOrbitErrorf("unable to parse the template file %s. Details:\n%s", g.context.TemplateFilePath, err)
	}

	tmpl.Option("missingkey=error")

	orbitData := &orbitData{
		Orbit: g.context.Payload,
	}

	if err := tmpl.Execute(&data, orbitData); err != nil {
		return data, OrbitError.NewOrbitErrorf("unable to execute the template file %s. Details:\n%s", g.context.TemplateFilePath, err)
	}

	logger.Debugf("template file %s has been parsed and the following data have been retrieved:\n%s", g.context.TemplateFilePath, data.String())

	return data, nil
}

/*
Flush writes bytes into a file or to Stdout if no output path given.

This function should be called after Execute function.
*/
func (g *OrbitGenerator) Flush(outputPath string, data bytes.Buffer) error {
	if outputPath != "" {
		return flushToFile(outputPath, data)
	}

	// ok, no output file given, let's flush the result to Stdout.
	flushToStdout(data)
	return nil
}

/*
flushToFile writes bytes into a file.

If the file does not exist, this function will create it.
*/
func flushToFile(outputPath string, data bytes.Buffer) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return OrbitError.NewOrbitErrorf("unable to create the output file %s. Details:\n%s", outputPath, err)
	}

	defer file.Close()

	_, err = file.Write(data.Bytes())
	if err != nil {
		return OrbitError.NewOrbitErrorf("unable to flushToFile into the output file %s. Details:\n%s", outputPath, err)
	}

	logger.Infof("output file %s has been created", outputPath)

	return nil
}

// flushToStdout writes bytes to Stdout.
func flushToStdout(data bytes.Buffer) {
	logger.Infof("no output file given, printing the result to Stdout")
	fmt.Println(string(data.Bytes()))
}
