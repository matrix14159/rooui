//go:generate go run ht/gen_html.go
//go:generate go fmt html.go
//go:generate go run st/gen/gen_styles.go
//go:generate go fmt st/styles.go

package ui

import (
	dom "honnef.co/go/js/dom/v2"
)

type Event = dom.Event

// UI represent an ui element, use to construct ui-tree
type UI interface {
	// Kind return UI element kind
	Kind() Kind

	// GetUIElementId return the unique id to flag UI element
	GetUIElementId() string

	// SetUIElementId set the id for UI element
	SetUIElementId(id string)

	// TurnOnDisplay will modify or restore style display to show UI
	TurnOnDisplay()

	// TurnOffDisplay will set style display=none to hide UI
	TurnOffDisplay()

	// GetSelfDomElement return the raw html element according UI itself
	// if UI is a controller, then return its first child's dome element
	GetSelfDomElement() dom.Element

	// BuildTreeDomElement creates raw html elements for UI itself and children body
	BuildTreeDomElement() []dom.Element

	// OnCreated will be called when UI element is created
	OnCreated(f func())
	doCreated()

	// OnMounted set f which will be called when UI element added to ui-tree
	OnMounted(f func())
	doMounted()

	// OnUnmounted set f which will be called when UI element removed from ui-tree
	OnUnmounted(f func())
	doUnmounted()
}

type HtmlUI interface {
	UI
}

type ControllerUI interface {
	UI
}

// Comp is the base component interface
type Comp interface {
	ControllerUI

	setBody(child []UI)

	// Render return the component template for rending
	Render() UI
}

type ConditionController interface {
	ControllerUI
}

type SliceController interface {
	ControllerUI
}

type MapController interface {
	ControllerUI
}

type getBody interface {
	getBody() []UI
}

// Kind represent ui element type
type Kind int

const (
	UnknownElem Kind = iota

	HtmlElem

	ComponentElem

	ControllerElem

	ConditionControlElem

	RangeSliceElem

	RangeMapElem
)

func (p Kind) String() string {
	switch p {
	case HtmlElem:
		return "html"
	case ComponentElem:
		return "component"
	case ControllerElem:
		return "controller"
	case ConditionControlElem:
		return "condition"
	case RangeSliceElem:
		return "slice"
	case RangeMapElem:
		return "map"
	}
	return "undefined"
}
