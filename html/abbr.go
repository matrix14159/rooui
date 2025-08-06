package html

type AbbrElement struct {
	BaseElement
}

func Abbr() *AbbrElement {
	return &AbbrElement{InitBaseElement("abbr")}
}
