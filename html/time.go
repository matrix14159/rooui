package html

type TimeElement struct {
	BaseElement
}

func Time() *TimeElement {
	return &TimeElement{InitBaseElement("time")}
}
