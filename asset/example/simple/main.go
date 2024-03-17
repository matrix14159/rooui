package main

import (
	"log/slog"

	"github.com/matrix14159/rooui/vdom"
	"github.com/matrix14159/sharp"
	"honnef.co/go/js/dom/v2"
)

// run steps:
// 1. enter example/simple directory, compile to wasm: GOOS=js GOARCH=wasm go build -o=./server/public/main.wasm
// 2. enter example/simple/server, start http server: go run main.go
// 3. open chrome, navigate to http://localhost:12000/
// 4. check chrome console
func main() {
	initLog()

	root := insertDiv("root")

	p := vdom.NewPatcher(vdom.NewStandardDomApi(), vdom.EmptyNodeAt(root))

	vnode := vdom.H("div", nil, "hello", nil)
	err := p.Patch(vnode)
	if err != nil {
		slog.Error("path failed.", "error", err)
	}
}

func initLog() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler := sharp.NewSimpleHandler(sharp.NewWasmWriter(), opts, "2006-01-02T15:04:05.000", false)

	slog.SetDefault(slog.New(handler))
}

func insertDiv(id string) dom.Element {
	api := vdom.NewStandardDomApi()
	elms := api.Document.GetElementsByTagName("body")
	if len(elms) != 1 {
		slog.Error("html page miss body tag")
		return nil
	}
	body := elms[0]

	div := api.CreateElement("div")
	div.SetID(id)
	body.InsertBefore(div, nil)

	newDiv := api.Document.GetElementByID(id)
	if newDiv == nil {
		slog.Error("can't find the div", "id", id)
	}
	slog.Info("insertDiv done.", "id", id)
	return newDiv
}
