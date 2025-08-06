package html

type SourceElement struct {
	BaseElement
}

func Source() *SourceElement {
	return &SourceElement{InitBaseElement("source")}
}
