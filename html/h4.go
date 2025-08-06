package html

type H4Element struct {
	BaseElement
}

func H4() *H4Element {
	return &H4Element{InitBaseElement("h4")}
}
