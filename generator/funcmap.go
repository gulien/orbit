package generator

import (
	"runtime"

	"github.com/gulien/orbit/logger"
)

/*
getOS returns the OS name at runtime.

This function is available in
a data-driven template by using "os".
*/
func getOS() string {
	return runtime.GOOS
}

/*
getOS returns the OS name at runtime.

This function is available in
a data-driven template by using "debug".
*/
func isDebug() bool {
	return !logger.IsSilent()
}
