package runner

import (
	"path/filepath"
	"testing"

	"github.com/gulien/orbit/context"
)

// testRunner is the instance of OrbitRunner used in this test suite.
var testRunner *OrbitRunner

// init instantiates the OrbitRunner testRunner.
func init() {
	configFilePath, err := filepath.Abs("../.assets/tests/orbit.yml")
	if err != nil {
		panic(err)
	}

	ctx, err := context.NewOrbitContext(configFilePath, "", "")
	if err != nil {
		panic(err)
	}

	r, err := NewOrbitRunner(ctx)
	if err != nil {
		panic(err)
	}

	testRunner = r
}

// Tests if Orbit command "glide_x" does not exist.
func TestOrbitCommandDoesNotExist(t *testing.T) {
	t.Log("Tests if Orbit command \"glide_x\" does not exist...")

	if err := testRunner.Exec("glide_x"); err == nil {
		t.Error("\"glide_x\" should not exist!")
	}
}

// Tests Orbit command "glide_1".
func TestOrbitCommand(t *testing.T) {
	t.Log("Tests Orbit command \"glide_1\"...")

	if err := testRunner.Exec("glide_1"); err != nil {
		t.Error("\"glide_1\" should have been executed!")
	}
}

// Tests nested Orbit commands "glide_1" and "glide_2".
func TestNestedOrbitCommands(t *testing.T) {
	t.Log("Tests nested Orbit commands \"glide_1\" and \"glide_2\"...")

	if err := testRunner.Exec("glide_1", "glide_2"); err != nil {
		t.Error("\"glide_1\" and \"glide_2\" should have been executed!")
	}
}
