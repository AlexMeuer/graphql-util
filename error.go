package hasura

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const noErrCode = 0

// Error is Hasura-compliant error response.
type Error struct {
	Message string
	Code    int
}

// Shorthand for ErrC(message, 0).
func Err(message string) *Error {
	return ErrC(message, noErrCode)
}

// Err creates a new Error from a string, without a code.
// Shorthand for NewError(message, code).
func ErrC(message string, code int) *Error {
	return NewError(message, code)
}

// ErrFrom create a new Error from an existing error with a code.
// Generally used to wrap an error when sending it back to GraphQL as an action/event response.
// Shorthand for NewErrorFrom(err).
func ErrFrom(err error) *Error {
	return ErrCFrom(err, noErrCode)
}

// ErrFrom create a new Error from an existing error with a code.
// Generally used to wrap an error when sending it back to GraphQL as an action/event response.
// Shorthand for NewErrorFrom(err).
func ErrCFrom(err error, code int) *Error {
	return NewErrorCFrom(err, noErrCode)
}

// Err creates a new Error from a string and a code.
func NewError(message string, code int) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

// ErrFrom create a new Error from an existing error, without a code.
func NewErrorFrom(err error) *Error {
	return NewErrorCFrom(err, noErrCode)
}

// ErrFrom create a new Error from an existing error with a code.
func NewErrorCFrom(err error, code int) *Error {
	if err == nil {
		return NewError("", code)
	}
	return NewError(err.Error(), code)
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	if e.Code == 0 {
		return e.Message
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *Error) MarshalJSON() ([]byte, error) {
	// Hasura expects `code` to be a string, and will be unhappy if it's an int,
	// despite expecting it to be in the range of http response codes.
	m := map[string]interface{}{
		"message": e.Message,
		"code":    strconv.Itoa(e.Code),
	}
	return json.Marshal(m)
}
