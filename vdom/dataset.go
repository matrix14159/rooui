package vdom

import (
	"fmt"
	"strings"
)

type Dataset map[string]string

func NewDataset() Dataset {
	return make(Dataset)
}

type DatasetModule struct {
}

func NewDatasetModule() *DatasetModule {
	return new(DatasetModule)
}

func (p *DatasetModule) updateDataset(oldVNode, vnode *VNode) {
	oldDataset, isNew1 := getDataset(oldVNode)
	newDataset, isNew2 := getDataset(vnode)
	if isNew1 && isNew2 {
		return
	}

	elm := vnode.Elm

	for key, _ := range oldDataset {
		if _, found := newDataset[key]; !found {
			name := fmt.Sprintf("data-%s", strings.ToLower(key))
			elm.RemoveAttribute(name)
		}
	}

	for key, val := range newDataset {
		if oldDataset[key] != newDataset[key] {
			name := fmt.Sprintf("data-%s", strings.ToLower(key))
			elm.SetAttribute(name, val)
		}
	}
}

func getDataset(vnode *VNode) (dataset Dataset, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.Dataset == nil {
		vnode.Data.Dataset = NewDataset()
		isNew = true
	}
	dataset = vnode.Data.Dataset
	return
}

func (p *DatasetModule) Pre() {
}

func (p *DatasetModule) Create(empty, vnode *VNode) {
	p.updateDataset(empty, vnode)
}

func (p *DatasetModule) Update(oldVNode, vnode *VNode) {
	p.updateDataset(oldVNode, vnode)
}

func (p *DatasetModule) Destroy(vnode *VNode) {
}

func (p *DatasetModule) Remove(vnode *VNode, removeCallback func()) {
	removeCallback()
}

func (p *DatasetModule) Post() {
}
