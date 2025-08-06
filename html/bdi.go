package html

type BdiElement struct {
	BaseElement
}

func Bdi() *BdiElement {
	return &BdiElement{InitBaseElement("bdi")}
}
