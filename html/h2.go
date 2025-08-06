package html

type H2Element struct {
	BaseElement
}

func H2() *H2Element {
	return &H2Element{InitBaseElement("h2")}
}
