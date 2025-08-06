package html

type SlotElement struct {
	BaseElement
}

func Slot() *SlotElement {
	return &SlotElement{InitBaseElement("slot")}
}
