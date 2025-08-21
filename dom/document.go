package dom

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

import (
	"syscall/js"
)

type document struct {
	js.Value
}

// Head return head element by shortcut
func (p *document) Head() *Object {
	elms := p.GetElementsByTagName("head")
	if elms.Length() > 0 {
		return elms.Item(0)
	}
	return nil
}

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

// https://developer.mozilla.org/en-US/docs/Web/API/Document/styleSheets
func (p *document) StyleSheets() *StyleSheetList {
	return &StyleSheetList{p.Get("styleSheets")}
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
	v := p.Call("getElementById", id)
	if v.IsNull() || v.IsUndefined() {
		return nil
	}
	return &Object{v}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (p *document) GetElementsByTagName(name string) *HTMLCollection {
	return &HTMLCollection{p.Call("getElementsByTagName", name)}
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/Document/querySelectorAll
func (p *document) QuerySelectorAll(selectors string) *NodeList {
	return &NodeList{p.Call("querySelectorAll", selectors)}
}
