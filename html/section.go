package html

type SectionElement struct {
	BaseElement
}

func Section() *SectionElement {
	return &SectionElement{InitBaseElement("section")}
}
