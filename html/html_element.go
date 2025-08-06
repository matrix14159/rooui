package html

type HtmlElement struct {
	BaseElement
}

func Html() *HtmlElement {
	return &HtmlElement{InitBaseElement("html")}
}
