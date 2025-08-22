package html

import (
	"log/slog"

	"github.com/matrix14159/rooui/dom"
	"github.com/matrix14159/rooui/vdom"
)

type DialogElement struct {
	BaseElement
}

func Dialog() *DialogElement {
	return &DialogElement{InitBaseElement("dialog")}
}

func (p *DialogElement) Open() *DialogElement {
	p.props["open"] = true
	return p
}

func (p *DialogElement) CloseByMask() *DialogElement {
	p.attributes["closeByMask"] = true
	return p
}

func (p *DialogElement) OnClose(f func(event dom.Event, options ...any), options ...any) *DialogElement {
	handlers := p.events["close"]
	handlers = append(handlers, vdom.EventHandler{
		Options: options,
		Handler: f,
	})
	p.events["close"] = handlers
	return p
}

func (p *DialogElement) OnCancel(f func(event dom.Event, options ...any), options ...any) *DialogElement {
	handlers := p.events["cancel"]
	handlers = append(handlers, vdom.EventHandler{
		Options: options,
		Handler: f,
	})
	p.events["cancel"] = handlers
	return p
}

func (p *DialogElement) Show(left, top string) {
	dlg := dom.Document.GetElementById(p.GetId())
	if dlg == nil {
		slog.Warn("dialog show failed. can't find element.", "id", p.GetId())
		return
	}
	dlg.Style().SetProperty("left", left)
	dlg.Style().SetProperty("top", top)
	dlg.Call("show")
}

func (p *DialogElement) ShowModal(left, top string) {
	dlg := dom.Document.GetElementById(p.GetId())
	if dlg == nil {
		slog.Warn("dialog show modal failed. can't find element.", "id", p.GetId())
		return
	}
	dlg.Style().SetProperty("left", left)
	dlg.Style().SetProperty("top", top)
	dlg.Call("showModal")
}

func (p *DialogElement) Close() {
	dlg := dom.Document.GetElementById(p.GetId())
	if dlg == nil {
		slog.Warn("dialog close failed. can't find element.", "id", p.GetId())
		return
	}
	dlg.Call("close")
}
