package dom

import (
	"syscall/js"
)

type Event struct {
	js.Value
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (p Event) Target() *Object {
	return &Object{p.Get("target")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (p Event) EventType() string {
	return p.Get("type").String()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault
func (p Event) PreventDefault() {
	p.Call("preventDefault")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopImmediatePropagation
func (p Event) StopImmediatePropagation() {
	p.Call("stopImmediatePropagation")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation
func (p Event) StopPropagation() {
	p.Call("stopPropagation")
}
