package html

type TableElement struct {
	BaseElement
}

func Table() *TableElement {
	return &TableElement{InitBaseElement("table")}
}
