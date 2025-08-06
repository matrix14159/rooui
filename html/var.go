package html

type VarElement struct {
	BaseElement
}

func Var() *VarElement {
	return &VarElement{InitBaseElement("var")}
}
