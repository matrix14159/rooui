package html

type OutputElement struct {
	BaseElement
}

func Output() *OutputElement {
	return &OutputElement{InitBaseElement("output")}
}
