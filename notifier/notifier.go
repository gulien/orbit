// Package notifier implements a simple helper for displaying output to users.
// Credits: this package has been heavily inspired by https://github.com/Masterminds/glide/blob/master/msg/msg.go.
package notifier

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// OrbitNotifier provides the underlying implementation that displays output to users.
type OrbitNotifier struct {
	sync.Mutex

	// Stdout is the location where this prints output.
	Stdout io.Writer

	// Stderr is the location where this prints logs.
	Stderr io.Writer
}

// newOrbitNotifier creates a default OrbitNotifier to display output.
func newOrbitNotifier() *OrbitNotifier {
	return &OrbitNotifier{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// Houston contains the default OrbitNotifier used by the application.
var Houston = newOrbitNotifier()

// Info logs information using the Houston notifier.
func Info(notification string, args ...interface{}) {
	prefix := "[INF]\t"
	Houston.notify(prefix+notification, nil, args...)
}

// Error logs error information using the Houston notifier.
func Error(err error) {
	prefix := "[ERR]\t"
	Houston.notify(prefix+err.Error(), err)
}

// notify prints a notification with optional parameters.
// If err is not nil, prints the notification to Stderr.
func (n *OrbitNotifier) notify(notification string, err error, args ...interface{}) {
	n.Lock()
	defer n.Unlock()

	if !strings.HasSuffix(notification, "\n") {
		notification += "\n"
	}

	if err != nil {
		fmt.Fprint(n.Stderr, notification)
	} else {
		fmt.Fprintf(n.Stdout, notification, args...)
	}
}
