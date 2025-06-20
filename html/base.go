package html

import (
	"fmt"
	"strings"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/dom"
	"github.com/matrix14159/rooui/vdom"
)

type Element interface {
	Tag() string

	GetText() string

	GetEvents() map[string][]vdom.EventHandler

	GetBody() []Element

	GetClasses() []string

	GetStyles() []css.Style
}

type BaseElement struct {
	tag string

	text string

	body []Element

	classes []string

	styles []css.Style

	events map[string][]vdom.EventHandler
}

func InitBaseElement(tag string) BaseElement {
	p := BaseElement{tag: tag}
	p.events = make(map[string][]vdom.EventHandler)
	return p
}

func (p *BaseElement) Tag() string {
	return p.tag
}

func (p *BaseElement) GetText() string {
	return p.text
}

func (p *BaseElement) GetBody() []Element {
	return p.body
}

func (p *BaseElement) GetClasses() []string {
	return p.classes
}

func (p *BaseElement) GetStyles() []css.Style {
	return p.styles
}

func (p *BaseElement) GetEvents() map[string][]vdom.EventHandler {
	return p.events
}

func (p *BaseElement) Body(child ...Element) *BaseElement {
	p.body = append(p.body, child...)
	return p
}

func (p *BaseElement) Class(name ...string) *BaseElement {
	for _, one := range name {
		p.classes = append(p.classes, strings.TrimLeft(strings.TrimSpace(one), "."))
	}
	return p
}

func (p *BaseElement) Style(style ...css.Style) *BaseElement {
	for _, one := range style {
		p.styles = append(p.styles, one)
	}
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
