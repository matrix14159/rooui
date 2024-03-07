package vdom

import (
	"log/slog"

	"github.com/matrix14159/sharp"
)

func init() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler := sharp.NewSimpleHandler(sharp.NewWasmWriter(), opts, "2006-01-02T15:04:05.000", false)
	slog.SetDefault(slog.New(handler))
}
