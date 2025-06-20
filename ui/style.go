package ui

import (
	"fmt"
	"strings"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/dom"
	"github.com/vishalkuo/bimap"
)

const c_rooui_style_id = "rooui_style"

type styleMan struct {
	// selector: index
	ruleMap *bimap.BiMap[string, int]
}

var Style = &styleMan{
	ruleMap: bimap.NewBiMap[string, int](),
}

// Add insert style associate with selector
func (p *styleMan) Add(selector string, style ...css.Style) {
	if len(style) == 0 {
		return
	}
	if p.ruleMap.Exists(selector) {
		return
	}

	el := dom.Document.GetElementById(c_rooui_style_id)
	if el == nil {
		el = dom.Document.CreateElement("style")
		el.SetId(c_rooui_style_id)
		el = dom.Document.Head().AppendChild(el)
	}

	b := strings.Builder{}
	for _, one := range style {
		b.WriteString(fmt.Sprintf("%v:%v; ", one.Name, one.Value))
	}
	rule := fmt.Sprintf("%v {%v}", selector, b.String())
	index := el.Sheet().CssRules().Length()
	index = el.Sheet().InsertRule(rule, index) // rule example: "#blanc { color: white; background-color: gray; }"
	p.ruleMap.Insert(selector, index)
}

// Remove delete style associate with selector
func (p *styleMan) Remove(selector string) {
	index, found := p.ruleMap.Get(selector)
	if !found {
		return
	}

	el := dom.Document.GetElementById(c_rooui_style_id)
	if el == nil {
		return
	}

	el.Sheet().DeleteRule(index)
}

func AddLinkCSS(url string) {

}
