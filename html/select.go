package html

type SelectElement struct {
	BaseElement
}

func Select() *SelectElement {
	return &SelectElement{InitBaseElement("select")}
}
