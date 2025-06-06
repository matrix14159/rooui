package dom

// This file implements CSSStyleDeclaration interface
// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration

import (
	"syscall/js"
)

type CSSStyleDeclaration struct {
	js.Value
}

// Properties

func (s *CSSStyleDeclaration) CssText() string {
	return s.Get("cssText").String()
}

func (s *CSSStyleDeclaration) Length() int {
	return s.Get("length").Int()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration/item
func (p *CSSStyleDeclaration) Item(index int) string {
	return p.Call("item", index).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration/removeProperty
func (p *CSSStyleDeclaration) RemoveProperty(name string) string {
	return p.Call("removeProperty", name).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration/setProperty
func (p *CSSStyleDeclaration) SetProperty(name, value string) {
	p.Call("setProperty", value)
}
