package html

type PElement struct {
	BaseElement
}

func P() *PElement {
	return &PElement{InitBaseElement("p")}
}
