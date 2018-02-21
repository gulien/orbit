package generator

import (
	"runtime"
	"testing"

	"github.com/gulien/orbit/logger"
)

// A dumb test to improve code coverage.
func TestGetOS(t *testing.T) {
	if getOS() != runtime.GOOS {
		t.Error("Dumb test should have been successful!")
	}
}

// A dumb test to improve code coverage.
func TestIsDebug(t *testing.T) {
	if isDebug() != logger.IsSilent() {
		t.Error("Dumb test should have been successful!")
	}
}
