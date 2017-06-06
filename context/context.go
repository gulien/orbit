package context

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/gulien/orbit/helpers"

	"encoding/json"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type (
	// OrbitContext struct contains the data necessary for generating a file from a template.
	OrbitContext struct {
		// TemplateFilePath is the path of the template.
		TemplateFilePath string
		// Values map contains data from YAML files.
		Values map[string]map[interface{}]interface{}
		// EnvFiles map contains pairs from .env files.
		EnvFiles map[string]map[string]string
		// Env map contains pairs from environment variables.
		Env map[string]string
		// Os is the OS name at runtime.
		Os string
	}

	// OrbitFileMap struct represents a section of some JSON parameter given to an Orbit command.
	// Flags: -v --values, -e --env
	OrbitFileMap struct {
		// Name is the given name of the file.
		Name string `json:"name"`
		// Path is the path of the file.
		Path string `json:"path"`
	}
)

// NewOrbitContext function instantiates a new OrbitContext.
func NewOrbitContext(templateFilePath string, valuesFiles string, envFiles string) (*OrbitContext, error) {
	// as the template is mandatory, we must check its validity.
	if templateFilePath == "" || helpers.FileDoesNotExist(templateFilePath) {
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
func getValuesMap(valuesFiles string) (map[string]map[interface{}]interface{}, error) {
	filesMap, err := getFilesMap(valuesFiles)
	if err != nil {
		return nil, err
	}

	valuesMap := make(map[string]map[interface{}]interface{})
	for _, f := range filesMap {
		// first, checks if the file exists
		if helpers.FileDoesNotExist(f.Path) {
			return nil, fmt.Errorf("Values file %s does not exist", f.Path)
		}

		// the file containing values must be a valid YAML file.
		if !helpers.IsYAML(f.Path) {
			return nil, fmt.Errorf("Values file %s is not a valid YAML file", f.Path)
		}

		// alright, let's read it to retrieve its data!
		data, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return nil, fmt.Errorf("Failed to read the values file %s:\n%s", f.Path, err)
		}

		// last but not least, parses the YAML.
		valuesMap[f.Name] = make(map[interface{}]interface{})
		if err := yaml.Unmarshal(data, &valuesMap); err != nil {
			return nil, fmt.Errorf("Values file %s is not a valid YAML file:\n%s", f.Path, err)
		}
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
		if helpers.FileDoesNotExist(f.Path) {
			return nil, fmt.Errorf("Env file %s does not exist", f.Path)
		}

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

	// checks if the given string is in JSON format:
	// if not, considers the string as a path.
	// otherwise tries to populate an array of OrbitFileMap instances.
	if !helpers.IsJSONString(s) {
		filesMap = append(filesMap, &OrbitFileMap{"default", s})
	} else if err := json.Unmarshal([]byte(s), &filesMap); err != nil {
		return filesMap, fmt.Errorf("Unable to read JSON %s:\n%s", s, err)
	}

	return filesMap, nil
}
