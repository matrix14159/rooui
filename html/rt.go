package html

type RtElement struct {
	BaseElement
}

func Rt() *RtElement {
	return &RtElement{InitBaseElement("rt")}
}
