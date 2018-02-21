package runner

import (
	"path/filepath"
	"testing"

	"github.com/gulien/orbit/context"
)

// Tests if initializing an OrbitRunner throws an error
// with a wrong/broken configuration file or no error with a correct
// configuration file.
func TestNewOrbitRunner(t *testing.T) {
	// case 1: uses a wrong configuration file.
	wrongTemplateFilePath, _ := filepath.Abs("../_tests/.env")
	ctx, _ := context.NewOrbitContext(wrongTemplateFilePath, "")
	if _, err := NewOrbitRunner(ctx); err == nil {
		t.Error("OrbitRunner should not have been instantiated!")
	}

	// case 2: uses a broken configuration file.
	brokenTemplateFilePath, _ := filepath.Abs("../_tests/broken-template.yml")
	ctx, _ = context.NewOrbitContext(brokenTemplateFilePath, "")
	if _, err := NewOrbitRunner(ctx); err == nil {
		t.Error("OrbitRunner should not have been instantiated!")
	}

	// case 2: uses a correct configuration file.
	templateFilePath, _ := filepath.Abs("../_tests/orbit.yml")
	ctx, _ = context.NewOrbitContext(templateFilePath, "")
	if _, err := NewOrbitRunner(ctx); err != nil {
		t.Error("OrbitRunner should have been instantiated!")
	}
}

// A dumb test to improve code coverage.
func TestPrint(t *testing.T) {
	templateFilePath, _ := filepath.Abs("../_tests/orbit.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "")
	r, _ := NewOrbitRunner(ctx)

	r.Print()
}

// Tests Exec function by running different kind of commands.
func TestExec(t *testing.T) {
	templateFilePath, _ := filepath.Abs("../_tests/orbit.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "")
	r, _ := NewOrbitRunner(ctx)

	// case 1: uses a non existing Orbit command.
	if err := r.Exec("discovery"); err == nil {
		t.Error("Orbit command should not exist!")
	}

	// case 2: uses an Orbit command which has a non existing
	// external command.
	if err := r.Exec("challenger"); err == nil {
		t.Error("Orbit command should have failed!")
	}

	// case 3: uses a correct Orbit command.
	if err := r.Exec("explorer"); err != nil {
		t.Error("Orbit command should have been executed!")
	}

	// case 4: uses nested Orbit commands.
	if err := r.Exec("explorer", "sputnik"); err != nil {
		t.Error("Nested Orbit commands should have been executed!")
	}
}
