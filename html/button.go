package html

type ButtonElement struct {
	text string
}

func Button() *ButtonElement {
	return new(ButtonElement)
}

func (p *ButtonElement) Tag() string {
	return "button"
}

func (p *ButtonElement) Text(text string) *ButtonElement {
	p.text = text
	return p
}
