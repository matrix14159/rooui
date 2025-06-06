package dom

// This file implements useful functions not in DOM API

func (p *Object) RemoveAllChildNodes() {
	for p.HasChildNodes() {
		p.RemoveChild(p.LastChild())
	}
}

// element.AppendBefore(newElement) - insert newElement before element.
// Some people call this as *insert before*, but this calling is confusing
// because the meaning of insertBefore() in DOM API is different.
// http://stackoverflow.com/a/32135318
func (p *Object) AppendBefore(n *Object) {
	p.ParentNode().InsertBefore(n, p)
}

// element.AppendAfter(newElement) - insert newElement after element.
// Some people call this as *insert after*.
// http://stackoverflow.com/a/32135318
func (p *Object) AppendAfter(n *Object) {
	p.ParentNode().InsertBefore(n, p.NextSibling())
}

// Check if the element is currently focused
func (p *Object) IsFocused() bool {
	return p.IsEqualNode(Document.ActiveElement())
}
