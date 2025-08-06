package html

type TrElement struct {
	BaseElement
}

func Tr() *TrElement {
	return &TrElement{InitBaseElement("tr")}
}
