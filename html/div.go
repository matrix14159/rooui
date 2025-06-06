package html

import (
	"github.com/matrix14159/rooui/dom"
)

type DivElement struct {
	BaseElement
}

func Div() *DivElement {
	p := new(DivElement)
	p.init()
	return p
}

func (p *DivElement) Tag() string {
	return "div"
}

func (p *DivElement) Text(text any) *DivElement {
	p.BaseElement.Text(text)
	return p
}

func (p *DivElement) Body(child ...Element) *DivElement {
	p.BaseElement.Body(child...)
	return p
}

func (p *DivElement) Class(name ...string) *DivElement {
	p.BaseElement.Class(name...)
	return p
}

func (p *DivElement) OnClick(f func(event dom.Event, options ...any), options ...any) *DivElement {
	p.BaseElement.OnClick(f, options)
	return p
}
