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
	RootComponent = c

	initLog()

	js.Global().Set("MountTo", js.FuncOf(mountToFunc))

	defaultUpdateFlow = new(updateFlow)
	defaultUpdateFlow.RunUpdateLoop()

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

func mountToFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(mountTo(args[0].String()))
}

func mountTo(root string) string {
	if RootComponent == nil {
		slog.Error("root component not set")
		return ""
	}

	r := dom.Document.GetElementById(root)

	oldVNode := vdom.EmptyNodeAt(r)
	RootComponent.updateVNode(oldVNode)

	Update(RootComponent)

	slog.Info("root component mount done")
	return ""
}
