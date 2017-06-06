package notifier

import (
	"fmt"
	"os"

	jww "github.com/spf13/jwalterweatherman"
)

// Start function sets the stdout verbosity to info level.
func Start() {
	jww.SetStdoutThreshold(jww.LevelInfo)
}

// Errorf function prints an error message in a given format.
func Errorf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	Error(message)
}

// Error function prints an error message.
func Error(message interface{}) {
	jww.ERROR.Println(message)
	// always exit 1 on error.
	os.Exit(1)
}

// Infof function prints an info message in a given format.
func Infof(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	Info(message)
}

// Info function prints an info message.
func Info(message interface{}) {
	jww.INFO.Println(message)
}
