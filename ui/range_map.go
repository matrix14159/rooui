package ui

import (
	"fmt"

	dom "honnef.co/go/js/dom/v2"
)

type RangeMap interface {
	MapController
}

type rangeMap[K comparable, V any] struct {
	baseController

	// map or *RefMap[K,V]
	source any

	refSource map[K]mapRefSource[K, V]

	// generate construct UI by each element of source
	iterFunc func(*Ref[K], *Ref[V]) UI

	// body ui
	subs  map[K]UI
	guard UI
}

type mapRefSource[K comparable, V any] struct {
	refKey *Ref[K]
	refVal *Ref[V]
}

func (p *rangeMap[K, V]) Kind() Kind {
	return RangeMapElem
}

func (p *rangeMap[K, V]) getBody() []UI {
	ret := make([]UI, 0, len(p.subs)+1)
	for _, sub := range p.subs {
		ret = append(ret, sub)
	}
	ret = append(ret, p.guard)
	return ret
}

func (p *rangeMap[K, V]) iter() {
	var data map[K]V
	switch p.source.(type) {
	case map[K]V:
		data = p.source.(map[K]V)
	case *RefMap[K, V]:
		data = p.source.(*RefMap[K, V]).data
	}

	p.refSource = make(map[K]mapRefSource[K, V], len(data))
	for key, val := range data {
		p.refSource[key] = mapRefSource[K, V]{
			refKey: ToRef(key),
			refVal: ToRef(val),
		}
	}

	p.subs = make(map[K]UI, len(data))
	for rawKey, one := range p.refSource {
		p.subs[rawKey] = p.iterFunc(one.refKey, one.refVal)
	}
}

func (p *rangeMap[K, V]) findGuardElement() (guard, parent dom.Element) {
	w := dom.GetWindow()
	d := w.Document()
	guard = d.GetElementByID(p.GetUIElementId() + "-map-guard")
	if guard == nil {
		Console.Error("map (%v) guard missing\n", p.GetUIElementId())
		return
	}
	parent = guard.ParentElement()
	if parent == nil {
		Console.Error("map (%v) guard parent missing\n", p.GetUIElementId())
		return
	}
	return
}

func (p *rangeMap[K, V]) handleUpserted(key K, val V) {
	// the ui for {key,val} is already exist, update value
	if _, found := p.subs[key]; found {
		p.refSource[key].refVal.Set(val)
		return
	}

	// new data, just push to end
	guard, parent := p.findGuardElement()
	if guard == nil || parent == nil {
		return
	}

	refItem := mapRefSource[K, V]{
		refKey: ToRef(key),
		refVal: ToRef(val),
	}
	p.refSource[key] = refItem

	child := p.iterFunc(refItem.refKey, refItem.refVal)
	p.subs[key] = child

	els := BuildHtml(child)
	for _, el := range els {
		parent.InsertBefore(el, guard)
	}
}

func (p *rangeMap[K, V]) handleDeleted(keys []K) {
	for _, key := range keys {
		if child := p.subs[key]; child != nil {
			child.doUnmounted()
		}
		delete(p.subs, key)
		delete(p.refSource, key)
	}
}

func Map[K comparable, V any](data any, f func(*Ref[K], *Ref[V]) UI) RangeMap {
	p := new(rangeMap[K, V])
	p.id = NewId()
	p.source = data
	p.iterFunc = f

	switch data.(type) {
	case map[K]V:
	case *RefMap[K, V]:
		rm := data.(*RefMap[K, V])
		rm.AddUpsertHandler(p.handleUpserted, p.GetUIElementId()+"upserted")
		rm.AddDeletedHandler(p.handleDeleted, p.GetUIElementId()+"deleted")
	default:
		fmt.Printf("data must be a map or *RefMap\n")
	}
	return p
}

func (p *rangeMap[K, V]) BuildTreeDomElement() []dom.Element {
	fmt.Printf("map (%v) BuildTreeDomElement\n", p.GetUIElementId())
	p.iter()
	ret := make([]dom.Element, 0, len(p.subs)+1)
	for _, child := range p.subs {
		els := BuildHtml(child)
		ret = append(ret, els...)
	}
	p.guard = Div().Id(p.GetUIElementId() + "-map-guard") // last element for guard
	ret = append(ret, BuildHtml(p.guard)...)
	return ret
}

func (p *rangeMap[K, V]) doCreated() {
	for _, child := range p.subs {
		child.doCreated()
	}
	p.guard.doCreated()

	if p.onCreatedHandler != nil {
		p.onCreatedHandler()
	}
}

func (p *rangeMap[K, V]) doMounted() {
	if p.onMountedHandler != nil {
		p.onMountedHandler()
	}

	// NOTE: condition doesn't use baseController.body
	for _, child := range p.subs {
		child.doMounted()
	}
	p.guard.doMounted()
}

func (p *rangeMap[K, V]) doUnmounted() {
	// child first
	p.guard.doUnmounted()
	p.guard = nil

	for _, child := range p.subs {
		child.doUnmounted()
	}
	p.subs = nil

	//
	if rm, ok := p.source.(*RefMap[K, V]); ok {
		rm.RemoveUpsertHandler(p.GetUIElementId() + "upserted")
		rm.RemoveDeletedHandler(p.GetUIElementId() + "deleted")
	}
	p.source = nil
	p.refSource = nil

	if p.onUnmountedHandler != nil {
		p.onUnmountedHandler()
	}
	fmt.Printf("map (%v) unmounted is done.\n", p.GetUIElementId())
}
