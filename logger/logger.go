/*
Package logger implements a simple helper for displaying output to users.
*/
package logger

import (
	"os"

	"github.com/gulien/orbit/errors"

	"github.com/sirupsen/logrus"
)

// OrbitLogger provides the underlying implementation that displays output to users.
type OrbitLogger struct {
	// Log is an instance of logrus logger.
	Log *logrus.Logger

	// Silent disables logging if true.
	Silent bool
}

// newOrbitLogged creates an instance of OrbitLogger.
func newOrbitLogger() *OrbitLogger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel

	return &OrbitLogger{
		Log:    log,
		Silent: false,
	}
}

// Houston is the OrbitLogger instance used by the application.
var Houston = newOrbitLogger()

// Mute disables logging.
func Mute() {
	Houston.Silent = true
}

// Debugf logs information using the Houston logger.
func Debugf(message string, args ...interface{}) {
	if !Houston.Silent {
		Houston.Log.Debugf(message, args...)
	}
}

// Error logs error information using the Houston logger.
func Error(err error) {
	if !Houston.Silent {
		Houston.Log.Error(err.Error())
	}
}

// NotifyOrbitError prints an error which have to be displayed to the user.
func NotifyOrbitError(err *errors.OrbitError) {
	Houston.Log.Error(err.Error())
}
