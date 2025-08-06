package html

type RubyElement struct {
	BaseElement
}

func Ruby() *RubyElement {
	return &RubyElement{InitBaseElement("ruby")}
}
