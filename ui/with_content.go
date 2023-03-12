package ui

type WithContent struct {
	Content []UI
}

func (p *WithContent) AppendContent(elems ...UI) {
	p.Content = append(p.Content, elems...)
}
