package context

import (
	"path/filepath"
	"reflect"
	"testing"
)

// Tests if populating an orbitPayload from a file throws an error
// with a wrong parameter or no error if the parameter is correct.
func TestPopulateFromFile(t *testing.T) {
	// case 1: uses an empty file path to force the usage of the default payload file.
	p := &orbitPayload{}
	if err := p.populateFromFile(""); err != nil {
		t.Error("orbitPayload should have been skipped!")
	}

	// case 2: uses a broken payload file.
	brokenPayloadFilePath, _ := filepath.Abs("../_tests/broken-orbit-payload.yml")
	p = &orbitPayload{}
	if err := p.populateFromFile(brokenPayloadFilePath); err == nil {
		t.Error("orbitPayload should not have been populated!")
	}

	// case 3: uses an existing payload file path.
	customPayloadFilePath, _ := filepath.Abs("../_tests/orbit-payload.yml")
	p = &orbitPayload{}
	if err := p.populateFromFile(customPayloadFilePath); err != nil {
		t.Error("orbitPayload should have been populated!")
	}
}

// Tests if populating an orbitPayload from a string throws an error
// with a wrong parameter or no error if the parameter is correct.
func TestPopulateFromString(t *testing.T) {
	// case 1: uses an empty payload string.
	p := &orbitPayload{}
	if err := p.populateFromString(""); err != nil {
		t.Error("orbitPayload should have been skipped!")
	}

	// case 2: uses a broken payload string.
	p = &orbitPayload{}
	if err := p.populateFromString("key"); err == nil {
		t.Error("orbitPayload should not have been populated!")
	}

	// case 3: uses a correct payload string.
	p = &orbitPayload{}
	if err := p.populateFromString("key,value"); err != nil {
		t.Error("orbitPayload should have been populated!")
	}
}

// Tests if retrieving data from an orbitPayload throws an error
// with a wrong payload entry or no error if the payload entry is correct.
func TestRetrieveData(t *testing.T) {
	// case 1: uses a broken payload entry.
	brokenDataSourceFilePath, _ := filepath.Abs("../_tests/broken-data-source.yml")
	p := &orbitPayload{}
	p.populateFromString("key," + brokenDataSourceFilePath)
	if _, err := p.retrieveData(); err == nil {
		t.Error("orbitPayload should not have been hable to retrieve data!")
	}

	// case 2: uses a correct payload entry.
	dataSourceFilePath, _ := filepath.Abs("../_tests/data-source.yml")
	p = &orbitPayload{}
	p.populateFromString("key," + dataSourceFilePath)
	if _, err := p.retrieveData(); err != nil {
		t.Error("orbitPayload should have been hable to retrieve data!")
	}
}

// Tests if for a given value the function getDecoder returns a correct
// instance of decoder.
func TestGetDecoder(t *testing.T) {
	// case 1: should returns an instance of orbitDumbDecoder
	d := getDecoder("raw data")
	dumbDecoder := &orbitDumbDecoder{}
	if reflect.TypeOf(d) != reflect.TypeOf(dumbDecoder) {
		t.Error("Decoder should have been an instance of orbitDumbDecoder!")
	}

	// case 2: should returns an instance of orbitYAMLDecoder
	YAMLDataSourceFilePath, _ := filepath.Abs("../_tests/data-source.yml")
	d = getDecoder(YAMLDataSourceFilePath)
	YAMLDecoder := &orbitYAMLDecoder{}
	if reflect.TypeOf(d) != reflect.TypeOf(YAMLDecoder) {
		t.Error("Decoder should have been an instance of orbitYAMLDecoder!")
	}

	YAMLDataSourceFilePath, _ = filepath.Abs("../_tests/data-source.yaml")
	d = getDecoder(YAMLDataSourceFilePath)
	YAMLDecoder = &orbitYAMLDecoder{}
	if reflect.TypeOf(d) != reflect.TypeOf(YAMLDecoder) {
		t.Error("Decoder should have been an instance of orbitYAMLDecoder!")
	}

	// case 3: should returns an instance of orbitTOMLDecoder
	TOMLDataSourceFilePath, _ := filepath.Abs("../_tests/data-source.toml")
	d = getDecoder(TOMLDataSourceFilePath)
	TOMLDecoder := &orbitTOMLDecoder{}
	if reflect.TypeOf(d) != reflect.TypeOf(TOMLDecoder) {
		t.Error("Decoder should have been an instance of orbitTOMLDecoder!")
	}

	// case 4: should returns an instance of orbitJSONDecoder
	JSONDataSourceFilePath, _ := filepath.Abs("../_tests/data-source.json")
	d = getDecoder(JSONDataSourceFilePath)
	JSONDecoder := &orbitJSONDecoder{}
	if reflect.TypeOf(d) != reflect.TypeOf(JSONDecoder) {
		t.Error("Decoder should have been an instance of orbitJSONDecoder!")
	}

	// case 5: should returns an instance of orbitEnvFileDecoder
	envFileDataSourceFilePath, _ := filepath.Abs("../_tests/.env")
	d = getDecoder(envFileDataSourceFilePath)
	envFileDecoder := &orbitEnvFileDecoder{}
	if reflect.TypeOf(d) != reflect.TypeOf(envFileDecoder) {
		t.Error("Decoder should have been an instance of orbitEnvFileDecoder!")
	}
}
