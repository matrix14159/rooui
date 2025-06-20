package dom

import (
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CSSRuleList
type CSSRuleList struct {
	js.Value
}
