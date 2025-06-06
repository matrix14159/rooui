package dom

// This file implements Node interface
// https://developer.mozilla.org/en-US/docs/Web/API/Node

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (p *Object) ChildNodes() []*Object {
	nodeList := p.Get("childNodes")
	length := nodeList.Get("length").Int()
	var nodes []*Object
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Object{nodeList.Call("item", i)})
	}
	return nodes
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (p *Object) FirstChild() *Object {
	return &Object{p.Get("firstChild")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/lastChild
func (p *Object) LastChild() *Object {
	return &Object{p.Get("lastChild")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
func (p *Object) NextSibling() *Object {
	return &Object{p.Get("nextSibling")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (p *Object) NodeType() int {
	return p.Get("nodeType").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (p *Object) NodeValue() string {
	return p.Get("nodeValue").String()
}
func (p *Object) SetNodeValue(s string) {
	p.Set("nodeValue", s)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (p *Object) ParentNode() *Object {
	return &Object{p.Get("parentNode")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
// Returns the textual content of an element and all its descendants.
func (p *Object) TextContent() string {
	return p.Get("textContent").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
// Sets the textual content of an element and all its descendants.
func (p *Object) SetTextContent(s string) {
	p.Set("textContent", s)
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (p *Object) AppendChild(c *Object) *Object {
	if c == nil {
		return nil
	}
	return &Object{p.Call("appendChild", c.Value)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/contains
func (p *Object) Contains(n *Object) bool {
	return p.Call("contains", n).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (p *Object) HasChildNodes() bool {
	return p.Call("hasChildNodes").Bool()
}

// Inserts the first Node given in a parameter immediately before the second,
// child of this element, Node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (p *Object) InsertBefore(newNode, referenceNode *Object) *Object {
	if newNode == nil {
		return nil
	}
	var before interface{}
	if referenceNode != nil {
		before = referenceNode.Value
	}
	return &Object{p.Call("insertBefore", newNode.Value, before)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/isEqualNode
func (p *Object) IsEqualNode(n *Object) bool {
	return p.Call("isEqualNode", n).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (p *Object) IsSameNode(n *Object) bool {
	return p.Call("isSameNode", n).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/lookupPrefix
func (p *Object) LookupPrefix() string {
	return p.Call("lookupPrefix").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/normalize
func (p *Object) Normalize() {
	p.Call("normalize")
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (p *Object) RemoveChild(c *Object) *Object {
	if c == nil {
		return nil
	}
	return &Object{p.Call("removeChild", c.Value)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (p *Object) ReplaceChild(newChild, oldChild *Object) *Object {
	if oldChild == nil {
		return nil
	}
	var newCh interface{}
	if newChild != nil {
		newCh = newChild.Value
	}
	return &Object{p.Call("replaceChild", newCh, oldChild.Value)}
}
