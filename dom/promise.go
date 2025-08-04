package dom

import (
	"errors"
	"syscall/js"
)

// Promise is an instance of a JS promise.
// The zero value of this struct is not a valid Promise.
type Promise struct {
	js.Value
}

// FromJSValue turns a JS value to a Promise.
func (p *Promise) FromJSValue(value js.Value) error {
	p.Value = value
	return nil
}

// NewPromise returns a promise that is fulfilled or rejected when the provided handler returns.
// The handler is spawned in its own goroutine.
func NewPromise(handler func() (interface{}, error)) Promise {
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Invoke the handler in a new goroutine.
	go func() {
		result, err := handler()
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- result
	}()

	// Create a JS promise handler.
	var jsHandler js.Func
	jsHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			panic("not enough arguments are passed to the Promise constructor handler")
		}

		resolve := args[0]
		reject := args[1]

		if resolve.Type() != js.TypeFunction || reject.Type() != js.TypeFunction {
			panic("invalid type passed to Promise constructor handler")
		}

		go func() {
			select {
			case r := <-resultChan:
				resolve.Invoke(ToJSValue(r))
			case err := <-errChan:
				reject.Invoke(NewError(err))
			}

			// Free up resources now that we are done.
			jsHandler.Release()
		}()

		return nil
	})

	promise := js.Global().Get("Promise")
	return mustJSValueToPromise(promise.New(jsHandler))
}

// Await2 waits for the Promise. It unmarshals the resolved value to v. An error
// will be returned if unmarshalling is unsuccessful or the Promise rejects.
// It is implemented by calling then and catch on JS.
func (p Promise) Await2(v interface{}) error {
	err := make(chan error)
	p.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 && v != nil {
			err <- FromJSValue(args[0], v)
			return nil
		}
		err <- nil
		return nil
	}))
	p.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err <- errors.New(args[0].Call("toString").String())
		return nil
	}))
	return <-err
}

// Await waits for the Promise and return raw js.Value
func (p Promise) Await() (val js.Value, err error) {
	errCh := make(chan error)
	p.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			val = args[0]
			errCh <- nil
			return nil
		}
		errCh <- nil
		return nil
	}))
	p.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		errCh <- errors.New(args[0].Call("toString").String())
		return nil
	}))
	err = <-errCh
	return
}

// PromiseAll creates a promise that is fulfilled when all the provided promises have been fulfilled.
// The promise is rejected when any of the promises provided rejects.
// It is implemented by calling Promise.all on JS.
func PromiseAll(promise ...Promise) Promise {
	promiseAll := js.Global().Get("Promise").Get("all")

	pInterface := make([]interface{}, 0, len(promise))
	for _, v := range promise {
		pInterface = append(pInterface, v)
	}

	return mustJSValueToPromise(promiseAll.Invoke(pInterface))
}

// PromiseAllSettled creates a promise that is fulfilled when all the provided promises have been fulfilled or rejected.
// It is implemented by calling Promise.allSettled on JS.
func PromiseAllSettled(promise ...Promise) Promise {
	promiseAllSettled := js.Global().Get("Promise").Get("allSettled")

	pInterface := make([]interface{}, 0, len(promise))
	for _, v := range promise {
		pInterface = append(pInterface, v)
	}

	return mustJSValueToPromise(promiseAllSettled.Invoke(pInterface))
}

// PromiseAny creates a promise that is fulfilled when any of the provided promises have been fulfilled.
// The promise is rejected when all of the provided promises gets rejected.
// It is implemented by calling Promise.any on JS.
func PromiseAny(promise ...Promise) Promise {
	promiseAny := js.Global().Get("Promise").Get("any")

	pInterface := make([]interface{}, 0, len(promise))
	for _, v := range promise {
		pInterface = append(pInterface, v)
	}

	return mustJSValueToPromise(promiseAny.Invoke(pInterface))
}

// PromiseRace creates a promise that is fulfilled or rejected when one of the provided promises fulfill or reject.
// It is implemented by calling Promise.race on JS.
func PromiseRace(promise ...Promise) Promise {
	promiseRace := js.Global().Get("Promise").Get("race")

	pInterface := make([]interface{}, 0, len(promise))
	for _, v := range promise {
		pInterface = append(pInterface, v)
	}

	return mustJSValueToPromise(promiseRace.Invoke(pInterface))
}

func mustJSValueToPromise(v js.Value) Promise {
	var p Promise
	err := p.FromJSValue(v)
	if err != nil {
		panic("Expected a Promise from JS standard library")
	}

	return p
}
