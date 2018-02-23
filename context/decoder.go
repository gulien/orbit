package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gulien/orbit/errors"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type (
	// orbitDecoder is an interface which is implemented for each data source
	// we wish to decode.
	orbitDecoder interface {
		// decode is implemented by each decoder according to its data source.
		decode() (interface{}, error)
	}

	// orbitDumbDecoder is a basic implementation of orbitDecoder.
	orbitDumbDecoder struct {
		// value represents a raw data.
		value string
	}

	// orbitYAMLDecoder allows to decode a YAML file.
	orbitYAMLDecoder struct {
		// value represents a path to a YAML file.
		value string
	}

	// orbitTOMLDecoder allows to decode a TOML file.
	orbitTOMLDecoder struct {
		// value represents a path to a TOML file.
		value string
	}

	// orbitJSONDecoder allows to decode a JSON file.
	orbitJSONDecoder struct {
		// value represents a path to a JSON file.
		value string
	}

	// orbitEnvFileDecoder allows to decode a .env file.
	orbitEnvFileDecoder struct {
		// value represents a path to a .env file.
		value string
	}
)

// decode from orbitDumbDecoder directly returns the value associated
// with the decoder. Its main goal is to be used for raw data.
func (d *orbitDumbDecoder) decode() (interface{}, error) {
	return d.value, nil
}

// decode from orbitYAMLDecoder reads a YAML file and retrieves its data.
func (d *orbitYAMLDecoder) decode() (interface{}, error) {
	data, err := ioutil.ReadFile(d.value)
	if err != nil {
		return nil, errors.NewOrbitErrorf("unable to read the file %s. Details:\n%s", d.value, err)
	}

	var decoded interface{}
	if err := yaml.Unmarshal(data, &decoded); err != nil {
		return nil, errors.NewOrbitErrorf("unable to decode the YAML file %s. Details:\n%s", d.value, err)
	}

	var result interface{}
	cleanup(decoded, &result)

	return result, nil
}

// decode from orbitTOMLDecoder reads a TOML file and retrieves its data.
func (d *orbitTOMLDecoder) decode() (interface{}, error) {
	var decoded interface{}
	if _, err := toml.DecodeFile(d.value, &decoded); err != nil {
		return nil, errors.NewOrbitErrorf("unable to decode the TOML file %s. Details:\n%s", d.value, err)
	}

	var result interface{}
	cleanup(decoded, &result)

	return result, nil
}

// decode from orbitJSONDecoder reads a JSON file and retrieves its data.
func (d *orbitJSONDecoder) decode() (interface{}, error) {
	data, err := ioutil.ReadFile(d.value)
	if err != nil {
		return nil, errors.NewOrbitErrorf("unable to read the file %s. Details:\n%s", d.value, err)
	}

	var decoded interface{}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return nil, errors.NewOrbitErrorf("unable to decode the JSON file %s. Details:\n%s", d.value, err)
	}

	var result interface{}
	cleanup(decoded, &result)

	return result, nil
}

// decode from orbitEnvFileDecoder reads a .env file and retrieves its data.
func (d *orbitEnvFileDecoder) decode() (interface{}, error) {
	result, err := godotenv.Read(d.value)
	if err != nil {
		return nil, errors.NewOrbitErrorf("unable to decode the .env file %s. Details:\n%s", d.value, err)
	}

	return result, nil
}

/*
cleanup parses an interface recursively to find and update
map[interface{}]interface{} to map[string]interface{}.

The goal here is to allow the use of dictionaries and dict functions
from Sprig library in a data-driven template.
*/
func cleanup(decoded interface{}, out interface{}) {
	*out.(*interface{}) = cleanupMapValue(decoded)
}

// cleanupMapValue parses an interface recursively to find and update
// map[interface{}]interface{} to map[string]interface{}.
func cleanupMapValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return cleanupInterfaceArray(v)
	case map[interface{}]interface{}:
		return cleanupInterfaceMap(v)
	default:
		return v
	}
}

// cleanupInterfaceArray parses an array of interfaces to find and update
// map[interface{}]interface{} to map[string]interface{}.
func cleanupInterfaceArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))

	for i, v := range in {
		res[i] = cleanupMapValue(v)
	}

	return res
}

// cleanupInterfaceArray transforms a map[interface{}]interface{} to map[string]interface{}.
// It also finds and updates its children of type map[interface{}]interface{} to map[string]interface{}.
func cleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = cleanupMapValue(v)
	}

	return res
}
