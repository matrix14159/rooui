package dom

import (
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleSheet
type CSSStyleSheet struct {
	js.Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleSheet/cssRules
func (p *CSSStyleSheet) CssRules() *CSSRuleList {
	return &CSSRuleList{p.Get("cssRules")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleSheet/insertRule
func (p *CSSStyleSheet) InsertRule(rule string, index int) (newlyIndex int) {
	return p.Call("insertRule", rule, index).Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleSheet/deleteRule
func (p *CSSStyleSheet) DeleteRule(index int) {
	p.Call("deleteRule", index)
}
