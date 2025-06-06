package dom

// This file implements DOMTokenList interface
// https://developer.mozilla.org/en/docs/Web/API/DOMTokenList

import (
	"syscall/js"
)

type DOMTokenList struct {
	js.Value
}

func (p *DOMTokenList) Length() int {
	return p.Get("length").Int()
}

func (p *DOMTokenList) Contains(s string) bool {
	return p.Call("contains", s).Bool()
}

func (p *DOMTokenList) Add(s string) {
	p.Call("add", s)
}

func (p *DOMTokenList) Remove(s string) {
	p.Call("remove", s)
}

func (p *DOMTokenList) Toggle(s string) {
	p.Call("toggle", s)
}
