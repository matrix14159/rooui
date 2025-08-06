package html

type CiteElement struct {
	BaseElement
}

func Cite() *CiteElement {
	return &CiteElement{InitBaseElement("cite")}
}
