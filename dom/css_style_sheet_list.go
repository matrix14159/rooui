package dom

import (
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/StyleSheetList
type StyleSheetList struct {
	js.Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/StyleSheetList/length
func (p *StyleSheetList) Length() int {
	return p.Get("length").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/StyleSheetList/item
func (p *StyleSheetList) Item(index int) *CSSStyleSheet {
	return &CSSStyleSheet{p.Call("item", index)}
}
