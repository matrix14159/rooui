//go:build wasm

package main

import (
	"github.com/matrix14159/rooui/cs/hselect"
	"github.com/matrix14159/rooui/ui"
)

func main() {
	ui.Init(NewDemo())
	ui.WaitExit()
}

type Demo struct {
	ui.Component

	value   *ui.Ref[string]
	options *ui.RefSlice[hselect.Option]
}

func NewDemo() *Demo {
	p := &Demo{
		value: ui.ToRef("1"),
	}

	opts := []hselect.Option{
		{Label: "unset", Value: "", Disabled: true},
		{Label: "hello", Value: "1"},
		{Label: "world", Value: "2"},
	}
	p.options = ui.ToRefSlice(opts)
	return p
}

func (p *Demo) onSelectChanged() {
	ui.Console.Log("on changed. value:%v\n", p.value.Value())
}

func (p *Demo) Render() ui.UI {
	return hselect.New().Value(p.value).Options(p.options).OnChanged(p.onSelectChanged)
}
