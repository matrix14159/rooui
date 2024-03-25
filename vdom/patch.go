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
	oldVnode := p.curVNode

	p.inserted = make([]*VNode, 0)
	if SameVNode(oldVnode, newNode) {
		p.patchVNode(oldVnode, newNode)
	} else {
		elm := oldVnode.Elm
		parent := p.api.ParentNode(elm)

		newNode.Elm = p.createElm(newNode)

		slog.Info("not the same vnode, replace current element",
			"oldVnode", *oldVnode, "parent", parent.NodeName(), "newNode", newNode)

		if parent != nil {
			p.api.InsertBefore(parent, newNode.Elm, p.api.NextSibling(newNode.Elm))
			p.removeVNodes(parent, []*VNode{oldVnode}, 0, 0)
		}
	}
	p.curVNode = newNode
	return
}

func (p *Patcher) patchVNode(oldVnode, vnode *VNode) {
	vnode.Elm = oldVnode.Elm
	if vnode.Text == "" {
		switch {
		case len(vnode.Children) > 0 && len(oldVnode.Children) > 0:
			p.updateChildren(vnode.Elm, oldVnode.Children, vnode.Children)

		case len(vnode.Children) > 0:
			if oldVnode.Text != "" {
				p.api.SetTextContent(vnode.Elm, "")
			}
			p.addVNodes(vnode.Elm, nil, vnode.Children, 0, len(vnode.Children)-1)

		case len(oldVnode.Children) > 0:
			p.removeVNodes(vnode.Elm, oldVnode.Children, 0, len(oldVnode.Children)-1)

		case oldVnode.Text != "":
			p.api.SetTextContent(vnode.Elm, "")
		}
	} else if vnode.Text != oldVnode.Text {
		if len(oldVnode.Children) > 0 {
			p.removeVNodes(vnode.Elm, oldVnode.Children, 0, len(oldVnode.Children)-1)
		}
		p.api.SetTextContent(vnode.Elm, "")
	}
}

func (p *Patcher) updateChildren(parentElm dom.Node, oldCh, newCh []*VNode) {
	oldStartIdx := 0
	oldEndIdx := len(oldCh) - 1
	oldStartVnode := oldCh[0]
	oldEndVnode := oldCh[oldEndIdx]

	newStartIdx := 0
	newEndIdx := len(newCh) - 1
	newStartVnode := newCh[0]
	newEndVnode := newCh[newEndIdx]

	oldKeyToIdx := make(map[string]int)

	for oldStartIdx <= oldEndIdx && newStartIdx <= newEndIdx {
		switch {
		case oldStartVnode == nil: // Vnode might have been moved left
			oldStartIdx++
			oldStartVnode = oldCh[oldStartIdx]

		case oldEndVnode == nil:
			oldEndIdx--
			oldEndVnode = oldCh[oldEndIdx]

		case newStartVnode == nil:
			newStartIdx++
			newStartVnode = newCh[newStartIdx]

		case newEndVnode == nil:
			newEndIdx--
			newEndVnode = newCh[newEndIdx]

		case SameVNode(oldStartVnode, newStartVnode):
			p.patchVNode(oldStartVnode, newStartVnode)
			oldStartIdx++
			oldStartVnode = oldCh[oldStartIdx]
			newStartIdx++
			newStartVnode = newCh[newStartIdx]

		case SameVNode(oldEndVnode, newEndVnode):
			p.patchVNode(oldEndVnode, newEndVnode)
			oldEndIdx--
			oldEndVnode = oldCh[oldEndIdx]
			newEndIdx--
			newEndVnode = newCh[newEndIdx]

		case SameVNode(oldStartVnode, newEndVnode):
			// Vnode moved right
			p.patchVNode(oldStartVnode, newEndVnode)
			p.api.InsertBefore(parentElm, oldStartVnode.Elm, p.api.NextSibling(oldEndVnode.Elm))
			oldStartIdx++
			oldStartVnode = oldCh[oldStartIdx]
			newEndIdx--
			newEndVnode = newCh[newEndIdx]

		case SameVNode(oldEndVnode, newStartVnode):
			// Vnode moved left
			p.patchVNode(oldEndVnode, newStartVnode)
			p.api.InsertBefore(parentElm, oldEndVnode.Elm, oldStartVnode.Elm)
			oldEndIdx--
			oldEndVnode = oldCh[oldEndIdx]
			newStartIdx++
			newStartVnode = newCh[newStartIdx]

		default:
			if len(oldKeyToIdx) == 0 {
				oldKeyToIdx = p.createKeyToOldIdx(oldCh, oldStartIdx, oldEndIdx)
			}
			idxInOld, found := oldKeyToIdx[newStartVnode.Key]
			if found {
				elmToMove := oldCh[idxInOld]
				if elmToMove.Sel != newStartVnode.Sel {
					p.api.InsertBefore(parentElm, p.createElm(newStartVnode), oldStartVnode.Elm)
				} else {
					p.patchVNode(elmToMove, newStartVnode)
					oldCh[idxInOld] = nil
					p.api.InsertBefore(parentElm, elmToMove.Elm, oldStartVnode.Elm)
				}
			} else {
				// New element
				p.api.InsertBefore(parentElm, p.createElm(newStartVnode), oldStartVnode.Elm)
			}

			newStartIdx++
			newStartVnode = newCh[newStartIdx]
		}
	}

	if newStartIdx <= newEndIdx {
		var before dom.Node = nil
		if newCh[newEndIdx+1] != nil {
			before = newCh[newEndIdx+1].Elm
		}
		p.addVNodes(parentElm, before, newCh, newStartIdx, newEndIdx)
	}

	if oldStartIdx <= oldEndIdx {
		p.removeVNodes(parentElm, oldCh, oldStartIdx, oldEndIdx)
	}
}

func (p *Patcher) createKeyToOldIdx(children []*VNode, beginIdx, endIdx int) map[string]int {
	m := make(map[string]int)
	for i := beginIdx; i < endIdx; i++ {
		key := children[i].Key
		m[key] = i
	}
	return m
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

func (p *Patcher) addVNodes(parentElm, before dom.Node, vnodes []*VNode, startIdx, endIdx int) {
	for i := startIdx; i <= endIdx; i++ {
		ch := vnodes[i]
		if ch == nil {
			continue
		}
		p.api.InsertBefore(parentElm, p.createElm(ch), before)
	}
}

func (p *Patcher) removeVNodes(parentElm dom.Node, vnodes []*VNode, startIdx, endIdx int) {
	for i := startIdx; i <= endIdx; i++ {
		ch := vnodes[i]
		if ch == nil {
			continue
		}
		p.removeVNodes(ch.Elm, ch.Children, 0, len(ch.Children)-1)
		p.api.RemoveChild(parentElm, ch.Elm)
	}
}
