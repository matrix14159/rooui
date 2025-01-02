package ui

import (
	"github.com/matrix14159/rooui/vdom"
)

// Comp is the base component interface
type Comp interface {
	Render

	// updateVNode update refer vnode for current component
	updateVNode(vnode *vdom.VNode)
}

type Component struct {
	// vdom node for current component
	vnode *vdom.VNode
}

func (p *Component) Render() Element {
	return nil
}

func (p *Component) updateVNode(vnode *vdom.VNode) {
	p.vnode = vnode
}

func (p *Component) LinkCSS(url string) {

}
