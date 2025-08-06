package html

type LegendElement struct {
	BaseElement
}

func Legend() *LegendElement {
	return &LegendElement{InitBaseElement("legend")}
}
