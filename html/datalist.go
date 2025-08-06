package html

type DatalistElement struct {
	BaseElement
}

func Datalist() *DatalistElement {
	return &DatalistElement{InitBaseElement("datalist")}
}
