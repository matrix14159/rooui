package dom

// This file implements Element interface
// https://developer.mozilla.org/en-US/docs/Web/API/Element

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (p *Object) ClassList() *DOMTokenList {
	return &DOMTokenList{p.Get("classList")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (p *Object) InnerHTML() string {
	return p.Get("innerHTML").String()
}
func (p *Object) SetInnerHTML(html string) {
	p.Set("innerHTML", html)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (p *Object) OuterHTML() string {
	return p.Get("outerHTML").String()
}
func (p *Object) SetOuterHTML(html string) {
	p.Set("outerHTML", html)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (p *Object) TagName() string {
	return p.Get("tagName").String()
}

// Methods

// Call HasAttribute to check if the attribute exists or not before using this
// method (GetAttribute).
// todo: Return (string, bool) to indicate the existence of the attribute?
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (p *Object) GetAttribute(name string) string {
	return p.Call("getAttribute", name).String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute
func (p *Object) HasAttribute(name string) bool {
	return p.Call("hasAttribute", name).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttribute
func (p *Object) SetAttribute(name string, value any) {
	p.Call("setAttribute", name, value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttributeNS
func (p *Object) SetAttributeNS(namespace, name string, value any) {
	p.Call("setAttributeNS", namespace, name, value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/removeAttribute
func (p *Object) RemoveAttribute(name string) {
	p.Call("removeAttribute", name)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (p *Object) GetBoundingClientRect() *DOMRect {
	return &DOMRect{p.Call("getBoundingClientRect")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (p *Object) GetElementsByTagName(tagName string) []*Object {
	nodeList := p.Call("getElementsByTagName", tagName)
	length := nodeList.Get("length").Int()
	var nodes []*Object
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Object{nodeList.Call("item", i)})
	}
	return nodes
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (p *Object) QuerySelector(selectors string) *Object {
	return &Object{p.Call("querySelector", selectors)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (p *Object) QuerySelectorAll(selectors string) []*Object {
	nodeList := p.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()
	var nodes []*Object
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Object{nodeList.Call("item", i)})
	}
	return nodes
}
