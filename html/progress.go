package html

type ProgressElement struct {
	BaseElement
}

func Progress() *ProgressElement {
	return &ProgressElement{InitBaseElement("progress")}
}
