package dom

import (
	"sync"
	"sync/atomic"
	"syscall/js"
)

// Promise is an instance of a JS promise.
// The zero value of this struct is not a valid Promise.
type Promise struct {
	js.Value
}

// NewRawPromise wrap JS promise to a Promise
func NewRawPromise(promise js.Value) *Promise {
	return &Promise{Value: promise}
}

// NewPromise 创建新的Promise实例
func NewPromise(executor func(resolve func(interface{}), reject func(interface{}))) *Promise {
	promiseCtor := js.Global().Get("Promise")

	jsExecutor := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsNativeResolve := args[0]
		jsNativeReject := args[1]

		jsResolve := func(v interface{}) {
			jsVal := ToJSValue(v)
			jsNativeResolve.Invoke(jsVal) // 调用JS原生resolve
		}

		jsReject := func(v interface{}) {
			jsVal := ToJSValue(v)
			jsNativeReject.Invoke(jsVal) // 调用JS原生reject
		}

		executor(jsResolve, jsReject)
		return nil
	})
	defer jsExecutor.Release() // 释放JS函数资源

	// 创建并返回封装的Promise实例
	return &Promise{Value: promiseCtor.New(jsExecutor)}
}

// Then 封装Promise的then方法
func (p *Promise) Then(onFulfilled func(js.Value)) *Promise {
	jsCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result := js.Undefined()
		if len(args) > 0 {
			result = args[0]
		}
		onFulfilled(result)
		return nil
	})
	// 注意：不能在这里释放 jsCallback，因为它需要在 Promise 异步完成时被调用
	// JS 引擎会在 Promise 链不再需要时自动垃圾回收

	thenPromise := p.Call("then", jsCallback)
	return &Promise{Value: thenPromise}
}

// Catch 封装Promise的catch方法
func (p *Promise) Catch(onRejected func(js.Value)) *Promise {
	jsCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := js.Undefined()
		if len(args) > 0 {
			err = args[0]
		}
		onRejected(err)
		return nil
	})
	// 注意：不能在这里释放 jsCallback，因为它需要在 Promise 异步完成时被调用
	// JS 引擎会在 Promise 链不再需要时自动垃圾回收

	catchPromise := p.Call("catch", jsCallback)
	return &Promise{Value: catchPromise}
}

// Await 模拟JS的await效果，同步等待Promise完成
func (p *Promise) Await() (result js.Value, err js.Value) {
	successChan := make(chan js.Value, 1)
	errorChan := make(chan js.Value, 1)

	p.Then(func(res js.Value) {
		successChan <- res
		close(successChan)
	})

	p.Catch(func(e js.Value) {
		errorChan <- e
		close(errorChan)
	})

	select {
	case result = <-successChan:
		err = js.Undefined()
	case err = <-errorChan:
		result = js.Undefined()
	}

	return
}

// Resolve 静态方法：创建已成功的Promise
func Resolve(value interface{}) *Promise {
	return NewPromise(func(resolve func(interface{}), reject func(interface{})) {
		resolve(value)
	})
}

// Reject 静态方法：创建已失败的Promise
func Reject(reason interface{}) *Promise {
	return NewPromise(func(resolve func(interface{}), reject func(interface{})) {
		reject(reason)
	})
}

// PromiseAll 静态方法：模拟JS Promise.all()（无goroutine，适配tinygo WASM）
func PromiseAll(promises []*Promise) *Promise {
	return NewPromise(func(resolve func(interface{}), reject func(interface{})) {
		// 边界：空切片直接返回空数组
		if len(promises) == 0 {
			resolve([]interface{}{})
			return
		}

		// 初始化变量（所有操作均在主线程回调中执行）
		promiseCount := int32(len(promises))          // 总Promise数
		completedCount := int32(0)                    // 已完成计数（原子操作）
		results := make([]interface{}, len(promises)) // 保持结果顺序
		var mu sync.Mutex                             // 保护rejected标记和results写入
		rejected := false                             // 标记是否已触发失败（快速失败）

		// 遍历所有Promise，监听结果
		for i, p := range promises {
			// 捕获当前循环的索引和Promise（避免闭包共享循环变量）
			index := i
			promise := p

			// 监听Promise成功
			promise.Then(func(res js.Value) {
				mu.Lock()
				defer mu.Unlock()

				// 快速失败：若已触发reject，直接忽略
				if rejected {
					return
				}

				// 存储结果（保持输入顺序）
				results[index] = res

				// 原子增加已完成计数
				completed := atomic.AddInt32(&completedCount, 1)

				// 检查是否所有Promise都完成 → 触发resolve
				if completed == promiseCount {
					resolve(results)
				}
			})

			// 监听Promise失败（快速失败）
			promise.Catch(func(err js.Value) {
				mu.Lock()
				defer mu.Unlock()

				// 只触发一次reject（快速失败）
				if !rejected {
					rejected = true
					reject(err) // 立即返回第一个错误
				}

				// 原子增加已完成计数（不影响逻辑，仅保证计数完整）
				atomic.AddInt32(&completedCount, 1)
			})
		}
	})
}
