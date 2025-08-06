package html

type OptionElement struct {
	BaseElement
}

func Option() *OptionElement {
	return &OptionElement{InitBaseElement("option")}
}
