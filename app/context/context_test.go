package context

import (
	"path/filepath"
	"testing"
)

// Tests if initializing an OrbitContext throws an error
// with wrong parameters or no error if the parameters are OK.
func TestNewOrbitContext(t *testing.T) {
	// case 1: uses an empty template file path.
	if _, err := NewOrbitContext("", "", "", nil); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 2: uses a non existing template file path.
	if _, err := NewOrbitContext("non_existing_file", "", "", nil); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 3: uses an existing template file path.
	templateFilePath, _ := filepath.Abs("../../_tests/template.yml")
	if _, err := NewOrbitContext(templateFilePath, "", "", nil); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 4: uses a broken payload string.
	if _, err := NewOrbitContext(templateFilePath, "key", "", nil); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 5: uses a correct payload string.
	if _, err := NewOrbitContext(templateFilePath, "key,value", "", nil); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 6: uses a broken payload entry.
	brokenPayloadEntryFilePath, _ := filepath.Abs("../../_tests/broken-data-source.yml")
	if _, err := NewOrbitContext(templateFilePath, "key,"+brokenPayloadEntryFilePath, "", nil); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 7: uses a correct payload entry.
	payloadEntryFilePath, _ := filepath.Abs("../../_tests/data-source.yml")
	if _, err := NewOrbitContext(templateFilePath, "key,"+payloadEntryFilePath, "", nil); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 8: uses a nil template delimiter set
	if _, err := NewOrbitContext(templateFilePath, "", "", nil); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 9: uses a valid, two-element delimiter set
	templateDelimiters := []string{"a", "b"}
	if _, err := NewOrbitContext(templateFilePath, "", "", templateDelimiters); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}

	// case 10: uses an invalid, empty delimiter set
	templateDelimiters = []string{}
	if _, err := NewOrbitContext(templateFilePath, "", "", templateDelimiters); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 11: uses an invalid, one-element delimiter set
	templateDelimiters = []string{"a"}
	if _, err := NewOrbitContext(templateFilePath, "", "", templateDelimiters); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	// case 12: uses an invalid, three-element delimiter set
	templateDelimiters = []string{"a", "b", "c"}
	if _, err := NewOrbitContext(templateFilePath, "", "", templateDelimiters); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}
}
