package html

type IElement struct {
	BaseElement
}

func I() *IElement {
	return &IElement{InitBaseElement("i")}
}
