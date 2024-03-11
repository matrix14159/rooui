package vdom

import (
	"testing"
)

// GOOS=js GOARCH=wasm go test -run TestPatch
func TestPatch(t *testing.T) {
	root := insertDiv(t, "root")
	p := NewPatcher(NewStandardDomApi(), EmptyNodeAt(root))

	vnode := H("div", nil, "hello", nil)
	err := p.Patch(vnode)
	if err != nil {
		return
	}
}

// GOOS=js GOARCH=wasm go test -run TestPatch2
func TestPatch2(t *testing.T) {
	root := insertDiv(t, "root")
	p := NewPatcher(NewStandardDomApi(), EmptyNodeAt(root))

	vnode := H("div#app", nil, "hello", nil)
	err := p.Patch(vnode)
	if err != nil {
		return
	}
}

// GOOS=js GOARCH=wasm go test -run TestPatch3
func TestPatch3(t *testing.T) {
	root := insertDiv(t, "root")
	p := NewPatcher(NewStandardDomApi(), EmptyNodeAt(root))

	vnode := H("div.btn.mart24", nil, "hello", nil)
	err := p.Patch(vnode)
	if err != nil {
		return
	}
}
