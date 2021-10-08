package errors

import "fmt"

/*
Generalized constructor functions don't really jive well with the analysis tool yet (at least as far as I can figure out).
Might be worth trying to make something like these later.
For now: to be friendly to the analyser, we'll make a constructor function for each specific code.
(That's not a bad thing anyway: lets us get more specific about args it should have!)


func New(code string, msg string) *Error {
	return &Error{
		code,
		msg,
		nil,
		nil,
	}
}

func Newf(code string, format string, args ...interface{}) *Error {
	return &Error{
		code,
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
*/

type Error struct {
	TheCode    string            `json:"code"`
	TheMessage string            `json:"msg,omitempty"`
	TheDetails map[string]string `json:"details,omitempty"`
	TheCause   error             `json:"cause,omitempty"`
}

func (e *Error) Code() string               { return e.TheCode }
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
