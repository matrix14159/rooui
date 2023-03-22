package hselect

import (
	"github.com/matrix14159/rooui/ui"
)

type Select struct {
	ui.Component

	content ui.WithContent
	class   ui.WithClass
	style   ui.WithStyle

	value   *ui.Ref[string]
	options *ui.RefSlice[Option]

	onChanged func()
}

type Option struct {
	Label    string
	Value    string
	Disabled bool
}

func New() *Select {
	p := &Select{
		value:   ui.ToRef(""),
		options: ui.ToRefSlice([]Option{}),
	}
	return p
}

func (p *Select) Body(elems ...ui.UI) *Select {
	p.content.AppendContent(elems...)
	return p
}

func (p *Select) Style(items ...*ui.Ref[ui.StyleItem]) *Select {
	p.style.SetStyle(items...)
	return p
}

func (p *Select) Class(name string, items ...*ui.Ref[ui.StyleItem]) *Select {
	p.class.SetClass(name, items...)
	return p
}

func (p *Select) Options(opts *ui.RefSlice[Option]) *Select {
	if opts == nil {
		opts = ui.ToRefSlice([]Option{})
	}
	p.options = opts
	return p
}

func (p *Select) Value(val *ui.Ref[string]) *Select {
	p.value = val
	return p
}

func (p *Select) handleChanged(event ui.Event) {
	defer func() {
		if err := recover(); err != nil {
			ui.Console.Error("select handle changed error:%v\n", err)
		}
	}()
	idx := event.Target().Underlying().Get("selectedIndex").Int()
	p.value.Set(p.options.Get(idx).Value)

	if p.onChanged != nil {
		p.onChanged()
	}
}

func (p *Select) OnChanged(handler func()) *Select {
	p.onChanged = handler
	return p
}

func (p *Select) isOptionSelected(index int) bool {
	if index < 0 || index >= p.options.Len() {
		return false
	}
	return p.options.Get(index).Value == p.value.Value()
}

func (p *Select) Render() ui.UI {
	ret := ui.Select().
		Style(p.style.StyleItems...).
		Body(
			ui.Slice(p.options, func(i *ui.Ref[int], v *ui.Ref[Option]) ui.UI {
				return ui.Option().
					Value(v.Value().Value).
					Text(v.Value().Label).
					Disabled(v.Value().Disabled).
					Selected(p.isOptionSelected(i.Value()))
			}),
		).
		Body(p.content.Content...).
		OnChange(p.handleChanged)

	for _, group := range p.class.Classes {
		ret.Class(group.ClassName, group.Items...)
	}
	return ret
}
