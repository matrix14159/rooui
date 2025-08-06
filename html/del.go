package html

type DelElement struct {
	BaseElement
}

func Del() *DelElement {
	return &DelElement{InitBaseElement("del")}
}
