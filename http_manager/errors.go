package http_manager

import "fmt"

type badRequestError struct {
	Message string
}

func newBadRequestError(format string, a ...interface{}) *badRequestError {
	return &badRequestError{Message: fmt.Sprintf(format, a...)}
}

func (bre *badRequestError) Error() string {
	return bre.Message
}
