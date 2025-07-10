package ui

import (
	"fmt"
	"strings"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/dom"
)

const c_rooui_style_id = "rooui_style"

type styleMan struct {
	// selector: index
	ruleMap map[string]int

	// url: struct{}
	linkMap map[string]any
}

var Style = &styleMan{
	ruleMap: make(map[string]int),
	linkMap: make(map[string]any),
}

// Add insert style associate with selector
// inspect by browser console: document.getElementById("rooui_style").sheet
func (p *styleMan) Add(selector string, style ...css.Style) {
	if len(style) == 0 {
		return
	}
	if _, found := p.ruleMap[selector]; found {
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
	p.ruleMap[selector] = index
}

// Remove delete style associate with selector
func (p *styleMan) Remove(selector string) {
	index, found := p.ruleMap[selector]
	if !found {
		return
	}

	el := dom.Document.GetElementById(c_rooui_style_id)
	if el == nil {
		return
	}

	el.Sheet().DeleteRule(index)
}

func (p *styleMan) LinkCSS(url string, media string) {
	if len(url) == 0 {
		return
	}
	if _, found := p.linkMap[url]; found {
		return
	}

	link := dom.Document.CreateElement("link")
	link.SetAttribute("type", "text/css")
	link.SetAttribute("rel", "stylesheet")
	link.SetAttribute("href", url)
	if media != "" {
		link.SetAttribute("media", media)
	}
	link = dom.Document.Head().AppendChild(link)

	p.linkMap[url] = struct{}{}
}
