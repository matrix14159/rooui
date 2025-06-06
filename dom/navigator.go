package dom

import (
	"syscall/js"
)

type Navigator struct {
	js.Value
}

// Properties

func (p *Navigator) Language() string {
	return p.Get("language").String()
}

func (p *Navigator) Languages() string {
	return p.Get("languages").String()
}
