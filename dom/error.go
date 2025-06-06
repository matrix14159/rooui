package dom

import (
	"reflect"
	"syscall/js"
)

// NewError returns a JS Error with the provided Go error's error message.
func NewError(goErr error) js.Value {
	errConstructor := js.Global().Get("Error")
	return errConstructor.New(goErr.Error())
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()
