package css

/// 盒模型与尺寸

func Width(v string) Style {
	return Style{Name: "width", Value: v}
}

func Height(v string) Style {
	return Style{Name: "height", Value: v}
}

func MinWidth(v string) Style {
	return Style{Name: "min-width", Value: v}
}

func MaxWidth(v string) Style {
	return Style{Name: "max-width", Value: v}
}

func MinHeight(v string) Style {
	return Style{Name: "min-height", Value: v}
}

func MaxHeight(v string) Style {
	return Style{Name: "max-height", Value: v}
}

func AspectRatio(v string) Style {
	return Style{Name: "aspect-ratio", Value: v}
}

func Margin(v string) Style {
	return Style{Name: "margin", Value: v}
}

func MarginTop(v string) Style {
	return Style{Name: "margin-top", Value: v}
}

func MarginRight(v string) Style {
	return Style{Name: "margin-right", Value: v}
}

func MarginBottom(v string) Style {
	return Style{Name: "margin-bottom", Value: v}
}

func MarginLeft(v string) Style {
	return Style{Name: "margin-left", Value: v}
}

func Padding(v string) Style {
	return Style{Name: "padding", Value: v}
}

func PaddingTop(v string) Style {
	return Style{Name: "padding-top", Value: v}
}

func PaddingRight(v string) Style {
	return Style{Name: "padding-right", Value: v}
}

func PaddingBottom(v string) Style {
	return Style{Name: "padding-bottom", Value: v}
}

func PaddingLeft(v string) Style {
	return Style{Name: "padding-left", Value: v}
}

// 可选值: content-box | border-box | inherit | initial | unset
func BoxSizing(v string) Style {
	return Style{Name: "box-sizing", Value: v}
}

func BoxShadow(v string) Style {
	return Style{Name: "box-shadow", Value: v}
}

// 可选值: slice | clone | inherit | initial | unset
func BoxDecorationBreak(v string) Style {
	return Style{Name: "box-decoration-break", Value: v}
}

/// 背景与边框

func BackgroundColor(v string) Style {
	return Style{Name: "background-color", Value: v}
}

func BackgroundImage(v string) Style {
	return Style{Name: "background-image", Value: v}
}

// 可选值: repeat | repeat-x | repeat-y | no-repeat | space | round | inherit
func BackgroundRepeat(v string) Style {
	return Style{Name: "background-repeat", Value: v}
}

func BackgroundPosition(v string) Style {
	return Style{Name: "background-position", Value: v}
}

func BackgroundSize(v string) Style {
	return Style{Name: "background-size", Value: v}
}

// 可选值: scroll | fixed | local | inherit
func BackgroundAttachment(v string) Style {
	return Style{Name: "background-attachment", Value: v}
}

// 可选值: border-box | padding-box | content-box | text | inherit
func BackgroundClip(v string) Style {
	return Style{Name: "background-clip", Value: v}
}

// 可选值: border-box | padding-box | content-box | inherit
func BackgroundOrigin(v string) Style {
	return Style{Name: "background-origin", Value: v}
}

// 可选值: normal | multiply | screen | overlay | darken | lighten | color-dodge |
// color-burn | hard-light | soft-light | difference | exclusion | hue |
// saturation | color | luminosity | inherit
func BackgroundBlendMode(v string) Style {
	return Style{Name: "background-blend-mode", Value: v}
}

func Border(v string) Style {
	return Style{Name: "border", Value: v}
}

func BorderWidth(v string) Style {
	return Style{Name: "border-width", Value: v}
}

// 可选值: none | hidden | dotted | dashed | solid | double | groove | ridge | inset | outset | inherit
func BorderStyle(v string) Style {
	return Style{Name: "border-style", Value: v}
}

func BorderColor(v string) Style {
	return Style{Name: "border-color", Value: v}
}

func BorderRadius(v string) Style {
	return Style{Name: "border-radius", Value: v}
}

func BorderImage(v string) Style {
	return Style{Name: "border-image", Value: v}
}

func BorderImageSource(v string) Style {
	return Style{Name: "border-image-source", Value: v}
}

func BorderImageSlice(v string) Style {
	return Style{Name: "border-image-slice", Value: v}
}

func BorderImageWidth(v string) Style {
	return Style{Name: "border-image-width", Value: v}
}

func BorderImageOutset(v string) Style {
	return Style{Name: "border-image-outset", Value: v}
}

// 可选值: stretch | repeat | round | space | inherit
func BorderImageRepeat(v string) Style {
	return Style{Name: "border-image-repeat", Value: v}
}

/// 文本与字体

func Color(v string) Style {
	return Style{Name: "color", Value: v}
}

func FontFamily(v string) Style {
	return Style{Name: "font-family", Value: v}
}

func FontSize(v string) Style {
	return Style{Name: "font-size", Value: v}
}

// 可选值: normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
func FontWeight(v string) Style {
	return Style{Name: "font-weight", Value: v}
}

// 可选值: normal | italic | oblique | inherit
func FontStyle(v string) Style {
	return Style{Name: "font-style", Value: v}
}

// 可选值: normal | small-caps | all-small-caps | petite-caps | all-petite-caps |
// titling-caps | unicase | inherit
func FontVariant(v string) Style {
	return Style{Name: "font-variant", Value: v}
}

func LineHeight(v string) Style {
	return Style{Name: "line-height", Value: v}
}

// 可选值: left | right | center | justify | start | end | match-parent | inherit
func TextAlign(v string) Style {
	return Style{Name: "text-align", Value: v}
}

// 可选值: none | underline | overline | line-through | blink | inherit
func TextDecoration(v string) Style {
	return Style{Name: "text-decoration", Value: v}
}

// 可选值: none | capitalize | uppercase | lowercase | full-width | inherit
func TextTransform(v string) Style {
	return Style{Name: "text-transform", Value: v}
}

func TextIndent(v string) Style {
	return Style{Name: "text-indent", Value: v}
}

func TextShadow(v string) Style {
	return Style{Name: "text-shadow", Value: v}
}

// 可选值: clip | ellipsis | inherit
func TextOverflow(v string) Style {
	return Style{Name: "text-overflow", Value: v}
}

func LetterSpacing(v string) Style {
	return Style{Name: "letter-spacing", Value: v}
}

func WordSpacing(v string) Style {
	return Style{Name: "word-spacing", Value: v}
}

// 可选值: normal | nowrap | pre | pre-wrap | pre-line | break-spaces | inherit
func WhiteSpace(v string) Style {
	return Style{Name: "white-space", Value: v}
}

// 可选值: normal | break-all | keep-all | break-word | inherit
func WordBreak(v string) Style {
	return Style{Name: "word-break", Value: v}
}

// 可选值: normal | break-word | anywhere | inherit
func OverflowWrap(v string) Style {
	return Style{Name: "overflow-wrap", Value: v}
}

// 可选值: none | manual | auto | inherit
func Hyphens(v string) Style {
	return Style{Name: "hyphens", Value: v}
}

/// 布局与定位

// 可选值: block | inline | inline-block | none | flex | grid | table |
// inline-table | list-item | inherit
func Display(v string) Style {
	return Style{Name: "display", Value: v}
}

// 可选值: static | relative | absolute | fixed | sticky | inherit
func Position(v string) Style {
	return Style{Name: "position", Value: v}
}

func Top(v string) Style {
	return Style{Name: "top", Value: v}
}

func Right(v string) Style {
	return Style{Name: "right", Value: v}
}

func Bottom(v string) Style {
	return Style{Name: "bottom", Value: v}
}

func Left(v string) Style {
	return Style{Name: "left", Value: v}
}

func Inset(v string) Style {
	return Style{Name: "inset", Value: v}
}

func ZIndex(v string) Style {
	return Style{Name: "z-index", Value: v}
}

// 可选值: none | left | right | inherit
func Float(v string) Style {
	return Style{Name: "float", Value: v}
}

// 可选值: none | left | right | both | inherit
func Clear(v string) Style {
	return Style{Name: "clear", Value: v}
}

// 可选值: visible | hidden | collapse | inherit
func Visibility(v string) Style {
	return Style{Name: "visibility", Value: v}
}

func ClipPath(v string) Style {
	return Style{Name: "clip-path", Value: v}
}

func ShapeOutside(v string) Style {
	return Style{Name: "shape-outside", Value: v}
}

func ShapeMargin(v string) Style {
	return Style{Name: "shape-margin", Value: v}
}

/// Flexbox 布局属性方法

// 可选值: row | row-reverse | column | column-reverse | inherit
func FlexDirection(v string) Style {
	return Style{Name: "flex-direction", Value: v}
}

// 可选值: nowrap | wrap | wrap-reverse | inherit
func FlexWrap(v string) Style {
	return Style{Name: "flex-wrap", Value: v}
}

func FlexFlow(v string) Style {
	return Style{Name: "flex-flow", Value: v}
}

// 可选值: flex-start | flex-end | center | space-between | space-around | space-evenly | inherit
func JustifyContent(v string) Style {
	return Style{Name: "justify-content", Value: v}
}

// 可选值: stretch | flex-start | flex-end | center | baseline | first baseline |
// last baseline | start | end | self-start | self-end | safe | unsafe | inherit
func AlignItems(v string) Style {
	return Style{Name: "align-items", Value: v}
}

// 可选值: flex-start | flex-end | center | space-between | space-around | stretch | inherit
func AlignContent(v string) Style {
	return Style{Name: "align-content", Value: v}
}

func FlexGrow(v string) Style {
	return Style{Name: "flex-grow", Value: v}
}

func FlexShrink(v string) Style {
	return Style{Name: "flex-shrink", Value: v}
}

func FlexBasis(v string) Style {
	return Style{Name: "flex-basis", Value: v}
}

func Flex(v string) Style {
	return Style{Name: "flex", Value: v}
}

// 可选值: auto | flex-start | flex-end | center | baseline | stretch | safe | unsafe | inherit
func AlignSelf(v string) Style {
	return Style{Name: "align-self", Value: v}
}

func Order(v string) Style {
	return Style{Name: "order", Value: v}
}

/// Grid 布局属性方法

func GridTemplateColumns(v string) Style {
	return Style{Name: "grid-template-columns", Value: v}
}

func GridTemplateRows(v string) Style {
	return Style{Name: "grid-template-rows", Value: v}
}

func GridTemplateAreas(v string) Style {
	return Style{Name: "grid-template-areas", Value: v}
}

func GridAutoColumns(v string) Style {
	return Style{Name: "grid-auto-columns", Value: v}
}

func GridAutoRows(v string) Style {
	return Style{Name: "grid-auto-rows", Value: v}
}

// 可选值: row | column | dense | row dense | column dense | inherit
func GridAutoFlow(v string) Style {
	return Style{Name: "grid-auto-flow", Value: v}
}

func GridColumn(v string) Style {
	return Style{Name: "grid-column", Value: v}
}

func GridRow(v string) Style {
	return Style{Name: "grid-row", Value: v}
}

func GridArea(v string) Style {
	return Style{Name: "grid-area", Value: v}
}

func Gap(v string) Style {
	return Style{Name: "gap", Value: v}
}

func RowGap(v string) Style {
	return Style{Name: "row-gap", Value: v}
}

func ColumnGap(v string) Style {
	return Style{Name: "column-gap", Value: v}
}

func PlaceItems(v string) Style {
	return Style{Name: "place-items", Value: v}
}

func PlaceContent(v string) Style {
	return Style{Name: "place-content", Value: v}
}

func PlaceSelf(v string) Style {
	return Style{Name: "place-self", Value: v}
}

// 可选值: masonry | pack | next | ordered | layered | inherit
func MasonryAutoFlow(v string) Style {
	return Style{Name: "masonry-auto-flow", Value: v}
}

/// 动画与变换属性方法

func Transition(v string) Style {
	return Style{Name: "transition", Value: v}
}

func TransitionProperty(v string) Style {
	return Style{Name: "transition-property", Value: v}
}

func TransitionDuration(v string) Style {
	return Style{Name: "transition-duration", Value: v}
}

// 可选值: linear | ease | ease-in | ease-out | ease-in-out | step-start | step-end |
// cubic-bezier(n,n,n,n) | inherit
func TransitionTimingFunction(v string) Style {
	return Style{Name: "transition-timing-function", Value: v}
}

func TransitionDelay(v string) Style {
	return Style{Name: "transition-delay", Value: v}
}

func Animation(v string) Style {
	return Style{Name: "animation", Value: v}
}

func AnimationName(v string) Style {
	return Style{Name: "animation-name", Value: v}
}

func AnimationDuration(v string) Style {
	return Style{Name: "animation-duration", Value: v}
}

// 可选值: linear | ease | ease-in | ease-out | ease-in-out | step-start | step-end |
// cubic-bezier(n,n,n,n) | inherit
func AnimationTimingFunction(v string) Style {
	return Style{Name: "animation-timing-function", Value: v}
}

func AnimationDelay(v string) Style {
	return Style{Name: "animation-delay", Value: v}
}

// 可选值: infinite | <number> | inherit
func AnimationIterationCount(v string) Style {
	return Style{Name: "animation-iteration-count", Value: v}
}

// 可选值: normal | reverse | alternate | alternate-reverse | inherit
func AnimationDirection(v string) Style {
	return Style{Name: "animation-direction", Value: v}
}

// 可选值: none | forwards | backwards | both | inherit
func AnimationFillMode(v string) Style {
	return Style{Name: "animation-fill-mode", Value: v}
}

// 可选值: running | paused | inherit
func AnimationPlayState(v string) Style {
	return Style{Name: "animation-play-state", Value: v}
}

func Transform(v string) Style {
	return Style{Name: "transform", Value: v}
}

func TransformOrigin(v string) Style {
	return Style{Name: "transform-origin", Value: v}
}

// 可选值: flat | preserve-3d | inherit
func TransformStyle(v string) Style {
	return Style{Name: "transform-style", Value: v}
}

func Perspective(v string) Style {
	return Style{Name: "perspective", Value: v}
}

func PerspectiveOrigin(v string) Style {
	return Style{Name: "perspective-origin", Value: v}
}

/// 响应式与交互属性方法

func Media(v string) Style {
	return Style{Name: "@media", Value: v}
}

// 可选值: auto | default | none | context-menu | help | pointer | progress | wait |
// cell | crosshair | text | vertical-text | alias | copy | move | no-drop |
// not-allowed | e-resize | n-resize | ne-resize | nw-resize | s-resize |
// se-resize | sw-resize | w-resize | ew-resize | ns-resize | nesw-resize |
// nwse-resize | col-resize | row-resize | all-scroll | zoom-in | zoom-out |
// grab | grabbing | inherit
func Cursor(v string) Style {
	return Style{Name: "cursor", Value: v}
}

// 可选值: auto | none | text | contain | all | element | inherit
func UserSelect(v string) Style {
	return Style{Name: "user-select", Value: v}
}

// 可选值: auto | none | visiblePainted | visibleFill | visibleStroke | visible |
// painted | fill | stroke | all | inherit
func PointerEvents(v string) Style {
	return Style{Name: "pointer-events", Value: v}
}

// 可选值: none | both | horizontal | vertical | block | inline | inherit
func Resize(v string) Style {
	return Style{Name: "resize", Value: v}
}

// 可选值: auto | smooth | inherit
func ScrollBehavior(v string) Style {
	return Style{Name: "scroll-behavior", Value: v}
}

// 可选值: none | x | y | block | inline | both | mandatory | proximity | inherit
func ScrollSnapType(v string) Style {
	return Style{Name: "scroll-snap-type", Value: v}
}

// 可选值: start | end | center | none | inherit
func ScrollSnapAlign(v string) Style {
	return Style{Name: "scroll-snap-align", Value: v}
}

// 可选值: fill | contain | cover | none | scale-down | inherit
func ObjectFit(v string) Style {
	return Style{Name: "object-fit", Value: v}
}

func ObjectPosition(v string) Style {
	return Style{Name: "object-position", Value: v}
}

func Srcset(v string) Style {
	return Style{Name: "srcset", Value: v}
}

/// 其他

func Opacity(v string) Style {
	return Style{Name: "opacity", Value: v}
}

func Filter(v string) Style {
	return Style{Name: "filter", Value: v}
}

func BackdropFilter(v string) Style {
	return Style{Name: "backdrop-filter", Value: v}
}

// 可选值: normal | multiply | screen | overlay | darken | lighten | color-dodge |
// color-burn | hard-light | soft-light | difference | exclusion | hue |
// saturation | color | luminosity | inherit
func MixBlendMode(v string) Style {
	return Style{Name: "mix-blend-mode", Value: v}
}

func WillChange(v string) Style {
	return Style{Name: "will-change", Value: v}
}

func Content(v string) Style {
	return Style{Name: "content", Value: v}
}

func Quotes(v string) Style {
	return Style{Name: "quotes", Value: v}
}

func CounterReset(v string) Style {
	return Style{Name: "counter-reset", Value: v}
}

func CounterIncrement(v string) Style {
	return Style{Name: "counter-increment", Value: v}
}

func All(v string) Style {
	return Style{Name: "all", Value: v}
}

func AccentColor(v string) Style {
	return Style{Name: "accent-color", Value: v}
}

// 可选值: auto | isolate | inherit
func Isolation(v string) Style {
	return Style{Name: "isolation", Value: v}
}
