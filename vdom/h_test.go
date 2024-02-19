package vdom

import (
	"testing"
)

// GOOS=js GOARCH=wasm go test -run TestEmptyNodeAt
func TestEmptyNodeAt(t *testing.T) {
	api := NewStandardDomApi()
	elms := api.document.GetElementsByTagName("body")
	if len(elms) != 1 {
		t.Fatalf("html page miss body tag")
	}
	body := elms[0]

	div := api.CreateElement("div")
	div.SetID("root")
	body.InsertBefore(div, nil)

	root := api.document.GetElementByID("root")
	if root == nil {
		t.Fatalf("can't find the root div")
	}

	vnode := EmptyNodeAt(root)
	expectSel := "div#root"
	if vnode.Sel != expectSel {
		t.Fatalf("vnode.Sel must be:%v, but got:%v", expectSel, vnode.Sel)
	}
	if vnode.Elm != root {
		t.Fatalf("vnode.Elm must be the same as root")
	}
}
