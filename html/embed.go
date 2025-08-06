package html

type EmbedElement struct {
	BaseElement
}

func Embed() *EmbedElement {
	return &EmbedElement{InitBaseElement("embed")}
}
