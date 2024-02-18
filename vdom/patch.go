package vdom

import (
	"honnef.co/go/js/dom/v2"
)

type Patcher struct {
	api DOMAPI

	curVNode *VNode

	inserted []*VNode
}

func NewPatcher(api DOMAPI, vnode *VNode) *Patcher {
	return &Patcher{
		api:      api,
		curVNode: vnode,
		inserted: make([]*VNode, 0),
	}
}

// CurrentVNode return current state of virtual-dom
func (p *Patcher) CurrentVNode() *VNode {
	return p.curVNode
}

func (p *Patcher) Patch(vnode *VNode) (err error) {
	if SameVNode(p.curVNode, vnode) {
		p.patchVNode(vnode)
	} else {
		elm := p.curVNode.Elm
		parent := p.api.ParentNode(elm)

		vnode.Elm = p.createElm(vnode)

		if parent != nil {
			p.api.InsertBefore(parent, vnode.Elm, p.api.NextSibling(vnode.Elm))
			p.removeVNodes(parent, []*VNode{p.curVNode}, 0, 0)
		}
	}
	p.curVNode = vnode
	return
}

func (p *Patcher) patchVNode(vnode *VNode) {

}

func (p *Patcher) createElm(vnode *VNode) dom.Node {
	return nil
}

func (p *Patcher) removeVNodes(parentElm dom.Node, vnodes []*VNode, startIdx, endIdx int) {

}
