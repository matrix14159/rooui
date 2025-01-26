package html

import (
	"honnef.co/go/js/dom/v2"
)

type ButtonElement struct {
	BaseElement
}

func Button() *ButtonElement {
	p := new(ButtonElement)
	p.init()
	return p
}

func (p *ButtonElement) Tag() string {
	return "button"
}

func (p *ButtonElement) Text(text any) *ButtonElement {
	p.BaseElement.Text(text)
	return p
}

func (p *ButtonElement) Body(child ...Element) *ButtonElement {
	p.BaseElement.Body(child...)
	return p
}

func (p *ButtonElement) OnClick(f func(event dom.Event, options ...any), options ...any) *ButtonElement {
	p.BaseElement.OnClick(f, options)
	return p
}
