package html

type ThElement struct {
	BaseElement
}

func Th() *ThElement {
	return &ThElement{InitBaseElement("th")}
}
