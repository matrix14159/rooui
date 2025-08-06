package html

type BElement struct {
	BaseElement
}

func B() *BElement {
	return &BElement{InitBaseElement("b")}
}
