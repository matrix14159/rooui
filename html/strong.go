package html

type StrongElement struct {
	BaseElement
}

func Strong() *StrongElement {
	return &StrongElement{InitBaseElement("strong")}
}
