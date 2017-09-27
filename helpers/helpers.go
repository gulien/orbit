// Package helpers implements simple functions used across the application.
package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// FileExists returns true if the specified path exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Unmarshal YAML to map[string]interface{} instead of map[interface{}]interface{}.
// The goal here is to allow the use of dictionaries and dict functions
// from Sprig library in a template.
func Unmarshal(in []byte, out interface{}) error {
	var res interface{}

	if err := yaml.Unmarshal(in, &res); err != nil {
		return err
	}

	*out.(*interface{}) = cleanupMapValue(res)

	return nil
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
