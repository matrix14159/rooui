package vdom

import (
	"syscall/js"

	"github.com/matrix14159/rooui/dom"
)

// Listener represent event listener
type Listener struct {
	Events map[string][]EventHandler

	// 事件注册后得到的js方法存根，后续可以用来移除监听器
	stubs map[string]js.Func
}

type EventHandler struct {
	Options []any
	Handler func(event dom.Event, options ...any)
}

func NewEventListener() *Listener {
	return &Listener{
		Events: make(map[string][]EventHandler),
		stubs:  make(map[string]js.Func),
	}
}

func (p *Listener) handle(event dom.Event) {
	name := event.EventType()
	handlers := p.Events[name]
	for _, one := range handlers {
		one.Handler(event, one.Options...)
	}
}

type EventModule struct {
}

func NewEventModule() *EventModule {
	return new(EventModule)
}

func (p *EventModule) updateEventListeners(oldVNode, vnode *VNode) {
	oldListener, isNew1 := getEventListener(oldVNode)
	newListener, isNew2 := getEventListener(vnode)
	if isNew1 && isNew2 {
		return
	}

	oldElm := oldVNode.Elm
	for name, _ := range oldListener.Events {
		if _, found := newListener.Events[name]; !found {
			oldElm.RemoveEventListener(name, false, oldListener.stubs[name])
			delete(oldListener.stubs, name)
		}
	}

	newElm := vnode.Elm
	for name, _ := range newListener.Events {
		if _, found := oldListener.Events[name]; !found {
			f := newElm.AddEventListener(name, false, newListener.handle)
			newListener.stubs[name] = f
		}
	}
}

func getEventListener(vnode *VNode) (on *Listener, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.Listener == nil {
		vnode.Data.Listener = NewEventListener()
		isNew = true
	}
	on = vnode.Data.Listener
	return
}

func (p *EventModule) Pre() {
}

func (p *EventModule) Create(empty, vnode *VNode) {
	p.updateEventListeners(empty, vnode)
}

func (p *EventModule) Update(oldVNode, vnode *VNode) {
	p.updateEventListeners(oldVNode, vnode)
}

func (p *EventModule) Destroy(vnode *VNode) {
	p.updateEventListeners(vnode, emptyNode)
}

func (p *EventModule) Remove(vnode *VNode, removeCallback func()) {
	removeCallback()
}

func (p *EventModule) Post() {
}
