package html

type SubElement struct {
	BaseElement
}

func Sub() *SubElement {
	return &SubElement{InitBaseElement("sub")}
}
