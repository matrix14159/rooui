package html

type QElement struct {
	BaseElement
}

func Q() *QElement {
	return &QElement{InitBaseElement("q")}
}
