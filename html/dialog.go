package html

type DialogElement struct {
	BaseElement
}

func Dialog() *DialogElement {
	return &DialogElement{InitBaseElement("dialog")}
}
