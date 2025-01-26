package ui

import (
	"github.com/matrix14159/rooui/core"
	"github.com/matrix14159/rooui/vdom"
)

// Comp is the base component interface
type Comp interface {
	Render() Element

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
}

func (p *Component) Render() Element {
	return nil
}

// Use uses atom as current component's Element
func (p *Component) Use(atom core.HtmlElement) Element {
	for _, child := range atom.GetBody() {
		el, ok := child.(Element)
		if !ok {
			continue
		}
		el.setParent(p)
	}
	p.element = &compElement{
		HtmlElement: atom,
		comp:        p,
	}
	return p.element
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

func (p *Component) LinkCSS(url string) {

}
