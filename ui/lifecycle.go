package ui

import (
	"fmt"
	"syscall/js"

	dom "honnef.co/go/js/dom/v2"
)

var RootComponent Comp

func Init(root Comp) {
	RootComponent = root
	initConsole()
	initLog()
	js.Global().Set("MountTo", js.FuncOf(mountToFunc))
}

func mountToFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(mountTo(args[0].String()))
}

func mountTo(root string) string {
	if RootComponent == nil {
		return fmt.Sprintf("root-component not set")
	}

	w := dom.GetWindow()
	d := w.Document()
	p := d.GetElementByID(root)

	hui := BuildHtml(RootComponent)
	RootComponent.doCreated()

	for _, u := range hui {
		p.AppendChild(u)
	}
	RootComponent.doMounted()
	fmt.Println("mount is done")
	return ""
}

func WaitExit() {
	select {}
}
