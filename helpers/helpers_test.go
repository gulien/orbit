package helpers

import (
	"path/filepath"
	"testing"
)

// Tests if file "orbit.yml" exists.
func TestFileExists(t *testing.T) {
	t.Log("Tests if file \"orbit.yml\" exists...")

	file, _ := filepath.Abs("../.assets/tests/orbit.yml")
	if !FileExists(file) {
		t.Error("\"orbit.yml\" should exist!")
	}
}

// Tests if file "foo.yml" does not exist.
func TestFileDoesNotExist(t *testing.T) {
	t.Log("Tests if file \"orbit.yml\" does not exist...")

	if FileExists("foo.yml") {
		t.Error("\"foo.yml\" should not exist!")
	}
}
