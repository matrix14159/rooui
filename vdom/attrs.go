package vdom

import (
	"fmt"
	"reflect"

	"honnef.co/go/js/dom/v2"
)

const xlinkNS = "http://www.w3.org/1999/xlink"
const xmlnsNS = "http://www.w3.org/2000/xmlns/"
const xmlNS = "http://www.w3.org/XML/1998/namespace"

type Attrs map[string]any // type of value must be one of the bool, number, string

func NewAttributes() Attrs {
	return make(Attrs)
}

type AttrModule struct {
}

func (p *AttrModule) updateAttrs(oldVNode, vnode *VNode) {
	oldAttrs, isNew1 := getAttrs(oldVNode)
	newAttrs, isNew2 := getAttrs(vnode)
	if isNew1 && isNew2 {
		return
	}

	elm := vnode.Elm.(dom.HTMLElement)

	// update modified attributes, add new attributes
	for key, val := range newAttrs {
		oldVal := oldAttrs[key]
		if reflect.DeepEqual(val, oldVal) {
			continue
		}

		switch val.(type) {
		case bool:
			if b := val.(bool); b {
				elm.SetAttribute(key, "")
			} else {
				elm.RemoveAttribute(key)
			}
		case string:
			s := val.(string)
			if len(s) > 0 && s[0] != 'x' {
				elm.SetAttribute(key, s)
			} else if len(s) >= 3 && s[3] == ':' {
				// Assume xml namespace
				elm.SetAttributeNS(xmlNS, key, s)
			} else if len(s) >= 5 && s[5] == ':' {
				// Assume 'xmlns' or 'xlink' namespace
				if s[0] == 'm' {
					elm.SetAttributeNS(xmlnsNS, key, s)
				} else {
					elm.SetAttributeNS(xlinkNS, key, s)
				}
			} else {
				elm.SetAttribute(key, s)
			}
		default:
			elm.SetAttribute(key, fmt.Sprintf("%v", val))
		}
	}

	// remove removed attributes
	for key, _ := range oldAttrs {
		if _, found := newAttrs[key]; !found {
			elm.RemoveAttribute(key)
		}
	}
}

func getAttrs(vnode *VNode) (attrs Attrs, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.Attrs == nil {
		vnode.Data.Attrs = NewAttributes()
		isNew = true
	}
	attrs = vnode.Data.Attrs
	return
}

func (p *AttrModule) Pre() {
}

func (p *AttrModule) Create(empty, vnode *VNode) {
	p.updateAttrs(empty, vnode)
}

func (p *AttrModule) Update(oldVNode, vnode *VNode) {
	p.updateAttrs(oldVNode, vnode)
}

func (p *AttrModule) Destroy(vnode *VNode) {
}

func (p *AttrModule) Remove(vnode *VNode, removeCallback func()) {
	removeCallback()
}

func (p *AttrModule) Post() {
}
