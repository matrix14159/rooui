package vdom

type PreHook func()
type InitHook func(vnode *VNode)
type CreateHook func(empty, vnode *VNode)
type InsertHook func(vnode *VNode)
type PrePatchHook func(oldVNode, vnode *VNode)
type UpdateHook func(oldVNode, vnode *VNode)
type PostPatchHook func(oldVNode, vnode *VNode)
type DestroyHook func(vnode *VNode)
type RemoveHook func(vnode *VNode, removeCallback func())
type PostHook func()

type Hooks struct {
	Pre       PreHook
	Init      InitHook
	Create    CreateHook
	Insert    InsertHook
	PrePatch  PrePatchHook
	Update    UpdateHook
	PostPatch PostPatchHook
	Destroy   DestroyHook
	Remove    RemoveHook
	Post      PostHook
}
