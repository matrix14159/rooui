//go:build wasm

package main

import (
	"github.com/matrix14159/rooui/ui"
	"github.com/matrix14159/rooui/ui/st"
)

func main() {
	ui.Init(NewDemo())
	ui.WaitExit()
}

type Demo struct {
	ui.Component

	active *ui.Ref[string]
	tabs   *ui.RefSlice[string]
}

func NewDemo() *Demo {
	p := &Demo{
		active: ui.ToRef("hello"),
	}

	tabs := []string{"hello", "world"}
	p.tabs = ui.ToRefSlice(tabs)
	return p
}

func (p *Demo) isActive(name string) *ui.Ref[bool] {
	return ui.Effect(p.active, func(active string) bool {
		return name == active
	})
}

func (p *Demo) setActiveTab(event ui.Event) {
	if p.active.Value() == "hello" {
		p.active.Set("world")
	} else {
		p.active.Set("hello")
	}
}

func (p *Demo) Render() ui.UI {
	return ui.Div().Style(
		st.Display("flex"),
		st.ColumnGap("5px"),
	).Body(
		ui.Slice(p.tabs, func(i *ui.Ref[int], v *ui.Ref[string]) ui.UI {
			return ui.Div().StyleIf(p.isActive(v.Value()),
				st.Color("red"),
			).Body(
				ui.Span().Text(v),
			).OnClick(p.setActiveTab)
		}),
	)
}
