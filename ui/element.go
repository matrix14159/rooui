package ui

import (
	"github.com/matrix14159/rooui/core"
)

type Element interface {
	core.HtmlElement

	setComp(comp Comp)
	getComp() Comp

	setParent(parent Comp)
	getParent() Comp
}

type compElement struct {
	core.HtmlElement

	parent Comp

	comp Comp
}

func (p *compElement) setComp(comp Comp) {
	p.comp = comp
}

func (p *compElement) getComp() Comp {
	return p.comp
}

func (p *compElement) setParent(parent Comp) {
	p.parent = parent
}

func (p *compElement) getParent() Comp {
	return p.parent
}
