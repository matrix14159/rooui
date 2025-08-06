package html

type BlockquoteElement struct {
	BaseElement
}

func Blockquote() *BlockquoteElement {
	return &BlockquoteElement{InitBaseElement("blockquote")}
}
