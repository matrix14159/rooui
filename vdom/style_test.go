package vdom

import (
	"log/slog"
	"testing"
	"time"
)

// GOOS=js GOARCH=wasm go test -run TestRaf
func TestNextFrame(t *testing.T) {
	f := func() {
		slog.Info("call next frame")
	}
	NextFrame(f)
	time.Sleep(1 * time.Second)
}

// GOOS=js GOARCH=wasm go test -run TestPatchStyle
func TestPatchStyle(t *testing.T) {
	root := insertDiv(t, "root")
	oldVnode := EmptyNodeAt(root)

	s := NewVNodeStyle()
	s.Style["color"] = "red"
	vnode := H("div", &VNodeData{Style: s}, "hello", nil)

	p := NewPatcher(NewStandardDomApi())
	_, err := p.Patch(oldVnode, vnode)
	if err != nil {
		t.Fatalf("patch failed. error:%v", err)
	}
	if vnode.Elm == nil {
		t.Fatalf("after patch, vnode.Elm must not be nil")
	}
	time.Sleep(1 * time.Second)
}
