package vdom

import (
	"github.com/matrix14159/rooui/gs"
)

type VNodeStyle struct {
	Style map[string]string

	Delayed map[string]string

	Remove map[string]string
}

func NewVNodeStyle() *VNodeStyle {
	return &VNodeStyle{
		Style:   make(map[string]string),
		Delayed: make(map[string]string),
		Remove:  make(map[string]string),
	}
}

func raf(f func()) {
	gs.RequestAnimationFrame(f)
}

func NextFrame(f func()) {
	raf(func() {
		raf(f)
	})
}
