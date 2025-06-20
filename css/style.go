package css

import (
	"fmt"
)

type Style struct {
	Name  string
	Value string
}

func (p Style) ToCSS() string {
	return fmt.Sprintf("%v: %v;", p.Name, p.Value)
}
