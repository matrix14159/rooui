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

func a() {
	//data := VNodeData{
	//	Props: Props{
	//		"":  1,
	//		"a": 2,
	//	},
	//	Events: Events{
	//		"click": func(vn VNode, event dom.Event) {
	//
	//		},
	//		"dbclick": func(vn VNode, event dom.Event) {
	//
	//		},
	//	},
	//}
}
