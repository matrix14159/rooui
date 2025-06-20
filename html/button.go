package html

type ButtonElement struct {
	BaseElement
}

func Button() *ButtonElement {
	return &ButtonElement{InitBaseElement("button")}
}
