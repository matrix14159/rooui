package vdom

import (
	"log/slog"

	"honnef.co/go/js/dom/v2"
)

type Patcher struct {
	api DOMAPI

	curVNode *VNode

	inserted []*VNode
}

// NewPatcher create a patcher for dom's elm
func NewPatcher(api DOMAPI, elm dom.Element) *Patcher {
	return &Patcher{
		api:      api,
		curVNode: EmptyNodeAt(elm),
		inserted: make([]*VNode, 0),
	}
}

// CurrentVNode return current state of virtual-dom
func (p *Patcher) CurrentVNode() *VNode {
	return p.curVNode
}

func (p *Patcher) Patch(newNode *VNode) (err error) {
	if SameVNode(p.curVNode, newNode) {
		p.patchVNode(newNode)
	} else {
		slog.Info("not the same vnode, replace current element", "curVNode", *p.curVNode)
		elm := p.curVNode.Elm
		parent := p.api.ParentNode(elm)

		newNode.Elm = p.createElm(newNode)

		if parent != nil {
			p.api.InsertBefore(parent, newNode.Elm, p.api.NextSibling(newNode.Elm))
			p.removeVNodes(parent, []*VNode{p.curVNode}, 0, 0)
		}
	}
	p.curVNode = newNode
	return
}

func (p *Patcher) patchVNode(vnode *VNode) {

}

func (p *Patcher) createElm(vnode *VNode) dom.Node {
	if vnode.Sel == "" {
		return p.api.CreateTextNode(vnode.Text)
	}

	//hashIdx := strings.Index(vnode.Sel, "#")
	//dotIdx := strings.Index(vnode.Sel, ".")

	return nil
}

func (p *Patcher) removeVNodes(parentElm dom.Node, vnodes []*VNode, startIdx, endIdx int) {

}
