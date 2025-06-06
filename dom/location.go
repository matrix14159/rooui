package dom

import (
	"syscall/js"
)

type Location struct {
	js.Value
}

// Properties

func (p *Location) Host() string {
	return p.Get("host").String()
}

func (p *Location) Hostname() string {
	return p.Get("hostname").String()
}

func (p *Location) Href() string {
	return p.Get("href").String()
}

func (p *Location) Origin() string {
	return p.Get("origin").String()
}

func (p *Location) Pathname() string {
	return p.Get("pathname").String()
}

func (p *Location) Port() string {
	return p.Get("port").String()
}

func (p *Location) Protocol() string {
	return p.Get("protocol").String()
}

func (p *Location) Search() string {
	return p.Get("search").String()
}
