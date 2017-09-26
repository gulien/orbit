/*
Package errors provides an implementation of the error interface used
across the application.
*/
package errors

import "fmt"

// OrbitError is a dead simple implementation of the error interface.
type OrbitError struct {
	// message is the information text of the error.
	message string
}

// NewOrbitError creates an instance of OrbitError using a simple message.
func NewOrbitError(message string) *OrbitError {
	return &OrbitError{
		message: message,
	}
}

// NewOrbitErrorf creates an instance of OrbitError using a parametrized message.
func NewOrbitErrorf(message string, args ...interface{}) *OrbitError {
	return &OrbitError{
		message: fmt.Sprintf(message, args...),
	}
}

// Error is the implementation of the function Error from the error interface.
func (e *OrbitError) Error() string {
	return e.message
}
