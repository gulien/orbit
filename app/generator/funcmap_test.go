package generator

import (
	"runtime"
	"testing"
)

// A dumb test to improve code coverage.
func TestGetOS(t *testing.T) {
	if getOS() != runtime.GOOS {
		t.Error("Dumb test should have been successful!")
	}
}

// A dumb test to improve code coverage.
func TestIsVerbose(t *testing.T) {
	if isVerbose() != false {
		t.Error("Dumb test should have been successful!")
	}
}

// A dumb test to improve code coverage.
func TestIsDebug(t *testing.T) {
	if isDebug() != false {
		t.Error("Dumb test should have been successful!")
	}
}

// Tests run function to check if it returns a well-formed string.
func TestRun(t *testing.T) {
	if run("explorer", "falcon") != "run@explorer,falcon" {
		t.Error("String returned by run function is malformated!")
	}
}
