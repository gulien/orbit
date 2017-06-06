package helpers

import (
	"os"
	"path"
)

// FileExist function returns true if the specified path exists.
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileDoesNotExist function returns true is the specified path does not exist.
func FileDoesNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// IsYAML function returns true is the specified file name has .yml or .yaml extension.
func IsYAML(fileName string) bool {
	fileExt := path.Ext(fileName)
	return fileExt == ".yml" || fileExt == ".yaml"
}
