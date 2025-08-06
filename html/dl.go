package html

type DlElement struct {
	BaseElement
}

func Dl() *DlElement {
	return &DlElement{InitBaseElement("dl")}
}
