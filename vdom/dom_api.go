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
	RemoveChild(node dom.Node, child dom.Node)
	SetTextContent(node dom.Node, text string)
}

type StandardDomApi struct {
	Window   dom.Window
	Document dom.Document
}

func NewStandardDomApi() *StandardDomApi {
	p := new(StandardDomApi)
	p.Window = dom.GetWindow()
	p.Document = p.Window.Document()
	return p
}

func (p *StandardDomApi) CreateElement(name string) dom.HTMLElement {
	return p.Document.CreateElement(name).(dom.HTMLElement)
}

func (p *StandardDomApi) CreateElementNS(namespace, name string) dom.HTMLElement {
	return p.Document.CreateElementNS(namespace, name).(dom.HTMLElement)
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
	return p.Document.CreateTextNode(text)
}

func (p *StandardDomApi) AppendChild(node dom.Node, child dom.Node) {
	node.AppendChild(child)
}

func (p *StandardDomApi) RemoveChild(node dom.Node, child dom.Node) {
	node.RemoveChild(child)
}

func (p *StandardDomApi) SetTextContent(node dom.Node, text string) {
	node.SetTextContent(text)
}
