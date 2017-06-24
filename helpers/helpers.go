// Package helpers implements simple functions used across the application.
package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// FileExists returns true if the specified path exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

/*
Abs returns the absolute path of a given relative path.

This function should only be used in tests!
*/
func Abs(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return absPath
}

/*
ReadYAML returns the data of a YAML file.

This function should only be used in tests!
*/
func ReadYAML(path string) interface{} {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var value interface{}
	if err := yaml.Unmarshal(data, &value); err != nil {
		panic(err)
	}

	return value
}
