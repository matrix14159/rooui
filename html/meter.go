package html

type MeterElement struct {
	BaseElement
}

func Meter() *MeterElement {
	return &MeterElement{InitBaseElement("meter")}
}
