package html

type LinkElement struct {
	BaseElement
}

func Link() *LinkElement {
	return &LinkElement{InitBaseElement("link")}
}
