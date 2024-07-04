package vdom

import (
	"testing"
)

// GOOS=js GOARCH=wasm go test -run TestPatch
func TestPatch(t *testing.T) {
	root := insertDiv(t, "root")
	oldVnode := EmptyNodeAt(root)

	vnode := H("div", "hello", nil, nil)

	p := NewPatcher(NewStandardDomApi())
	_, err := p.Patch(oldVnode, vnode)
	if err != nil {
		t.Fatalf("patch failed. error:%v", err)
	}
	if vnode.Elm == nil {
		t.Fatalf("after patch, vnode.Elm must not be nil")
	}
}

// GOOS=js GOARCH=wasm go test -run TestPatch2
func TestPatch2(t *testing.T) {
	root := insertDiv(t, "root")
	oldVnode := EmptyNodeAt(root)

	vnode := H("div#app", "hello", nil, nil)

	p := NewPatcher(NewStandardDomApi())
	_, err := p.Patch(oldVnode, vnode)
	if err != nil {
		t.Fatalf("patch failed. error:%v", err)
	}
	if vnode.Elm == nil {
		t.Fatalf("after patch, vnode.Elm must not be nil")
	}
}

// GOOS=js GOARCH=wasm go test -run TestPatch3
func TestPatch3(t *testing.T) {
	root := insertDiv(t, "root")
	oldVnode := EmptyNodeAt(root)

	vnode := H("div.btn.mart24", "hello", nil, nil)

	p := NewPatcher(NewStandardDomApi())
	_, err := p.Patch(oldVnode, vnode)
	if err != nil {
		t.Fatalf("patch failed. error:%v", err)
	}
	if vnode.Elm == nil {
		t.Fatalf("after patch, vnode.Elm must not be nil")
	}
}
