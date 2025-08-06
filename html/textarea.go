package html

type TextareaElement struct {
	BaseElement
}

func Textarea() *TextareaElement {
	return &TextareaElement{InitBaseElement("textarea")}
}
