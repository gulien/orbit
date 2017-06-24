package runner

import (
	"testing"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/helpers"
)

// testRunner is the instance of OrbitRunner used in this test suite.
var testRunner *OrbitRunner

// init instantiates the OrbitRunner testRunner.
func init() {
	config := helpers.Abs("../.assets/tests/orbit.yml")

	ctx, err := context.NewOrbitContext(config, "", "")
	if err != nil {
		panic(err)
	}

	testRunner, err = NewOrbitRunner(ctx)
	if err != nil {
		panic(err)
	}
}

/*
Tests if calling an unknown Orbit command throws an error.

Expects an error.
*/
func TestNotFound(t *testing.T) {
	if err := testRunner.Exec("discovery"); err == nil {
		t.Error("Orbit command should not exist!")
	}
}

/*
Tests a simple run.

Expects no error.
*/
func TestRun(t *testing.T) {
	if err := testRunner.Exec("explorer"); err != nil {
		t.Error("Orbit command should have run!")
	}
}

/*
Tests a nested run.

Expects no error.
*/
func TestNestedRun(t *testing.T) {
	if err := testRunner.Exec("explorer", "sputnik"); err != nil {
		t.Error("Nested Orbit commands should have run!")
	}
}
