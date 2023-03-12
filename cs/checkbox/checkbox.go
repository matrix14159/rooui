package checkbox

import (
	"sunny/rooui/ui"
	"sunny/rooui/ui/st"
)

type Checkbox struct {
	ui.Component

	inputId string

	checked *ui.Ref[bool]
	label   any
}

func New(checked *ui.Ref[bool]) *Checkbox {
	p := &Checkbox{
		inputId: ui.NewId(),
		checked: checked,
	}
	return p
}

// Label set the text for checkbox
// text should be a string or *ui.Ref[string]
func (p *Checkbox) Label(text any) *Checkbox {
	p.label = text
	return p
}

func (p *Checkbox) Render() ui.UI {
	return ui.Div().Style(
		st.Display("flex"),
	).Body(
		ui.Input().Id(p.inputId).Type("checkbox").Checked(p.checked),
		ui.Label().For(p.inputId).Text(p.label),
	)
}
