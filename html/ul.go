package html

type UlElement struct {
	BaseElement
}

func Ul() *UlElement {
	return &UlElement{InitBaseElement("ul")}
}
