package ui

import (
	"fmt"

	dom "honnef.co/go/js/dom/v2"
)

type Condition interface {
	ConditionController

	ElseIf(expr any, elems ...UI) Condition

	Else(elems ...UI) UI
}

type condition struct {
	baseController

	statements []conditionStatement
}

type conditionStatement struct {
	statement string // if, elseif, else
	expr      any
	elems     []UI
}

func (p *condition) Kind() Kind {
	return ConditionControlElem
}

func (p *condition) getBody() []UI {
	ret := make([]UI, 0, p.elemLen())
	for _, s := range p.statements {
		ret = append(ret, s.elems...)
	}
	return ret
}

func (p *condition) elems() []UI {
	ret := make([]UI, 0, p.elemLen())
	for _, s := range p.statements {
		ret = append(ret, s.elems...)
	}
	return ret
}

func (p *condition) elemLen() int {
	num := 0
	for _, s := range p.statements {
		num += len(s.elems)
	}
	return num
}

func (p *condition) match() {
	satisfiedIdx := -1
	for i, s := range p.statements {
		satisfied := false
		switch s.expr.(type) {
		case bool:
			satisfied = s.expr.(bool)
		case *Ref[bool]:
			satisfied = s.expr.(*Ref[bool]).Value()
		case func() bool:
			satisfied = (s.expr.(func() bool))()
		}
		if satisfied {
			satisfiedIdx = i
			break // get the first satisfied statement
		}
	}
	if satisfiedIdx < 0 {
		if p.statements[len(p.statements)-1].statement == "else" {
			satisfiedIdx = len(p.statements) - 1
		}
	}

	for i, s := range p.statements {
		if i == satisfiedIdx {
			p.turnOn(s)
		} else {
			p.turnOff(s)
		}
	}
}

func (p *condition) turnOn(s conditionStatement) {
	for _, u := range s.elems {
		u.TurnOnDisplay()
	}
}

func (p *condition) turnOff(s conditionStatement) {
	for _, u := range s.elems {
		u.TurnOffDisplay()
	}
}

// If return a Condition element by filter expression
// expr result must be bool or *Ref[bool], or func() bool
func If(expr any, elems ...UI) Condition {
	p := new(condition)
	p.id = NewId()
	s := conditionStatement{
		statement: "if",
		expr:      expr,
		elems:     elems,
	}
	p.statements = make([]conditionStatement, 1)
	p.statements[0] = s

	if r, ok := expr.(*Ref[bool]); ok {
		r.AddUpdatedHandler(func(oldVal, newVal any) {
			p.match()
		}, p.GetUIElementId()+s.statement)
	}
	return p
}

// ElseIf return a Condition element by filter expression
// expr result must be bool or *Ref[bool], or func() bool
func (p *condition) ElseIf(expr any, elems ...UI) Condition {
	s := conditionStatement{
		statement: "elseif" + NewId(),
		expr:      expr,
		elems:     elems,
	}
	p.statements = append(p.statements, s)

	if r, ok := expr.(*Ref[bool]); ok {
		r.AddUpdatedHandler(func(oldVal, newVal any) {
			p.match()
		}, p.GetUIElementId()+s.statement)
	}
	return p
}

func (p *condition) Else(elems ...UI) UI {
	statement := conditionStatement{
		statement: "else",
		expr:      nil,
		elems:     elems,
	}
	p.statements = append(p.statements, statement)
	return p
}

func (p *condition) GetSelfDomElement() dom.Element {
	for _, child := range p.elems() {
		if el := child.GetSelfDomElement(); el != nil {
			return el
		}
	}
	return nil
}

func (p *condition) BuildTreeDomElement() []dom.Element {
	//fmt.Printf("condition (%v) BuildTreeDomElement\n", p.GetUIElementId())
	ret := make([]dom.Element, 0, p.elemLen())
	for _, child := range p.elems() {
		els := BuildHtml(child)
		ret = append(ret, els...)
	}
	p.match()
	return ret
}

func (p *condition) doCreated() {
	for _, child := range p.elems() {
		child.doCreated()
	}
	if p.onCreatedHandler != nil {
		p.onCreatedHandler()
	}
}

func (p *condition) doMounted() {
	if p.onMountedHandler != nil {
		p.onMountedHandler()
	}

	// NOTE: condition doesn't use baseController.body
	for _, child := range p.elems() {
		child.doMounted()
	}
}

func (p *condition) doUnmounted() {
	// child first
	for _, child := range p.elems() {
		child.doUnmounted()
	}

	//
	for _, s := range p.statements {
		if r, ok := s.expr.(*Ref[bool]); ok {
			r.RemoveUpdatedHandler(p.GetUIElementId() + s.statement)
		}
	}
	if p.onUnmountedHandler != nil {
		p.onUnmountedHandler()
	}
	fmt.Printf("condition (%v) unmounted is done.\n", p.GetUIElementId())
}
