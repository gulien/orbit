package context

import (
	"path/filepath"
	"testing"
)

// Tests if initializing an OrbitContext with wrong parameters
// throws an errors or no errors if the parameters are OK.
func TestNewOrbitContext(t *testing.T) {
	if _, err := NewOrbitContext("", "", "", ""); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	if _, err := NewOrbitContext("random_file.yml", "", "", ""); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	template, err := filepath.Abs("../.assets/tests/template.yml")
	if err != nil {
		panic(err)
	}

	if _, err := NewOrbitContext(template, "wrong_values_parameter", "", ""); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	if _, err := NewOrbitContext(template, "", "wrong_env_files_parameter", ""); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	if _, err := NewOrbitContext(template, "", "", "wrong_raw_data_parameter"); err == nil {
		t.Error("OrbitContext should not have been instantiated!")
	}

	values, err := filepath.Abs("../.assets/tests/values.yml")
	if err != nil {
		panic(err)
	}

	envFile, err := filepath.Abs("../.assets/tests/.env")
	if err != nil {
		panic(err)
	}

	if _, err := NewOrbitContext(template, values, envFile, "key=data"); err != nil {
		t.Error("OrbitContext should have been instantiated!")
	}
}

// Tests getValuesMap function.
func TestOrbitContext_getValuesMap(t *testing.T) {
	if _, err := getValuesMap("key1,values;key2"); err == nil {
		t.Error("getValuesMap should not have been able to parse \"key1,values;key2\"!")
	}

	if _, err := getValuesMap("values"); err == nil {
		t.Error("getValuesMap should not have been able to find \"values\" as a file!")
	}

	values, err := filepath.Abs("../.assets/tests/wrong_values.yml")
	if err != nil {
		panic(err)
	}

	if _, err := getValuesMap(values); err == nil {
		t.Error("getValuesMap should not have been able to parse \"wrong_values.yml\"!")
	}

	values, err = filepath.Abs("../.assets/tests/values.yml")
	if err != nil {
		panic(err)
	}

	if _, err := getValuesMap(values); err != nil {
		t.Error("getValuesMap should have been able to retrieve values from \"values.yml\"!")
	}
}

// Tests getEnvFilesMap function.
func TestOrbitContext_getEnvFilesMap(t *testing.T) {
	if _, err := getEnvFilesMap("key1,.env;key2"); err == nil {
		t.Error("getEnvFilesMap should not have been able to parse \"key1,.env;key2\"!")
	}

	if _, err := getEnvFilesMap(".wrong_env"); err == nil {
		t.Error("getEnvFilesMap should not have been able to find \".wrong_env\" as a file!")
	}

	envFile, err := filepath.Abs("../.assets/tests/values.yml")
	if err != nil {
		panic(err)
	}

	if _, err := getEnvFilesMap(envFile); err == nil {
		t.Error("getEnvFilesMap should not have been able to parse \"values.yml\"!")
	}

	envFile, err = filepath.Abs("../.assets/tests/.env")
	if err != nil {
		panic(err)
	}

	if _, err := getEnvFilesMap(envFile); err != nil {
		t.Error("getEnvFilesMap should have been able to retrieve values from \".env\"!")
	}
}

// Tests getFilesMap function.
func TestOrbitContext_getFilesMap(t *testing.T) {
	if _, err := getFilesMap("key1,file;key2"); err == nil {
		t.Error("getFilesMap should not have been able to parse \"key1,file;key2\"!")
	}

	filesMap, err := getFilesMap("test.yml")
	if err != nil {
		panic(err)
	}

	if len(filesMap) != 1 {
		t.Error("Files map should have a length of one!")
	}

	if filesMap[0].Name != "default" {
		t.Error("Item of files map at index 0 should have \"default\" as name!")
	}

	if filesMap[0].Path != "test.yml" {
		t.Error("Item of files map at index 0 should have \"test.yml\" as path!")
	}

	filesMap, err = getFilesMap("first,first.yml;last,last.yml")
	if err != nil {
		panic(err)
	}

	if len(filesMap) != 2 {
		t.Error("Files map should have a length of two!")
	}

	if filesMap[0].Name != "first" {
		t.Error("Item of files map at index 0 should have \"first\" as name!")
	}

	if filesMap[0].Path != "first.yml" {
		t.Error("Item of files map at index 0 should have \"first.yml\" as path!")
	}

	if filesMap[1].Name != "last" {
		t.Error("Item of files map at index 1 should have \"last\" as name!")
	}

	if filesMap[1].Path != "last.yml" {
		t.Error("Item of files map at index 0 should have \"last.yml\" as path!")
	}
}

// Tests getRawDataMap function.
func TestOrbitContext_getRawDataMap(t *testing.T) {
	if _, err := getRawDataMap("key1=data;key2"); err == nil {
		t.Error("getRawDataMap should not have been able to parse \"key1=data;key2\"!")
	}

	rawDataMap, err := getRawDataMap("key=data")
	if err != nil {
		panic(err)
	}

	if len(rawDataMap) != 1 {
		t.Error("Raw data map should have a length of one!")
	}

	if rawDataMap["key"] != "data" {
		t.Error("Raw data at key \"key\" should have \"data\" as value!")
	}

	rawDataMap, err = getRawDataMap("key1=data1;key2=data2")
	if err != nil {
		panic(err)
	}

	if len(rawDataMap) != 2 {
		t.Error("Raw data map should have a length of two!")
	}

	if rawDataMap["key1"] != "data1" {
		t.Error("Raw data at key \"key1\" should have \"data1\" as value!")
	}

	if rawDataMap["key2"] != "data2" {
		t.Error("Raw data at key \"key2\" should have \"data2\" as value!")
	}
}
