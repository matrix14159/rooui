package html

type SupElement struct {
	BaseElement
}

func Sup() *SupElement {
	return &SupElement{InitBaseElement("sup")}
}
