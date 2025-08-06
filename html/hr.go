package html

type HrElement struct {
	BaseElement
}

func Hr() *HrElement {
	return &HrElement{InitBaseElement("hr")}
}
