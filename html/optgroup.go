package html

type OptgroupElement struct {
	BaseElement
}

func Optgroup() *OptgroupElement {
	return &OptgroupElement{InitBaseElement("optgroup")}
}
