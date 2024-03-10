package vdom

import (
	"testing"
)

// GOOS=js GOARCH=wasm go test -run TestPatch
func TestPatch(t *testing.T) {
	root := insertDiv(t, "root")
	p := NewPatcher(NewStandardDomApi(), EmptyNodeAt(root))
	err := p.Patch(&VNode{})
	if err != nil {
		return
	}
}
