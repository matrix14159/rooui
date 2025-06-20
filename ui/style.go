package ui

import (
	"fmt"
	"strings"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/dom"
)

func AddStyle(selector string, style ...css.Style) {
	if len(style) == 0 {
		return
	}

	el := dom.Document.GetElementById("rooui_style")
	if el == nil {
		el = dom.Document.CreateElement("style")
		el.SetId("rooui_style")
		el = dom.Document.Head().AppendChild(el)
	}

	b := strings.Builder{}
	for _, one := range style {
		b.WriteString(fmt.Sprintf("%v:%v; ", one.Name, one.Value))
	}
	rule := fmt.Sprintf("%v {%v}", selector, b.String())
	el.Sheet().InsertRule(rule, 0) // rule example: "#blanc { color: white; background-color: gray; }"
}

func RemoveStyle(selector string) {

}

func AddLinkCSS(url string) {

}
