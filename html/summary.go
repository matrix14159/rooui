package html

type SummaryElement struct {
	BaseElement
}

func Summary() *SummaryElement {
	return &SummaryElement{InitBaseElement("summary")}
}
