/*
Package errors provides an implementation of the error interface used
across the application.
*/
package errors

import "fmt"

// OrbitError is a dead simple implementation of the error interface.
type OrbitError struct {
	Message string
}

// NewOrbitError creates an instance of OrbitError using a simple text message.
func NewOrbitError(message string) *OrbitError {
	return &OrbitError{
		Message: message,
	}
}

// NewOrbitErrorf creates an instance of OrbitError using a parametrized text message.
func NewOrbitErrorf(message string, args ...interface{}) *OrbitError {
	return &OrbitError{
		Message: fmt.Sprintf(message, args...),
	}
}

// Error is the implementation of the function Error from error interface.
func (e *OrbitError) Error() string {
	return e.Message
}
