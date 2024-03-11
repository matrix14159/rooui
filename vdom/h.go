package vdom

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
	"honnef.co/go/js/dom/v2"
)

// H creates a vnode with sel tag
func H(sel string, data *VNodeData, text string, children VNodeChildren) (vnode *VNode) {
	vnode = &VNode{}
	if data == nil {
		data = &VNodeData{}
	}

	if strings.HasPrefix(sel, "svg") &&
		(len(sel) == 3 || sel[3] == '.' || sel[3] == '#') {
		addNS(sel, data, children)
	}
	vnode = &VNode{
		Sel:      sel,
		Data:     data,
		Children: children,
		Elm:      nil,
		Text:     text,
		Key:      ulid.Make().String(),
	}
	return
}

func addNS(sel string, data *VNodeData, children VNodeChildren) {
	data.Namespace = "http://www.w3.org/2000/svg"
	if sel == "foreignObject" {
		return
	}
	for _, child := range children {
		if child.Data != nil {
			addNS(child.Sel, child.Data, child.Children)
		}
	}
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
