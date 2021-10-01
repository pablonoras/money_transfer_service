package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	Status  int
	Code    string
	Message string
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.Code,
		Message: e.Message,
	})
}

func (e *Error) StatusCode() int {
	return e.Status
}

// Error returns a string message of the error. It is a concatenation of Code and Message fields.
// This means the Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError creates a new error with the given status code and message.
func NewError(statusCode int, message string) error {
	return NewErrorf(statusCode, message)
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func NewErrorf(status int, format string, args ...interface{}) error {
	return &Error{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}
}
