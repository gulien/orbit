package helpers

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// Tests FileExists function.
func TestFileExists(t *testing.T) {
	file, _ := filepath.Abs("../.assets/tests/orbit.yml")
	if !FileExists(file) {
		t.Error("File should exist!")
	}
}

// Tests Unmarshal function.
func TestUnmarshal(t *testing.T) {
	path, err := filepath.Abs("../.assets/tests/values.yml")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var values interface{}
	if err := Unmarshal(data, &values); err != nil {
		t.Error("File should have been parsed as YAML!")
	}

	path, err = filepath.Abs("../.assets/tests/wrong_values.yml")
	if err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := Unmarshal(data, &values); err == nil {
		t.Error("File should not have been parsed as YAML!")
	}
}
