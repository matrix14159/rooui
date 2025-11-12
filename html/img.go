package html

type ImgElement struct {
	BaseElement
}

func Img() *ImgElement {
	return &ImgElement{InitBaseElement("img")}
}

func (p *ImgElement) Src(v string) *ImgElement {
	p.props["src"] = v
	return p
}

func (p *ImgElement) Alt(v string) *ImgElement {
	p.props["alt"] = v
	return p
}
