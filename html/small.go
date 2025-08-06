package html

type SmallElement struct {
	BaseElement
}

func Small() *SmallElement {
	return &SmallElement{InitBaseElement("small")}
}
