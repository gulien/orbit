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
func TesIsVerbose(t *testing.T) {
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
