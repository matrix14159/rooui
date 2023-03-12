package ui

import (
	dom "honnef.co/go/js/dom/v2"
)

func BuildHtml(u UI) (els []dom.Element) {
	switch u.(type) {
	case Comp:
		c := u.(Comp)
		c.setBody([]UI{c.Render()})
		els = c.BuildTreeDomElement()

	default:
		els = u.BuildTreeDomElement()
	}
	return
}

func FindUIElement(id string) UI {
	return FindUIElementFrom(RootComponent, id)
}

func FindUIElementFrom(from UI, id string) UI {
	if from == nil {
		return nil
	}
	if from.GetUIElementId() == id {
		return from
	}

	if b, ok := from.(getBody); ok {
		for _, sub := range b.getBody() {
			if ret := FindUIElementFrom(sub, id); ret != nil {
				return ret
			}
		}
	}
	return nil
}
