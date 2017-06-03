package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gulien/orbit/context"
)

type (
	// OrbitGenerator struct helps to generate a file from a template.
	OrbitGenerator struct {
		// context is an instance of OrbitContext.
		context *context.OrbitContext
	}
)

// NewOrbitGenerator func instantiates a new instance of OrbitGenerator.
func NewOrbitGenerator(context *context.OrbitContext) *OrbitGenerator {
	return &OrbitGenerator{
		context,
	}
}

// Parse function parses a template and populates it.
func (g *OrbitGenerator) Parse() (bytes.Buffer, error) {
	tmpl := template.New(g.context.TemplateFilePath)
	tmpl, err := tmpl.ParseFiles(g.context.TemplateFilePath)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse the template file %s:\n%s", g.context.TemplateFilePath, err)
	}

	var data bytes.Buffer
	if err := tmpl.Execute(&data, g.context); err != nil {
		return nil, fmt.Errorf("Failed to execute the template file %s:\n%s", g.context.TemplateFilePath, err)
	}

	return data, nil
}

// WriteOutputFile function write data into a file.
func (g *OrbitGenerator) WriteOutputFile(outputPath string, data bytes.Buffer) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Failed to create the output file %s:\n%s", outputPath, err)
	}

	_, err = file.Write(data.Bytes())
	if err != nil {
		return fmt.Errorf("Failed to write into the output file %s:\n%s", outputPath, err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("Failed to close the writer:\n%s", err)
	}

	return nil
}
