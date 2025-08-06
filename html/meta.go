package html

type MetaElement struct {
	BaseElement
}

func Meta() *MetaElement {
	return &MetaElement{InitBaseElement("meta")}
}
