package dom

// This file implements HTMLInputElement interface
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
// https://developer.mozilla.org/en/docs/Web/API/HTMLInputElement

// Properties

func (p *Object) InputValue() string {
	return p.Get("value").String()
}

func (p *Object) SetValue(s string) {
	p.Set("value", s)
}
