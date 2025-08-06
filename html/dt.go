package html

type DtElement struct {
	BaseElement
}

func Dt() *DtElement {
	return &DtElement{InitBaseElement("dt")}
}
