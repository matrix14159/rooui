# RooUI 使用教程

RooUI 是一个用于 Go WebAssembly 的声明式 UI 框架，采用虚拟 DOM (Virtual DOM) 技术来实现高效的 UI 更新。

## 目录

1. [快速开始](#快速开始)
2. [核心概念](#核心概念)
3. [组件开发](#组件开发)
4. [HTML 元素](#html-元素)
5. [样式管理](#样式管理)
6. [事件处理](#事件处理)
7. [状态更新](#状态更新)
8. [高级特性](#高级特性)
9. [完整示例](#完整示例)

## 快速开始

### 1. 创建组件

创建一个简单的计数器组件：

```go
package main

import (
	"github.com/matrix14159/rooui/ui"
)

type Counter struct {
	ui.Component
	num int
}

func NewCounter() *Counter {
	return &Counter{num: 0}
}

func (p *Counter) Render() ui.Element {
	b := ui.Button().Text(p.num).OnClick(p.increase)
	return ui.Render(p, b)
}

func (p *Counter) increase(event ui.Event, options ...any) {
	p.num++
	ui.Update(p)
}
```

### 2. 启动应用

在 `main` 函数中启动应用：

```go
func main() {
	ui.Run(NewCounter())
}
```

### 3. 编译为 WASM

```bash
GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o=main.wasm
```

### 4. HTML 页面

创建一个 HTML 页面来加载 WASM：

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>RooUI App</title>
  </head>
  <body>
    <div id="app"></div>
    <script src="wasm_exec.js"></script>
    <script>
      const go = new Go();
      WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
        .then((result) => {
          go.run(result.instance);
          MountTo("app");
        });
    </script>
  </body>
</html>
```

## 核心概念

### Component 接口

所有组件必须实现 `ui.Component` 接口：

```go
type Comp interface {
	Render() Element          // 返回组件的 UI 元素
	OnUpdated()              // 组件更新后的回调
	GetId() string           // 获取组件 ID
	getElement() Element     // 内部方法
	updateElement(el Element) // 内部方法
	getVNode() *vdom.VNode   // 内部方法
	updateVNode(vnode *vdom.VNode) // 内部方法
	findVNodeChild(child *vdom.VNode) int // 内部方法
	replaceVNodeChild(idx int, oldChild, newChild *vdom.VNode) // 内部方法
}
```

通常，你只需要嵌入 `ui.Component` 并实现 `Render()` 方法：

```go
type MyComponent struct {
	ui.Component
	// 你的状态字段
}

func (p *MyComponent) Render() ui.Element {
	// 返回 UI 元素
	return ui.Render(p, ui.Div().Text("Hello"))
}
```

### Render 方法

`Render()` 方法必须返回一个 `ui.Element`，通常使用 `ui.Render()` 函数来包装：

```go
func (p *MyComponent) Render() ui.Element {
	el := ui.Div().Text("Hello World")
	return ui.Render(p, el)
}
```

`ui.Render()` 函数将组件与元素关联，并设置父子关系。

## 组件开发

### 基本组件结构

```go
type MyComponent struct {
	ui.Component

	// 组件状态
	title string
	count int
}

func NewMyComponent() *MyComponent {
	return &MyComponent{
		title: "My Component",
		count: 0,
	}
}

func (p *MyComponent) Render() ui.Element {
	return ui.Render(p, ui.Div().Text(p.title))
}
```

### 组件生命周期

#### OnUpdated 回调

组件更新完成后会调用 `OnUpdated()` 方法：

```go
func (p *MyComponent) OnUpdated() {
	log.Println("Component updated")
}
```

### 嵌套组件

组件可以包含子组件：

```go
type List struct {
	ui.Component
	items []*ListItem
}

func (p *List) Render() ui.Element {
	d := ui.Div()
	for _, item := range p.items {
		d.Body(item.Render())
	}
	return ui.Render(p, d)
}
```

## HTML 元素

RooUI 提供了所有标准 HTML 元素的创建函数，都在 `ui` 包中：

### 常用元素

```go
// 容器元素
ui.Div()
ui.Span()
ui.Section()
ui.Article()
ui.Header()
ui.Footer()
ui.Main()

// 文本元素
ui.P()
ui.H1(), ui.H2(), ui.H3(), ui.H4(), ui.H5(), ui.H6()
ui.Span()
ui.Strong()
ui.Em()

// 表单元素
ui.Button()
ui.Input()
ui.Textarea()
ui.Select()
ui.Label()
ui.Form()

// 媒体元素
ui.Img()
ui.Video()
ui.Audio()

// 其他
ui.A()  // 链接
ui.Ul(), ui.Ol(), ui.Li()  // 列表
ui.Table(), ui.Tr(), ui.Td()  // 表格
ui.Dialog()  // 对话框
```

### 元素方法链

所有元素都支持方法链式调用：

```go
ui.Div().
	Id("my-div").
	Class("container", nil).
	Style(css.Color("red")).
	Text("Hello").
	Body(
		ui.P().Text("Paragraph 1"),
		ui.P().Text("Paragraph 2"),
	)
```

### 设置属性

```go
// ID
el.Id("my-id")

// 文本内容
el.Text("Hello World")
el.Text(123)  // 自动转换为字符串

// 子元素
el.Body(child1, child2, child3)

// HTML 属性
el.Attribute("data-id", "123")
el.Attribute("role", "button")

// DOM 属性
el.Prop("disabled", true)
el.Prop("value", "input value")

// CSS 类名
el.Classes("class1", "class2", "class3")

// 内联样式
el.Style(
	css.Color("red"),
	css.BackgroundColor("blue"),
	css.Padding("10px"),
)
```

## 样式管理

### CSS 样式对象

使用 `css` 包创建样式：

```go
import "github.com/matrix14159/rooui/css"

// 创建样式
styles := []css.Style{
	css.Color("blue"),
	css.BackgroundColor("gray"),
	css.Padding("10px"),
	css.Margin("5px"),
	css.Width("100%"),
	css.Height("50px"),
	css.Display("flex"),
	css.FlexDirection("column"),
	css.JustifyContent("space-between"),
	css.AlignItems("center"),
}
```

### 内联样式

直接在元素上设置样式：

```go
ui.Div().Style(
	css.Color("red"),
	css.BackgroundColor("white"),
	css.Padding("20px"),
)
```

### CSS 类样式

为元素添加 CSS 类并定义样式：

```go
// 定义类样式
ui.Style.Add(".blue",
	css.Color("blue"),
	css.BackgroundColor("gray"),
)

// 使用类
ui.Div().Classes("blue")
```

或者使用 `Class()` 方法直接定义：

```go
ui.Div().Class("my-class",
	css.Color("blue"),
	css.BackgroundColor("gray"),
)
```

当元素设置了id时，class-name可以忽略，例如：

```go
ui.Div().Id("my-div").Class("",
	css.Color("blue"),
	css.BackgroundColor("gray"),
)
```

### 全局样式管理

使用 `ui.Style` 管理全局样式：

```go
// 添加样式规则
ui.Style.Add(".container",
	css.Display("flex"),
	css.FlexDirection("column"),
)

// 移除样式规则
ui.Style.Remove(".container")

// 链接外部 CSS
ui.Style.LinkCSS("https://example.com/style.css", "")

// 定义动画关键帧
ui.Style.SetKeyframes("fadeIn", []ui.Keyframe{
	{Selector: "0%", Styles: []css.Style{css.Opacity("0")}},
	{Selector: "100%", Styles: []css.Style{css.Opacity("1")}},
})
```

## 事件处理

### 事件处理器签名

所有事件处理器的签名都是：

```go
func(event ui.Event, options ...any)
```

### 常用事件

```go
// 点击事件
el.OnClick(func(event ui.Event, options ...any) {
	// 处理点击
})

// 鼠标事件
el.OnMouseDown(...)
el.OnMouseUp(...)
el.OnMouseMove(...)
el.OnMouseEnter(...)
el.OnMouseLeave(...)

// 键盘事件
el.OnKeyDown(...)
el.OnKeyUp(...)
el.OnKeyPress(...)

// 输入事件
el.OnInput(...)
el.OnFocus(...)
el.OnBlur(...)

// 其他事件
el.OnScroll(...)
el.OnTransitionEnd(...)
```

### 事件对象

`ui.Event` 提供了丰富的 API：

```go
func (p *MyComponent) handleClick(event ui.Event, options ...any) {
	// 获取事件目标元素
	target := event.Target()

	// 获取坐标
	x := event.X()
	y := event.Y()

	// 阻止事件传播
	event.StopPropagation()

	// 阻止默认行为
	event.PreventDefault()

	// 获取元素位置信息
	rect := target.GetBoundingClientRect()
	left := rect.Left()
	top := rect.Top()
	width := rect.Width()
	height := rect.Height()
}
```

### 自定义事件

使用 `AddEventListener` 添加自定义事件：

```go
el.AddEventListener("custom-event", func(event ui.Event, options ...any) {
	// 处理自定义事件
})
```

## 状态更新

### 更新组件

当组件状态改变时，调用 `ui.Update()` 来触发重新渲染：

```go
func (p *Counter) increase(event ui.Event, options ...any) {
	p.num++
	ui.Update(p)  // 触发重新渲染
}
```

### 更新流程

1. 调用 `ui.Update(component)`
2. 框架调用组件的 `Render()` 方法
3. 生成新的虚拟 DOM
4. 与旧的虚拟 DOM 进行 diff
5. 更新实际的 DOM
6. 调用 `OnUpdated()` 回调

### 异步更新

可以在 goroutine 中更新组件：

```go
func (p *List) loadData() {
	go func() {
		for i := 0; i < 10; i++ {
			p.items = append(p.items, newItem(i))
			time.Sleep(100 * time.Millisecond)
			ui.Update(p)  // 在 goroutine 中安全更新
		}
	}()
}
```

## 高级特性

### 唯一 ID 生成

使用 `Uid()` 方法生成唯一 ID：

```go
func (p *MyComponent) Render() ui.Element {
	id := p.Uid("button")  // 生成类似 "button_123" 的唯一 ID
	return ui.Render(p, ui.Button().Id(id))
}
```

### Dialog 对话框

创建和使用对话框：

```go
type MyComponent struct {
	ui.Component
	dialog *html.DialogElement
}

func NewMyComponent() *MyComponent {
	p := &MyComponent{}
	p.dialog = ui.Dialog()
	p.dialog.Id("my-dialog")
	p.dialog.CloseByMask()  // 点击遮罩关闭
	p.dialog.Body(
		ui.Div().Text("Dialog Content"),
		ui.Button().Text("Close").OnClick(p.closeDialog),
	)
	return p
}

func (p *MyComponent) Render() ui.Element {
	return ui.Render(p, ui.Div().Body(
		ui.Button().Text("Open Dialog").OnClick(p.openDialog),
		p.dialog,
	))
}

func (p *MyComponent) openDialog(event ui.Event, options ...any) {
	event.StopPropagation()
	rect := event.Target().GetBoundingClientRect()
	x, y := ui.ToPx(rect.Left()), ui.ToPx(rect.Bottom()+10)
	p.dialog.ShowModal(x, y)
}

func (p *MyComponent) closeDialog(event ui.Event, options ...any) {
	event.StopPropagation()
	p.dialog.Close()
}
```

### 动态样式更新

在渲染后动态更新样式：

```go
func (p *MyComponent) updateStyle() {
	el := ui.Div().Id("dynamic-element")
	// ... 渲染元素

	// 渲染后更新样式
	el.SetStyleProperty("color", "red")
	el.SetStyleProperty("background-color", "blue")
}
```

注意：元素必须有 ID 才能使用 `SetStyleProperty`。

### 工具函数

```go
// 转换为 px 单位
px := ui.ToPx(100)  // "100px"
px := ui.ToPx(50.5) // "50.5px"
```

## 完整示例

### 示例 1: 计数器

```go
package main

import (
	"github.com/matrix14159/rooui/ui"
)

type Counter struct {
	ui.Component
	num int
}

func NewCounter() *Counter {
	return &Counter{num: 0}
}

func (p *Counter) Render() ui.Element {
	b := ui.Button().Text(p.num).OnClick(p.increase)
	return ui.Render(p, b)
}

func (p *Counter) increase(event ui.Event, options ...any) {
	p.num++
	ui.Update(p)
}

func main() {
	ui.Run(NewCounter())
}
```

### 示例 2: 列表组件

```go
package main

import (
	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/ui"
)

type List struct {
	ui.Component
	items []int
}

func NewList() *List {
	p := &List{}
	for i := 0; i < 10; i++ {
		p.items = append(p.items, i)
	}
	return p
}

func (p *List) Render() ui.Element {
	d := ui.Div().Class("",
		css.Display("flex"),
		css.FlexDirection("column"),
		css.Gap("10px"),
	)

	d.Body(
		ui.Button().Text("Increase All").OnClick(p.increase),
		ui.Button().Text("Decrease All").OnClick(p.decrease),
	)

	for _, item := range p.items {
		d.Body(ui.Button().Text(item))
	}

	return ui.Render(p, d)
}

func (p *List) increase(event ui.Event, options ...any) {
	for i := range p.items {
		p.items[i]++
	}
	ui.Update(p)
}

func (p *List) decrease(event ui.Event, options ...any) {
	for i := range p.items {
		p.items[i]--
	}
	ui.Update(p)
}

func main() {
	ui.Run(NewList())
}
```

### 示例 3: 嵌套组件

```go
package main

import (
	"github.com/matrix14159/rooui/ui"
)

// 列表项组件
type ListItem struct {
	ui.Component
	Num int
}

func (p *ListItem) Render() ui.Element {
	d := ui.Div().Body(
		ui.Div().Text(p.Num),
		ui.Button().Text("Increase").OnClick(p.increase),
	)
	return ui.Render(p, d)
}

func (p *ListItem) increase(event ui.Event, options ...any) {
	p.Num++
	ui.Update(p)
}

// 列表组件
type List struct {
	ui.Component
	items []*ListItem
}

func NewList() *List {
	p := &List{}
	for i := 0; i < 3; i++ {
		p.items = append(p.items, &ListItem{Num: i})
	}
	return p
}

func (p *List) Render() ui.Element {
	d := ui.Div().Body(
		ui.Button().Text("Increase All").OnClick(p.increaseAll),
		ui.Button().Text("Decrease All").OnClick(p.decreaseAll),
	)

	for _, item := range p.items {
		d.Body(item.Render())
	}

	return ui.Render(p, d)
}

func (p *List) increaseAll(event ui.Event, options ...any) {
	for _, item := range p.items {
		item.Num++
	}
	ui.Update(p)
}

func (p *List) decreaseAll(event ui.Event, options ...any) {
	for _, item := range p.items {
		item.Num--
	}
	ui.Update(p)
}

func main() {
	ui.Run(NewList())
}
```

### 示例 4: 带样式的复杂组件

```go
package main

import (
	"moss/pop"

	"github.com/matrix14159/rooui/css"
	"github.com/matrix14159/rooui/html"
	"github.com/matrix14159/rooui/ui"
)

type Page struct {
	ui.Component
	dialog *html.DialogElement
}

func NewPage() *Page {
	p := &Page{}
	p.dialog = ui.Dialog()
	p.dialog.Id("my-dialog")
	p.dialog.Class("", pop.Style()...)
	p.dialog.CloseByMask()
	p.dialog.Body(
		ui.Div().Class("my-class",
			css.Padding("1rem"),
		).Body(
			ui.P().Text("Dialog Content"),
			ui.Button().Text("Close").OnClick(p.closeDialog),
		),
	)
	return p
}

func (p *Page) Render() ui.Element {
	d := ui.Div().Id("page").Class("",
		css.Display("flex"),
		css.FlexDirection("column"),
		css.Height("100vh"),
	)

	d.Body(
		ui.Div().Id("header").Class("",
			css.Height("50px"),
			css.BorderBottom("1px solid #d0d7de"),
			css.Padding("0 30px"),
			css.BackgroundColor("#f6f7fa"),
			css.Display("flex"),
			css.AlignItems("center"),
		).Body(
			ui.Button().Text("Open Dialog").OnClick(p.openDialog),
		),
		ui.Div().Class("class-2",
			css.Flex("1"),
			css.Padding("20px"),
		).Body(
			ui.P().Text("Page Content"),
		),
		p.dialog,
	)

	return ui.Render(p, d)
}

func (p *Page) openDialog(event ui.Event, options ...any) {
	event.StopPropagation()
	rect := event.Target().GetBoundingClientRect()
	x, y := ui.ToPx(rect.Left()), ui.ToPx(rect.Bottom()+10)
	p.dialog.ShowModal(x, y)
}

func (p *Page) closeDialog(event ui.Event, options ...any) {
	event.StopPropagation()
	p.dialog.Close()
}

func main() {
	ui.Run(NewPage())
}
```

## 最佳实践

1. **组件设计**
   - 保持组件小而专注
   - 使用组合而非继承
   - 将状态提升到合适的层级

2. **性能优化**
   - 只在必要时调用 `ui.Update()`
   - 避免在 `Render()` 中执行耗时操作
   - 使用 `Uid()` 生成稳定的 ID

3. **代码组织**
   - 将相关组件放在同一包中
   - 使用有意义的组件和变量名
   - 保持 `Render()` 方法简洁

4. **样式管理**
   - 优先使用 CSS 类而非内联样式
   - 使用 `ui.Style.Add()` 定义可复用的样式
   - 保持样式的一致性

5. **事件处理**
   - 使用 `event.StopPropagation()` 防止事件冒泡
   - 在事件处理器中更新状态后调用 `ui.Update()`
   - 避免在事件处理器中执行耗时操作

## 总结

RooUI 提供了一个简洁而强大的方式来构建 Go WebAssembly 应用的用户界面。通过组件化、虚拟 DOM 和声明式 API，你可以高效地构建复杂的交互式应用。

主要特点：
- ✅ 声明式 UI 编程
- ✅ 虚拟 DOM 高效更新
- ✅ 完整的 HTML 元素支持
- ✅ 灵活的样式管理
- ✅ 丰富的事件处理
- ✅ 组件化架构
- ✅ Type-safe API

开始使用 RooUI，享受 Go + WebAssembly 的强大组合！

