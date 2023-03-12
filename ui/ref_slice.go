package ui

import (
	"fmt"
)

func ToRefSlice[T any](data []T) (ret *RefSlice[T]) {
	return &RefSlice[T]{data: data}
}

type RefSlice[T any] struct {
	data []T

	insertHandlers []refSliceInsertedHandler[T]
	deleteHandlers []refSliceDeletedHandler
	updateHandlers []refSliceUpdatedHandler[T]
}

type refSliceInsertedHandler[T any] struct {
	id      string
	handler func(pos int, val []T)
}

type refSliceDeletedHandler struct {
	id      string
	handler func(pos, num int)
}

type refSliceUpdatedHandler[T any] struct {
	id      string
	handler func(pos int, val T)
}

func (p RefSlice[T]) Len() int {
	return len(p.data)
}

func (p RefSlice[T]) Get(index int) T {
	if index < 0 || index >= len(p.data) {
		panic(fmt.Sprintf("get index out of range [%d] with length %d", index, len(p.data)))
	}
	return p.data[index]
}

func (p *RefSlice[T]) Set(index int, val T) {
	if index < 0 || index >= len(p.data) {
		panic(fmt.Sprintf("set index out of range [%d] with length %d", index, len(p.data)))
	}
	p.data[index] = val
	p.emitUpdated(index, val)
}

func (p *RefSlice[T]) emitUpdated(pos int, val T) {
	for _, h := range p.updateHandlers {
		if h.id != "" && h.handler != nil {
			h.handler(pos, val)
		}
	}
}

// Insert inserts val at index
func (p *RefSlice[T]) Insert(pos int, val ...T) {
	if pos >= len(p.data) {
		p.Push(val...)
		return
	}
	if pos < 0 {
		pos = 0
	}
	p.data = append(p.data[:pos], append(val, p.data[pos:]...)...)
	p.emitInserted(pos, val)
}

// Push pushes val to the end
func (p *RefSlice[T]) Push(val ...T) {
	pos := len(p.data)
	p.data = append(p.data, val...)
	p.emitInserted(pos, val)
}

func (p *RefSlice[T]) emitInserted(pos int, val []T) {
	for _, h := range p.insertHandlers {
		if h.id != "" && h.handler != nil {
			h.handler(pos, val)
		}
	}
}

// Delete deletes element from pos to pos+num
func (p *RefSlice[T]) Delete(pos int, num int) {
	if num <= 0 {
		return
	}
	if pos < 0 || pos >= len(p.data) || pos+num > len(p.data) {
		fmt.Printf("delete index out of range [%d] with length %d\n", pos, len(p.data))
		return
	}
	p.data = append(p.data[:pos], p.data[pos+num:]...)
	p.emitDeleted(pos, num)
}

// Pop delete last element of data slice
func (p *RefSlice[T]) Pop() {
	if len(p.data) == 0 {
		return
	}
	pos := len(p.data) - 1
	p.data = p.data[:pos]
	p.emitDeleted(pos, 1)
}

func (p *RefSlice[T]) emitDeleted(pos int, num int) {
	if len(p.deleteHandlers) > 0 {
		for _, h := range p.deleteHandlers {
			if h.id != "" && h.handler != nil {
				h.handler(pos, num)
			}
		}
	}
}

func (p *RefSlice[T]) AddUpdatedHandler(handler func(pos int, val T), id ...string) (handlerId string, num int) {
	if len(id) > 0 && id[0] != "" {
		handlerId = id[0]
	} else {
		handlerId = NewId()
	}
	p.updateHandlers = append(p.updateHandlers, refSliceUpdatedHandler[T]{id: handlerId, handler: handler})
	num = len(p.updateHandlers)
	return
}

func (p *RefSlice[T]) RemoveUpdatedHandler(handlerId string) (ok bool, num int) {
	for i, h := range p.updateHandlers {
		if handlerId == h.id {
			ok = true
			p.updateHandlers = append(p.updateHandlers[:i], p.updateHandlers[i+1:]...)
			break
		}
	}
	num = len(p.updateHandlers)
	return
}

func (p *RefSlice[T]) AddInsertedHandler(handler func(pos int, val []T), id ...string) (handlerId string, num int) {
	if len(id) > 0 && id[0] != "" {
		handlerId = id[0]
	} else {
		handlerId = NewId()
	}
	p.insertHandlers = append(p.insertHandlers, refSliceInsertedHandler[T]{id: handlerId, handler: handler})
	num = len(p.insertHandlers)
	return
}

func (p *RefSlice[T]) RemoveInsertedHandler(handlerId string) (ok bool, num int) {
	for i, h := range p.insertHandlers {
		if handlerId == h.id {
			ok = true
			p.insertHandlers = append(p.insertHandlers[:i], p.insertHandlers[i+1:]...)
			break
		}
	}
	num = len(p.insertHandlers)
	return
}

func (p *RefSlice[T]) AddDeletedHandler(handler func(pos, num int), id ...string) (handlerId string, num int) {
	if len(id) > 0 && id[0] != "" {
		handlerId = id[0]
	} else {
		handlerId = NewId()
	}
	p.deleteHandlers = append(p.deleteHandlers, refSliceDeletedHandler{id: handlerId, handler: handler})
	num = len(p.deleteHandlers)
	return
}

func (p *RefSlice[T]) RemoveDeletedHandler(handlerId string) (ok bool, num int) {
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
