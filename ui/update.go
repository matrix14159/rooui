package ui

import (
	"log/slog"
	"time"

	"github.com/matrix14159/rooui/core"
	"github.com/matrix14159/rooui/vdom"
)

func Update(c Comp, opts ...UpdateOption) {
	now := time.Now()
	defer func() {
		since := time.Now().Sub(now)
		slog.Info("update done.", "time", since.Milliseconds())
	}()

	cfg := &UpdateConfig{Mode: M_Self}
	for _, one := range opts {
		one(cfg)
	}

	element := c.Render()
	if element == nil {
		return
	}

	vnode := buildVNode(element)

	p := vdom.NewPatcher(vdom.NewStandardDomApi(),
		vdom.NewAttrModule(),
		vdom.NewClassModule(),
		vdom.NewDatasetModule(),
		vdom.NewEventModule(),
		vdom.NewPropsModule(),
		vdom.NewStyleModule(),
	)
	old, err := p.Patch(c.getVNode(), vnode)
	if err != nil {
		slog.Error("update component patch failed.", "error", err)
		return
	}
	c.updateVNode(old)
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

type UpdateConfig struct {
	Mode UpdateMode
}

type UpdateOption func(config *UpdateConfig)

type UpdateMode int

const (
	M_Self UpdateMode = 1
	M_Tree UpdateMode = 2
)

func WithUpdateMode(m UpdateMode) UpdateOption {
	return func(config *UpdateConfig) {
		config.Mode = m
	}
}
