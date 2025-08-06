package html

type IframeElement struct {
	BaseElement
}

func Iframe() *IframeElement {
	return &IframeElement{InitBaseElement("iframe")}
}
