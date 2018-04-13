package helpers

import (
	"path/filepath"
	"testing"
)

// Tests FileExists function.
func TestFileExists(t *testing.T) {
	file, _ := filepath.Abs("../../_tests/template.yml")
	if !FileExists(file) {
		t.Error("File should exist!")
	}
}
