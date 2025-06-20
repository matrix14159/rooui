package vdom

import (
	"testing"

	"github.com/matrix14159/rooui/dom"
)

// GOOS=js GOARCH=wasm go test -run TestH
func TestH(t *testing.T) {
	sel := "div"
	vnode := H(sel, "hello", nil, nil)
	if vnode.Sel != sel {
		t.Fatalf("sel must be %v, but got:%v", sel, vnode.Sel)
	}
	if vnode.Text != "hello" {
		t.Fatalf("text must be hello, but got:%v", vnode.Text)
	}
}

// GOOS=js GOARCH=wasm go test -run TestEmptyNodeAt
func TestEmptyNodeAt(t *testing.T) {
	api := NewStandardDomApi()
	elms := dom.Document.GetElementsByTagName("body")
	if elms.Length() != 1 {
		t.Fatalf("html page miss body tag")
	}
	body := elms.Item(0)

	div := api.CreateElement("div")
	div.SetId("root")
	body.InsertBefore(div, nil)

	root := dom.Document.GetElementById("root")
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
