package ui

import (
	"github.com/matrix14159/rooui/core"
)

type Element = core.Element

type Render interface {
	Render() Element
}
