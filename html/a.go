package html

type AElement struct {
	BaseElement
}

func A() *AElement {
	return &AElement{InitBaseElement("a")}
}
