package vdom

import (
	"log/slog"

	"github.com/matrix14159/rooui/gs"
	"honnef.co/go/js/dom/v2"
)

type VNodeStyle struct {
	Style map[string]string

	Delayed map[string]string

	Removes map[string]string
}

func NewVNodeStyle() *VNodeStyle {
	return &VNodeStyle{
		Style:   make(map[string]string),
		Delayed: make(map[string]string),
		Removes: make(map[string]string),
	}
}

func raf(f func()) {
	gs.RequestAnimationFrame(f)
}

func NextFrame(f func()) {
	raf(func() {
		raf(f)
	})
}

func setNextFrameStyle(el dom.HTMLElement, prop string, val any) {
	el.Style().Set(prop, val)
}

type StyleModule struct {
	reflowForced bool
}

func NewStyleModule() *StyleModule {
	return new(StyleModule)
}

func (p *StyleModule) updateStyle(oldVnode, vnode *VNode) {
	slog.Info("updateStyle", slog.Any("old", oldVnode), slog.Any("new", vnode))
	oldStyle, isNew1 := getStyle(oldVnode)
	newStyle, isNew2 := getStyle(vnode)
	if isNew1 && isNew2 {
		return
	}
	if oldStyle == newStyle {
		return
	}

	elm := vnode.Elm.(dom.HTMLElement)
	_, oldHasDel := oldStyle.Style["delayed"]

	for name, _ := range oldStyle.Style {
		if _, found := newStyle.Style[name]; !found {
			if len(name) >= 2 && name[0] == '-' && name[1] == '-' {
				elm.Style().RemoveProperty(name)
			} else {
				elm.Style().Set(name, "")
			}
		}
	}
	for name, value := range newStyle.Style {
		if name == "delayed" && len(newStyle.Delayed) > 0 {
			for name2, cur := range newStyle.Delayed {
				if !oldHasDel || cur != oldStyle.Delayed[name2] {
					setNextFrameStyle(elm, name2, cur)
				}
			}
		} else if name != "remove" && value != oldStyle.Style[name] {
			if len(name) >= 2 && name[0] == '-' && name[1] == '-' {
				elm.Style().SetProperty(name, value, "")
			} else {
				elm.Style().Set(name, value)
			}
		}
	}
}

func (p *StyleModule) applyDestroyStyle(vnode *VNode) {
	s, isNew := getStyle(vnode)
	if isNew {
		return
	}
	elm := vnode.Elm.(dom.HTMLElement)
	for name, value := range s.Style {
		elm.Style().Set(name, value)
	}
}

func (p *StyleModule) applyRemoveStyle(vnode *VNode, removeCallback func()) {
	//slog.Info("style module remove", slog.Any("key", vnode.Key))
	defer func() {
		removeCallback()
	}()

	s, isNew := getStyle(vnode)
	if isNew {
		return
	}
	elm := vnode.Elm.(dom.HTMLElement)

	if !p.reflowForced {
		p.reflowForced = true
	}

	for name, value := range s.Removes {
		elm.Style().Set(name, value)
	}

	//amount := 0
	//cs := dom.GetWindow().GetComputedStyle(elm, "")
	//csMap := cs.ToMap()
	//props := strings.Split(csMap["transition-property"], ",")
	//for _, one := range props {
	//	if _, found := s.Removes[one]; found {
	//		amount++
	//	}
	//}
	//
	//elm.AddEventListener("transitionend", false, func(ev dom.Event) {
	//	if ev.Target() == elm {
	//		amount--
	//	}
	//	if amount == 0 {
	//		removeCallback()
	//	}
	//})
}

func getStyle(vnode *VNode) (style *VNodeStyle, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.Style == nil {
		vnode.Data.Style = NewVNodeStyle()
		isNew = true
	}
	style = vnode.Data.Style
	return
}

func (p *StyleModule) Pre() {
	p.reflowForced = false
}

func (p *StyleModule) Create(empty, vnode *VNode) {
	p.updateStyle(empty, vnode)
}

func (p *StyleModule) Update(oldVNode, vnode *VNode) {
	p.updateStyle(oldVNode, vnode)
}

func (p *StyleModule) Destroy(vnode *VNode) {
	p.applyDestroyStyle(vnode)
}

func (p *StyleModule) Remove(vnode *VNode, removeCallback func()) {
	p.applyRemoveStyle(vnode, removeCallback)
}

func (p *StyleModule) Post() {}
