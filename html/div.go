package html

import (
	"fmt"

	"github.com/matrix14159/rooui/core"
	"github.com/matrix14159/rooui/vdom"
	"honnef.co/go/js/dom/v2"
)

type DivElement struct {
	text string

	body []core.Element

	events map[string][]vdom.EventHandler
}

func Div() *DivElement {
	p := new(DivElement)
	p.events = make(map[string][]vdom.EventHandler)
	return p
}

func (p *DivElement) Tag() string {
	return "div"
}

func (p *DivElement) GetText() string {
	return p.text
}

func (p *DivElement) GetEvents() map[string][]vdom.EventHandler {
	return p.events
}

func (p *DivElement) GetBody() []core.Element {
	return p.body
}

func (p *DivElement) Text(text any) *DivElement {
	p.text = fmt.Sprintf("%v", text)
	return p
}

func (p *DivElement) Body(child ...core.Element) *DivElement {
	p.body = child
	return p
}

func (p *DivElement) OnClick(f func(event dom.Event, options ...any), options ...any) *DivElement {
	handlers := p.events["click"]
	handlers = append(handlers, vdom.EventHandler{
		Options: options,
		Handler: f,
	})
	p.events["click"] = handlers
	return p
}
