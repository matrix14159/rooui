package html

type TdElement struct {
	BaseElement
}

func Td() *TdElement {
	return &TdElement{InitBaseElement("td")}
}
