package html

type TrackElement struct {
	BaseElement
}

func Track() *TrackElement {
	return &TrackElement{InitBaseElement("track")}
}
