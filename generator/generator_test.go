package generator

import (
	"path/filepath"

	"github.com/gulien/orbit/context"
	"testing"
)

var (
	// ctxDefault is an instance of OrbitContext used in this test suite which contains one values file and one
	// .env file.
	ctxDefault *context.OrbitContext

	// ctxMany is an instance of OrbitContext used in this test suite which contains two values files and two
	// .env files.
	ctxMany *context.OrbitContext
)

// init instantiates the OrbitContext ctxDefault and ctxMany.
func init() {
	defaultTemplateFilePath := getAbsPath("../.tests/template.txt")
	manyTemplateFilePath := getAbsPath("../.tests/template_many.txt")
	valuesRu := getAbsPath("../.tests/values_ru.yml")
	valuesUsa := getAbsPath("../.tests/values_usa.yml")
	envFileRu := getAbsPath("../.tests/.env_ru")
	envFileUsa := getAbsPath("../.tests/.env_usa")

	c, err := context.NewOrbitContext(defaultTemplateFilePath, valuesUsa, envFileRu)
	if err != nil {
		panic(err)
	}

	cMany, err := context.NewOrbitContext(manyTemplateFilePath, "ru," + valuesRu + ";usa," + valuesUsa, "ru," + envFileRu + ";usa," + envFileUsa)
	if err != nil {
		panic(err)
	}

	ctxDefault = c
	ctxMany = cMany
}

// getAbsPath is a simple wrapper which returns the absolute path of a given relative path.
func getAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return absPath
}

// Tests to parse the template "template.txt" and generate a resulting file "result.txt".
func TestDefaultTemplate(t *testing.T) {

}