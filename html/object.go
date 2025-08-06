package html

type ObjectElement struct {
	BaseElement
}

func Object() *ObjectElement {
	return &ObjectElement{InitBaseElement("object")}
}
