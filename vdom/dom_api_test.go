package vdom

import (
	"testing"
)

// GOOS=js GOARCH=wasm go test -run TestCreateTag
func TestCreateTag(t *testing.T) {
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
}
