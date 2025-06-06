package vdom

import (
	"github.com/matrix14159/rooui/dom"
)

type VNode struct {
	Sel      string
	Data     *VNodeData
	Children []*VNode
	Elm      *dom.Object
	Text     string
	Key      string
}

type VNodeChildren []*VNode

type VNodeData struct {
	Props     Props
	Attrs     Attrs
	Class     Classes
	Style     *VNodeStyle
	Dataset   Dataset
	On        *On
	Hooks     *Hooks
	Key       string
	Namespace string // for SVGs
}

func SameVNode(vnode1, vnode2 *VNode) bool {
	sameKey := vnode1.Key == vnode2.Key
	sameSel := vnode1.Sel == vnode2.Sel
	return sameKey && sameSel
}
