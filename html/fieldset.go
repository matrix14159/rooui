package html

type FieldsetElement struct {
	BaseElement
}

func Fieldset() *FieldsetElement {
	return &FieldsetElement{InitBaseElement("fieldset")}
}
