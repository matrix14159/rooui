package vdom

type PreHook func()
type InitHook func(vnode *VNode)
type CreateHook func(empty, vnode *VNode)
type InsertHook func(vnode *VNode)
type PrePatchHook func(oldVNode, vnode *VNode)
type UpdateHook func(oldVNode, vnode *VNode)
type PostPatchHook func(oldVNode, vnode *VNode)
type DestroyHook func(vnode *VNode)
type RemoveHook func(vnode *VNode)
type PostHook func()

type Hooks interface {
	Pre()
	Init(vnode *VNode)
	Create(empty, vnode *VNode)
	Insert(vnode *VNode)
	PrePatch(oldVNode, vnode *VNode)
	Update(oldVNode, vnode *VNode)
	PostPatch(oldVNode, vnode *VNode)
	Destroy(vnode *VNode)
	Remove(vnode *VNode)
	Post()
}
