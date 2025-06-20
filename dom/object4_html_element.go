package dom

// This file implements HTMLElement interface
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/style
func (p *Object) Style() *CSSStyleDeclaration {
	return &CSSStyleDeclaration{p.Get("style")}
}

func (p *Object) Sheet() *CSSStyleSheet {
	return &CSSStyleSheet{p.Get("sheet")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset
func (p *Object) Dataset() *Object {
	return &Object{p.Get("dataset")}
}

// Methods

// Remove keyboard focus from the current element
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
func (p *Object) Blur() {
	p.Call("blur")
}

// Set focus on the specified element, if it can be focused.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
func (p *Object) Focus() {
	p.Call("focus")
}
