package vdom

import (
	"fmt"
	"testing"

	"github.com/oklog/ulid/v2"
)

// GOOS=js GOARCH=wasm go test -run TestPatch
func TestPatch(t *testing.T) {
	p := NewPatcher(NewStandardDomApi(), &VNode{})
	err := p.Patch(&VNode{})
	if err != nil {
		return
	}

	fmt.Println("ulid: ", ulid.Make().String())
}
