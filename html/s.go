package html

type SElement struct {
	BaseElement
}

func S() *SElement {
	return &SElement{InitBaseElement("s")}
}
