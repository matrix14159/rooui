package html

type InputElement struct {
	BaseElement
}

func Input() *InputElement {
	return &InputElement{InitBaseElement("input")}
}
