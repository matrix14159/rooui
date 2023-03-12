package ui

import (
	"fmt"

	dom "honnef.co/go/js/dom/v2"
)

type baseController struct {
	id string

	body []UI

	onCreatedHandler   func()
	onMountedHandler   func()
	onUnmountedHandler func()
}

func (p *baseController) Kind() Kind {
	return ControllerElem
}

func (p *baseController) GetUIElementId() string {
	if p.id == "" {
		p.id = NewId()
	}
	return p.id
}

func (p *baseController) SetUIElementId(id string) {
	p.id = id
}

func (p *baseController) getBody() []UI {
	return p.body
}

func (p *baseController) GetSelfDomElement() dom.Element {
	for _, child := range p.body {
		if el := child.GetSelfDomElement(); el != nil {
			return el
		}
	}
	return nil
}

func (p *baseController) BuildTreeDomElement() []dom.Element {
	ret := make([]dom.Element, 0, len(p.body))
	for _, child := range p.body {
		els := BuildHtml(child)
		ret = append(ret, els...)
	}
	return ret
}

func (p *baseController) TurnOnDisplay() {
	for _, child := range p.body {
		child.TurnOnDisplay()
	}
}

func (p *baseController) TurnOffDisplay() {
	for _, child := range p.body {
		child.TurnOffDisplay()
	}
}

func (p *baseController) doCreated() {
	for _, child := range p.body {
		child.doCreated()
	}
	if p.onCreatedHandler != nil {
		p.onCreatedHandler()
	}
}

func (p *baseController) OnCreated(f func()) {
	p.onCreatedHandler = f
}

func (p *baseController) doMounted() {
	if p.onMountedHandler != nil {
		p.onMountedHandler()
	}
	for _, child := range p.body {
		child.doMounted()
	}
}

func (p *baseController) OnMounted(f func()) {
	p.onMountedHandler = f
}

func (p *baseController) doUnmounted() {
	for _, child := range p.body {
		child.doUnmounted()
	}
	if p.onUnmountedHandler != nil {
		p.onUnmountedHandler()
	}
	fmt.Printf("baseController (%v) unmounted is done.\n", p.GetUIElementId())
}

func (p *baseController) OnUnmounted(f func()) {
	p.onUnmountedHandler = f
}
