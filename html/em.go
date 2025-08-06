package html

type EmElement struct {
	BaseElement
}

func Em() *EmElement {
	return &EmElement{InitBaseElement("em")}
}
