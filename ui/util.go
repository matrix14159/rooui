package ui

import (
	"fmt"
)

// ToPx convert value v to px
// example: ToPx(3) -> "3px"
func ToPx[T int | float64](v T) string {
	return fmt.Sprintf("%vpx", v)
}
