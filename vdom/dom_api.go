package vdom

import (
	"github.com/matrix14159/rooui/dom"
)

type DOMAPI interface {
	CreateElement(name string) *dom.Object
	CreateElementNS(namespace, name string) *dom.Object
	CreateTextNode(text string) *dom.Object
	ParentNode(node *dom.Object) *dom.Object
	InsertBefore(parentNode, newNode, referenceNode *dom.Object)
	NextSibling(node *dom.Object) *dom.Object
	AppendChild(node *dom.Object, child *dom.Object)
	RemoveChild(node *dom.Object, child *dom.Object)
	SetTextContent(node *dom.Object, text string)
}

type StandardDomApi struct {
}

func NewStandardDomApi() *StandardDomApi {
	p := new(StandardDomApi)
	return p
}

func (p *StandardDomApi) CreateElement(name string) *dom.Object {
	return dom.Document.CreateElement(name)
}

func (p *StandardDomApi) CreateElementNS(namespace, name string) *dom.Object {
	return dom.Document.CreateElementNS(namespace, name)
}

func (p *StandardDomApi) CreateTextNode(text string) *dom.Object {
	return dom.Document.CreateTextNode(text)
}

func (p *StandardDomApi) ParentNode(node *dom.Object) *dom.Object {
	return node.ParentNode()
}

func (p *StandardDomApi) InsertBefore(parentNode, newNode, referenceNode *dom.Object) {
	parentNode.InsertBefore(newNode, referenceNode)
}

func (p *StandardDomApi) NextSibling(node *dom.Object) *dom.Object {
	return node.NextSibling()
}

func (p *StandardDomApi) AppendChild(node *dom.Object, child *dom.Object) {
	node.AppendChild(child)
}

func (p *StandardDomApi) RemoveChild(node *dom.Object, child *dom.Object) {
	node.RemoveChild(child)
}

func (p *StandardDomApi) SetTextContent(node *dom.Object, text string) {
	node.SetTextContent(text)
}
