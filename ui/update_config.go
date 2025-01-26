package ui

import (
	"github.com/matrix14159/rooui/vdom"
)

type UpdateConfig struct {
	Patcher *vdom.Patcher
}

type UpdateOption func(config *UpdateConfig)

func WithUpdatePatcher(p *vdom.Patcher) UpdateOption {
	return func(config *UpdateConfig) {
		config.Patcher = p
	}
}
