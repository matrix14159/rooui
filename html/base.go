package html

import (
	"fmt"
	"strings"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/dom"
	"github.com/matrix14159/rooui/vdom"
)

type Element interface {
	Tag() string

	GetId() string

	GetText() string

	GetEvents() map[string][]vdom.EventHandler

	GetBody() []Element

	GetProps() map[string]any

	GetClasses() []string

	GetStyles() []css.Style

	GetClassStyle() map[string][]css.Style
}

type BaseElement struct {
	tag string

	id string

	text string

	body []Element

	props map[string]any

	classes []string

	styles []css.Style

	classStyle map[string][]css.Style

	events map[string][]vdom.EventHandler
}

func InitBaseElement(tag string) BaseElement {
	p := BaseElement{tag: tag}
	p.props = make(map[string]any)
	p.events = make(map[string][]vdom.EventHandler)
	return p
}

func (p *BaseElement) Tag() string {
	return p.tag
}

func (p *BaseElement) GetId() string {
	return p.id
}

func (p *BaseElement) GetText() string {
	return p.text
}

func (p *BaseElement) GetBody() []Element {
	return p.body
}

func (p *BaseElement) GetProps() map[string]any {
	return p.props
}

func (p *BaseElement) GetClasses() []string {
	return p.classes
}

func (p *BaseElement) GetStyles() []css.Style {
	return p.styles
}

func (p *BaseElement) GetClassStyle() map[string][]css.Style {
	return p.classStyle
}

func (p *BaseElement) GetEvents() map[string][]vdom.EventHandler {
	return p.events
}

func (p *BaseElement) Id(id string) *BaseElement {
	p.id = id
	return p
}

func (p *BaseElement) Body(child ...Element) *BaseElement {
	p.body = append(p.body, child...)
	return p
}

func (p *BaseElement) Classes(name ...string) *BaseElement {
	for _, one := range name {
		p.classes = append(p.classes, strings.TrimPrefix(strings.TrimSpace(one), "."))
	}
	return p
}

func (p *BaseElement) Class(name string, styles []css.Style) *BaseElement {
	if p.classStyle == nil {
		p.classStyle = make(map[string][]css.Style)
	}
	cls := strings.TrimPrefix(strings.TrimSpace(name), ".")
	p.classStyle[cls] = styles
	return p
}

func (p *BaseElement) Style(style ...css.Style) *BaseElement {
	for _, one := range style {
		p.styles = append(p.styles, one)
	}
	return p
}

func (p *BaseElement) Text(text any) *BaseElement {
	p.text = fmt.Sprintf("%v", text)
	return p
}

func (p *BaseElement) AddEventListener(name string, f func(event dom.Event, options ...any), options ...any) *BaseElement {
	handlers := p.events[name]
	handlers = append(handlers, vdom.EventHandler{
		Options: options,
		Handler: f,
	})
	p.events[name] = handlers
	return p
}

func (p *BaseElement) OnClick(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("click", f, options...)
}

func (p *BaseElement) OnAnimationCancel(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("animationcancel", f, options...)
}

func (p *BaseElement) OnAnimationEnd(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("animationend", f, options...)
}

func (p *BaseElement) OnAnimationIteration(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("animationiteration", f, options...)
}

func (p *BaseElement) OnAnimationStart(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("animationstart", f, options...)
}

func (p *BaseElement) OnAuxClick(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("auxclick", f, options...)
}

func (p *BaseElement) OnBeforeInput(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("beforeinput", f, options...)
}

func (p *BaseElement) OnBeforeMatch(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("beforematch", f, options...)
}

func (p *BaseElement) OnBeforeXrSelect(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("beforexrselect", f, options...)
}

func (p *BaseElement) OnBlur(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("blur", f, options...)
}

func (p *BaseElement) OnCompositionEnd(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("compositionend", f, options...)
}

func (p *BaseElement) OnCompositionStart(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("compositionstart", f, options...)
}

func (p *BaseElement) OnCompositionUpdate(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("compositionupdate", f, options...)
}

func (p *BaseElement) OnContentVisibilityAutoStateChange(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("contentvisibilityautostatechange", f, options...)
}

func (p *BaseElement) OnContextMenu(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("contextmenu", f, options...)
}

func (p *BaseElement) OnCopy(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("copy", f, options...)
}

func (p *BaseElement) OnCut(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("cut", f, options...)
}

func (p *BaseElement) OnDblclick(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("dblclick", f, options...)
}

func (p *BaseElement) OnFocus(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("focus", f, options...)
}

func (p *BaseElement) OnFocusin(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("focusin", f, options...)
}

func (p *BaseElement) OnFocusout(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("focusout", f, options...)
}

func (p *BaseElement) OnFullScreenChange(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("fullscreenchange", f, options...)
}

func (p *BaseElement) OnFullScreenError(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("fullscreenerror", f, options...)
}

func (p *BaseElement) OnGestureChange(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("gesturechange", f, options...)
}

func (p *BaseElement) OnGestureEnd(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("gestureend", f, options...)
}

func (p *BaseElement) OnGestureStart(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("gesturestart", f, options...)
}

func (p *BaseElement) OnGotPointerCapture(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("gotpointercapture", f, options...)
}

func (p *BaseElement) OnInput(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("input", f, options...)
}

func (p *BaseElement) OnKeyDown(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("keydown", f, options...)
}

func (p *BaseElement) OnKeyPress(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("keypress", f, options...)
}

func (p *BaseElement) OnKeyUp(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("keyup", f, options...)
}

func (p *BaseElement) OnLostPointerCapture(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("lostpointercapture", f, options...)
}

func (p *BaseElement) OnMouseDown(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mousedown", f, options...)
}

func (p *BaseElement) OnMouseEnter(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mouseenter", f, options...)
}

func (p *BaseElement) OnMouseLeave(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mouseleave", f, options...)
}

func (p *BaseElement) OnMouseMove(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mousemove", f, options...)
}

func (p *BaseElement) OnMouseOut(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mouseout", f, options...)
}

func (p *BaseElement) OnMouseOver(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mouseover", f, options...)
}

func (p *BaseElement) OnMouseUp(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mouseup", f, options...)
}

func (p *BaseElement) OnMouseWheel(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("mousewheel", f, options...)
}

func (p *BaseElement) OnMozMousePixelScroll(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("MozMousePixelScroll", f, options...)
}

func (p *BaseElement) OnPaste(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("paste", f, options...)
}

func (p *BaseElement) OnPointerCancel(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointercancel", f, options...)
}

func (p *BaseElement) OnPointerDown(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerdown", f, options...)
}

func (p *BaseElement) OnPointerEnter(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerenter", f, options...)
}

func (p *BaseElement) OnPointerLeave(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerleave", f, options...)
}

func (p *BaseElement) OnPointerMove(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointermove", f, options...)
}

func (p *BaseElement) OnPointerOut(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerout", f, options...)
}

func (p *BaseElement) OnPointerOver(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerover", f, options...)
}

func (p *BaseElement) OnPointerRawUpdate(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerrawupdate", f, options...)
}

func (p *BaseElement) OnPointerUp(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("pointerup", f, options...)
}

func (p *BaseElement) OnScroll(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("scroll", f, options...)
}

func (p *BaseElement) OnScrollEnd(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("scrollend", f, options...)
}

func (p *BaseElement) OnScrollSnapChange(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("scrollsnapchange", f, options...)
}

func (p *BaseElement) OnScrollSnapChanging(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("scrollsnapchanging", f, options...)
}

func (p *BaseElement) OnSecurityPolicyViolation(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("securitypolicyviolation", f, options...)
}

func (p *BaseElement) OnTouchCancel(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("touchcancel", f, options...)
}

func (p *BaseElement) OnTouchEnd(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("touchend", f, options...)
}

func (p *BaseElement) OnTouchMove(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("touchmove", f, options...)
}

func (p *BaseElement) OnTouchStart(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("touchstart", f, options...)
}

func (p *BaseElement) OnTransitionCancel(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("transitioncancel", f, options...)
}

func (p *BaseElement) OnTransitionEnd(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("transitionend", f, options...)
}

func (p *BaseElement) OnTransitionRun(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("transitionrun", f, options...)
}

func (p *BaseElement) OnTransitionStart(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("transitionstart", f, options...)
}

func (p *BaseElement) OnWebkitMouseforceChanged(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("webkitmouseforcechanged", f, options...)
}

func (p *BaseElement) OnWebkitMouseForceDown(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("webkitmouseforcedown", f, options...)
}

func (p *BaseElement) OnWebkitMouseForceJp(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("webkitmouseforceup", f, options...)
}

func (p *BaseElement) OnWebkitMouseForceWillBegin(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("webkitmouseforcewillbegin", f, options...)
}

func (p *BaseElement) OnWheel(f func(event dom.Event, options ...any), options ...any) *BaseElement {
	return p.AddEventListener("wheel", f, options...)
}
