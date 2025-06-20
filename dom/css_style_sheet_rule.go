package dom

import (
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CSSRule
type CSSRule struct {
	js.Value
}
