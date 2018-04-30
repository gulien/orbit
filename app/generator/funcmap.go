package generator

import (
	"runtime"

	"github.com/gulien/orbit/app/logger"

	"github.com/sirupsen/logrus"
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
isVerbose returns true if the logs are set to info level.

This function is available in
a data-driven template by using "verbose".
*/
func isVerbose() bool {
	return logger.GetLevel() == logrus.InfoLevel
}

/*
isDebug returns true if the logs are set to debug level.

This function is available in
a data-driven template by using "debug".
*/
func isDebug() bool {
	return logger.GetLevel() == logrus.DebugLevel
}
