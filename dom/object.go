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

func (p *Object) GetId() string {
	return p.Get("id").String()
}

func (p *Object) Id(id string) {
	p.Set("id", id)
}
