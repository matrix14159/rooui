package dom

import (
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CSSRuleList
type CSSRuleList struct {
	js.Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSRuleList/length
func (p *CSSRuleList) Length() int {
	return p.Get("length").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSRuleList/item
func (p *CSSRuleList) Item(index int) *CSSRule {
	return &CSSRule{p.Call("item", index)}
}
