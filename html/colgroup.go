package html

type ColgroupElement struct {
	BaseElement
}

func Colgroup() *ColgroupElement {
	return &ColgroupElement{InitBaseElement("colgroup")}
}
