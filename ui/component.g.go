package ui

type Component struct {
	baseController
}

func (p *Component) Kind() Kind {
	return ComponentElem
}

func (p *Component) setBody(child []UI) {
	p.body = child
}

func (p *Component) Render() UI {
	return nil
}
