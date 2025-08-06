package html

type H6Element struct {
	BaseElement
}

func H6() *H6Element {
	return &H6Element{InitBaseElement("h6")}
}
