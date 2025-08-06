package html

type HeaderElement struct {
	BaseElement
}

func Header() *HeaderElement {
	return &HeaderElement{InitBaseElement("header")}
}
