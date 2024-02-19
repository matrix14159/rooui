package vdom

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
	"honnef.co/go/js/dom/v2"
)

func H(sel string, data *VNodeData, children VNodeChildren) *VNode {
	ulid.Make().String()
	return &VNode{}
}

func EmptyNodeAt(elm dom.Element) *VNode {
	id := elm.ID()
	if id != "" {
		id = fmt.Sprintf("#%v", id)
	}

	// elm.className doesn't return a string when elm is an SVG element inside a shadowRoot.
	// https://stackoverflow.com/questions/29454340/detecting-classname-of-svganimatedstring
	classes := elm.GetAttribute("class")

	c := ""
	if classes != "" {
		cs := strings.Split(classes, " ")
		c = strings.Join(cs, ".")
	}

	return &VNode{
		Sel:      strings.ToLower(elm.TagName()) + id + c,
		Data:     nil,
		Children: nil,
		Elm:      elm,
		Text:     "",
		Key:      ulid.Make().String(),
	}
}
