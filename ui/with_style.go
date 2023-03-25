package ui

import (
	dom "honnef.co/go/js/dom/v2"
)

type WithStyle struct {
	StyleItems []*Ref[StyleItem]

	ConditionStyles []conditionStyle

	target UI
}

type conditionStyle struct {
	Condition  *Ref[bool]
	StyleItems []*Ref[StyleItem]
}

func (p *WithStyle) SetStyle(items ...*Ref[StyleItem]) {
	p.StyleItems = append(p.StyleItems, items...)
}

func (p *WithStyle) SetStyleIf(v *Ref[bool], items ...*Ref[StyleItem]) {
	handler := func(oldVal, newVal any) {
		if p.target == nil {
			return
		}

		if newVal.(bool) {
			p.bindUpdateHandler(p.target, items...)
		} else {
			p.unbindUpdateHandler(p.target, items...)
		}
	}
	v.AddUpdatedHandler(handler)
	p.ConditionStyles = append(p.ConditionStyles, conditionStyle{Condition: v, StyleItems: items})
}

func (p *WithStyle) bindUpdateHandler(target UI, items ...*Ref[StyleItem]) {
	p.target = target
	for _, item := range items {
		handlerId := target.GetUIElementId() + item.data.Name() + item.data.Id()
		handler := func(oldVal, newVal any) {
			s := newVal.(StyleItem)
			w := dom.GetWindow()
			d := w.Document()
			el := d.GetElementByID(target.GetUIElementId()).(dom.HTMLElement)
			if s.Value() != "" {
				el.Style().SetProperty(s.Name(), s.ValueWithUnit(), "")
			} else {
				el.Style().RemoveProperty(s.Name())
			}
		}
		item.AddUpdatedHandler(handler, handlerId)
	}
}

func (p *WithStyle) unbindUpdateHandler(target UI, items ...*Ref[StyleItem]) {
	w := dom.GetWindow()
	d := w.Document()
	el := d.GetElementByID(target.GetUIElementId()).(dom.HTMLElement)
	for _, one := range items {
		one.Value().RemoveValueUpdatedHandler()
		one.RemoveUpdatedHandler(target.GetUIElementId() + one.data.Name() + one.data.Id())
		el.Style().RemoveProperty(one.Value().Name())
	}
}

func (p *WithStyle) clear(target UI) {
	p.unbindUpdateHandler(target, p.StyleItems...)
	for _, group := range p.ConditionStyles {
		p.unbindUpdateHandler(target, group.StyleItems...)
	}
	p.StyleItems = nil
}
