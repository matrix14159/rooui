package html

type MapElement struct {
	BaseElement
}

func Map() *MapElement {
	return &MapElement{InitBaseElement("map")}
}
