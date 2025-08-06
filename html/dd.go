package html

type DdElement struct {
	BaseElement
}

func Dd() *DdElement {
	return &DdElement{InitBaseElement("dd")}
}
