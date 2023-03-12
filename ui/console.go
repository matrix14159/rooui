package ui

import (
	"fmt"
	"syscall/js"
)

// Console represent as browser Window.Console
var Console *console

type console struct {
	js.Value
}

func initConsole() {
	Console = &console{js.Global().Get("console")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/assert
func (p *console) Assert(condition bool, a ...any) {
	p.Call("assert", condition, js.ValueOf(a))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/clear
func (p *console) Clear() {
	p.Call("clear")
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/count
func (p *console) Count(label string) {
	p.Call("count", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/countReset
func (p *console) CountReset(label string) {
	p.Call("countReset", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/debug
func (p *console) Debug(msg string, a ...any) {
	p.Call("debug", fmt.Sprintf(msg, a...))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/dir
func (p *console) Dir(obj any) {
	p.Call("dir", js.ValueOf(obj))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/dirxml
func (p *console) Dirxml(obj any) {
	p.Call("dirxml", obj)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/error
func (p *console) Error(msg string, a ...any) {
	p.Call("error", fmt.Sprintf(msg, a...))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/group
func (p *console) Group(label string) {
	p.Call("group", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/groupCollapsed
func (p *console) GroupCollapsed(label string) {
	p.Call("groupCollapsed", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/groupEnd
func (p *console) groupEnd() {
	p.Call("groupEnd")
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/info
func (p *console) Info(msg string, a ...any) {
	p.Call("info", fmt.Sprintf(msg, a...))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/log
func (p *console) Log(msg string, a ...any) {
	p.Call("log", fmt.Sprintf(msg, a...))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/table
func (p *console) Table(data any) {
	p.Call("table", js.ValueOf(data))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/time
func (p *console) Time(label string) {
	p.Call("time", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/timeEnd
func (p *console) TimeEnd(label string) {
	p.Call("timeEnd", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/timeLog
func (p *console) TimeLog(label string) {
	p.Call("timeLog", label)
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/trace
func (p *console) Trace(msg string, a ...any) {
	p.Call("trace", fmt.Sprintf(msg, a...))
}

// https://developer.mozilla.org/en-US/docs/Web/API/console/warn
func (p *console) Warn(msg string, a ...any) {
	p.Call("warn", fmt.Sprintf(msg, a...))
}
