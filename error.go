package gql

import "fmt"

// Error is Hasura-compliant error response.
type Error struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

// Err creates a new Error from a string, without a code.
// Shorthand for NewError(message).
func Err(message string) *Error {
	return NewError(message)
}

// ErrFrom create a new Error from an existing error, without a code.
// Generally used to wrap an error when sending it back to GraphQL as an action/event response.
// Shorthand for NewErrorFrom(err).
func ErrFrom(err error) *Error {
	return NewErrorFrom(err)
}

// Err creates a new Error from a string, without a code.
func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}

// ErrFrom create a new Error from an existing error, without a code.
// Generally used to wrap an error when sending it back to GraphQL as an action/event response.
func NewErrorFrom(err error) *Error {
	if err == nil {
		return NewError("")
	}
	return NewError(err.Error())
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	if e.Code == "" {
		return e.Message
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
