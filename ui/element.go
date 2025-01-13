package ui

import (
	"github.com/matrix14159/rooui/core"
)

type Element interface {
	core.HtmlElement

	setComp(comp Comp)

	getComp() Comp
}

type compElement struct {
	core.HtmlElement

	comp Comp
}

func (p *compElement) setComp(comp Comp) {
	p.comp = comp
}

func (p *compElement) getComp() Comp {
	return p.comp
}
