package dom

// This file implements HTMLCollection interface
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection

import (
	"syscall/js"
)

type HTMLCollection struct {
	js.Value
}

func (p *HTMLCollection) Length() int {
	return p.Get("length").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection/item
// return object as Element
func (p *HTMLCollection) Item(index int) *Object {
	return &Object{p.Call("item", index)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection/namedItem
// return object as Element
func (p *HTMLCollection) NamedItem(index int) *Object {
	return &Object{p.Call("namedItem", index)}
}
