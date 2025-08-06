package html

type AsideElement struct {
	BaseElement
}

func Aside() *AsideElement {
	return &AsideElement{InitBaseElement("aside")}
}
