package html

type PreElement struct {
	BaseElement
}

func Pre() *PreElement {
	return &PreElement{InitBaseElement("pre")}
}
