package vdom

type Module interface {
	Pre()
	Create(empty, vnode *VNode)
	Update(oldVNode, vnode *VNode)
	Destroy(vnode *VNode)
	Remove(vnode *VNode)
	Post()
}
