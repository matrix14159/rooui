package vdom

import (
	"log/slog"
	"math"
	"strings"

	"honnef.co/go/js/dom/v2"
)

type Patcher struct {
	api DOMAPI

	curVNode *VNode

	// new vnode when patch
	inserted []*VNode
}

// NewPatcher create a patcher for dom's elm
func NewPatcher(api DOMAPI, cur *VNode) *Patcher {
	return &Patcher{
		api:      api,
		curVNode: cur,
		inserted: make([]*VNode, 0),
	}
}

// CurrentVNode return current state of virtual-dom
func (p *Patcher) CurrentVNode() *VNode {
	return p.curVNode
}

func (p *Patcher) Patch(newNode *VNode) (err error) {
	p.inserted = make([]*VNode, 0)
	if SameVNode(p.curVNode, newNode) {
		p.patchVNode(newNode)
	} else {
		elm := p.curVNode.Elm
		parent := p.api.ParentNode(elm)

		newNode.Elm = p.createElm(newNode)

		slog.Info("not the same vnode, replace current element",
			"curVNode", *p.curVNode, "parent", parent.NodeName(), "newNode", newNode)

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

	hashIdx := strings.Index(vnode.Sel, "#")
	hash := len(vnode.Sel)
	if hashIdx > 0 {
		hash = hashIdx
	}

	dotIdx := strings.Index(vnode.Sel, ".")
	dot := len(vnode.Sel)
	if dotIdx > 0 {
		dot = dotIdx
	}

	tag := vnode.Sel
	if hashIdx != -1 || dotIdx != -1 {
		minLen := math.Min(float64(hash), float64(dot))
		tag = vnode.Sel[0:int(minLen)]
	}
	slog.Info("createElm", "tag", tag)

	var elm dom.HTMLElement
	if vnode.Data != nil && vnode.Data.Namespace != "" {
		elm = p.api.CreateElementNS(vnode.Data.Namespace, tag)
	} else {
		elm = p.api.CreateElement(tag)
	}
	vnode.Elm = elm

	if hash < dot {
		slog.Info("createElm", "id", vnode.Sel[hash+1:dot])
		elm.SetAttribute("id", vnode.Sel[hash+1:dot])
	}

	if dotIdx > 0 {
		cls := strings.Replace(vnode.Sel[dot+1:], ".", " ", -1)
		slog.Info("createElm", "raw-class", vnode.Sel[dot+1:], "use-class", cls)
		elm.SetAttribute("class", cls)
	}

	if vnode.Text != "" && len(vnode.Children) == 0 {
		// allow h1 and similar nodes to be created w/ text and empty child list
		p.api.AppendChild(elm, p.api.CreateTextNode(vnode.Text))
	}
	for _, child := range vnode.Children {
		if child == nil {
			continue
		}
		c := p.createElm(child)
		p.api.AppendChild(elm, c)
	}
	p.inserted = append(p.inserted, vnode)
	return vnode.Elm
}

func (p *Patcher) removeVNodes(parentElm dom.Node, vnodes []*VNode, startIdx, endIdx int) {

}
