package ui

import (
	"log/slog"
	"syscall/js"

	"github.com/matrix14159/rooui/vdom"
	"github.com/matrix14159/sharp"
	"honnef.co/go/js/dom/v2"
)

var RootComponent Comp

func Run(c Comp) {
	RootComponent = c

	initLog()

	js.Global().Set("MountTo", js.FuncOf(mountToFunc))

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

	w := dom.GetWindow()
	d := w.Document()
	r := d.GetElementByID(root)

	oldVNode := vdom.EmptyNodeAt(r)

	element := RootComponent.Render()
	if element == nil {
		slog.Error("root component render nil")
		return ""
	}
	vnode := vdom.H(element.Tag(), "", nil, nil)

	p := vdom.NewPatcher(vdom.NewStandardDomApi())
	old, err := p.Patch(oldVNode, vnode)
	if err != nil {
		slog.Error("mount patch failed.", "error", err)
		return ""
	}
	RootComponent.updateVNode(old)

	slog.Info("mount is done")
	return ""
}
