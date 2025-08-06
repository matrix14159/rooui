package html

type CaptionElement struct {
	BaseElement
}

func Caption() *CaptionElement {
	return &CaptionElement{InitBaseElement("caption")}
}
