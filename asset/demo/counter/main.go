//go:build wasm

package main

import (
	"github.com/matrix14159/rooui/ui"
)

func main() {
	ui.Init(NewCounter())
	ui.WaitExit()
}

type Counter struct {
	ui.Component

	count *ui.Ref[int]
}

func NewCounter() *Counter {
	p := &Counter{
		count: ui.ToRef(0),
	}
	return p
}

func (p *Counter) increase(event ui.Event, options ...any) {
	p.count.Set(p.count.Value() + 1)
}

func (p *Counter) Render() ui.UI {
	return ui.Button().Text(p.count).OnClick(p.increase)
}
