package html

type DetailsElement struct {
	BaseElement
}

func Details() *DetailsElement {
	return &DetailsElement{InitBaseElement("details")}
}
