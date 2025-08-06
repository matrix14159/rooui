package html

type H5Element struct {
	BaseElement
}

func H5() *H5Element {
	return &H5Element{InitBaseElement("h5")}
}
