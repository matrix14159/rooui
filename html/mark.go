package html

type MarkElement struct {
	BaseElement
}

func Mark() *MarkElement {
	return &MarkElement{InitBaseElement("mark")}
}
