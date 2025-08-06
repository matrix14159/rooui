package html

type HeadElement struct {
	BaseElement
}

func Head() *HeadElement {
	return &HeadElement{InitBaseElement("head")}
}
