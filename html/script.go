package html

type ScriptElement struct {
	BaseElement
}

func Script() *ScriptElement {
	return &ScriptElement{InitBaseElement("script")}
}
