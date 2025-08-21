package dom

import (
	"syscall/js"
)

type NodeList struct {
	js.Value
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/NodeList/length
func (p *NodeList) Length() int {
	return p.Get("length").Int()
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/NodeList/item
// return object as Element
func (p *NodeList) Item(index int) *Object {
	return &Object{p.Call("item", index)}
}
