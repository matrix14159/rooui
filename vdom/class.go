package vdom

import (
	"maps"
)

type Classes map[string]bool

func NewClasses() Classes {
	return make(Classes)
}

type ClassModule struct {
}

func NewClassModule() *ClassModule {
	return new(ClassModule)
}

func (p *ClassModule) updateClass(oldVNode, vnode *VNode) {
	oldClass, isNew1 := getClass(oldVNode)
	newClass, isNew2 := getClass(vnode)
	if isNew1 && isNew2 {
		return
	}
	if maps.Equal(oldClass, newClass) {
		return
	}

	elm := vnode.Elm

	for name, val := range oldClass {
		if val {
			if _, ok := newClass[name]; !ok {
				elm.ClassList().Remove(name)
			}
		}
	}
	for name, val := range newClass {
		if val == oldClass[name] {
			continue
		}
		if val {
			elm.ClassList().Add(name)
		} else {
			elm.ClassList().Remove(name)
		}
	}
}

func getClass(vnode *VNode) (cls Classes, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.Class == nil {
		vnode.Data.Class = NewClasses()
		isNew = true
	}
	cls = vnode.Data.Class
	return
}

func (p *ClassModule) Pre() {
}

func (p *ClassModule) Create(empty, vnode *VNode) {
	p.updateClass(empty, vnode)
}

func (p *ClassModule) Update(oldVNode, vnode *VNode) {
	p.updateClass(oldVNode, vnode)
}

func (p *ClassModule) Destroy(vnode *VNode) {
}

func (p *ClassModule) Remove(vnode *VNode, removeCallback func()) {
	removeCallback()
}

func (p *ClassModule) Post() {
}
