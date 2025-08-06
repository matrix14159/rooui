package html

type FigureElement struct {
	BaseElement
}

func Figure() *FigureElement {
	return &FigureElement{InitBaseElement("figure")}
}
