package runner

import (
	"path/filepath"
	"testing"

	"github.com/gulien/orbit/context"
)

// Tests if initializing an OrbitRunner with a wrong configuration file
// throws an errors.
func TestNewOrbitRunner(t *testing.T) {
	config, err := filepath.Abs("../.assets/tests/wrong_template.yml")
	if err != nil {
		panic(err)
	}

	ctx, err := context.NewOrbitContext(config, "", "", "")
	if err != nil {
		panic(err)
	}

	if _, err := NewOrbitRunner(ctx); err == nil {
		t.Error("OrbitRunner should not have been instantiated!")
	}

	config, err = filepath.Abs("../.assets/tests/.env")
	if err != nil {
		panic(err)
	}

	ctx, err = context.NewOrbitContext(config, "", "", "")
	if err != nil {
		panic(err)
	}

	if _, err := NewOrbitRunner(ctx); err == nil {
		t.Error("OrbitRunner should not have been instantiated!")
	}
}

// Tests Exec function.
func TestOrbitRunner_Exec(t *testing.T) {
	config, err := filepath.Abs("../.assets/tests/orbit.yml")
	if err != nil {
		panic(err)
	}

	ctx, err := context.NewOrbitContext(config, "", "", "")
	if err != nil {
		panic(err)
	}

	r, err := NewOrbitRunner(ctx)
	if err != nil {
		panic(err)
	}

	// tests if calling an unknown Orbit command throws an errors.
	if err := r.Exec("discovery"); err == nil {
		t.Error("Orbit command should not exist!")
	}

	// tests if calling an Orbit command with a non-existing external command
	// throws an errors.
	if err := r.Exec("challenger"); err == nil {
		t.Error("Orbit command should have failed!")
	}

	// tests a simple exec.
	if err := r.Exec("explorer"); err != nil {
		t.Error("Orbit command should have been executed!")
	}

	// tests a nested exec.
	if err := r.Exec("explorer", "sputnik"); err != nil {
		t.Error("Nested Orbit commands should been executed!")
	}
}
