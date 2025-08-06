package html

type ParamElement struct {
	BaseElement
}

func Param() *ParamElement {
	return &ParamElement{InitBaseElement("param")}
}
