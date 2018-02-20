package context

import (
	"path/filepath"
	"testing"
)

// Tests if an orbitDumpDecoder instance is able to decode a simple value.
func TestOrbitDumpDecoderDecode(t *testing.T) {
	// case 1: uses a simple value.
	d := &orbitDumbDecoder{value: "raw data"}
	if _, err := d.decode(); err != nil {
		t.Error("orbitDumpDecoder should have been able to decode its value!")
	}
}

// Tests if an orbitYAMLDecoder instance is able to decode a YAML file
// or fails if the YAML file is broken.
func TestOrbitYAMLDecoderDecode(t *testing.T) {
	// case 1: uses a broken YAML file.
	brokenDataSourceFilePath, err := filepath.Abs("../_tests/broken-data-source.yml")
	if err != nil {
		panic(err)
	}

	d := &orbitYAMLDecoder{value: brokenDataSourceFilePath}
	if _, err := d.decode(); err == nil {
		t.Error("orbitYAMLDecoder should not have been able to decode its value!")
	}

	// case 2: uses a correct YAML file.
	dataSourceFilePath, err := filepath.Abs("../_tests/data-source.yml")
	if err != nil {
		panic(err)
	}

	d = &orbitYAMLDecoder{value: dataSourceFilePath}
	if _, err := d.decode(); err != nil {
		t.Error("orbitYAMLDecoder should have been able to decode its value!")
	}
}

// Tests if an orbitTOMLDecoder instance is able to decode a TOML file
// or fails if the TOML file is broken.
func TestOrbitTOMLDecoderDecode(t *testing.T) {
	// case 1: uses a broken TOML file.
	brokenDataSourceFilePath, err := filepath.Abs("../_tests/broken-data-source.toml")
	if err != nil {
		panic(err)
	}

	d := &orbitTOMLDecoder{value: brokenDataSourceFilePath}
	if _, err := d.decode(); err == nil {
		t.Error("orbitTOMLDecoder should not have been able to decode its value!")
	}

	// case 2: uses a correct TOML file.
	dataSourceFilePath, err := filepath.Abs("../_tests/data-source.toml")
	if err != nil {
		panic(err)
	}

	d = &orbitTOMLDecoder{value: dataSourceFilePath}
	if _, err := d.decode(); err != nil {
		t.Error("orbitTOMLDecoder should have been able to decode its value!")
	}
}

// Tests if an orbitJSONDecoder instance is able to decode a JSON file
// or fails if the JSON file is broken.
func TestOrbitJSONDecoderDecode(t *testing.T) {
	// case 1: uses a broken JSON file.
	brokenDataSourceFilePath, err := filepath.Abs("../_tests/broken-data-source.json")
	if err != nil {
		panic(err)
	}

	d := &orbitJSONDecoder{value: brokenDataSourceFilePath}
	if _, err := d.decode(); err == nil {
		t.Error("orbitJSONDecoder should not have been able to decode its value!")
	}

	// case 2: uses a correct JSON file.
	dataSourceFilePath, err := filepath.Abs("../_tests/data-source.json")
	if err != nil {
		panic(err)
	}

	d = &orbitJSONDecoder{value: dataSourceFilePath}
	if _, err := d.decode(); err != nil {
		t.Error("orbitJSONDecoder should have been able to decode its value!")
	}
}

// Tests if an orbitEnvFileDecoder instance is able to decode a .env file
// or fails if the .env file is broken.
func TestOrbitEnvFileDecoderDecode(t *testing.T) {
	// case 1: uses a broken .env file.
	brokenDataSourceFilePath, err := filepath.Abs("../_tests/.broken-env")
	if err != nil {
		panic(err)
	}

	d := &orbitEnvFileDecoder{value: brokenDataSourceFilePath}
	if _, err := d.decode(); err == nil {
		t.Error("orbitEnvFileDecoder should not have been able to decode its value!")
	}

	// case 2: uses a correct .env file.
	dataSourceFilePath, err := filepath.Abs("../_tests/.env")
	if err != nil {
		panic(err)
	}

	d = &orbitEnvFileDecoder{value: dataSourceFilePath}
	if _, err := d.decode(); err != nil {
		t.Error("orbitEnvFileDecoder should have been able to decode its value!")
	}
}
