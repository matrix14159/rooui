package html

type InsElement struct {
	BaseElement
}

func Ins() *InsElement {
	return &InsElement{InitBaseElement("ins")}
}
