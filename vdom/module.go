package vdom

type Module interface {
	Pre()
	Create(empty, vnode *VNode)
	Update(oldVNode, vnode *VNode)
	Destroy(vnode *VNode)
	Remove(vnode *VNode, removeCallback func())
	Post()
}

type HookGroup struct {
	Pres     []PreHook
	Creates  []CreateHook
	Updates  []UpdateHook
	Destroys []DestroyHook
	Removes  []RemoveHook
	Posts    []PostHook
}
