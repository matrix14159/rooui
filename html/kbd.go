package html

type KbdElement struct {
	BaseElement
}

func Kbd() *KbdElement {
	return &KbdElement{InitBaseElement("kbd")}
}
