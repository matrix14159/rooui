package html

type VideoElement struct {
	BaseElement
}

func Video() *VideoElement {
	return &VideoElement{InitBaseElement("video")}
}
