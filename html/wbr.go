package html

type WbrElement struct {
	BaseElement
}

func Wbr() *WbrElement {
	return &WbrElement{InitBaseElement("wbr")}
}
