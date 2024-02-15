package vdom

import (
	"honnef.co/go/js/dom/v2"
)

type Listener func(vn VNode, event dom.Event)

type Events map[string]Listener
