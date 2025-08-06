package html

type AudioElement struct {
	BaseElement
}

func Audio() *AudioElement {
	return &AudioElement{InitBaseElement("audio")}
}
