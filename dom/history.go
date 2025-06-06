package dom

import (
	"syscall/js"
)

type History struct {
	js.Value
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/History/length
func (p *History) Length() int {
	return p.Get("length").Int()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/History/back
func (p *History) Back() {
	p.Call("back")
}

// https://developer.mozilla.org/en-US/docs/Web/API/History/forward
func (p *History) Forward() {
	p.Call("forward")
}

// https://developer.mozilla.org/en-US/docs/Web/API/History/go
func (p *History) Go(delta int) {
	p.Call("go", delta)
}

// https://developer.mozilla.org/en-US/docs/Web/API/History_API#The_pushState()_method
func (p *History) PushState(stateObj interface{}, title, url string) {
	p.Call("pushState", stateObj, title, url)
}

// https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState
func (p *History) ReplaceState(stateObj interface{}, title, url string) {
	p.Call("replaceState", stateObj, title, url)
}
