package html

type RpElement struct {
	BaseElement
}

func Rp() *RpElement {
	return &RpElement{InitBaseElement("rp")}
}
