package context

import (
	"path/filepath"
	"testing"
)

// Tests if initializing an OrbitContext throws an error
// with wrong parameters or no error if the parameters are OK.
func TestNewOrbitContext(t *testing.T) {
	// case 1: uses an empty template file path.
	if _, err := NewOrbitContext("", ""); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 2: uses a non existing template file path.
	if _, err := NewOrbitContext("non_existing_file", ""); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 3: uses an existing template file path.
	templateFilePath, _ := filepath.Abs("../_tests/template.yml")
	if _, err := NewOrbitContext(templateFilePath, ""); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 4: uses a broken payload file.
	brokenPayloadFilePath, _ := filepath.Abs("../_tests/broken-orbit-payload.yml")
	if _, err := NewOrbitContext(templateFilePath, "key,"+brokenPayloadFilePath); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 5: uses a correct payload file.
	payloadFilePath, _ := filepath.Abs("../_tests/orbit-payload.yml")
	if _, err := NewOrbitContext(templateFilePath, "key,"+payloadFilePath); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 6: uses a broken payload string.
	if _, err := NewOrbitContext(templateFilePath, "key"); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 7: uses a correct payload string.
	if _, err := NewOrbitContext(templateFilePath, "key,value"); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 8: uses a broken payload entry.
	brokenPayloadEntryFilePath, _ := filepath.Abs("../_tests/broken-data-source.yml")
	if _, err := NewOrbitContext(templateFilePath, "key,"+brokenPayloadEntryFilePath); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 8: uses a correct payload entry.
	payloadEntryFilePath, _ := filepath.Abs("../_tests/data-source.yml")
	if _, err := NewOrbitContext(templateFilePath, "key,"+payloadEntryFilePath); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}
}
