package dom

// This file implements XMLHttpRequest (XHR)
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest

import (
	"syscall/js"
)

type XMLHttpRequest struct {
	js.Value
}

func NewXMLHttpRequest() *XMLHttpRequest {
	return &XMLHttpRequest{Window.Get("XMLHttpRequest").New()}
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseText
func (p *XMLHttpRequest) ResponseText() string {
	return p.Get("responseText").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseURL
func (p *XMLHttpRequest) ResponseURL() string {
	return p.Get("responseURL").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseXML
func (p *XMLHttpRequest) ResponseXML() *Object {
	return &Object{p.Get("responseXML")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/statusText
func (p *XMLHttpRequest) StatusText() string {
	return p.Get("statusText").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials
func (p *XMLHttpRequest) WithCredentials() bool {
	return p.Get("withCredentials").Bool()
}

func (p *XMLHttpRequest) SetWithCredentials(c bool) {
	p.Set("withCredentials", c)
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/abort
func (p *XMLHttpRequest) Abort() {
	p.Call("abort")
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/open
//
// Optional parameters 'user' and 'password' are not supported.
func (p *XMLHttpRequest) Open(method, url string, args ...interface{}) {
	if len(args) == 1 {
		// args[0] should be bool async
		p.Call("open", method, url, args[0])
	} else {
		p.Call("open", method, url)
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/overrideMimeType
func (p *XMLHttpRequest) OverrideMimeType(mimeType string) {
	p.Call("overrideMimeType", mimeType)
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/send
func (p *XMLHttpRequest) Send(args ...interface{}) {
	if len(args) == 1 {
		p.Call("send", args[0])
	} else {
		p.Call("send")
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/setRequestHeader
func (p *XMLHttpRequest) SetRequestHeader(header, value string) {
	p.Call("setRequestHeader", header, value)
}
