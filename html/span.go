package html

type SpanElement struct {
	BaseElement
}

func Span() *SpanElement {
	return &SpanElement{InitBaseElement("span")}
}
