package html

type H3Element struct {
	BaseElement
}

func H3() *H3Element {
	return &H3Element{InitBaseElement("h3")}
}
