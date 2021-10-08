package shared

import "github.com/ipld/go-ipldtool/errors"

const (
	ErrCode_InvalidArgs = "ipldtool-error-invalid-args"
)

// ErrInvalidArgs is an error constructor function.
//
// Errors:
//
//   - ipldtool-error-invalid-args -- with the message and cause you provide.
func ErrInvalidArgs(msg string, cause error) error {
	if cause != nil {
		msg += ": " + cause.Error()
	}
	return &errors.Error{
		ErrCode_InvalidArgs,
		msg,
		nil,
		cause,
	}
}
