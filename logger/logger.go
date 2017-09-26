/*
Package logger implements a simple helper for displaying output to the user.
*/
package logger

import (
	"os"

	"github.com/gulien/orbit/errors"

	"github.com/sirupsen/logrus"
)

// OrbitLogger provides the underlying implementation that displays output to the user.
type OrbitLogger struct {
	// log is an instance of logrus logger.
	log *logrus.Logger

	// silent disables logging if true.
	silent bool
}

// newOrbitLogged creates an instance of OrbitLogger.
func newOrbitLogger() *OrbitLogger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel

	return &OrbitLogger{
		log:    log,
		silent: false,
	}
}

// Houston is the OrbitLogger instance used by the application.
var Houston = newOrbitLogger()

// Mute disables logging.
func Mute() {
	Houston.silent = true
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
func NotifyOrbitError(err *errors.OrbitError) {
	Houston.log.Error(err.Error())
}
