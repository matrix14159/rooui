package html

type LiElement struct {
	BaseElement
}

func Li() *LiElement {
	return &LiElement{InitBaseElement("li")}
}
