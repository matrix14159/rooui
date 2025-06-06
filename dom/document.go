package dom

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

import (
	"syscall/js"
)

type document struct {
	js.Value
}

// Window equivalent to window object in JavaScript DOM API.
var Document = &document{Window.Get("document")}

// Body return body element by shortcut
func (p *document) Body() *Object {
	elms := p.GetElementsByTagName("body")
	if elms.Length() > 0 {
		return elms.Item(0)
	}
	return nil
}

// Properties

// Returns the currently focused element
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (p *document) ActiveElement() *Object {
	return &Object{p.Get("activeElement")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (p *document) DocumentElement() *Object {
	return &Object{p.Get("documentElement")}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (p *document) CreateElement(tag string) *Object {
	return &Object{p.Call("createElement", tag)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElementNS
func (p *document) CreateElementNS(namespace, tag string) *Object {
	return &Object{p.Call("createElementNS", namespace, tag)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode
func (p *document) CreateTextNode(textContent string) *Object {
	return &Object{p.Call("createTextNode", textContent)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (p *document) GetElementById(id string) *Object {
	return &Object{p.Call("getElementById", id)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (p *document) GetElementsByTagName(name string) *HTMLCollection {
	return &HTMLCollection{p.Call("getElementsByTagName", name)}
}
