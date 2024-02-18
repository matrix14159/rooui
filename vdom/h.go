package vdom

import (
	"github.com/oklog/ulid/v2"
	"honnef.co/go/js/dom/v2"
)

func H(sel string, data *VNodeData, children VNodeChildren) *VNode {
	ulid.Make().String()
	return &VNode{}
}

func EmptyNodeAt(elm dom.Element) *VNode {
	return nil
}
