package ui

import (
	"log/slog"
	"time"

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

	body := []*vdom.VNode{}
	for _, child := range element.GetBody() {
		node := vdom.H(child.Tag(), child.GetText(), nil, nil)
		body = append(body, node)
	}

	on := vdom.NewEventListener(nil)
	on.Events = element.GetEvents()
	data := &vdom.VNodeData{On: on}
	vnode := vdom.H(element.Tag(), element.GetText(), data, body)

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
