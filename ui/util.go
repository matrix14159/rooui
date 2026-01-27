package ui

import (
	"fmt"
	"syscall/js"
)

// AsyncTask schedules a function to be executed asynchronously using JavaScript Promise microtask
// This avoids the goroutine limitation in TinyGo WebAssembly
func AsyncTask(f func()) {
	js.Global().Get("Promise").Call("resolve", nil).Call("then", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
		f()
		return nil
	}))
}

// ToPx convert value v to px
// example: ToPx(3) -> "3px"
func ToPx[T int | float64](v T) string {
	return fmt.Sprintf("%vpx", v)
}
