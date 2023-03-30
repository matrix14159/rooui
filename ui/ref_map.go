package ui

func ToRefMap[K comparable, V any](data map[K]V) (ret *RefMap[K, V]) {
	ret = new(RefMap[K, V])
	ret.data = data
	if ret.data == nil {
		ret.data = make(map[K]V)
	}
	return ret
}

type RefMap[K comparable, V any] struct {
	data map[K]V

	upsertHandlers []refMapUpsertHandler[K, V]
	deleteHandlers []refMapDeleteHandler[K]
}

type refMapUpsertHandler[K comparable, V any] struct {
	id      string
	handler func(K, V)
}

type refMapDeleteHandler[K comparable] struct {
	id      string
	handler func(keys []K)
}

func (p RefMap[K, V]) RawData() map[K]V {
	return p.data
}

func (p RefMap[K, V]) Len() int {
	return len(p.data)
}

func (p RefMap[K, V]) Get(k K) V {
	return p.data[k]
}

func (p *RefMap[K, V]) Set(key K, val V) {
	p.data[key] = val
	p.emitUpserted(key, val)
}

func (p *RefMap[K, V]) Delete(keys ...K) {
	for _, key := range keys {
		delete(p.data, key)
	}
	p.emitDeleted(keys)
}

func (p *RefMap[K, V]) emitUpserted(key K, val V) {
	for _, h := range p.upsertHandlers {
		if h.id != "" && h.handler != nil {
			h.handler(key, val)
		}
	}
}

func (p *RefMap[K, V]) emitDeleted(keys []K) {
	for _, h := range p.deleteHandlers {
		if h.id != "" && h.handler != nil {
			h.handler(keys)
		}
	}
}

func (p *RefMap[K, V]) AddUpsertHandler(handler func(K, V), id ...string) (handlerId string, num int) {
	if len(id) > 0 && id[0] != "" {
		handlerId = id[0]
	} else {
		handlerId = NewId()
	}
	p.upsertHandlers = append(p.upsertHandlers, refMapUpsertHandler[K, V]{id: handlerId, handler: handler})
	num = len(p.upsertHandlers)
	return
}

func (p *RefMap[K, V]) RemoveUpsertHandler(handlerId string) (ok bool, num int) {
	for i, h := range p.upsertHandlers {
		if handlerId == h.id {
			ok = true
			p.upsertHandlers = append(p.upsertHandlers[:i], p.upsertHandlers[i+1:]...)
			break
		}
	}
	num = len(p.upsertHandlers)
	return
}

func (p *RefMap[K, V]) AddDeletedHandler(handler func(K []K), id ...string) (handlerId string, num int) {
	if len(id) > 0 && id[0] != "" {
		handlerId = id[0]
	} else {
		handlerId = NewId()
	}
	p.deleteHandlers = append(p.deleteHandlers, refMapDeleteHandler[K]{id: handlerId, handler: handler})
	num = len(p.deleteHandlers)
	return
}

func (p *RefMap[K, V]) RemoveDeletedHandler(handlerId string) (ok bool, num int) {
	for i, h := range p.deleteHandlers {
		if handlerId == h.id {
			ok = true
			p.deleteHandlers = append(p.deleteHandlers[:i], p.deleteHandlers[i+1:]...)
			break
		}
	}
	num = len(p.deleteHandlers)
	return
}
