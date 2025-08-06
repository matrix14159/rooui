package html

type MainElement struct {
	BaseElement
}

func Main() *MainElement {
	return &MainElement{InitBaseElement("main")}
}
