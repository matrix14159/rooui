package vdom

import (
	"log/slog"
	"testing"

	"github.com/matrix14159/rooui/dom"
)

// GOOS=js GOARCH=wasm go test -run TestCreateTag
func TestCreateTag(t *testing.T) {
	insertDiv(t, "root")
}

func insertDiv(t *testing.T, id string) *dom.Object {
	api := NewStandardDomApi()
	elms := dom.Document.GetElementsByTagName("body")
	if elms.Length() != 1 {
		t.Fatalf("html page miss body tag")
	}
	body := elms.Item(0)

	div := api.CreateElement("div")
	div.SetId(id)
	body.InsertBefore(div, nil)

	newDiv := dom.Document.GetElementById(id)
	if newDiv == nil {
		t.Fatalf("can't find the %v div", id)
	}
	slog.Info("insertDiv done.", "id", id)
	return newDiv
}
