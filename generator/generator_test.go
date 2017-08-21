package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gulien/orbit/context"

	"gopkg.in/yaml.v2"
)

// Tests Parse function.
func TestOrbitGenerator_Parse(t *testing.T) {
	template, err := filepath.Abs("../.assets/tests/wrong_template.yml")
	if err != nil {
		panic(err)
	}

	ctx, err := context.NewOrbitContext(template, "", "", "")
	if err != nil {
		panic(err)
	}

	g := NewOrbitGenerator(ctx)

	if _, err := g.Parse(); err == nil {
		t.Error("OrbitGenerator should not have been able to parse the template \"wrong_template.yml\"!")
	}
}

// Tests WriteOutputFile function.
func TestOrbitGenerator_WriteOutputFile(t *testing.T) {
	expectedResultPath, err := filepath.Abs("../.assets/tests/expected_result.yml")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(expectedResultPath)
	if err != nil {
		panic(err)
	}

	var expectedResult interface{}
	if err := yaml.Unmarshal(data, &expectedResult); err != nil {
		panic(err)
	}

	defaultTmpl, err := filepath.Abs("../.assets/tests/template.yml")
	if err != nil {
		panic(err)
	}

	values, err := filepath.Abs("../.assets/tests/values.yml")
	if err != nil {
		panic(err)
	}

	envFile, err := filepath.Abs("../.assets/tests/.env")
	if err != nil {
		panic(err)
	}

	rawData := "author=Julien Neuhart;comment=A simple file for testing purpose"

	ctx, err := context.NewOrbitContext(defaultTmpl, values, envFile, rawData)
	if err != nil {
		panic(err)
	}

	g := NewOrbitGenerator(ctx)

	dataDefaultTmpl, err := g.Parse()
	if err != nil {
		t.Error("Failed to parse the default template!")
	}

	if err := g.WriteOutputFile("result.yml", dataDefaultTmpl); err != nil {
		t.Error("Failed to write the outpout file from the default template!")
	}

	dataDefaultResult, err := ioutil.ReadFile("result.yml")
	if err != nil {
		panic(err)
	}

	var defaultResult interface{}
	if err := yaml.Unmarshal(dataDefaultResult, &defaultResult); err != nil {
		panic(err)
	}

	os.Remove("result.yml")

	if !reflect.DeepEqual(defaultResult, expectedResult) {
		t.Error("Result file from the default template should be equal to the expected result!")
	}

	manyTmpl, err := filepath.Abs("../.assets/tests/template_many.yml")
	if err != nil {
		panic(err)
	}

	ctx, err = context.NewOrbitContext(manyTmpl, "ru,"+values+";usa,"+values, "ru,"+envFile+";usa,"+envFile, rawData)
	if err != nil {
		panic(err)
	}

	g = NewOrbitGenerator(ctx)

	dataManyTmpl, err := g.Parse()
	if err != nil {
		t.Error("Failed to parse the many template!")
	}

	if err := g.WriteOutputFile("result.yml", dataManyTmpl); err != nil {
		t.Error("Failed to write the outpout file from the many template!")
	}

	dataManyResult, err := ioutil.ReadFile("result.yml")
	if err != nil {
		panic(err)
	}

	var manyResult interface{}
	if err := yaml.Unmarshal(dataManyResult, &manyResult); err != nil {
		panic(err)
	}

	os.Remove("result.yml")

	if !reflect.DeepEqual(manyResult, expectedResult) {
		t.Error("Result file from the many template should be equal to the expected result!")
	}

	if err := g.WriteOutputFile("/...", dataManyTmpl); err == nil {
		t.Error("WriteOutputFile should not been able to to write the outpout file \"/...\"!")
	}
}
