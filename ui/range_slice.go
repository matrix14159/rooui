package ui

import (
	"fmt"

	dom "honnef.co/go/js/dom/v2"
)

type RangeSlice interface {
	SliceController
}

type rangeSlice[T any] struct {
	baseController

	// slice data
	source any

	refSource []sliceRefSource[T]

	// generate construct UI by each element of source
	iterFunc func(*Ref[int], *Ref[T]) UI
}

type sliceRefSource[T any] struct {
	refIndex *Ref[int]
	refValue *Ref[T]
}

func (p *rangeSlice[T]) Kind() Kind {
	return RangeSliceElem
}

func (p *rangeSlice[T]) iter() (ret []UI) {
	var data []T
	switch p.source.(type) {
	case []T:
		data = p.source.([]T)
	case *RefSlice[T]:
		data = p.source.(*RefSlice[T]).data
	}

	p.refSource = make([]sliceRefSource[T], len(data))
	for i, one := range data {
		p.refSource[i] = sliceRefSource[T]{
			refIndex: ToRef(i),
			refValue: ToRef(one),
		}
	}

	ret = make([]UI, 0, len(data)+1)
	for _, one := range p.refSource {
		u := p.iterFunc(one.refIndex, one.refValue)
		ret = append(ret, u)
	}
	ret = append(ret, Div().Id(p.GetUIElementId()+"-slice-guard")) // last element for guard
	p.body = ret
	return ret
}

func (p *rangeSlice[T]) findElementByIndex(pos int) (el, parent dom.Element) {
	if pos < 0 || pos >= len(p.body) {
		Console.Error("slice (%v) find element index out of range [%d] with length %d\n", p.GetUIElementId(), pos, len(p.body))
		return
	}

	u := p.body[pos]
	el = u.GetSelfDomElement()
	if el == nil {
		Console.Error("slice (%v) can't find dom element for item:%v\n", p.GetUIElementId(), u.GetUIElementId())
		return
	}
	parent = el.ParentElement()
	if parent == nil {
		Console.Error("slice (%v) can't find parent dom element for item:%v\n", p.GetUIElementId(), u.GetUIElementId())
		return
	}
	return
}

func (p *rangeSlice[T]) handleInserted(pos int, val []T) {
	if len(val) == 0 {
		return
	}
	if pos < 0 || pos >= len(p.body) {
		Console.Error("slice (%v) handle inserted with bad pos:%v\n", p.GetUIElementId(), pos)
		return
	}

	refItems := make([]sliceRefSource[T], len(val))
	for i, one := range val {
		refItems[i] = sliceRefSource[T]{
			refIndex: ToRef(i),
			refValue: ToRef(one),
		}
	}

	p.refSource = append(p.refSource[:pos], append(refItems, p.refSource[pos:]...)...)
	for i := pos + len(val); i < len(p.refSource); i++ {
		p.refSource[i].refIndex.Set(i)
	}

	p.insertNews(pos, refItems)
}

func (p *rangeSlice[T]) insertNews(pos int, refItems []sliceRefSource[T]) {
	before, parent := p.findElementByIndex(pos)
	if before == nil || parent == nil {
		return
	}

	news := make([]UI, 0, len(refItems))
	for i, item := range refItems {
		item.refIndex.Set(pos + i)
		child := p.iterFunc(item.refIndex, item.refValue)
		news = append(news, child)

		els := BuildHtml(child)
		for _, el := range els {
			parent.InsertBefore(el, before)
		}
	}
	p.body = append(p.body[:pos], append(news, p.body[pos:]...)...)
}

func (p *rangeSlice[T]) handleDeleted(pos int, num int) {
	//fmt.Printf("slice (%v) handle deleted. pos:%v, num:%v\n", p.GetUIElementId(), pos, num)
	if num == 0 {
		return
	}

	if pos < 0 || pos >= len(p.body) || pos+num > len(p.body) {
		fmt.Printf("slice (%v) handle deleted index out of range [%d] with length %d\n", p.GetUIElementId(), pos, len(p.body))
		return
	}

	p.refSource = append(p.refSource[:pos], p.refSource[pos+num:]...)
	for i := pos; i < len(p.refSource); i++ {
		p.refSource[i].refIndex.Set(i)
	}

	dels := make([]UI, 0, num)
	for i := 0; i < num; i++ {
		dels = append(dels, p.body[pos+i])
	}
	p.body = append(p.body[:pos], p.body[pos+num:]...)

	for _, one := range dels {
		one.doUnmounted()
	}
}

func (p *rangeSlice[T]) handleCleared() {
	for _, child := range p.body {
		child.doUnmounted()
	}
	p.refSource = make([]sliceRefSource[T], 0)
	p.body = make([]UI, 0)
}

func (p *rangeSlice[T]) handleUpdated(pos int, val T) {
	if pos < 0 || pos >= len(p.body) {
		fmt.Printf("slice (%v) handle updated index out of range [%d] with length %d\n", p.GetUIElementId(), pos, len(p.body))
		return
	}
	p.refSource[pos].refValue.Set(val)
}

// Slice will range data by function f to create ui element
func Slice[T any](data any, f func(i *Ref[int], v *Ref[T]) UI) RangeSlice {
	p := new(rangeSlice[T])
	p.id = NewId()
	p.source = data
	p.iterFunc = f

	switch data.(type) {
	case []T:
	case *RefSlice[T]:
		rs := data.(*RefSlice[T])
		rs.AddInsertedHandler(p.handleInserted, p.GetUIElementId()+"inserted")
		rs.AddDeletedHandler(p.handleDeleted, p.GetUIElementId()+"deleted")
		rs.AddUpdatedHandler(p.handleUpdated, p.GetUIElementId()+"updated")
		rs.AddClearedHandler(p.handleCleared, p.GetUIElementId()+"cleared")
	default:
		fmt.Printf("data must be a slice or *RefSlice\n")
	}
	return p
}

func (p *rangeSlice[T]) BuildTreeDomElement() []dom.Element {
	//fmt.Printf("slice (%v) BuildTreeDomElement\n", p.GetUIElementId())
	p.body = p.iter()
	ret := make([]dom.Element, 0, len(p.body))
	for _, child := range p.body {
		els := BuildHtml(child)
		ret = append(ret, els...)
	}
	return ret
}

func (p *rangeSlice[T]) doUnmounted() {
	// child first
	for _, child := range p.body {
		child.doUnmounted()
	}
	p.body = nil

	//
	if rs, ok := p.source.(*RefSlice[T]); ok {
		rs.RemoveInsertedHandler(p.GetUIElementId() + "inserted")
		rs.RemoveDeletedHandler(p.GetUIElementId() + "deleted")
		rs.RemoveUpdatedHandler(p.GetUIElementId() + "updated")
		rs.RemoveClearedHandler(p.GetUIElementId() + "cleared")
	}
	if p.onUnmountedHandler != nil {
		p.onUnmountedHandler()
	}
	fmt.Printf("slice (%v) unmounted is done.\n", p.GetUIElementId())
}
