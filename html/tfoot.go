package html

type TfootElement struct {
	BaseElement
}

func Tfoot() *TfootElement {
	return &TfootElement{InitBaseElement("tfoot")}
}
