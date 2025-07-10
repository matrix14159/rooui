package ui

import (
	"log/slog"
	"sync"
	"time"

	"github.com/matrix14159/rooui/html"
	"github.com/matrix14159/rooui/vdom"
)

var defaultUpdateFlow *updateFlow

// Update updates component c
func Update(c Comp, opts ...UpdateOption) {
	defaultUpdateFlow.Accept(c, opts)
}

type updateFlow struct {
	ch chan updateContext

	patch *vdom.Patcher

	ucPool sync.Pool
}

type updateContext struct {
	c Comp

	opts []UpdateOption
}

func (p *updateFlow) Accept(c Comp, opts []UpdateOption) {
	uc := p.ucPool.Get().(updateContext)
	uc.c = c
	uc.opts = opts
	p.ch <- uc
}

func (p *updateFlow) RunUpdateLoop() {
	p.ch = make(chan updateContext, 1024)
	p.patch = vdom.NewPatcher(vdom.NewStandardDomApi(),
		vdom.NewAttrModule(),
		vdom.NewClassModule(),
		vdom.NewDatasetModule(),
		vdom.NewEventModule(),
		vdom.NewPropsModule(),
		vdom.NewStyleModule(),
	)
	p.ucPool.New = func() any {
		return updateContext{}
	}

	go p.runUpdateLoop()
}

func (p *updateFlow) runUpdateLoop() {
	for uc := range p.ch {
		p.handleUpdate(uc)
		p.ucPool.Put(uc)
	}
}

// Update updates component c
func (p *updateFlow) handleUpdate(uc updateContext) {
	now := time.Now()
	defer func() {
		since := time.Now().Sub(now)
		slog.Info("update done.", "time", since.Milliseconds())
	}()

	cfg := &UpdateConfig{Patcher: p.patch}
	for _, one := range uc.opts {
		one(cfg)
	}

	oldVn := uc.c.getVNode()
	oldEl := uc.c.getElement()

	element := uc.c.Render()
	if element == nil {
		return
	}

	vnode := p.buildVNode(element)
	newVn, err := cfg.Patcher.Patch(oldVn, vnode)
	if err != nil {
		slog.Error("update component patch failed.", "error", err)
		return
	}
	p.updateCompVNode(element, newVn)

	if oldEl != nil {
		parent := oldEl.getParent()
		element.setParent(parent)

		if parent != nil {
			idx := parent.findVNodeChild(oldVn)
			parent.replaceVNodeChild(idx, oldVn, newVn)
		}
	}
}

func (p *updateFlow) buildVNode(element html.Element) *vdom.VNode {
	listener := vdom.NewEventListener()
	listener.Events = element.GetEvents()

	classes := element.GetClasses()
	classMap := make(vdom.Classes, len(classes))
	for _, class := range classes {
		classMap[class] = true
	}

	classStyleMap := element.GetClassStyle()
	for cls, styles := range classStyleMap {
		Style.Add("."+cls, styles...)
		classMap[cls] = true
	}

	styles := element.GetStyles()
	vStyle := vdom.NewVNodeStyle()
	for _, item := range styles {
		vStyle.Style[item.Name] = item.Value
	}

	props := make(vdom.Props)
	if id := element.GetId(); len(id) > 0 {
		props["id"] = id
	}

	data := &vdom.VNodeData{Props: props, Style: vStyle, Class: classMap, Listener: listener}

	body := make([]*vdom.VNode, 0, len(element.GetBody()))
	for _, child := range element.GetBody() {
		node := p.buildVNode(child)
		body = append(body, node)
	}

	vnode := vdom.H(element.Tag(), element.GetText(), data, body)
	return vnode
}

func (p *updateFlow) updateCompVNode(element Element, vnode *vdom.VNode) {
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
		p.updateCompVNode(el, vnode.Children[i])
	}
}
