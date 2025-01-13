package core

import (
	"github.com/matrix14159/rooui/vdom"
)

type HtmlElement interface {
	Tag() string

	GetText() string

	GetEvents() map[string][]vdom.EventHandler

	GetBody() []HtmlElement
}
