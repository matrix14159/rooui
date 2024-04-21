package vdom

import (
	"log/slog"

	"github.com/matrix14159/sharp"
)

// 在执行wasm单元测试之前，需先安装 github.com/agnivade/wasmbrowsertest
// 然后重命名为 go_js_wasm_exec.exe
func init() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler := sharp.NewSimpleHandler(sharp.NewWasmWriter(), opts, "2006-01-02T15:04:05.000", false)
	slog.SetDefault(slog.New(handler))
}
