package dom

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// NewError returns a JS Error with the provided Go error's error message.
func NewError(goErr error) js.Value {
	errConstructor := js.Global().Get("Error")
	return errConstructor.New(goErr.Error())
}

// TypeMismatchError is returned when a function is called with a js.Value that has the incorrect type.
type TypeMismatchError struct {
	Expected js.Type
	Actual   js.Type
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf("expected %v type, got %v type instead", e.Expected, e.Actual)
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()
