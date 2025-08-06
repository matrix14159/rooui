package html

type StyleElement struct {
	BaseElement
}

func Style() *StyleElement {
	return &StyleElement{InitBaseElement("style")}
}
