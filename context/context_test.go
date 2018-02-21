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

	// case 4: uses also a payload.
	if _, err := NewOrbitContext(templateFilePath, "key,value"); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}
}
