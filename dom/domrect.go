package dom

// This file implements DOMRect interface
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRect

import (
	"fmt"
	"syscall/js"
)

type DOMRect struct {
	js.Value
}

func (p *DOMRect) X() float64 {
	return p.Get("x").Float()
}

func (p *DOMRect) Y() float64 {
	return p.Get("y").Float()
}

func (p *DOMRect) Width() float64 {
	return p.Get("width").Float()
}

func (p *DOMRect) Height() float64 {
	return p.Get("height").Float()
}

func (p *DOMRect) Top() float64 {
	return p.Get("top").Float()
}

func (p *DOMRect) Right() float64 {
	return p.Get("right").Float()
}

func (p *DOMRect) Bottom() float64 {
	return p.Get("bottom").Float()
}

func (p *DOMRect) Left() float64 {
	return p.Get("left").Float()
}

func (p *DOMRect) String() string {
	return fmt.Sprintf("{X:%v, Y:%v, W:%v, H:%v, Top:%v, Right:%v, Bottom:%v, Left:%v}",
		p.X(), p.Y(), p.Width(), p.Height(), p.Top(), p.Right(), p.Bottom(), p.Left())
}
