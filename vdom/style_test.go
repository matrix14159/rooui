package vdom

import (
	"log/slog"
	"testing"
	"time"
)

// GOOS=js GOARCH=wasm go test -run TestRaf
func TestRaf(t *testing.T) {
	f := func() {
		slog.Info("call raf")
	}
	NextFrame(f)
	time.Sleep(1 * time.Second)
}
