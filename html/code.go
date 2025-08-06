package html

type CodeElement struct {
	BaseElement
}

func Code() *CodeElement {
	return &CodeElement{InitBaseElement("code")}
}
