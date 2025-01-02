package ui

func Update(c Render, opts ...UpdateOption) {
	cfg := &UpdateConfig{Mode: M_Self}
	for _, one := range opts {
		one(cfg)
	}

	c.Render()
}

type UpdateConfig struct {
	Mode UpdateMode
}

type UpdateOption func(config *UpdateConfig)

type UpdateMode int

const (
	M_Self UpdateMode = 1
	M_Tree UpdateMode = 2
)

func WithUpdateMode(m UpdateMode) UpdateOption {
	return func(config *UpdateConfig) {
		config.Mode = m
	}
}
