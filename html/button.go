package html

import (
	"fmt"
)

type ButtonElement struct {
	text string
}

func Button() *ButtonElement {
	return new(ButtonElement)
}

func (p *ButtonElement) Tag() string {
	return "button"
}

func (p *ButtonElement) GetText() string {
	return p.text
}

func (p *ButtonElement) Text(text any) *ButtonElement {
	p.text = fmt.Sprintf("%v", text)
	return p
}
