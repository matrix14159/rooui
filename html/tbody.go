package html

type TbodyElement struct {
	BaseElement
}

func Tbody() *TbodyElement {
	return &TbodyElement{InitBaseElement("tbody")}
}
