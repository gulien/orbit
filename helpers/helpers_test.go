package helpers

import (
	"path/filepath"
	"testing"
)

// Tests the function FileExists.
func TestFileExists(t *testing.T) {
	file, _ := filepath.Abs("../.assets/tests/orbit.yml")
	if !FileExists(file) {
		t.Error("File should exist!")
	}
}
