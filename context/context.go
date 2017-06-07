package context

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/gulien/orbit/helpers"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type (
	// OrbitContext struct contains the data necessary for generating a file from a template.
	OrbitContext struct {
		// TemplateFilePath is the path of the template.
		TemplateFilePath string
		// Values map contains data from YAML files.
		Values map[string]interface{}
		// EnvFiles map contains pairs from .env files.
		EnvFiles map[string]map[string]string
		// Env map contains pairs from environment variables.
		Env map[string]string
		// Os is the OS name at runtime.
		Os string
	}

	// OrbitFileMap struct represents a parameter given to some flags of an Orbit command.
	// Flags: -v --values, -e --env
	OrbitFileMap struct {
		// Name is the given name of the file.
		Name string
		// Path is the path of the file.
		Path string
	}
)

// NewOrbitContext function instantiates a new OrbitContext.
func NewOrbitContext(templateFilePath string, valuesFiles string, envFiles string) (*OrbitContext, error) {
	// as the template is mandatory, we must check its validity.
	if templateFilePath == "" || !helpers.FileExist(templateFilePath) {
		return nil, fmt.Errorf("Template file %s does not exist", templateFilePath)
	}

	// let's instantiates our OrbitContext!
	ctx := &OrbitContext{
		TemplateFilePath: templateFilePath,
		Os:               runtime.GOOS,
	}

	// checks if a file with values has been specified.
	if valuesFiles != "" {
		data, err := getValuesMap(valuesFiles)
		if err != nil {
			return nil, err
		}

		ctx.Values = data
	}

	// checks if a .env file has been specified.
	if envFiles != "" {
		data, err := getEnvFilesMap(envFiles)
		if err != nil {
			return nil, err
		}

		ctx.EnvFiles = data
	}

	// last but not least, populates the Env map.
	ctx.Env = getEnvMap()

	return ctx, nil
}

// getValuesMap function retrieves values from YAML files.
func getValuesMap(valuesFiles string) (map[string]interface{}, error) {
	filesMap, err := getFilesMap(valuesFiles)
	if err != nil {
		return nil, err
	}

	valuesMap := make(map[string]interface{})
	for _, f := range filesMap {
		// first, checks if the file exists
		if !helpers.FileExist(f.Path) {
			return nil, fmt.Errorf("Values file %s does not exist", f.Path)
		}

		// alright, let's read it to retrieve its data!
		data, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return nil, fmt.Errorf("Failed to read the values file %s:\n%s", f.Path, err)
		}

		// last but not least, parses the YAML.
		var values interface{}
		if err := yaml.Unmarshal(data, &values); err != nil {
			return nil, fmt.Errorf("Values file %s is not a valid YAML file:\n%s", f.Path, err)
		}

		valuesMap[f.Name] = values
	}

	return valuesMap, nil
}

// getEnvFilesMap function retrieves pairs from .env files.
func getEnvFilesMap(envFiles string) (map[string]map[string]string, error) {
	filesMap, err := getFilesMap(envFiles)
	if err != nil {
		return nil, err
	}

	envFilesMap := make(map[string]map[string]string)
	for _, f := range filesMap {
		// first, checks if the file exists
		if !helpers.FileExist(f.Path) {
			return nil, fmt.Errorf("Env file %s does not exist", f.Path)
		}

		// then parses the .env file to retrieve pairs.
		envFilesMap[f.Name], err = godotenv.Read(f.Path)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse the env file %s:\n%s", f.Path, err)
		}
	}

	return envFilesMap, nil
}

// getEnvMap function retrieves all pairs from environment variables.
func getEnvMap() map[string]string {
	envMap := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		envMap[pair[0]] = pair[1]
	}

	return envMap
}

// getFilesMap function reads a string and populates an array of OrbitFileMap instances.
func getFilesMap(s string) ([]*OrbitFileMap, error) {
	var filesMap []*OrbitFileMap

	// checks if the given string is a map of files:
	// if not, considers the string as a path.
	// otherwise tries to populate an array of OrbitFileMap instances.
	parts := strings.Split(s, ";")
	if len(parts) == 1 && len(strings.Split(s, ",")) == 1 {
		filesMap = append(filesMap, &OrbitFileMap{"default", s})
	} else {
		for _, part := range parts {
			data := strings.Split(part, ",")
			if len(data) != 2 {
				return filesMap, fmt.Errorf("Unable to process the files map %s", s)
			}

			filesMap = append(filesMap, &OrbitFileMap{data[0], data[1]})
		}
	}

	return filesMap, nil
}
