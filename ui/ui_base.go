package ui

import (
	"fmt"
	"strings"
	"syscall/js"

	dom "honnef.co/go/js/dom/v2"
)

type htmlBaseElement struct {
	id  string
	tag string

	domEl dom.HTMLElement

	// backup style display for turn on/off
	bakDisplay string

	body []UI
	text any

	classes WithClass
	style   WithStyle

	// key: attribute name
	attributes map[string]any

	// key: property name
	properties map[string]any

	// key: event name (example: click)
	events map[string][]htmlEventHandler

	onCreatedHandler   func()
	onMountedHandler   func()
	onUnmountedHandler func()
}

type htmlEventHandler struct {
	name       string                            // event name
	handler    func(event Event, options ...any) // handler function
	useCapture bool
	options    []any
}

func (p htmlEventHandler) listener() func(event Event) {
	return func(event Event) {
		p.handler(event, p.options...)
	}
}

func (p *htmlBaseElement) Kind() Kind {
	return HtmlElem
}

func (p *htmlBaseElement) GetUIElementId() string {
	if p.id == "" {
		p.id = NewId()
	}
	return p.id
}

func (p *htmlBaseElement) Id(val string) HtmlUI {
	p.id = val
	return p
}

func (p *htmlBaseElement) SetUIElementId(id string) {
	p.id = id
}

func (p *htmlBaseElement) getBody() []UI {
	return p.body
}

func (p *htmlBaseElement) GetAttribute(name string) (value any) {
	if p.domEl != nil {
		value = p.domEl.GetAttribute(name)
	}
	return
}

func (p *htmlBaseElement) SetAttribute(name string, value any) {
	if binder, ok := value.(Binder); ok {
		handlerId := p.GetUIElementId() + name
		handler := func(oldVal, newVal any) {
			if p.domEl != nil {
				p.domEl.SetAttribute(name, fmt.Sprintf("%v", newVal))
			}
		}
		binder.AddUpdatedHandler(handler, handlerId)
	}
	if p.attributes == nil {
		p.attributes = make(map[string]any)
	}
	p.attributes[name] = value
}

func (p *htmlBaseElement) SetProperty(name string, value any) {
	if binder, ok := value.(Binder); ok {
		handlerId := p.GetUIElementId() + name
		handler := func(oldVal, newVal any) {
			if p.domEl != nil {
				p.domEl.Underlying().Set(name, newVal)
			} else {
				Console.Error("html element set property {%v:%v} error. id:%v\n", name, value, p.GetUIElementId())
			}
		}
		binder.AddUpdatedHandler(handler, handlerId)
	}
	if p.properties == nil {
		p.properties = make(map[string]any)
	}
	p.properties[name] = value
}

// two way binding for <input/>, <textarea/> value
func (p *htmlBaseElement) bindInputValue(value any) {
	p.SetProperty("value", value)
	if r, ok := value.(*Ref[string]); ok {
		handler := func(event Event, options ...any) {
			r.Set(p.domEl.Underlying().Get("value").String())
		}
		p.registerEventHandler("input", handler)
	}
}

// data-*
func (p *htmlBaseElement) dataWith(name string, value any) {
	n := fmt.Sprintf("data-%s", strings.ToLower(name))
	p.SetAttribute(n, value)
}

func (p *htmlBaseElement) setText(val any) {
	if binder, ok := val.(Binder); ok {
		handlerId := p.GetUIElementId() + "-text"
		handler := func(oldVal, newVal any) {
			w := dom.GetWindow()
			d := w.Document()
			el := d.GetElementByID(p.id)
			el.SetInnerHTML(fmt.Sprintf("%v", newVal))
		}
		binder.AddUpdatedHandler(handler, handlerId)
	}
	p.text = val
}

func (p *htmlBaseElement) setBody(elems ...UI) {
	p.body = append(p.body, elems...)
}

func (p *htmlBaseElement) setStyle(items ...*Ref[StyleItem]) {
	p.style.SetStyle(items...)
	p.style.bindUpdateHandler(p, items...)
}

func (p *htmlBaseElement) setStyleIf(v *Ref[bool], items ...*Ref[StyleItem]) {
	p.style.SetStyleIf(v, items...)
	p.style.bindUpdateHandler(p, items...)
}

func (p *htmlBaseElement) setClass(name string, items ...*Ref[StyleItem]) {
	group := p.classes.SetClass(name, items...)
	p.classes.bindUpdateHandler(p, group)
}

func (p *htmlBaseElement) setClassIf(v *Ref[bool], name string, items ...*Ref[StyleItem]) {
	group := p.classes.SetClassIf(v, name, items...)
	p.classes.bindUpdateHandler(p, group)
}

func (p *htmlBaseElement) registerEventHandler(name string, handler func(event Event, options ...any), options ...any) {
	if p.events == nil {
		p.events = make(map[string][]htmlEventHandler)
	}
	p.events[name] = append(p.events[name], htmlEventHandler{
		name:       name,
		handler:    handler,
		useCapture: false,
		options:    options,
	})
}

func (p *htmlBaseElement) AddEventListener(name string, useCapture bool, handler func(Event, ...any), options ...any) js.Func {
	defer func() {
		if err := recover(); err != nil {
			Console.Error("%v", err)
		}
	}()
	evtHandler := htmlEventHandler{
		name:       name,
		handler:    handler,
		useCapture: useCapture,
		options:    options,
	}
	if p.domEl != nil {
		return p.domEl.AddEventListener(name, useCapture, evtHandler.listener())
	} else {
		panic(fmt.Sprintf("dom element is nil, can't add event listener. id:%v", p.GetUIElementId()))
	}
}

func (p *htmlBaseElement) RemoveEventListener(typ string, useCapture bool, listener js.Func) {
	if p.domEl != nil {
		p.domEl.RemoveEventListener(typ, useCapture, listener)
	}
}

func (p *htmlBaseElement) TurnOnDisplay() {
	if p.domEl != nil {
		if p.bakDisplay != "" && p.bakDisplay != "none" {
			p.domEl.Style().SetProperty("display", p.bakDisplay, "")
		} else {
			if p.domEl.Style().GetPropertyValue("display") == "none" {
				p.domEl.Style().RemoveProperty("display")
			}
		}
	}
}

func (p *htmlBaseElement) TurnOffDisplay() {
	if p.domEl != nil {
		p.bakDisplay = p.domEl.Style().GetPropertyValue("display")
		p.domEl.Style().SetProperty("display", "none", "")
	}
}

func (p *htmlBaseElement) GetSelfDomElement() dom.Element {
	if p.domEl != nil {
		return p.domEl
	}
	w := dom.GetWindow()
	d := w.Document()
	el := d.GetElementByID(p.GetUIElementId())
	if div, ok := el.(dom.HTMLElement); ok {
		p.domEl = div
	}
	return p.domEl
}

func (p *htmlBaseElement) BuildTreeDomElement() []dom.Element {
	if p.tag == "" {
		Console.Error("can't create html element because tag name is empty\n")
		return nil
	}

	//fmt.Printf("%v (%v) BuildTreeDomElement\n", p.tag, p.GetUIElementId())
	w := dom.GetWindow()
	d := w.Document()
	s := d.CreateElement(p.tag).(dom.HTMLElement)

	s.SetAttribute("Id", p.GetUIElementId())

	// set attributes
	for name, val := range p.attributes {
		data := val
		if binder, ok := val.(Binder); ok {
			data = binder.Data()
		}
		s.SetAttribute(name, fmt.Sprintf("%v", data))
	}

	// set properties
	for name, val := range p.properties {
		data := val
		if binder, ok := val.(Binder); ok {
			data = binder.Data()
		}
		s.Underlying().Set(name, data)
	}

	// set style
	for _, one := range p.style.StyleItems {
		item := one.Value()
		s.Style().SetProperty(item.Name(), item.ValueWithUnit(), "")
	}
	// set condition style
	for _, group := range p.style.ConditionStyles {
		if group.Condition.Value() {
			for _, one := range group.StyleItems {
				item := one.Value()
				s.Style().SetProperty(item.Name(), item.ValueWithUnit(), "")
			}
		}
	}

	// set class
	heads := d.GetElementsByTagName("head")
	if len(heads) > 0 {
		for classId, group := range p.classes.Classes {
			if len(group.Items) > 0 {
				st := d.CreateElement("style")
				st.SetAttribute("type", "text/css")
				st.SetID(classId)
				st.SetInnerHTML(group.ToClass())
				heads[0].(*dom.HTMLHeadElement).AppendChild(st)
			}

			if group.Use.Value() && strings.HasPrefix(group.ClassName, ".") {
				cn := strings.TrimLeft(group.ClassName, ".")
				s.Class().Add(cn)
			}
		}
	}

	for name, events := range p.events {
		for _, event := range events {
			s.AddEventListener(name, event.useCapture, event.listener())
		}
	}

	if p.text != nil {
		data := p.text
		if binder, ok := p.text.(Binder); ok {
			data = binder.Data()
		}
		s.SetInnerHTML(fmt.Sprintf("%v", data))
	}

	for _, child := range p.body {
		els := BuildHtml(child)
		for _, el := range els {
			s.AppendChild(el)
		}
	}

	p.domEl = s
	return []dom.Element{s}
}

func (p *htmlBaseElement) doCreated() {
	for _, child := range p.body {
		child.doCreated()
	}
	if p.onCreatedHandler != nil {
		p.onCreatedHandler()
	}
}

func (p *htmlBaseElement) OnCreated(f func()) {
	p.onCreatedHandler = f
}

func (p *htmlBaseElement) doMounted() {
	if p.onMountedHandler != nil {
		p.onMountedHandler()
	}
	for _, child := range p.body {
		child.doMounted()
	}
}

func (p *htmlBaseElement) OnMounted(f func()) {
	p.onMountedHandler = f
}

func (p *htmlBaseElement) doUnmounted() {
	for _, child := range p.body {
		child.doUnmounted()
	}
	p.body = nil

	for name, val := range p.attributes {
		if binder, ok := val.(Binder); ok {
			handlerId := p.GetUIElementId() + name
			binder.RemoveUpdatedHandler(handlerId)
		}
	}
	p.attributes = nil

	for name, val := range p.properties {
		if binder, ok := val.(Binder); ok {
			handlerId := p.GetUIElementId() + name
			binder.RemoveUpdatedHandler(handlerId)
		}
	}
	p.properties = nil

	if p.text != nil {
		if binder, ok := p.text.(Binder); ok {
			handlerId := p.GetUIElementId() + "-text"
			binder.RemoveUpdatedHandler(handlerId)
		}
		p.text = nil
	}

	p.style.clear(p)
	p.classes.clear(p)

	// remove from dom
	w := dom.GetWindow()
	d := w.Document()
	el := d.GetElementByID(p.id)
	if el != nil {
		if parent := el.ParentElement(); parent != nil {
			parent.RemoveChild(el)
			//fmt.Printf("%v (%v) is removed\n", p.tag, p.GetUIElementId())
		}
	}

	if p.onUnmountedHandler != nil {
		p.onUnmountedHandler()
	}
	fmt.Printf("%v (%v) unmounted is done.\n", p.tag, p.GetUIElementId())
}

func (p *htmlBaseElement) OnUnmounted(f func()) {
	p.onUnmountedHandler = f
}
