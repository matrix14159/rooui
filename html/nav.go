package html

type NavElement struct {
	BaseElement
}

func Nav() *NavElement {
	return &NavElement{InitBaseElement("nav")}
}
