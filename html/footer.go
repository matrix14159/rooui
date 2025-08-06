package html

type FooterElement struct {
	BaseElement
}

func Footer() *FooterElement {
	return &FooterElement{InitBaseElement("footer")}
}
