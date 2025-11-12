package ui

import (
	"fmt"
)

// ToPx convert value v to px
// example: ToPx(3) -> "3px"
func ToPx(v int) string {
	return fmt.Sprintf("%dpx", v)
}
