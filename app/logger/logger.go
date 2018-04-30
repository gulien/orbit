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
	// logger is an instance of logrus logger.
	logger *logrus.Logger
}

// newOrbitLogged creates an instance of orbitLogger.
func newOrbitLogger() *orbitLogger {
	l := logrus.New()
	l.Out = os.Stdout
	l.Level = logrus.ErrorLevel

	return &orbitLogger{
		logger: l,
	}
}

// houston is the logger instance used by the application.
var houston = newOrbitLogger()

// SetLevel updates the level of messages which will be logged.
func SetLevel(level logrus.Level) {
	houston.logger.SetLevel(level)
}

// GetLevel returns the current level of messages which are logged.
func GetLevel() logrus.Level {
	return houston.logger.Level
}

// Infof logs information using the Houston logger.
func Infof(message string, args ...interface{}) {
	houston.logger.Infof(message, args...)
}

// Debugf logs debug information using the Houston logger.
func Debugf(message string, args ...interface{}) {
	houston.logger.Debugf(message, args...)
}

// Error logs error information using the Houston logger.
func Error(err error) {
	if _, ok := err.(*OrbitError.OrbitError); ok {
		houston.logger.Error(err.Error())
	} else if GetLevel() == logrus.DebugLevel {
		// errors which are not "OrbitError" are not relevant unless we are
		// in debug mode.
		houston.logger.Error(err.Error())
	}
}
