package flex

import (
	"github.com/matrix14159/rooui/ui"
	"github.com/matrix14159/rooui/ui/st"
)

type Flex struct {
	ui.Component

	content ui.WithContent
	class   ui.WithClass
	style   ui.WithStyle

	direction string
}

func New() *Flex {
	p := &Flex{
		direction: "row",
	}
	return p
}

func (p *Flex) Direction(column bool) *Flex {
	p.direction = "row"
	if column {
		p.direction = "column"
	}
	return p
}

func (p *Flex) Body(elems ...ui.UI) *Flex {
	p.content.AppendContent(elems...)
	return p
}

func (p *Flex) Style(items ...*ui.Ref[ui.StyleItem]) *Flex {
	p.style.SetStyle(items...)
	return p
}

func (p *Flex) Class(name string, items ...*ui.Ref[ui.StyleItem]) *Flex {
	p.class.SetClass(name, items...)
	return p
}

func (p *Flex) Render() ui.UI {
	ret := ui.Div().
		Style(
			st.Display("flex"),
			st.FlexDirection(p.direction),
		).
		Style(
			p.style.StyleItems...,
		).
		Body(p.content.Content...)

	for _, group := range p.class.Classes {
		ret.Class(group.ClassName, group.Items...)
	}
	return ret
}
