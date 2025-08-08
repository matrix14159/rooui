package ui

import (
	"fmt"

	"github.com/matrix14159/rooui/html"
	"github.com/matrix14159/rooui/vdom"
	"github.com/rs/xid"
)

// Comp is the base component interface
type Comp interface {
	// Render return element for current component
	Render() Element

	GetId() string

	getElement() Element

	updateElement(el Element)

	// getVNode return current component refer vnode
	getVNode() *vdom.VNode

	// updateVNode update refer vnode for current component
	updateVNode(vnode *vdom.VNode)

	findVNodeChild(child *vdom.VNode) (idx int)
	replaceVNodeChild(idx int, oldChild, newChild *vdom.VNode)
}

type Component struct {
	// render element for current component
	element Element

	// vdom node for current component
	vnode *vdom.VNode

	// speed up for finding vnode children
	subVnodeMap map[*vdom.VNode]int

	//
	magicIdMap map[string]string
}

func (p *Component) Render() Element {
	return nil
}

// Use uses el as current component's render Element
func (p *Component) Use(el html.Element) Element {
	for _, child := range el.GetBody() {
		el, ok := child.(Element)
		if !ok {
			continue
		}
		el.setParent(p)
	}
	p.element = &compElement{
		Element: el,
		comp:    p,
	}
	return p.element
}

func (p *Component) GetId() string {
	return p.element.GetId()
}

// MagicId return a unique id for k
// k must begin with a~z if use for id or name
func (p *Component) MagicId(k string) string {
	if p.magicIdMap == nil {
		p.magicIdMap = make(map[string]string)
	}
	id, found := p.magicIdMap[k]
	if !found {
		id = fmt.Sprintf("%s%s", k, xid.New().String())
		p.magicIdMap[k] = id
	}
	return id
}

func (p *Component) getElement() Element {
	return p.element
}

func (p *Component) updateElement(el Element) {
	p.element = el
}

func (p *Component) getVNode() *vdom.VNode {
	return p.vnode
}

func (p *Component) updateVNode(vnode *vdom.VNode) {
	p.vnode = vnode
	p.subVnodeMap = nil
}

func (p *Component) findVNodeChild(child *vdom.VNode) (idx int) {
	if p.subVnodeMap == nil {
		p.subVnodeMap = make(map[*vdom.VNode]int, len(p.vnode.Children))
		for i, one := range p.vnode.Children {
			p.subVnodeMap[one] = i
		}
	}
	idx, found := p.subVnodeMap[child]
	if found {
		return
	}
	return -1
}

func (p *Component) replaceVNodeChild(idx int, oldChild, newChild *vdom.VNode) {
	if idx < 0 || idx >= len(p.vnode.Children) {
		return
	}
	p.vnode.Children[idx] = newChild
	delete(p.subVnodeMap, oldChild)
	p.subVnodeMap[newChild] = idx
}
