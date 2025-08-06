package html

type OlElement struct {
	BaseElement
}

func Ol() *OlElement {
	return &OlElement{InitBaseElement("ol")}
}
