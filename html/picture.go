package html

type PictureElement struct {
	BaseElement
}

func Picture() *PictureElement {
	return &PictureElement{InitBaseElement("picture")}
}
