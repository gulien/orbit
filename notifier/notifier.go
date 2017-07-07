/*
Package notifier implements a simple helper for displaying output to users.

Credits: this package has been inspired by https://github.com/Masterminds/glide/blob/master/msg/msg.go.
*/
package notifier

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/shiena/ansicolor"
)

// OrbitNotifier provides the underlying implementation that displays output to users.
type OrbitNotifier struct {
	sync.Mutex

	// stdout is the location where this prints output.
	stdout io.Writer

	// stderr is the location where this prints logs.
	stderr io.Writer

	// silent disables the notifications if true.
	silent bool
}

// newOrbitNotifier creates a default OrbitNotifier to display output.
func newOrbitNotifier() *OrbitNotifier {
	return &OrbitNotifier{
		stdout: ansicolor.NewAnsiColorWriter(os.Stdout),
		stderr: ansicolor.NewAnsiColorWriter(os.Stderr),
	}
}

// Houston is the OrbitNotifier instance used by the application.
var Houston = newOrbitNotifier()

// Mute disables the notifications.
func Mute() {
	Houston.silent = true
}

// Info logs information using the Houston notifier.
func Info(notification string, args ...interface{}) {
	if !Houston.silent {
		prefix := fmt.Sprintf("[%si%s] ", "\x1b[36m", "\x1b[0m")
		Houston.notify(prefix+notification, nil, args...)
	}
}

// Error logs error information using the Houston notifier.
func Error(err error) {
	if !Houston.silent {
		prefix := fmt.Sprintf("[%se%s] ", "\x1b[31m", "\x1b[0m")
		Houston.notify(prefix+err.Error(), err)
	}
}

/*
notify prints a notification with optional parameters.

If err is not nil, prints the notification to stderr.
*/
func (n *OrbitNotifier) notify(notification string, err error, args ...interface{}) {
	n.Lock()
	defer n.Unlock()

	if !strings.HasSuffix(notification, "\n") {
		notification += "\n"
	}

	if err != nil {
		fmt.Fprint(n.stderr, notification)
	} else {
		fmt.Fprintf(n.stdout, notification, args...)
	}
}
