package html

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/dom"
	"github.com/matrix14159/rooui/vdom"
)

type Element interface {
	Tag() string

	GetId() string

	GetText() string

	GetEvents() map[string][]vdom.EventHandler

	GetAttributes() map[string]any

	GetBody() []Element

	GetProps() map[string]any

	GetClasses() []string

	GetStyles() []css.Style

	GetClassStyle() map[string][]css.Style

	ReplaceChild(index int, el Element)
}

type BaseElement struct {
	tag string

	id string

	text string

	body []Element

	attributes map[string]any

	props map[string]any

	classes []string

	styles []css.Style

	classStyle map[string][]css.Style

	events map[string][]vdom.EventHandler
}

func InitBaseElement(tag string) BaseElement {
	p := BaseElement{tag: tag}
	p.attributes = make(map[string]any)
	p.props = make(map[string]any)
	p.events = make(map[string][]vdom.EventHandler)
	return p
}

func (p *BaseElement) Tag() string {
	return p.tag
}

func (p *BaseElement) GetId() string {
	return p.id
}

func (p *BaseElement) GetText() string {
	return p.text
}

func (p *BaseElement) GetBody() []Element {
	return p.body
}

func (p *BaseElement) GetAttributes() map[string]any {
	return p.attributes
}

func (p *BaseElement) GetProps() map[string]any {
	return p.props
}

func (p *BaseElement) GetClasses() []string {
	return p.classes
}

func (p *BaseElement) GetStyles() []css.Style {
	return p.styles
}

func (p *BaseElement) GetClassStyle() map[string][]css.Style {
	return p.classStyle
}

func (p *BaseElement) GetEvents() map[string][]vdom.EventHandler {
	return p.events
}

func (p *BaseElement) ReplaceChild(index int, el Element) {
	p.body[index] = el
}

func (p *BaseElement) Id(id string) *BaseElement {
	p.id = id
	return p
}

func (p *BaseElement) Body(child ...Element) *BaseElement {
	p.body = append(p.body, child...)
	return p
}

func (p *BaseElement) Attribute(name string, value any) *BaseElement {
	p.attributes[name] = value
	return p
}

func (p *BaseElement) Prop(name string, value any) *BaseElement {
	p.props[name] = value
	return p
}

func (p *BaseElement) Classes(name ...string) *BaseElement {
	for _, one := range name {
		if one == "" {
			continue
		}
		p.classes = append(p.classes, strings.TrimPrefix(strings.TrimSpace(one), "."))
	}
	return p
}

func (p *BaseElement) Class(name string, styles ...css.Style) *BaseElement {
	if p.classStyle == nil {
		p.classStyle = make(map[string][]css.Style)
	}
	if name == "" && p.id != "" {
		name = p.id
	}
	cls := strings.TrimPrefix(strings.TrimSpace(name), ".")
	p.classStyle[cls] = styles
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

// SetStyleProperty set style with name and v after render
// p must render with id and should void call in Component.Render
func (p *BaseElement) SetStyleProperty(name string, v string) {
	if p == nil || p.GetId() == "" {
		slog.Warn("can't find element: nil or id empty")
		return
	}
	el := dom.Document.GetElementById(p.GetId())
	el.Style().SetProperty(name, v)
}
