package html

type BodyElement struct {
	BaseElement
}

func Body() *BodyElement {
	return &BodyElement{InitBaseElement("body")}
}
