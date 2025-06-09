package dom

import (
	"syscall/js"
)

type window struct {
	js.Value
}

// Window equivalent to window object in JavaScript DOM API.
var Window = &window{js.Global()}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/document
var Document = &document{Window.Get("document")}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/location
func (p *window) Location() *Location {
	return &Location{p.Get("location")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/navigator
func (p *window) Navigator() *Navigator {
	return &Navigator{p.Get("navigator")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func (p *window) History() *History {
	return &History{p.Get("history")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (p *window) Alert(s string) {
	p.Call("alert", s)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/setInterval
func (p *window) SetInterval(f func(), delay int) {
	wrap := func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
	p.Call("setInterval", js.FuncOf(wrap), delay)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/setTimeout
func (p *window) SetTimeout(f func(), delay int) {
	wrap := func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
	p.Call("setTimeout", js.FuncOf(wrap), delay)
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (p *window) AddEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		p.Call("addEventListener", t, listener, args[0])
	} else {
		p.Call("addEventListener", t, listener)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (p *window) RemoveEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		p.Call("removeEventListener", t, listener, args[0])
	} else {
		p.Call("removeEventListener", t, listener)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/requestAnimationFrame
func (p *window) RequestAnimationFrame(f func()) {
	wrap := func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	}
	js.Global().Call("requestAnimationFrame", js.FuncOf(wrap))
}

// Properties of window object

// https://developer.mozilla.org/en-US/docs/Web/API/Window/pageXOffset
func (p *window) PageXOffset() float64 {
	return p.Get("pageXOffset").Float()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/pageYOffset
func (p *window) PageYOffset() float64 {
	return p.Get("pageYOffset").Float()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (p *window) ScrollX() float64 {
	return p.Get("scrollX").Float()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY
func (p *window) ScrollY() float64 {
	return p.Get("scrollY").Float()
}
