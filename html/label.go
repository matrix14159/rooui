package html

type LabelElement struct {
	BaseElement
}

func Label() *LabelElement {
	return &LabelElement{InitBaseElement("label")}
}
