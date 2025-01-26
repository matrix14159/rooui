package ui

import (
	"log/slog"
	"time"

	"github.com/matrix14159/rooui/core"
	"github.com/matrix14159/rooui/vdom"
)

var patch = vdom.NewPatcher(vdom.NewStandardDomApi(),
	vdom.NewAttrModule(),
	vdom.NewClassModule(),
	vdom.NewDatasetModule(),
	vdom.NewEventModule(),
	vdom.NewPropsModule(),
	vdom.NewStyleModule(),
)

// Update updates component c
func Update(c Comp, opts ...UpdateOption) {
	now := time.Now()
	defer func() {
		since := time.Now().Sub(now)
		slog.Info("update done.", "time", since.Milliseconds())
	}()

	cfg := &UpdateConfig{Patcher: patch}
	for _, one := range opts {
		one(cfg)
	}

	oldVn := c.getVNode()
	oldEl := c.getElement()

	element := c.Render()
	if element == nil {
		return
	}

	vnode := buildVNode(element)
	newVn, err := cfg.Patcher.Patch(oldVn, vnode)
	if err != nil {
		slog.Error("update component patch failed.", "error", err)
		return
	}
	updateCompVNode(element, newVn)

	if oldEl != nil {
		parent := oldEl.getParent()
		element.setParent(parent)

		if parent != nil {
			idx := parent.findVNodeChild(oldVn)
			parent.replaceVNodeChild(idx, oldVn, newVn)
		}
	}
}

func buildVNode(element core.HtmlElement) *vdom.VNode {
	on := vdom.NewEventListener()
	on.Events = element.GetEvents()
	data := &vdom.VNodeData{On: on}

	body := make([]*vdom.VNode, 0, len(element.GetBody()))
	for _, child := range element.GetBody() {
		node := buildVNode(child)
		body = append(body, node)
	}

	vnode := vdom.H(element.Tag(), element.GetText(), data, body)
	return vnode
}

func updateCompVNode(element Element, vnode *vdom.VNode) {
	element.getComp().updateVNode(vnode)
	len1 := len(element.GetBody())
	len2 := len(vnode.Children)
	if len1 != len2 {
		slog.Error("internal error: element body and vnode children doesn't match", "body", len1, "vnode", len2)
		return
	}
	for i, child := range element.GetBody() {
		el, ok := child.(Element)
		if !ok {
			continue
		}
		updateCompVNode(el, vnode.Children[i])
	}
}
