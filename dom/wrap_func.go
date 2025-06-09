package dom

import (
	"syscall/js"
)

func SetTimeout(f func(), delay int) {
	wrap := func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
	js.Global().Call("setTimeout", js.FuncOf(wrap), delay)
}

func RequestAnimationFrame(f func()) {
	wrap := func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
	js.Global().Call("requestAnimationFrame", js.FuncOf(wrap))
}
