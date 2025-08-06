package html

type FormElement struct {
	BaseElement
}

func Form() *FormElement {
	return &FormElement{InitBaseElement("form")}
}
