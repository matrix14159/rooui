package html

import (
	"fmt"

	"github.com/matrix14159/rooui/vdom"
	"honnef.co/go/js/dom/v2"
)

type Element interface {
	Tag() string

	GetText() string

	GetEvents() map[string][]vdom.EventHandler

	GetBody() []Element
}

type BaseElement struct {
	text string

	body []Element

	events map[string][]vdom.EventHandler
}

func BaseHtmlElement() *BaseElement {
	p := new(BaseElement)
	p.init()
	return p
}

func (p *BaseElement) init() {
	p.events = make(map[string][]vdom.EventHandler)
}

func (p *BaseElement) Tag() string {
	return "button"
}

func (p *BaseElement) GetText() string {
	return p.text
}

func (p *BaseElement) GetBody() []Element {
	return p.body
}

func (p *BaseElement) GetEvents() map[string][]vdom.EventHandler {
	return p.events
}

func (p *BaseElement) Body(child ...Element) *BaseElement {
	p.body = append(p.body, child...)
	return p
}

func (p *BaseElement) Text(text any) *BaseElement {
	p.text = fmt.Sprintf("%v", text)
	return p
}

func (p *BaseElement) OnClick(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	handlers := p.events["click"]
	handlers = append(handlers, vdom.EventHandler{
		Options: options,
		Handler: f,
	})
	p.events["click"] = handlers
	return p
}
