package ui

import (
	"fmt"

	"github.com/matrix14159/rooui/dom"
	"github.com/matrix14159/rooui/html"
	"github.com/matrix14159/rooui/vdom"
)

// Comp is the base component interface
type Comp interface {
	// Render return element for current component
	Render() Element

	// OnUpdated will be trigger when Render() is done
	OnUpdated()

	// GetName return component's name, only use for debug trace
	GetName() string

	// GetId return the component render element's id
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
	// debug trace
	Name string

	// render element for current component
	element Element

	// vdom node for current component
	vnode *vdom.VNode

	// speed up for finding vnode children
	subVnodeMap map[*vdom.VNode]int

	//
	magicIdMap map[string]string
}

// Render will render component c by element el
func Render(c Comp, el html.Element) Element {
	element := &compElement{
		Element: el,
		comp:    c,
	}
	c.updateElement(element)

	for i, child := range el.GetBody() {
		childEl, ok := child.(Element)
		if !ok {
			childComp := &Component{}
			if c.GetName() != "" {
				childComp.Name = fmt.Sprintf("%v-%v", c.GetName(), i)
			}
			childEl = Render(childComp, child)
			el.ReplaceChild(i, childEl)
		}
		childEl.setParent(c)
	}
	return element
}

func (p *Component) Render() Element {
	return nil
}

func (p *Component) OnUpdated() {
}

func (p *Component) GetName() string {
	return p.Name
}

func (p *Component) GetId() string {
	return p.element.GetId()
}

// Uid return a unique id for key k
// k must begin with a~z when use for id or name
func (p *Component) Uid(k string) string {
	if p.magicIdMap == nil {
		p.magicIdMap = make(map[string]string)
	}
	id, found := p.magicIdMap[k]
	if !found {
		maxId := dom.Window.Get("u_id").Int()
		maxId++
		dom.Window.Set("u_id", maxId)
		id = fmt.Sprintf("%s_%v", k, maxId)
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
