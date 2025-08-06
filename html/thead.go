package html

type TheadElement struct {
	BaseElement
}

func Thead() *TheadElement {
	return &TheadElement{InitBaseElement("thead")}
}
