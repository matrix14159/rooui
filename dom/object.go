package dom

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

import (
	"syscall/js"
)

// Object represents a DOM element or node
type Object struct {
	js.Value
}

func (p *Object) Id() string {
	return p.Get("id").String()
}

func (p *Object) SetId(id string) {
	p.Set("id", id)
}

func (p *Object) ObjectType() string {
	if p.InstanceOf(js.Global().Get("HTMLElement")) {
		return "HTMLElement"
	}
	if p.InstanceOf(js.Global().Get("Element")) {
		return "Element"
	}
	if p.InstanceOf(js.Global().Get("Node")) {
		return "Node"
	}
	if p.InstanceOf(js.Global().Get("EventTarget")) {
		return "EventTarget"
	}
	return ""
}
