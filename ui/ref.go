package ui

type Binder interface {
	Data() any

	AddUpdatedHandler(handler func(oldVal, newVal any), id ...string) (handlerId string, num int)
	RemoveUpdatedHandler(handlerId string) (ok bool, num int)
}

// ToRef create Ref by data
func ToRef[T any](data T) *Ref[T] {
	return &Ref[T]{data: data}
}

// Effect proxy a to b
// when a updated, f will be called and then update b
func Effect[T1, T2 any](a *Ref[T1], f func(T1) T2) (b *Ref[T2]) {
	b = ToRef(f(a.Value()))
	h := func(oldVal, newVal any) {
		b.Set(f(newVal.(T1)))
	}
	b.proxy = &refProxy{
		src:        a,
		updateFunc: h,
		handlerId:  "",
	}
	return
}

type Ref[T any] struct {
	data T

	// since tinygo(v0.26) panic: unimplemented: (reflect.Value).UnsafePointer()
	// when remove handler can't call RemoveUpdatedHandler(f func(o, n any))
	// so use id to flag a handler
	updateHandlers []refUpdatedHandler

	proxy *refProxy
}

type refUpdatedHandler struct {
	id      string
	handler func(oldVal, newVal any)
}

type refProxy struct {
	// Ref proxy source
	src Binder

	// updated handler for src
	updateFunc func(oldVal, newVal any)

	handlerId string
}

func (p Ref[T]) Value() T {
	return p.data
}

func (p *Ref[T]) Set(v T) {
	old := p.data
	p.data = v
	for _, h := range p.updateHandlers {
		if h.id != "" && h.handler != nil {
			h.handler(old, p.data)
		}
	}
}

// Data is the same as Value(), but return type is any
func (p *Ref[T]) Data() any {
	return p.data
}

// AddUpdatedHandler will add the handler to listen when p.data changed
// option to specify an id, otherwise generate a new id to flag handler
func (p *Ref[T]) AddUpdatedHandler(handler func(oldVal, newVal any), id ...string) (handlerId string, num int) {
	if len(id) > 0 && id[0] != "" {
		handlerId = id[0]
	} else {
		handlerId = NewId()
	}
	p.updateHandlers = append(p.updateHandlers, refUpdatedHandler{id: handlerId, handler: handler})
	num = len(p.updateHandlers)

	if p.proxy != nil && p.proxy.handlerId == "" {
		p.proxy.handlerId = handlerId
		p.proxy.src.AddUpdatedHandler(p.proxy.updateFunc, p.proxy.handlerId)
	}
	return
}

// RemoveUpdatedHandler will remove the listen handler
func (p *Ref[T]) RemoveUpdatedHandler(handlerId string) (ok bool, num int) {
	for i, h := range p.updateHandlers {
		if handlerId == h.id {
			ok = true
			p.updateHandlers = append(p.updateHandlers[:i], p.updateHandlers[i+1:]...)
			break
		}
	}
	num = len(p.updateHandlers)

	if p.proxy != nil && p.proxy.handlerId == handlerId {
		p.proxy.src.RemoveUpdatedHandler(p.proxy.handlerId)
	}
	return
}
