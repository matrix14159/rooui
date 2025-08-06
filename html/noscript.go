package html

type NoscriptElement struct {
	BaseElement
}

func Noscript() *NoscriptElement {
	return &NoscriptElement{InitBaseElement("noscript")}
}
