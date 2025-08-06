package html

type SampElement struct {
	BaseElement
}

func Samp() *SampElement {
	return &SampElement{InitBaseElement("samp")}
}
