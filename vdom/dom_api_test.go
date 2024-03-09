package vdom

import (
	"log/slog"
	"testing"

	"honnef.co/go/js/dom/v2"
)

// GOOS=js GOARCH=wasm go test -run TestCreateTag
func TestCreateTag(t *testing.T) {
	insertDiv(t, "root")
	//api := NewStandardDomApi()
	//elms := api.document.GetElementsByTagName("body")
	//if len(elms) != 1 {
	//	t.Fatalf("html page miss body tag")
	//}
	//body := elms[0]
	//
	//div := api.CreateElement("div")
	//div.SetID("root")
	//body.InsertBefore(div, nil)
	//
	//root := api.document.GetElementByID("root")
	//if root == nil {
	//	t.Fatalf("can't find the root div")
	//}
	//slog.Info("TestCreateTag done.")
}

func insertDiv(t *testing.T, id string) dom.Element {
	api := NewStandardDomApi()
	elms := api.document.GetElementsByTagName("body")
	if len(elms) != 1 {
		t.Fatalf("html page miss body tag")
	}
	body := elms[0]

	div := api.CreateElement("div")
	div.SetID(id)
	body.InsertBefore(div, nil)

	newDiv := api.document.GetElementByID(id)
	if newDiv == nil {
		t.Fatalf("can't find the %v div", id)
	}
	slog.Info("insertDiv done.", "id", id)
	return newDiv
}
