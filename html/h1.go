package html

type H1Element struct {
	BaseElement
}

func H1() *H1Element {
	return &H1Element{InitBaseElement("h1")}
}
