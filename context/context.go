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
		// Values map contains the data from a YAML file.
		Values map[interface{}]interface{}
		// EnvFile map contains the pairs from a .env file.
		EnvFile map[string]string
		// Env map contains the pairs from environments variables.
		Env map[string]string
		// Os is the OS name at runtime.
		Os string
	}
)

// NewOrbitContext function instantiates a new OrbitContext.
func NewOrbitContext(templateFilePath string, valuesFilePath string, envFilePath string) (*OrbitContext, error) {
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
	if valuesFilePath != "" && helpers.FileExist(valuesFilePath) {
		data, err := getValuesMap(valuesFilePath)
		if err != nil {
			return nil, err
		}

		ctx.Values = data
	}

	// checks if a .env file has been specified.
	if envFilePath != "" && helpers.FileExist(envFilePath) {
		data, err := getEnvFileMap(envFilePath)
		if err != nil {
			return nil, err
		}

		ctx.EnvFile = data
	}

	// last but not least, populates the Env map.
	ctx.Env = getEnvMap()

	return ctx, nil
}

// getValuesMap function retrieves values from a YAML file.
func getValuesMap(valuesFilePath string) (map[interface{}]interface{}, error) {
	// the file containing values must be a valid YAML file.
	if !helpers.IsYAML(valuesFilePath) {
		return nil, fmt.Errorf("Values file %s is not a valid YAML file", valuesFilePath)
	}

	// alright, let's read it to retrieve its data!
	data, err := ioutil.ReadFile(valuesFilePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the values file %s:\n%s", valuesFilePath, err)
	}

	// last but not least, parses the YAML.
	valuesMap := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(data, &valuesMap); err != nil {
		return nil, fmt.Errorf("Values file %s is not a valid YAML file:\n%s", valuesFilePath, err)
	}

	return valuesMap, nil
}

// getEnvFileMap function retrieves pairs from a .env file.
func getEnvFileMap(envFilePath string) (map[string]string, error) {
	envFileMap, err := godotenv.Read(envFilePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse the env file %s:\n%s", envFilePath, err)
	}

	return envFileMap, nil
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
