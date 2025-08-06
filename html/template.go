package html

type TemplateElement struct {
	BaseElement
}

func Template() *TemplateElement {
	return &TemplateElement{InitBaseElement("template")}
}
