package helpers

import "os"

// FileExist returns true if the specified path exists.
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
