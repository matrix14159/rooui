package html

type BdoElement struct {
	BaseElement
}

func Bdo() *BdoElement {
	return &BdoElement{InitBaseElement("bdo")}
}
