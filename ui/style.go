package ui

import (
	"fmt"
	"strings"
)

type StyleItem interface {
	Id() string
	Name() string
	Value() string
	ValueWithUnit() string
	ToStyle() string

	RemoveValueUpdatedHandler()
}

type StyleGroup struct {
	ClassId   string            // <style id=ClassId />
	ClassName string            // <style class=ClassName />
	Items     []*Ref[StyleItem] // style items
	Use       *Ref[bool]        // use or not, default is true
}

func (p StyleGroup) ToClass() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s {", p.ClassName))
	for _, item := range p.Items {
		b.WriteString(item.Value().ToStyle())
	}
	b.WriteString("}")
	return b.String()
}
