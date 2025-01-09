package vdom

import (
	"log/slog"
	"math"
	"strings"

	"honnef.co/go/js/dom/v2"
)

type Patcher struct {
	api DOMAPI

	modules []Module

	cbs HookGroup
}

// NewPatcher create a patcher for dom's elm
func NewPatcher(api DOMAPI, module ...Module) *Patcher {
	p := &Patcher{
		api:     api,
		modules: module,
	}
	for _, m := range p.modules {
		p.cbs.Pres = append(p.cbs.Pres, m.Pre)
		p.cbs.Creates = append(p.cbs.Creates, m.Create)
		p.cbs.Updates = append(p.cbs.Updates, m.Update)
		p.cbs.Destroys = append(p.cbs.Destroys, m.Destroy)
		p.cbs.Removes = append(p.cbs.Removes, m.Remove)
		p.cbs.Posts = append(p.cbs.Posts, m.Post)
	}
	return p
}

func (p *Patcher) Patch(oldVnode, vnode *VNode) (old *VNode, err error) {
	for _, pre := range p.cbs.Pres {
		pre()
	}
	defer func() {
		for _, post := range p.cbs.Posts {
			post()
		}
	}()

	if SameVNode(oldVnode, vnode) {
		p.patchVNode(oldVnode, vnode)
	} else {
		elm := oldVnode.Elm
		parent := p.api.ParentNode(elm)

		vnode.Elm = p.createElm(vnode)

		slog.Info("not the same vnode, replace current element",
			"oldVnode", *oldVnode, "parent", parent.NodeName(), "newNode", vnode)

		if parent != nil {
			p.api.InsertBefore(parent, vnode.Elm, p.api.NextSibling(vnode.Elm))
			p.removeVNodes(parent, []*VNode{oldVnode}, 0, 0)
		}
	}
	old = vnode
	return
}

func (p *Patcher) patchVNode(oldVnode, vnode *VNode) {
	vnode.Elm = oldVnode.Elm
	for _, update := range p.cbs.Updates {
		update(oldVnode, vnode)
	}

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
		p.api.SetTextContent(vnode.Elm, vnode.Text)
	}
}

func (p *Patcher) updateChildren(parentElm dom.Node, oldCh, newCh []*VNode) {
	oldStartIdx := 0
	oldEndIdx := len(oldCh) - 1
	newStartIdx := 0
	newEndIdx := len(newCh) - 1

	//slog.Info("1.updateChildren", "oldStartIdx", oldStartIdx, "oldEndIdx", oldEndIdx, "oldCh.len", len(oldCh))
	//slog.Info("2.updateChildren", "newStartIdx", newStartIdx, "newEndIdx", newEndIdx, "newCh.len", len(newCh))

	var oldStartVnode *VNode
	var oldEndVnode *VNode
	var newStartVnode *VNode
	var newEndVnode *VNode

	oldKeyToIdx := make(map[string]int)

	for oldStartIdx <= oldEndIdx && newStartIdx <= newEndIdx {
		oldStartVnode = oldCh[oldStartIdx]
		oldEndVnode = oldCh[oldEndIdx]
		newStartVnode = newCh[newStartIdx]
		newEndVnode = newCh[newEndIdx]

		switch {
		case oldStartVnode == nil: // Vnode might have been moved left
			oldStartIdx++

		case oldEndVnode == nil:
			oldEndIdx--

		case newStartVnode == nil:
			newStartIdx++

		case newEndVnode == nil:
			newEndIdx--

		case SameVNode(oldStartVnode, newStartVnode):
			p.patchVNode(oldStartVnode, newStartVnode)
			oldStartIdx++
			newStartIdx++

		case SameVNode(oldEndVnode, newEndVnode):
			p.patchVNode(oldEndVnode, newEndVnode)
			oldEndIdx--
			newEndIdx--

		case SameVNode(oldStartVnode, newEndVnode):
			// Vnode moved right
			p.patchVNode(oldStartVnode, newEndVnode)
			p.api.InsertBefore(parentElm, oldStartVnode.Elm, p.api.NextSibling(oldEndVnode.Elm))
			oldStartIdx++
			newEndIdx--

		case SameVNode(oldEndVnode, newStartVnode):
			// Vnode moved left
			p.patchVNode(oldEndVnode, newStartVnode)
			p.api.InsertBefore(parentElm, oldEndVnode.Elm, oldStartVnode.Elm)
			oldEndIdx--
			newStartIdx++

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
		}
	}

	//slog.Info("3.updateChildren", "oldStartIdx", oldStartIdx, "oldEndIdx", oldEndIdx, "oldCh.len", len(oldCh))
	//slog.Info("4.updateChildren", "newStartIdx", newStartIdx, "newEndIdx", newEndIdx, "newCh.len", len(newCh))

	if newStartIdx <= newEndIdx {
		var before dom.Node = nil
		if newCh[newEndIdx] != nil {
			before = newCh[newEndIdx].Elm
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

	for _, create := range p.cbs.Creates {
		create(emptyNode, vnode)
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

		switch {
		case ch.Sel != "":
			p.invokeDestroyHook(ch)
			listeners := len(p.cbs.Removes) + 1
			rm := p.createRmCb(ch.Elm, listeners)
			for _, remove := range p.cbs.Removes {
				remove(ch, rm)
			}
			rm()
			//slog.Info("removeVNodes.", slog.Any("listeners", listeners))

		case len(ch.Children) > 0:
			p.invokeDestroyHook(ch)
			p.removeVNodes(parentElm, ch.Children, 0, len(ch.Children)-1)

		default:
			// text node
			p.api.RemoveChild(parentElm, ch.Elm)
		}
	}
}

func (p *Patcher) createRmCb(childElm dom.Node, listeners int) func() {
	return func() {
		listeners--
		if listeners == 0 {
			parent := p.api.ParentNode(childElm)
			p.api.RemoveChild(parent, childElm)
		}
	}
}

func (p *Patcher) invokeDestroyHook(vnode *VNode) {

}
