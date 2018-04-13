package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gulien/orbit/app/context"

	"gopkg.in/yaml.v2"
)

// Tests if executing a data-driven template throws an error if it's broken
// or no error if it's correct.
func TestExecute(t *testing.T) {
	dataSourceFilePath, _ := filepath.Abs("../../_tests/data-source.yml")

	// case 1: uses a broken data-driven template.
	brokenTemplateFilePath, _ := filepath.Abs("../../_tests/broken-template.yml")
	ctx, _ := context.NewOrbitContext(brokenTemplateFilePath, "Values,"+dataSourceFilePath)
	g := NewOrbitGenerator(ctx)
	if _, err := g.Execute(); err == nil {
		t.Errorf("OrbitGenerator should not have been able to parse the data-driven template %s", brokenTemplateFilePath)
	}

	// case 2: uses a correct data-driven template.
	templateFilePath, _ := filepath.Abs("../../_tests/template.yml")
	ctx, _ = context.NewOrbitContext(templateFilePath, "Values,"+dataSourceFilePath)
	g = NewOrbitGenerator(ctx)
	if _, err := g.Execute(); err != nil {
		t.Errorf("OrbitGenerator should have been able to parse the data-driven template %s", templateFilePath)
	}
}

// Tests if flushing from raw data source works as expected.
func TestFlushFromRawDataSource(t *testing.T) {
	templateFilePath, _ := filepath.Abs("../../_tests/template-raw.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "SPACEX_LAUNCHERS,Falcon 9, Falcon Heavy;BLUE_ORIGIN_LAUNCHERS,New Shepard, New Glenn;ESA_LAUNCHERS,Ariane 5, Vega")
	g := NewOrbitGenerator(ctx)
	data, _ := g.Execute()

	// case 1: uses an empty output path.
	if err := g.Flush("", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to Stdout!")
	}

	// case 2: uses a broken output path.
	if err := g.Flush("/.../...", data); err == nil {
		t.Error("OrbitGenerator should not have been able to flush to result file /...!")
	}

	// case 3: uses a correct output path.
	if err := g.Flush("result.yml", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to result file result.yml!")
	}

	expectedResultFilePath, _ := filepath.Abs("../../_tests/expected-result-raw-env.yml")
	expectedResultFileData, _ := ioutil.ReadFile(expectedResultFilePath)

	var expectedResult interface{}
	yaml.Unmarshal(expectedResultFileData, &expectedResult)

	resultFileData, _ := ioutil.ReadFile("result.yml")

	var result interface{}
	yaml.Unmarshal(resultFileData, &result)

	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("result.yml should be equal to expected-result-raw-env.yml!")
	}
}

// Tests if flushing from YAML data source works as expected.
func TestFlushFromYAMLDataSource(t *testing.T) {
	dataSourceFilePath, _ := filepath.Abs("../../_tests/data-source.yml")
	templateFilePath, _ := filepath.Abs("../../_tests/template.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "Values,"+dataSourceFilePath)
	g := NewOrbitGenerator(ctx)
	data, _ := g.Execute()

	// case 1: uses an empty output path.
	if err := g.Flush("", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to Stdout!")
	}

	// case 2: uses a broken output path.
	if err := g.Flush("/.../...", data); err == nil {
		t.Error("OrbitGenerator should not have been able to flush to result file /...!")
	}

	// case 3: uses a correct output path.
	if err := g.Flush("result.yml", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to result file result.yml!")
	}

	expectedResultFilePath, _ := filepath.Abs("../../_tests/data-source.yml")
	expectedResultFileData, _ := ioutil.ReadFile(expectedResultFilePath)

	var expectedResult interface{}
	yaml.Unmarshal(expectedResultFileData, &expectedResult)

	resultFileData, _ := ioutil.ReadFile("result.yml")

	var result interface{}
	yaml.Unmarshal(resultFileData, &result)

	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("result.yml should be equal to data-source.yml!")
	}
}

// Tests if flushing from TOML data source works as expected.
func TestFlushFromTOMLDataSource(t *testing.T) {
	dataSourceFilePath, _ := filepath.Abs("../../_tests/data-source.toml")
	templateFilePath, _ := filepath.Abs("../../_tests/template.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "Values,"+dataSourceFilePath)
	g := NewOrbitGenerator(ctx)
	data, _ := g.Execute()

	// case 1: uses an empty output path.
	if err := g.Flush("", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to Stdout!")
	}

	// case 2: uses a broken output path.
	if err := g.Flush("/.../...", data); err == nil {
		t.Error("OrbitGenerator should not have been able to flush to result file /...!")
	}

	// case 3: uses a correct output path.
	if err := g.Flush("result.yml", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to result file result.yml!")
	}

	expectedResultFilePath, _ := filepath.Abs("../../_tests/data-source.yml")
	expectedResultFileData, _ := ioutil.ReadFile(expectedResultFilePath)

	var expectedResult interface{}
	yaml.Unmarshal(expectedResultFileData, &expectedResult)

	resultFileData, _ := ioutil.ReadFile("result.yml")

	var result interface{}
	yaml.Unmarshal(resultFileData, &result)

	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("result.yml should be equal to data-source.yml!")
	}
}

// Tests if flushing from JSON data source works as expected.
func TestFlushFromJSONDataSource(t *testing.T) {
	dataSourceFilePath, _ := filepath.Abs("../../_tests/data-source.json")
	templateFilePath, _ := filepath.Abs("../../_tests/template.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "Values,"+dataSourceFilePath)
	g := NewOrbitGenerator(ctx)
	data, _ := g.Execute()

	// case 1: uses an empty output path.
	if err := g.Flush("", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to Stdout!")
	}

	// case 2: uses a broken output path.
	if err := g.Flush("/.../...", data); err == nil {
		t.Error("OrbitGenerator should not have been able to flush to result file /...!")
	}

	// case 3: uses a correct output path.
	if err := g.Flush("result.yml", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to result file result.yml!")
	}

	expectedResultFilePath, _ := filepath.Abs("../../_tests/data-source.yml")
	expectedResultFileData, _ := ioutil.ReadFile(expectedResultFilePath)

	var expectedResult interface{}
	yaml.Unmarshal(expectedResultFileData, &expectedResult)

	resultFileData, _ := ioutil.ReadFile("result.yml")

	var result interface{}
	yaml.Unmarshal(resultFileData, &result)

	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("result.yml should be equal to data-source.yml!")
	}
}

// Tests if flushing from .env data source works as expected.
func TestFlushFromEnvFileDataSource(t *testing.T) {
	dataSourceFilePath, _ := filepath.Abs("../../_tests/.env")
	templateFilePath, _ := filepath.Abs("../../_tests/template-env.yml")
	ctx, _ := context.NewOrbitContext(templateFilePath, "Values,"+dataSourceFilePath)
	g := NewOrbitGenerator(ctx)
	data, _ := g.Execute()

	// case 1: uses an empty output path.
	if err := g.Flush("", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to Stdout!")
	}

	// case 2: uses a broken output path.
	if err := g.Flush("/.../...", data); err == nil {
		t.Error("OrbitGenerator should not have been able to flush to result file /...!")
	}

	// case 3: uses a correct output path.
	if err := g.Flush("result.yml", data); err != nil {
		t.Error("OrbitGenerator should have been able to flush to result file result.yml!")
	}

	expectedResultFilePath, _ := filepath.Abs("../../_tests/expected-result-raw-env.yml")
	expectedResultFileData, _ := ioutil.ReadFile(expectedResultFilePath)

	var expectedResult interface{}
	yaml.Unmarshal(expectedResultFileData, &expectedResult)

	resultFileData, _ := ioutil.ReadFile("result.yml")

	var result interface{}
	yaml.Unmarshal(resultFileData, &result)

	os.Remove("result.yml")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("result.yml should be equal to expected-result-raw-env.yml!")
	}
}
