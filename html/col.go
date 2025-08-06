package html

type ColElement struct {
	BaseElement
}

func Col() *ColElement {
	return &ColElement{InitBaseElement("col")}
}
