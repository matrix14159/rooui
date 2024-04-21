package vdom

import (
	"syscall/js"

	"honnef.co/go/js/dom/v2"
)

type EventHandler func(event dom.Event)

type EventTrigger struct {
	vnode *VNode
}

// On represent event listener
type On struct {
	Events map[string][]EventHandler

	// 事件注册后得到的js方法存根，后续可以用来移除监听器
	stubs map[string]js.Func

	// 在 Events 中注册的时间统一由 trigger 触发
	trigger *EventTrigger
}

func NewEventListener(vnode *VNode) *On {
	return &On{
		Events:  make(map[string][]EventHandler),
		stubs:   make(map[string]js.Func),
		trigger: &EventTrigger{vnode: vnode},
	}
}

func (p *EventTrigger) Handle(event dom.Event) {
	name := event.Type()
	listener, _ := getEventListener(p.vnode)
	handlers := listener.Events[name]
	for _, handler := range handlers {
		handler(event)
	}
}

type EventModule struct {
}

func (p *EventModule) updateEventListeners(oldVNode, vnode *VNode) {
	oldListener, isNew1 := getEventListener(oldVNode)
	newListener, isNew2 := getEventListener(vnode)
	if isNew1 && isNew2 {
		return
	}

	oldElm := oldVNode.Elm.(dom.HTMLElement)
	newElm := vnode.Elm.(dom.HTMLElement)

	for name, _ := range oldListener.Events {
		if _, found := newListener.Events[name]; !found {
			oldElm.RemoveEventListener(name, false, oldListener.stubs[name])
			delete(oldListener.stubs, name)
		}
	}

	for name, _ := range newListener.Events {
		if _, found := oldListener.Events[name]; !found {
			newListener.stubs[name] = newElm.AddEventListener(name, false, newListener.trigger.Handle)
		}
	}
}

func getEventListener(vnode *VNode) (on *On, isNew bool) {
	if vnode.Data == nil {
		vnode.Data = &VNodeData{}
	}
	if vnode.Data.On == nil {
		vnode.Data.On = NewEventListener(vnode)
		isNew = true
	}
	on = vnode.Data.On
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
