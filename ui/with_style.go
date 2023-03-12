package ui

import (
	dom "honnef.co/go/js/dom/v2"
)

type WithStyle struct {
	StyleItems []*Ref[StyleItem]
}

func (p *WithStyle) SetStyle(items ...*Ref[StyleItem]) {
	p.StyleItems = append(p.StyleItems, items...)
}

func (p *WithStyle) bindUpdateHandler(target UI, items ...*Ref[StyleItem]) {
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

func (p *WithStyle) unbindUpdateHandler(target UI) {
	for _, one := range p.StyleItems {
		one.Value().RemoveValueUpdatedHandler()
		one.RemoveUpdatedHandler(target.GetUIElementId() + one.data.Name() + one.data.Id())
	}
	p.StyleItems = nil
}
