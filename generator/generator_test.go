package generator

import (
	"os"
	"reflect"
	"testing"

	"github.com/gulien/orbit/context"
	"github.com/gulien/orbit/helpers"
)

var (
	// defaultGenerator is an instance of OrbitGenerator used in this test suite which contains one values file and one
	// .env file.
	defaultGenerator *OrbitGenerator

	// manyGenerator is an instance of OrbitGenerator used in this test suite which contains two values files and two
	// .env files.
	manyGenerator *OrbitGenerator

	// expectedResult contains the data from the file "expected_result.yml"
	expectedResult interface{}
)

// init instantiates expectedResult plus the OrbitGenerator defaultGenerator and manyGenerator.
func init() {
	// retrieves the data from the file "expected_result.yml".
	expectedResultPath := helpers.Abs("../.assets/tests/expected_result.yml")
	expectedResult = helpers.ReadYAML(expectedResultPath)

	// loads assets.
	defaultTmpl := helpers.Abs("../.assets/tests/template.yml")
	manyTmpl := helpers.Abs("../.assets/tests/template_many.yml")
	values := helpers.Abs("../.assets/tests/values.yml")
	envFile := helpers.Abs("../.assets/tests/.env")
	rawData := "author=Julien Neuhart;comment=A simple file for testing purpose"

	// last but not least, creates our OrbitGenerator instances.
	ctx, err := context.NewOrbitContext(defaultTmpl, values, envFile, rawData)
	if err != nil {
		panic(err)
	}

	defaultGenerator = NewOrbitGenerator(ctx)

	ctx, err = context.NewOrbitContext(manyTmpl, "ru,"+values+";usa,"+values, "ru,"+envFile+";usa,"+envFile, rawData)
	if err != nil {
		panic(err)
	}

	manyGenerator = NewOrbitGenerator(ctx)
}

/*
Tests to parse the template "template.yml" and generate a resulting file "result.yml".

Expects the file "result.yml" to be the same as "expected_result.yml".
*/
func TestDefaultTemplate(t *testing.T) {
	data, err := defaultGenerator.Parse()
	if err != nil {
		t.Error("Failed to parse the default template!")
	}

	if err := defaultGenerator.WriteOutputFile("result.yml", data); err != nil {
		t.Error("Failed to write the outpout file from the default template!")
	}

	result := helpers.ReadYAML("result.yml")
	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Result file from the default template should be equal to the expected result!")
	}
}

/*
Tests to parse the template "template_many.yml" and generate a resulting file "result.yml".

Expects the file "result.yml" to be the same as "expected_result.yml".
*/
func TestManyTemplate(t *testing.T) {
	data, err := manyGenerator.Parse()
	if err != nil {
		t.Error("Failed to parse the many template!")
	}

	if err := manyGenerator.WriteOutputFile("result.yml", data); err != nil {
		t.Error("Failed to write the outpout file from the many template!")
	}

	result := helpers.ReadYAML("result.yml")
	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Result file from the many template should be equal to the expected result!")
	}
}
