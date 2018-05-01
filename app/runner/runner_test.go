package runner

import (
	"path/filepath"
	"testing"

	"github.com/gulien/orbit/app/context"
)

// Tests if initializing an OrbitRunner throws an error
// with a wrong/broken configuration file or no error with a correct
// configuration file.
func TestNewOrbitRunner(t *testing.T) {
	// case 1: uses a wrong configuration file.
	wrongTemplateFilePath, _ := filepath.Abs("../../_tests/.env")
	ctx, _ := context.NewOrbitContext(wrongTemplateFilePath, "")
	if _, err := NewOrbitRunner(ctx); err == nil {
		t.Error("OrbitRunner should not have been instantiated!")
	}

	// case 2: uses a broken configuration file.
	brokenTemplateFilePath, _ := filepath.Abs("../../_tests/broken-template.yml")
	ctx, _ = context.NewOrbitContext(brokenTemplateFilePath, "")
	if _, err := NewOrbitRunner(ctx); err == nil {
		t.Error("OrbitRunner should not have been instantiated!")
	}

	// case 3 uses a correct configuration file.
	templateFilePath, _ := filepath.Abs("../../_tests/orbit.yml")
	ctx, _ = context.NewOrbitContext(templateFilePath, "")
	if _, err := NewOrbitRunner(ctx); err != nil {
		t.Error("OrbitRunner should have been instantiated!")
	}
}

// A dumb test to improve code coverage.
func TestPrint(t *testing.T) {
	templateFilePath, _ := filepath.Abs("../../_tests/orbit.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "")
	r, _ := NewOrbitRunner(ctx)

	r.Print()
}

// Tests Run function by running different kind of tasks.
func TestRun(t *testing.T) {
	templateFilePath, _ := filepath.Abs("../../_tests/orbit.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "")
	r, _ := NewOrbitRunner(ctx)

	// case 1: uses a non existing task.
	if err := r.Run("discovery"); err == nil {
		t.Error("Task should not exist!")
	}

	// case 2: uses a task which has a non existing command.
	if err := r.Run("challenger"); err == nil {
		t.Error("Task should have failed!")
	}

	// case 3: uses a correct task.
	if err := r.Run("explorer"); err != nil {
		t.Error("Task should have been run!")
	}

	// case 4: uses nested tasks.
	if err := r.Run("explorer", "sputnik"); err != nil {
		t.Error("Nested tasks should have been run!")
	}

	// case 5: uses custom shell with non-existent shell.
	if err := r.Run("zuma"); err == nil {
		t.Error("Custom shell task should have failed!")
	}

	// case 6: uses custom shell with existent shell.
	if err := r.Run("falcon"); err != nil {
		t.Error("Custom shell task should have been run!")
	}

	// case 7: uses custom shell without parameter.
	if err := r.Run("ariane"); err != nil {
		t.Error("Custom shell task should have been run!")
	}

	// case 8: uses a task which calls others tasks
	if err := r.Run("new shepard"); err != nil {
		t.Error("Task calling others tasks should have been run!")
	}

	// case 9: uses a task which calls a non existing task.
	if err := r.Run("new glenn"); err == nil {
		t.Error("Task calling another task should not have been run!")
	}
}
