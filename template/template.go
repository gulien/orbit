package template

import (
	"bytes"
	"html/template"
	"os"

	"github.com/gulien/orbit/context"

	jww "github.com/spf13/jwalterweatherman"
)

// ReadTemplate function parses a template and populates it.
func ReadTemplate(orbitContext *context.OrbitContext) bytes.Buffer {
	tmpl := template.New(orbitContext.TemplateFilePath)
	tmpl, err := tmpl.ParseFiles(orbitContext.TemplateFilePath)

	if err != nil {
		jww.FATAL.Println("Failed to parse the template file %s: %s", orbitContext.TemplateFilePath, err)
		os.Exit(1)
	}

	var data bytes.Buffer
	if err := tmpl.Execute(&data, orbitContext); err != nil {
		jww.FATAL.Println("Failed to execute the template file %s: %s", orbitContext.TemplateFilePath, err)
		os.Exit(1)
	}

	return data
}

// WriteOutputFile function write data into a file.
func WriteOutputFile(outputPath string, data bytes.Buffer) {
	file, err := os.Create(outputPath)
	if err != nil {
		jww.FATAL.Println("Failed to create the output file %s: %s", outputPath, err)
		os.Exit(1)
	}

	_, err = file.Write(data.Bytes())
	if err != nil {
		jww.FATAL.Println("Failed to write the output file %s: %s", outputPath, err)
		os.Exit(1)
	}

	err = file.Close()
	if err != nil {
		jww.FATAL.Println("Failed to close the writer: %s", outputPath, err)
		os.Exit(1)
	}
}
