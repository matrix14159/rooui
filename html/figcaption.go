package html

type FigcaptionElement struct {
	BaseElement
}

func Figcaption() *FigcaptionElement {
	return &FigcaptionElement{InitBaseElement("figcaption")}
}
