package ui

import (
	dom "honnef.co/go/js/dom/v2"
)

type WithClass struct {
	// key: StyleGroup.ClassId
	Classes map[string]StyleGroup
}

func (p *WithClass) SetClass(name string, items ...*Ref[StyleItem]) (group StyleGroup) {
	if name == "" {
		Console.Error("class name is empty")
		return
	}

	group = StyleGroup{
		ClassId:   NewId(),
		ClassName: name,
		Items:     items,
	}

	if p.Classes == nil {
		p.Classes = make(map[string]StyleGroup)
	}
	p.Classes[group.ClassId] = group
	return
}

func (p *WithClass) bindUpdateHandler(target UI, group StyleGroup) {
	if group.ClassName == "" {
		Console.Error("class name is empty")
		return
	}

	rUpdate := func(idx int, newVal any) {
		w := dom.GetWindow()
		d := w.Document()
		el := d.GetElementByID(group.ClassId)
		if el != nil {
			st := el.(*dom.HTMLStyleElement)
			tmp := p.Classes[group.ClassId]
			s := newVal.(StyleItem)
			tmp.Items[idx].data = s
			st.SetInnerHTML(tmp.ToClass())
		}
	}

	for i, item := range group.Items {
		idx := i
		handlerId := target.GetUIElementId() + item.data.Name() + item.data.Id()
		item.AddUpdatedHandler(func(oldVal, newVal any) {
			rUpdate(idx, newVal)
		}, handlerId)
	}
}

func (p *WithClass) unbindUpdateHandler(target UI) {
	for _, group := range p.Classes {
		for _, one := range group.Items {
			one.Value().RemoveValueUpdatedHandler()
			one.RemoveUpdatedHandler(target.GetUIElementId() + one.data.Name() + one.data.Id())
		}
		group.Items = nil
	}
	p.Classes = nil
}
