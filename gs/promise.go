package gs

import (
	"errors"
	"syscall/js"
)

type Promise struct {
	Value js.Value
}

func (p Promise) Await() (ret js.Value, err error) {
	done := make(chan any)
	p.Value.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		defer close(done)
		if len(args) > 0 {
			ret = args[0]
		}
		return nil
	}))
	p.Value.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
		defer close(done)
		err = errors.New(args[0].Call("toString").String())
		return nil
	}))
	<-done
	return
}
