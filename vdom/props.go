package vdom

import (
	"reflect"

	"honnef.co/go/js/dom/v2"
)

type Props map[string]any

func NewProps() Props {
	return make(Props)
}

type PropsModule struct {
}

func (p *PropsModule) updateProps(oldVNode, vnode *VNode) {
	oldProps, isNew1 := getProps(oldVNode)
	newProps, isNew2 := getProps(vnode)
	if isNew1 && isNew2 {
		return
	}

	elm := vnode.Elm.(dom.HTMLElement)
	for key, val := range newProps {
		oldVal := oldProps[key]
		if reflect.DeepEqual(val, oldVal) {
			continue
		}
		if key != "value" {
			elm.Underlying().Set(key, val)
		}
	}

	// remove removed props
	for key, _ := range oldProps {
		if _, found := newProps[key]; !found {
			elm.Underlying().Delete(key)
		}
	}
}

func getProps(vnode *VNode) (props Props, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.Props == nil {
		vnode.Data.Props = NewProps()
		isNew = true
	}
	props = vnode.Data.Props
	return
}

func (p *PropsModule) Pre() {
}

func (p *PropsModule) Create(empty, vnode *VNode) {
	p.updateProps(empty, vnode)
}

func (p *PropsModule) Update(oldVNode, vnode *VNode) {
	p.updateProps(oldVNode, vnode)
}

func (p *PropsModule) Destroy(vnode *VNode) {
}

func (p *PropsModule) Remove(vnode *VNode, removeCallback func()) {
	removeCallback()
}

func (p *PropsModule) Post() {
}
