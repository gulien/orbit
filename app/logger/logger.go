/*
Package logger implements a simple helper for displaying output to the user.
*/
package logger

import (
	"os"

	OrbitError "github.com/gulien/orbit/app/error"

	"github.com/sirupsen/logrus"
)

// orbitLogger provides the underlying implementation that displays output to the user.
type orbitLogger struct {
	// log is an instance of logrus logger.
	log *logrus.Logger

	// silent disables logging if true.
	silent bool
}

// newOrbitLogged creates an instance of orbitLogger.
func newOrbitLogger() *orbitLogger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel

	return &orbitLogger{
		log:    log,
		silent: false,
	}
}

// Houston is the logger instance used by the application.
var Houston = newOrbitLogger()

// Mute disables logging.
func Mute() {
	Houston.silent = true
}

// IsSilent returns true if logging is disabled.
func IsSilent() bool {
	return Houston.silent
}

// Debugf logs debug information using the Houston logger.
func Debugf(message string, args ...interface{}) {
	if !Houston.silent {
		Houston.log.Debugf(message, args...)
	}
}

// Error logs error information using the Houston logger.
func Error(err error) {
	if !Houston.silent {
		Houston.log.Error(err.Error())
	}
}

// NotifyOrbitError prints an error which have to be displayed to the user.
func NotifyOrbitError(err *OrbitError.OrbitError) {
	Houston.log.Error(err.Error())
}
