package ui

import (
	"log/slog"
	"sync"
	"time"

	"github.com/matrix14159/rooui/html"
	"github.com/matrix14159/rooui/vdom"
)

type updateFlow struct {
	queue      []updateContext
	processing bool
	mu         sync.Mutex

	patch *vdom.Patcher

	ucPool sync.Pool
}

type updateContext struct {
	c Comp

	opts []UpdateOption
}

var defaultUpdateFlow *updateFlow

func newUpdateFlow() *updateFlow {
	p := new(updateFlow)
	p.queue = make([]updateContext, 0, 1024)
	p.processing = false
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
	return p
}

// Update updates component c asynchronously using JavaScript event loop
func Update(c Comp, opts ...UpdateOption) {
	defaultUpdateFlow.accept(c, opts)
}

func (p *updateFlow) accept(c Comp, opts []UpdateOption) {
	uc := p.ucPool.Get().(updateContext)
	uc.c = c
	uc.opts = opts

	p.mu.Lock()
	p.queue = append(p.queue, uc)
	scheduleProcessing := !p.processing
	if scheduleProcessing {
		p.processing = true
	}
	p.mu.Unlock()

	if scheduleProcessing {
		AsyncTask(p.processQueue)
	}
}

func (p *updateFlow) processQueue() {
	for {
		p.mu.Lock()
		if len(p.queue) == 0 {
			p.processing = false
			p.mu.Unlock()
			break
		}

		// Take the first update from the queue
		uc := p.queue[0]
		p.queue = p.queue[1:]
		p.mu.Unlock()

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

	p.touchOnUpdated(element)
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

	attrs := element.GetAttributes()

	props := element.GetProps()
	if id := element.GetId(); len(id) > 0 {
		props["id"] = id
	}

	data := &vdom.VNodeData{Attrs: attrs, Props: props, Style: vStyle, Class: classMap, Listener: listener}

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

func (p *updateFlow) touchOnUpdated(element Element) {
	element.getComp().OnUpdated()
	for _, child := range element.GetBody() {
		el, ok := child.(Element)
		if !ok {
			continue
		}
		slog.Info("touchOnUpdated", "id", el.GetId())
		el.getComp().OnUpdated()
	}
}
