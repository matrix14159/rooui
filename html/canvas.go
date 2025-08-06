package html

type CanvasElement struct {
	BaseElement
}

func Canvas() *CanvasElement {
	return &CanvasElement{InitBaseElement("canvas")}
}
