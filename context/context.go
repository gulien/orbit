/*
Package context helps to populate the application context.

The main goal of the application context is to gather all the data which will be applied to a data-driven template.
*/
package context

import (
	"io/ioutil"
	"strings"

	"github.com/gulien/orbit/errors"
	"github.com/gulien/orbit/helpers"
	"github.com/gulien/orbit/logger"

	"github.com/joho/godotenv"
)

type (
	// OrbitContext contains the data necessary for executing a data-driven template.
	OrbitContext struct {
		// TemplateFilePath is the path of a data-driven template.
		TemplateFilePath string

		// Payload map contains data from various entries.
		Payload map[string]interface{}
	}

	// OrbitFileMap represents a value given to the --payload flag of generate and run commands.
	// Value format: name,path;name,path;...
	OrbitFileMap struct {
		// Name is the given name of the file.
		Name string

		// Path is the path of the file.
		Path string
	}
)

// NewOrbitContext creates an instance of OrbitContext.
func NewOrbitContext(templateFilePath string, payloadEntries string) (*OrbitContext, error) {
	// as the data-driven template is mandatory, we must check its validity.
	if templateFilePath == "" {
		return nil, errors.NewOrbitErrorf("no data-driven template given", templateFilePath)
	}

	if !helpers.FileExists(templateFilePath) {
		return nil, errors.NewOrbitErrorf("the data-driven template file %s does not exist", templateFilePath)
	}

	// let's instantiates our OrbitContext!
	ctx := &OrbitContext{
		TemplateFilePath: templateFilePath,
	}

	logger.Debugf("context has been instantiated with the data-driven template file %s", ctx.TemplateFilePath)

	// checks if files with values have been specified.
	if valuesFiles != "" {
		data, err := getValuesMap(valuesFiles)
		if err != nil {
			return nil, err
		}

		ctx.Values = data
	}

	logger.Debugf("context has been populated with values %s", ctx.Values)

	// checks if .env files have been specified.
	if envFiles != "" {
		data, err := getEnvFilesMap(envFiles)
		if err != nil {
			return nil, err
		}

		ctx.EnvFiles = data
	}

	logger.Debugf("context has been populated with env files' data %s", ctx.EnvFiles)

	// checks if raw data have been specified.
	if rawData != "" {
		data, err := getRawDataMap(rawData)
		if err != nil {
			return nil, err
		}

		ctx.RawData = data
	}

	logger.Debugf("context has been populated with raw data %s", ctx.RawData)

	return ctx, nil
}

// getValuesMap retrieves values from YAML files.
func getValuesMap(valuesFiles string) (map[string]interface{}, error) {
	filesMap, err := getFilesMap(valuesFiles)
	if err != nil {
		return nil, err
	}

	valuesMap := make(map[string]interface{})
	for _, f := range filesMap {
		// first, checks if the file exists
		if !helpers.FileExists(f.Path) {
			return nil, errors.NewOrbitErrorf("values file %s does not exist", f.Path)
		}

		// alright, let's read it to retrieve its data!
		data, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return nil, errors.NewOrbitErrorf("failed to read the values file %s. Details:\n%s", f.Path, err)
		}

		// last but not least, parses the YAML.
		var values interface{}
		if err := helpers.Unmarshal(data, &values); err != nil {
			return nil, errors.NewOrbitErrorf("unable to parse the values file %s. Details:\n%s", f.Path, err)
		}

		valuesMap[f.Name] = values
	}

	return valuesMap, nil
}

// getEnvFilesMap retrieves pairs from .env files.
func getEnvFilesMap(envFiles string) (map[string]map[string]string, error) {
	filesMap, err := getFilesMap(envFiles)
	if err != nil {
		return nil, err
	}

	envFilesMap := make(map[string]map[string]string)
	for _, f := range filesMap {
		// first, checks if the file exists
		if !helpers.FileExists(f.Path) {
			return nil, errors.NewOrbitErrorf("env file %s does not exist", f.Path)
		}

		// then parses the .env file to retrieve pairs.
		envFilesMap[f.Name], err = godotenv.Read(f.Path)
		if err != nil {
			return nil, errors.NewOrbitErrorf("unable to parse the env file %s. Details:\n%s", f.Path, err)
		}
	}

	return envFilesMap, nil
}

// getFilesMap reads a string and populates an array of OrbitFileMap instances.
func getFilesMap(s string) ([]*OrbitFileMap, error) {
	var filesMap []*OrbitFileMap

	// checks if the given string is a map of files:
	// if not, considers the string as a path.
	// otherwise tries to populate an array of OrbitFileMap instances.
	parts := strings.Split(s, ";")
	if len(parts) == 1 && len(strings.Split(s, ",")) == 1 {
		filesMap = append(filesMap, &OrbitFileMap{
			Name: "default",
			Path: s,
		})
	} else {
		for _, part := range parts {
			data := strings.Split(part, ",")
			if len(data) != 2 {
				return filesMap, errors.NewOrbitErrorf("unable to process the files map %s", s)
			}

			filesMap = append(filesMap, &OrbitFileMap{
				Name: data[0],
				Path: data[1],
			})
		}
	}

	return filesMap, nil
}

// getRawDataMap reads a string and populates a map of strings.
func getRawDataMap(s string) (map[string]string, error) {
	parts := strings.Split(s, ";")

	rawDataMap := make(map[string]string)
	for _, part := range parts {
		data := strings.Split(part, "=")
		if len(data) != 2 {
			return rawDataMap, errors.NewOrbitErrorf("unable to process the raw data %s", s)
		}

		rawDataMap[data[0]] = data[1]
	}

	return rawDataMap, nil
}
