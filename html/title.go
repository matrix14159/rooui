package html

type TitleElement struct {
	BaseElement
}

func Title() *TitleElement {
	return &TitleElement{InitBaseElement("title")}
}
