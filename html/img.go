package html

type ImgElement struct {
	BaseElement
}

func Img() *ImgElement {
	return &ImgElement{InitBaseElement("img")}
}
