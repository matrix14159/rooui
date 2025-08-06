package html

type UElement struct {
	BaseElement
}

func U() *UElement {
	return &UElement{InitBaseElement("u")}
}
