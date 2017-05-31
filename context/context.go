package context

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/gulien/orbit/helpers"

	"github.com/joho/godotenv"
	jww "github.com/spf13/jwalterweatherman"
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
func NewOrbitContext(templateFilePath string, valuesFilePath string, envFilePath string) *OrbitContext {
	// as the template is mandatory, we must check its validity.
	if templateFilePath == "" || helpers.FileDoesNotExist(templateFilePath) {
		jww.ERROR.Println("Template file %s does not exist", templateFilePath)
		os.Exit(1)
	}

	// let's instantiates our OrbitContext!
	orbitContext := &OrbitContext{
		TemplateFilePath: templateFilePath,
		Os:               runtime.GOOS,
	}

	// checks if a file with values has been specified.
	if valuesFilePath != "" && helpers.FileExist(valuesFilePath) {
		orbitContext.Values = getValuesMap(valuesFilePath)
	}

	// checks if a .env file has been specified.
	if envFilePath != "" && helpers.FileExist(envFilePath) {
		orbitContext.EnvFile = getEnvFileMap(envFilePath)
	}

	// last but not least, populates the Env map.
	orbitContext.Env = getEnvMap()

	return orbitContext
}

// getValuesMap function retrieves values from a YAML file.
func getValuesMap(valuesFilePath string) map[interface{}]interface{} {
	// the file containing values must be a valid YAML file.
	if !helpers.IsYAML(valuesFilePath) {
		jww.ERROR.Println("Values file %s is not a valid YAML file", valuesFilePath)
		os.Exit(1)
	}

	// alright, let's read it to retrieve its data!
	data, err := ioutil.ReadFile(valuesFilePath)
	if err != nil {
		jww.FATAL.Println("Failed to read the values file %s: %s", valuesFilePath, err)
		os.Exit(1)
	}

	// last but not least, parses the YAML.
	valuesMaps := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(data, &valuesMaps); err != nil {
		jww.ERROR.Println("Values file %s is not a valid YAML file: %s", valuesFilePath, err)
		os.Exit(1)
	}

	return valuesMaps
}

// getEnvFileMap function retrieves pairs from a .env file.
func getEnvFileMap(envFilePath string) map[string]string {
	envFileMap, err := godotenv.Read(envFilePath)
	if err != nil {
		jww.ERROR.Println("Unable to parse the env file %s: %s", envFilePath, err)
		os.Exit(1)
	}

	return envFileMap
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
