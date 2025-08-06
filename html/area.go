package html

type AreaElement struct {
	BaseElement
}

func Area() *AreaElement {
	return &AreaElement{InitBaseElement("area")}
}
