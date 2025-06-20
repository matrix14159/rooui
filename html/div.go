package html

type DivElement struct {
	BaseElement
}

func Div() *DivElement {
	return &DivElement{InitBaseElement("div")}
}
