package html

type DfnElement struct {
	BaseElement
}

func Dfn() *DfnElement {
	return &DfnElement{InitBaseElement("dfn")}
}
