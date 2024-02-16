package vdom

import (
	"testing"
)

// GOOS=js GOARCH=wasm go test -run TestPatch
func TestPatch(t *testing.T) {
	p := NewPatcher(NewStandardDomApi(), VNode{})
	err := p.Patch(VNode{})
	if err != nil {
		return
	}
}
