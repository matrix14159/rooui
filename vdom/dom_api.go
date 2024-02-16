package vdom

import (
	"honnef.co/go/js/dom/v2"
)

type DOMAPI interface {
	CreateElement(name string) dom.HTMLElement
	ParentNode(node dom.Node) dom.Node
	InsertBefore(parentNode, newNode, referenceNode dom.Node)
	NextSibling(node dom.Node) dom.Node
}

type StandardDomApi struct {
	window   dom.Window
	document dom.Document
}

func NewStandardDomApi() *StandardDomApi {
	p := new(StandardDomApi)
	p.window = dom.GetWindow()
	p.document = p.window.Document()
	return p
}

func (p *StandardDomApi) CreateElement(name string) dom.HTMLElement {
	return p.document.CreateElement(name).(dom.HTMLElement)
}

func (p *StandardDomApi) ParentNode(node dom.Node) dom.Node {
	return node.ParentNode()
}

func (p *StandardDomApi) InsertBefore(parentNode, newNode, referenceNode dom.Node) {
	parentNode.InsertBefore(newNode, referenceNode)
}

func (p *StandardDomApi) NextSibling(node dom.Node) dom.Node {
	return node.NextSibling()
}
