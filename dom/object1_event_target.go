package dom

import (
	"syscall/js"
)

// This file implements EventTarget interface
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (p *Object) AddEventListener(typ string, useCapture bool, listener func(Event)) js.Func {
	wrapper := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener(Event{args[0]})
		return nil
	})
	p.Call("addEventListener", typ, wrapper, useCapture)
	return wrapper
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (p *Object) RemoveEventListener(typ string, useCapture bool, listener js.Func) {
	p.Call("removeEventListener", typ, listener, useCapture)
	listener.Release()
}
