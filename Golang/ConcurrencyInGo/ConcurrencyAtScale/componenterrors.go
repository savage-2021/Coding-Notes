package main

import (
	"fmt"
	"runtime/debug"
)

type MyError struct {
	Inner error 
	Message string
	StackTrace string 
	Misc map[string] interface{}
}


func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError {
		Inner: err,
		Message: fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()), // Takes note of the stack trace 
		Misc: make(map[string]interface{}),
	}
}


func wrapError2(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError {
		Inner: error,
		Message: fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),

	}
}

func (err MyError) Error() string {
	return err.Message
}

func PostReport(id string) error {
	result, err := lowlevel.DoWork()
	if err != nil {
		// check to see if we have a wellformed error,
		// if we don't, then we ferry it up the stack
		if _, ok := err.(lowlevel.Error); ok {
			// By wrapping the error, we give some pertinent information for our module, amd give it a new type
			// this might involve hiding some low-level details that mayb not be important for the user within this context
			err = WrapErr(err, "Cannot post report with id %q", id)
		}
		return err
	}
	return nil
}
func main() {

}