package ui

import (
	"log/slog"
	"syscall/js"

	"github.com/matrix14159/rooui/dom"
	"github.com/matrix14159/rooui/vdom"
	"github.com/matrix14159/sharp"
)

var RootComponent Comp

func Run(c Comp) {
	initLog()

	RootComponent = c

	// init for Component.Uid()
	dom.Window.Set("u_id", 0)

	// Use js.FuncOf for better string parameter handling
	js.Global().Set("MountTo", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			MountTo(args[0].String())
		}
		return nil
	}))

	defaultUpdateFlow = newUpdateFlow()

	// Keep the program running to respond to JavaScript calls
	select {}
}

// if run by unit-test, before run please install "github.com/agnivade/wasmbrowsertest"
// then rename to go_js_wasm_exec.exe
func initLog() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler := sharp.NewSimpleHandler(sharp.NewWasmWriter(), opts, "2006-01-02T15:04:05.000", false)
	slog.SetDefault(slog.New(handler))
}

// MountTo mounts the root component to the specified element ID
func MountTo(root string) {
	if RootComponent == nil {
		slog.Error("root component not set")
		return
	}

	if root == "" {
		slog.Error("root parameter is empty")
		return
	}

	r := dom.Document.GetElementById(root)
	if r == nil {
		slog.Error("root element not found", "root", root)
		return
	}

	oldVNode := vdom.EmptyNodeAt(r)
	RootComponent.updateVNode(oldVNode)

	Update(RootComponent)

	slog.Info("root component mount done")
}
