package html

type BrElement struct {
	BaseElement
}

func Br() *BrElement {
	return &BrElement{InitBaseElement("br")}
}
