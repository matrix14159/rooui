package vdom

import (
	"honnef.co/go/js/dom/v2"
)

type DOMAPI interface {
	CreateElement(name string) dom.HTMLElement
	CreateElementNS(namespace, name string) dom.HTMLElement
	CreateTextNode(text string) *dom.Text
	ParentNode(node dom.Node) dom.Node
	InsertBefore(parentNode, newNode, referenceNode dom.Node)
	NextSibling(node dom.Node) dom.Node
	AppendChild(node dom.Node, child dom.Node)
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

func (p *StandardDomApi) CreateElementNS(namespace, name string) dom.HTMLElement {
	return p.document.CreateElementNS(namespace, name).(dom.HTMLElement)
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

func (p *StandardDomApi) CreateTextNode(text string) *dom.Text {
	return p.document.CreateTextNode(text)
}

func (p *StandardDomApi) AppendChild(node dom.Node, child dom.Node) {
	node.AppendChild(child)
}
