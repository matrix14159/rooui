package st

import (
	"fmt"

	"sunny/rooui/ui"
)

type styleItem struct {
	id     string
	name   string
	value  any
	unit   string
	pseudo bool

	vBinder ui.Binder // construct input is a binder
}

func (p styleItem) Id() string {
	return p.id
}

func (p styleItem) Name() string {
	return p.name
}

func (p styleItem) Value() string {
	return fmt.Sprintf("%v", p.value)
}

func (p styleItem) ValueWithUnit() string {
	return fmt.Sprintf("%v%v", p.value, p.unit)
}

func (p styleItem) ToStyle() string {
	return fmt.Sprintf("%v: %v%v;", p.name, p.value, p.unit)
}

func (p styleItem) RemoveValueUpdatedHandler() {
	if p.vBinder != nil {
		p.vBinder.RemoveUpdatedHandler(p.id)
	}
}

func makeRefStyle(name string, v any, u ...string) *ui.Ref[ui.StyleItem] {
	val := v
	binder, ok := v.(ui.Binder)
	if ok {
		val = binder.Data()
	}

	unit := ""
	if len(u) > 0 {
		unit = u[0]
	}

	data := &styleItem{
		id:    ui.NewId(),
		name:  name,
		value: val,
		unit:  unit,
	}
	p := ui.ToRef[ui.StyleItem](data)

	if ok {
		data.vBinder = binder
		binder.AddUpdatedHandler(func(oldVal, newVal any) {
			data.value = newVal
			p.Set(data)
		}, data.id)
	}
	return p
}
