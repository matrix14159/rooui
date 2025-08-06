package html

type DataElement struct {
	BaseElement
}

func Data() *DataElement {
	return &DataElement{InitBaseElement("data")}
}
