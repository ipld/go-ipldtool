package ipldtoolerr

// This package name is verbose,
// but using the package name "errors" would also be inconvenient
// because it's not uncommon to also need to import stdlib's "errors" package.

import "fmt"

// Some frequently used error constants are gathered here.
// This is not an exhaustive list;
// error code constants may also be defined locally and on-the-fly.
//
// (REVIEW: I'm not sure this actually provides value.
// One ends up replicating the constants in the docs anyway;
// so naming a constant in the code just gives you a *different* thing to copy around:
// that feels like boilerplate rather than value-added.
// The only value add is having something one can try to autocomplete with.)
// (... okay, and having the constants for equality checks in handling.  That's useful.)
const (
	ErrCode_InvalidArgs = "ipldtool-error-invalid-args"
)

// New constructs a new error value,
// taking an error code parameter and a freetext string as message.
//
// See Newf for a little more flexibility.
//
// Errors:
//
//    - param: errcode -- this constant will be the error's (analyzable!) code.
func New(errcode string, msg string) *Error {
	return &Error{
		errcode,
		msg,
		nil,
		nil,
	}
}

// Newf constructs a new error value,
// taking an error code parameter,
// and a format string and additional parameters in the style of fmt.Sprintf.
//
// If the last argument is an error,
// it will also be marked as the cause in the new error returned.
//
// Errors:
//
//    - param: errcode -- this constant will be the error's (analyzable!) code.
func Newf(errcode string, format string, args ...interface{}) *Error {
	return &Error{
		errcode,
		fmt.Sprintf(format, args...),
		nil,
		func() error {
			if len(args) == 0 {
				return nil
			}
			last := args[len(args)-1]
			if cast, ok := last.(error); ok {
				return cast
			}
			return nil
		}(),
	}
}

type Error struct {
	TheCode    string            `json:"code"`
	TheMessage string            `json:"msg,omitempty"`
	TheDetails map[string]string `json:"details,omitempty"`
	TheCause   error             `json:"cause,omitempty"`
}

func (e *Error) Code() string {
	if e == nil {
		return ""
	}
	return e.TheCode
}
func (e *Error) Message() string            { return e.TheMessage }
func (e *Error) Details() map[string]string { return e.TheDetails }
func (e *Error) Cause() error               { return e.TheCause }
func (e *Error) Error() string {
	switch {
	case e.TheCause == nil && e.TheMessage == "":
		return e.TheCode
	case e.TheCause == nil:
		return fmt.Sprintf("%s: %s", e.TheCode, e.TheMessage)
	case e.TheMessage == "":
		return fmt.Sprintf("%s: %s", e.TheCode, e.TheCause)
	}
	return fmt.Sprintf("%s: %s: %s", e.TheCode, e.TheMessage, e.TheCause)
}
