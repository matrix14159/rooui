package vdom

import (
	"honnef.co/go/js/dom/v2"
)

type VNode struct {
	Sel      string
	Data     VNodeData
	Children []VNode
	Elm      dom.Node
	Text     string
	Key      string
}

type VNodeData struct {
	Props     Props
	Attrs     Attrs
	Class     Classes
	Style     VNodeStyle
	Dataset   Dataset
	Events    Events
	Key       string
	Namespace string // for SVGs
}

func SameVNode(vnode1, vnode2 VNode) bool {
	sameKey := vnode1.Key == vnode2.Key
	sameSel := vnode1.Sel == vnode2.Sel
	return sameKey && sameSel
}
