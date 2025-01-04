package html

import (
	"fmt"

	"github.com/matrix14159/rooui/vdom"
	"honnef.co/go/js/dom/v2"
)

type ButtonElement struct {
	text string

	events map[string][]vdom.EventHandler
}

func Button() *ButtonElement {
	p := new(ButtonElement)
	p.events = make(map[string][]vdom.EventHandler)
	return p
}

func (p *ButtonElement) Tag() string {
	return "button"
}

func (p *ButtonElement) GetText() string {
	return p.text
}

func (p *ButtonElement) GetEvents() map[string][]vdom.EventHandler {
	return p.events
}

func (p *ButtonElement) Text(text any) *ButtonElement {
	p.text = fmt.Sprintf("%v", text)
	return p
}

func (p *ButtonElement) OnClick(f func(event dom.Event, options ...any), options ...any) *ButtonElement {
	handlers := p.events["click"]
	handlers = append(handlers, vdom.EventHandler{
		Options: options,
		Handler: f,
	})
	p.events["click"] = handlers
	return p
}
