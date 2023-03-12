////go:build ignore

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var tags []tag
var attributes = make(map[string]tagAttribute) // attributes and props
var events = make(map[string]tagEvent)
var defines = make(map[string]tagWriter)

func main() {
	initSubContent()
	initAttributes()
	initEvents()
	initTags()

	generateTags()
	fmt.Printf("gen html.go done\n")
}

type tag struct {
	Name string

	// the content, attributes, events needed for current tag
	items []string
}

type tagWriter struct {
	define    tagItemWriter
	implement tagItemWriter
}

type tagItemWriter func(tag) string

type tagAttribute struct {
	name string
	doc  string
}

type tagEvent struct {
	name string
	doc  string
}

func initTags() {
	tags = []tag{
		{
			Name:  "A",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), aOnly()),
		},
		{
			Name:  "Abbr",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Address",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Area",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), areaOnly()),
		},
		{
			Name:  "Article",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Aside",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Audio",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), audioOnly()),
		},
		{
			Name:  "B",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Base",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), baseOnly()),
		},
		{
			Name:  "Bdi",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Bdo",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Blockquote",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), blockquoteOnly()),
		},
		{
			Name:  "Body",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), bodyOnly()),
		},
		{
			Name:  "Br",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Button",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), buttonOnly()),
		},
		{
			Name:  "Canvas",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), canvasOnly()),
		},
		{
			Name:  "Caption",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Cite",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Code",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Col",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), colOnly()),
		},
		{
			Name:  "ColGroup",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), colGroupOnly()),
		},
		{
			Name:  "Data",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), dataOnly()),
		},
		{
			Name:  "DataList",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Dd",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Del",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), delOnly()),
		},
		{
			Name:  "Details",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), detailsOnly()),
		},
		{
			Name:  "Dfn",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Dialog",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), dialogOnly()),
		},
		{
			Name:  "Div",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Dl",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Dt",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Em",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Embed",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), embedOnly()),
		},
		{
			Name:  "FieldSet",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), fieldSetOnly()),
		},
		{
			Name:  "FigCaption",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Figure",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Footer",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Form",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), formSetOnly()),
		},
		{
			Name:  "H1",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "H2",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "H3",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "H4",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "H5",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "H6",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Head",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Header",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "HGroup",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Hr",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Html",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), htmlOnly()),
		},
		{
			Name:  "I",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "IFrame",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), iframeOnly()),
		},
		{
			Name:  "Img",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), imgOnly()),
		},
		{
			Name:  "Input",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), inputOnly()),
		},
		{
			Name:  "Ins",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), insOnly()),
		},
		{
			Name:  "Kbd",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Label",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), labelOnly()),
		},
		{
			Name:  "Legend",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Li",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), liOnly()),
		},
		{
			Name:  "Link",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), linkOnly()),
		},
		{
			Name:  "Main",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "HMap", // Map is used for range a map
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), mapOnly()),
		},
		{
			Name:  "Mark",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Menu",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Meta",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), metaOnly()),
		},
		{
			Name:  "Meter",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), meterOnly()),
		},
		{
			Name:  "Nav",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "NoScript",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Object",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), objectOnly()),
		},
		{
			Name:  "Ol",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), olOnly()),
		},
		{
			Name:  "OptGroup",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), optgroupOnly()),
		},
		{
			Name:  "Option",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), optionOnly()),
		},
		{
			Name:  "Output",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), outputOnly()),
		},
		{
			Name:  "P",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Picture",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Pre",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Progress",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), progressOnly()),
		},
		{
			Name:  "Q",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), qOnly()),
		},
		{
			Name:  "Rp",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Rt",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Ruby",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "S",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Samp",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Script",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), scriptOnly()),
		},
		{
			Name:  "Section",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Select",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), selectOnly()),
		},
		{
			Name:  "Slot",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), slotOnly()),
		},
		{
			Name:  "Small",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Source",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), sourceOnly()),
		},
		{
			Name:  "Span",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Strong",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Style",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), styleOnly()),
		},
		{
			Name:  "Sub",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Summary",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Sup",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Table",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "TBody",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Td",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), tdOnly()),
		},
		{
			Name:  "Template",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Textarea",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), textareaOnly()),
		},
		{
			Name:  "TFoot",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Th",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), thOnly()),
		},
		{
			Name:  "THead",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Time",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), timeOnly()),
		},
		{
			Name:  "Title",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "tr",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Track",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), trackOnly()),
		},
		{
			Name:  "U",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Ul",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Var",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
		{
			Name:  "Video",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents(), videoOnly()),
		},
		{
			Name:  "Wbr",
			items: combineTagItems(subContent(), globalAttributes(), commonEvents()),
		},
	}
}

func initSubContent() {
	defines["body"] = contentBody()
	defines["text"] = contentText()
}

func subContent() (ret []string) {
	return []string{"body", "text"}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes
func _globalAttributeList() (ret []string) {
	return []string{
		"AccessKey",
		"AutoCapitalize",
		"AutoFocus",
		"Class", //
		"ContentEditAble",
		"DataWith", // data-*
		"Dir",
		"DraggAble",
		"EnterKeyHint",
		"Hidden",
		"Id", //
		"Inert",
		"InputMode",
		"Is",
		"ItemId",
		"ItemProp",
		"ItemRef",
		"ItemScope",
		"ItemType",
		"Lang",
		"Nonce",
		"Part",
		"Slot",
		"SpellCheck",
		"Style", //
		"TabIndex",
		"Title",
		"Translate",
		"VirtualKeyBoardPolicy",
	}
}

func initAttributes() {
	globalAtts := _globalAttributeList()

	// add to attributes and defines
	for _, one := range globalAtts {
		key := fmt.Sprintf("%s", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/%s", strings.ToLower(one)),
		}

		switch one {
		case "Class":
			defines[key] = attributeClass()
		case "DataWith":
			defines[key] = attributeDataWith()
		case "Id":
			defines[key] = attributeId()
		case "Style":
			defines[key] = attributeStyle()
		default:
			defines[key] = getPropertyWriter(attributes[key])
		}
	}
}

func globalAttributes() (ret []string) {
	globalAtts := _globalAttributeList()

	// return attributes
	ret = make([]string, len(globalAtts))
	for i, one := range globalAtts {
		ret[i] = fmt.Sprintf("%v", strings.ToLower(one))
	}
	return ret
}

// event from Element: https://developer.mozilla.org/en-US/docs/Web/API/Element
func _commonEventList() (ret []string) {
	return []string{
		"AnimationCancel",
		"AnimationEnd",
		"AnimationIteration",
		"AnimationStart",
		"AuxClick",
		"Blur",
		"Click",
		"CompositionEnd",
		"CompositionStart",
		"CompositionUpdate",
		"ContentVisibilityAutoStateChange",
		"ContextMenu",
		"Copy",
		"Cut",
		"DblClick",
		"Error",
		"Focus",
		"FocusIn",
		"FocusOut",
		"FullScreenChange",
		"FullScreenError",
		"GotPointerCapture",
		"KeyDown",
		"KeyUp",
		"LostPointerCapture",
		"MouseDown",
		"MouseEnter",
		"MouseLeave",
		"MouseMove",
		"MouseOut",
		"MouseOver",
		"MouseUp",
		"Paste",
		"PointerCancel",
		"PointerDown",
		"PointerEnter",
		"PointerLeave",
		"PointerMove",
		"PointerOut",
		"PointerOver",
		"PointerUp",
		"Scroll",
		"ScrollEnd",
		"SecurityPolicyViolation",
		"TouchCancel",
		"TouchEnd",
		"TouchMove",
		"TouchStart",
		"TransitionCancel",
		"TransitionEnd",
		"TransitionRun",
		"TransitionStart",
		"Wheel",
	}
}

// event from HTMLElement: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement
func _commonEventList2() (ret []string) {
	return []string{
		"BeforeInput",
		"Change",
		"Drag",
		"DragEnd",
		"DragEnter",
		"DragLeave",
		"DragOver",
		"DragStart",
		"Drop",
		"Input",
	}
}

func initEvents() {
	// event from Element
	for _, one := range _commonEventList() {
		key := fmt.Sprintf("on-%s", strings.ToLower(one))
		events[key] = tagEvent{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/API/Element/%s_event", strings.ToLower(one)),
		}
		defines[key] = getEventWriter(events[key])
	}

	// event from HTMLElement
	for _, one := range _commonEventList2() {
		key := fmt.Sprintf("on-%s", strings.ToLower(one))
		events[key] = tagEvent{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/%s_event", strings.ToLower(one)),
		}
		defines[key] = getEventWriter(events[key])
	}
}

func commonEvents() (ret []string) {
	// all events from Element
	evts := _commonEventList()
	ret = make([]string, 0, len(evts))
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower(one)))
	}

	// pick events from HTMLElement
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("Drag")))
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("DragEnd")))
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("DragEnter")))
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("DragLeave")))
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("DragOver")))
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("DragStart")))
	ret = append(ret, fmt.Sprintf("on-%s", strings.ToLower("Drop")))
	return
}

func aOnly() (ret []string) {
	attrs := []string{
		"Download",
		"Href",
		"Ping",
		"ReferrerPolicy",
		"Rel",
		"Target",
		"Type",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-a", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-a", strings.ToLower(one))
	}
	return ret
}

func areaOnly() (ret []string) {
	attrs := []string{
		"Alt",
		"Coords",
		"Download",
		"Href",
		"Ping",
		"ReferrerPolicy",
		"Rel",
		"Shape",
		"Target",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-area", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-area", strings.ToLower(one))
	}
	return ret
}

func audioOnly() (ret []string) {
	attrs := []string{
		"AutoPlay",
		"Controls",
		"ControlsList",
		"CrossOrigin",
		"DisableRemotePlayback",
		"Loop",
		"Muted",
		"Preload",
		"Src",
	}

	evts := []string{
		"AudioProcess",
		"CanPlay",
		"CanPlayThrough",
		"Complete",
		"DurationChange",
		"Emptied",
		"Ended",
		"LoadedData",
		"LoadedMetadata",
		"Pause",
		"Play",
		"Playing",
		"RateChange",
		"Seeked",
		"Seeking",
		"Stalled",
		"Suspend",
		"TimeUpdate",
		"VolumeChange",
		"Waiting",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-audio", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-audio", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		events[eKey] = tagEvent{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/%s_event", strings.ToLower(one)),
		}
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-audio", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-audio", strings.ToLower(one)))
	}
	return ret
}

func baseOnly() (ret []string) {
	attrs := []string{
		"Href",
		"Target",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-base", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-base", strings.ToLower(one))
	}
	return ret
}

func blockquoteOnly() (ret []string) {
	attrs := []string{
		"Cite",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-blockquote", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-blockquote", strings.ToLower(one))
	}
	return ret
}

func bodyOnly() (ret []string) {
	attrs := []string{}

	evts := []string{
		"AfterPrint",
		"BeforePrint",
		"BeforeUnload",
		//"Blur",  // already include by common event
		//"Error", // already include by common event
		//"Focus", // already include by common event
		"HashChange",
		"LanguageChange",
		"Load",
		"Message",
		"Offline",
		"Online",
		"PopState",
		"Redo",
		"Resize",
		"Storage",
		"Undo",
		"Unload",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-body", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body#%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-body", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		events[eKey] = tagEvent{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body#%s", strings.ToLower(one)),
		}
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-body", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-body", strings.ToLower(one)))
	}
	return ret
}

func buttonOnly() (ret []string) {
	attrs := []string{
		"AutoComplete",
		"Disabled",
		"Form",
		"FormAction",
		"FormEnctype",
		"FormMethod",
		"FormNoValidate",
		"FormTarget",
		"Name",
		"Type",
		"Value",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-button", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-button", strings.ToLower(one))
	}
	return ret
}

func canvasOnly() (ret []string) {
	attrs := []string{
		"Height",
		"Width",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-canvas", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-canvas", strings.ToLower(one))
	}
	return ret
}

func colOnly() (ret []string) {
	attrs := []string{
		"Span",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-col", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-col", strings.ToLower(one))
	}
	return ret
}

func colGroupOnly() (ret []string) {
	attrs := []string{
		"Span",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-colgroup", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-colgroup", strings.ToLower(one))
	}
	return ret
}

func dataOnly() (ret []string) {
	attrs := []string{
		"Value",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-data", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-data", strings.ToLower(one))
	}
	return ret
}

func delOnly() (ret []string) {
	attrs := []string{
		"Cite",
		"DateTime",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-del", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-del", strings.ToLower(one))
	}
	return ret
}

func detailsOnly() (ret []string) {
	attrs := []string{
		"Open",
	}

	evts := []string{
		"Toggle",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-details", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details#%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-details", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		events[eKey] = tagEvent{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details#%s", strings.ToLower(one)),
		}
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-details", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-details", strings.ToLower(one)))
	}
	return ret
}

func dialogOnly() (ret []string) {
	attrs := []string{
		"Open",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-dialog", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-dialog", strings.ToLower(one))
	}
	return ret
}

func embedOnly() (ret []string) {
	attrs := []string{
		"Height",
		"Src",
		"Type",
		"Width",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-embed", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-embed", strings.ToLower(one))
	}
	return ret
}

func fieldSetOnly() (ret []string) {
	attrs := []string{
		"Disabled",
		"Form",
		"Name",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-fieldSet", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldSet#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-fieldSet", strings.ToLower(one))
	}
	return ret
}

func formSetOnly() (ret []string) {
	attrs := []string{
		//"Accept-Charset", // bad name
		//"AutoCapitalize ", // already include by global attributes
		"AutoComplete",
		"Name",
		"Rel",
		"Action",
		"Enctype",
		"Method",
		"NoValidate",
		"Target",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-form", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-form", strings.ToLower(one))
	}
	return ret
}

func htmlOnly() (ret []string) {
	attrs := []string{
		"Xmlns",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-html", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-html", strings.ToLower(one))
	}
	return ret
}

func iframeOnly() (ret []string) {
	attrs := []string{
		"Allow",
		"AllowFullScreen",
		"AllowPaymentRequest",
		"CredentialLess",
		"Csp",
		"Height",
		"Loading",
		"Name",
		"ReferrerPolicy",
		"Sandbox",
		"Src",
		"SrcDoc",
		"Width",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-iframe", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-iframe", strings.ToLower(one))
	}
	return ret
}

func imgOnly() (ret []string) {
	attrs := []string{
		"Alt",
		"CrossOrigin",
		"Decoding",
		"ElementTiming",
		"FetchPriority",
		"Height",
		"IsMap",
		"Loading",
		"ReferrerPolicy",
		"Sizes",
		"Src",
		"SrcSet",
		"Width",
		"UseMap",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-img", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img#%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-img", strings.ToLower(one))
	}
	return ret
}

func inputOnly() (ret []string) {
	attrs := []string{
		"Type",
		"Accept",
		"Alt",
		"AutoComplete",
		"Capture",
		"Checked",
		"Dirname",
		"Disabled",
		"Form",
		"FormAction",
		"FormEnctype",
		"FormMethod",
		"FormNoValidate",
		"FormTarget",
		"Height",
		"List",
		"Max",
		"MaxLength",
		"Min",
		"MinLength",
		"Multiple",
		"Name",
		"Pattern",
		"Placeholder",
		"Readonly",
		"Required",
		"Size",
		"Src",
		"Step",
		"Value",
		"Width",
	}

	evts := []string{
		"BeforeInput",
		"Change",
		"Input",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-input", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#%s", strings.ToLower(one)),
		}

		switch one {
		case "Value":
			defines[key] = propertyInputValue(attributes[key]) // two ways binding
		default:
			defines[key] = getPropertyWriter(attributes[key])
		}
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-input", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-input", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-input", strings.ToLower(one)))
	}
	return ret
}

func insOnly() (ret []string) {
	attrs := []string{
		"Cite",
		"Datetime",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-ins", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-ins", strings.ToLower(one))
	}
	return ret
}

func labelOnly() (ret []string) {
	attrs := []string{
		"For",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-label", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getAttributeWriter(attributes[key]) // getPropertyWriter doesn't work
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-label", strings.ToLower(one))
	}
	return ret
}

func liOnly() (ret []string) {
	attrs := []string{
		"Value",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-li", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-li", strings.ToLower(one))
	}
	return ret
}

func linkOnly() (ret []string) {
	attrs := []string{
		"As",
		"CrossOrigin",
		"FetchPriority",
		"Href",
		"HrefLang",
		"ImageSizes",
		"ImageSrcSet",
		"Integrity",
		"Media",
		"Prefetch",
		"ReferrerPolicy",
		"Rel",
		"Sizes",
		//"Title", // already include by global attribute
		"Type",
		"Blocking",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-link", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-link", strings.ToLower(one))
	}
	return ret
}

func mapOnly() (ret []string) {
	attrs := []string{
		"Name",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-map", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-map", strings.ToLower(one))
	}
	return ret
}

func metaOnly() (ret []string) {
	attrs := []string{
		"Charset",
		"Content",
		//"Http-Equiv", // bad name
		"Name",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-meta", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-meta", strings.ToLower(one))
	}
	return ret
}

func meterOnly() (ret []string) {
	attrs := []string{
		"Value",
		"Min",
		"Max",
		"Low",
		"High",
		"Optimum",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-meter", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-meter", strings.ToLower(one))
	}
	return ret
}

func objectOnly() (ret []string) {
	attrs := []string{
		"Data",
		"Form",
		"Height",
		"Name",
		"Type",
		"UseMap",
		"Width",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-object", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-object", strings.ToLower(one))
	}
	return ret
}

func olOnly() (ret []string) {
	attrs := []string{
		"Reversed",
		"Start",
		"Type",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-ol", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-ol", strings.ToLower(one))
	}
	return ret
}

func optgroupOnly() (ret []string) {
	attrs := []string{
		"Disabled",
		"Label",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-optgroup", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-optgroup", strings.ToLower(one))
	}
	return ret
}

func optionOnly() (ret []string) {
	attrs := []string{
		"Disabled",
		"Label",
		"Selected",
		"Value",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-option", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-option", strings.ToLower(one))
	}
	return ret
}

func outputOnly() (ret []string) {
	attrs := []string{
		"For",
		"Form",
		"Name",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-output", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-output", strings.ToLower(one))
	}
	return ret
}

func progressOnly() (ret []string) {
	attrs := []string{
		"Max",
		"Value",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-progress", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-progress", strings.ToLower(one))
	}
	return ret
}

func qOnly() (ret []string) {
	attrs := []string{
		"Cite",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-q", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-q", strings.ToLower(one))
	}
	return ret
}

func scriptOnly() (ret []string) {
	attrs := []string{
		"Async",
		"CrossOrigin",
		"Defer",
		"FetchPriority",
		"Integrity",
		"NoModule",
		//"Nonce", // already include by global attribute
		"ReferrerPolicy",
		"Src",
		"Type",
		"Blocking",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-script", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-script", strings.ToLower(one))
	}
	return ret
}

func selectOnly() (ret []string) {
	attrs := []string{
		"AutoComplete",
		//"AutoFocus", // already include by global attribute
		"Disabled",
		"Form",
		"Multiple",
		"Name",
		"Required",
		"Size",
	}

	evts := []string{
		"Change",
		"Input",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-select", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-select", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-select", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-select", strings.ToLower(one)))
	}
	return ret
}

func slotOnly() (ret []string) {
	attrs := []string{
		"Name",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-slot", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-slot", strings.ToLower(one))
	}
	return ret
}

func sourceOnly() (ret []string) {
	attrs := []string{
		"Type",
		"Src",
		"SrcSet",
		"Sizes",
		"Media",
		"Height",
		"Width",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-source", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return img element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-source", strings.ToLower(one))
	}
	return ret
}

func styleOnly() (ret []string) {
	attrs := []string{
		"Media",
		//"Nonce", // already include by global attribute
		//"Title", // already include by global attribute
		"Blocking",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-style", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return img element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-style", strings.ToLower(one))
	}
	return ret
}

func tdOnly() (ret []string) {
	attrs := []string{
		"ColSpan",
		"Headers",
		"RowSpan",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-td", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return img element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-td", strings.ToLower(one))
	}
	return ret
}

func textareaOnly() (ret []string) {
	attrs := []string{
		"AutoComplete",
		"AutoCorrect",
		// "AutoFocus", // already include by global attribute
		"Cols",
		"Disabled",
		"Form",
		"MaxLength",
		"MinLength",
		"Name",
		"Placeholder",
		"Readonly",
		"Required",
		"Rows",
		// "SpellCheck", // already include by global attribute
		"Value", //
		"Wrap",
	}

	evts := []string{
		"BeforeInput",
		"Change",
		"Input",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-textarea", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea#attr-%s", strings.ToLower(one)),
		}

		switch one {
		case "Value":
			defines[key] = propertyInputValue(attributes[key]) // two ways binding
		default:
			defines[key] = getPropertyWriter(attributes[key])
		}
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-textarea", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-textarea", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-textarea", strings.ToLower(one)))
	}
	return ret
}

func thOnly() (ret []string) {
	attrs := []string{
		"Abbr",
		"ColSpan",
		"Headers",
		"RowSpan",
		"Scope",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-th", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return img element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-th", strings.ToLower(one))
	}
	return ret
}

func timeOnly() (ret []string) {
	attrs := []string{
		"Datetime",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-time", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return img element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-time", strings.ToLower(one))
	}
	return ret
}

func trackOnly() (ret []string) {
	attrs := []string{
		"Default",
		"HKind", // Kind is used: UI.Kind()
		"Label",
		"Src",
		"SrcLang",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-track", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track#attr-%s", strings.ToLower(one)),
		}
		defines[key] = getPropertyWriter(attributes[key])
	}

	// return img element attributes and events
	ret = make([]string, len(attrs))
	for i, one := range attrs {
		ret[i] = fmt.Sprintf("%v-track", strings.ToLower(one))
	}
	return ret
}

func videoOnly() (ret []string) {
	attrs := []string{
		"AutoPlay",
		"AutoPictureInPicture",
		"Controls",
		"ControlsList",
		"CrossOrigin",
		"DisablePictureInPicture",
		"DisableRemotePlayback",
		"Height",
		"Loop",
		"Muted",
		"PlaysInline",
		"Poster",
		"Preload",
		"Src",
		"Width",
	}

	evts := []string{
		"CanPlay",
		"CanPlayThrough",
		"Complete",
		"DurationChange",
		"Emptied",
		"Ended",
		//"Error", // already include by common event
		"LoadedData",
		"LoadedMetadata",
		"Pause",
		"Play",
		"Playing",
		"Progress",
		"RateChange",
		"Seeked",
		"Seeking",
		"Stalled",
		"Suspend",
		"TimeUpdate",
		"VolumeChange",
		"Waiting",
	}

	// add to attribute and defines
	for _, one := range attrs {
		key := fmt.Sprintf("%s-video", strings.ToLower(one))
		attributes[key] = tagAttribute{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video#attr-%s", strings.ToLower(one)),
		}

		switch one {
		case "Value":
			defines[key] = propertyInputValue(attributes[key]) // two ways binding
		default:
			defines[key] = getPropertyWriter(attributes[key])
		}
	}

	// add events and defines
	for _, one := range evts {
		dKey := fmt.Sprintf("on-%s-video", strings.ToLower(one))
		eKey := fmt.Sprintf("on-%s", strings.ToLower(one))
		events[eKey] = tagEvent{
			name: one,
			doc:  fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/%s_event", strings.ToLower(one)),
		}
		defines[dKey] = getEventWriter(events[eKey])
	}

	// return element attributes and events
	ret = make([]string, 0, len(attrs)+len(evts))
	for _, one := range attrs {
		ret = append(ret, fmt.Sprintf("%v-video", strings.ToLower(one)))
	}
	for _, one := range evts {
		ret = append(ret, fmt.Sprintf("on-%s-video", strings.ToLower(one)))
	}
	return ret
}

func combineTagItems(items ...[]string) (ret []string) {
	ret = make([]string, 0)
	for _, list := range items {
		ret = append(ret, list...)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}

func generateTags() {
	f, err := os.Create("html.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer f.Sync()

	fmt.Fprintln(f, "package ui")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// this file is generated by go:generate, DO NOT EDIT.")

	for _, t := range tags {
		fmt.Fprintln(f, tagFunc(t))
		fmt.Fprintln(f, tagInterface(t))
		fmt.Fprintln(f, tagStruct(t))

		for _, item := range t.items {
			writer, found := defines[item]
			if !found {
				panic(fmt.Sprintf("can't find writer for %s.%s", t.Name, item))
			}
			fmt.Fprintln(f, writer.implement(t))
		}
	}
}

func tagFunc(t tag) string {
	return fmt.Sprintf(`
		// %s return a html <%s> element
		func %s() Html%s {
			return &html%s{htmlBaseElement: htmlBaseElement{tag: "%s"}}
		}`,
		t.Name, strings.ToLower(t.Name),
		t.Name, t.Name,
		t.Name, strings.ToLower(t.Name))
}

func tagInterface(t tag) string {
	buf := strings.Builder{}

	buf.WriteString(fmt.Sprintf(`
		// Html%s represent a html <%s> element
		type Html%s interface {
			HtmlUI

		`,
		t.Name, strings.ToLower(t.Name),
		t.Name))

	for _, item := range t.items {
		writer, found := defines[item]
		if !found {
			panic(fmt.Sprintf("can't find writer for %s.%s", t.Name, item))
		}
		buf.WriteString(writer.define(t))
	}
	buf.WriteString("}")
	return buf.String()
}

func tagStruct(t tag) string {
	return fmt.Sprintf(`
		type html%s struct {
			htmlBaseElement
		}

		`,
		t.Name)
}

func contentBody() tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// Body add elems as sub content for current element
				Body(elems ...UI) Html%s
		
				`,
				t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
		func (p *html%s) Body(elems ...UI) Html%s {
			p.setBody(elems...)
			return p
		}`,
				t.Name, t.Name)
		},
	}
}

func contentText() tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// Text set val as text conent for current element
				Text(val any) Html%s
		
				`,
				t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) Text(val any) Html%s {
					p.setText(val)
					return p
				}`,
				t.Name, t.Name)
		},
	}
}

// example: about div accesskey
// define:
//
//	 // https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/accesskey
//		AccessKey(char any) HtmlDiv
//
// implement:
//
//	func (p *htmlDiv) AccessKey(val any) HtmlDiv {
//		  p.SetAttribute("accesskey", val)
//		  return p
//	}
func getAttributeWriter(attr tagAttribute) tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// %s
				%s(val any) Html%s
		
				`,
				attr.doc, attr.name, t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) %s(val any) Html%s {
					p.SetAttribute("%s", val)
					return p
				}`,
				t.Name, attr.name, t.Name,
				strings.ToLower(attr.name))
		},
	}
}

// example: about input value
// define:
//
//	 // https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#value
//		Value(val any) HtmlInput
//
// implement:
//
//	func (p *htmlInput) Value(val any) HtmlInput {
//		  p.SetProperty("value", val)
//		  return p
//	}
func getPropertyWriter(attr tagAttribute) tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// %s
				%s(val any) Html%s
		
				`,
				attr.doc, attr.name, t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) %s(val any) Html%s {
					p.SetProperty("%s", val)
					return p
				}`,
				t.Name, attr.name, t.Name,
				strings.ToLower(attr.name))
		},
	}
}

func propertyInputValue(attr tagAttribute) tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// %s
				// val type should be string or *Ref[string] 
				%s(val any) Html%s
		
				`,
				attr.doc, attr.name, t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) %s(val any) Html%s {
					p.bindInputValue(val)
					return p
				}`,
				t.Name, attr.name, t.Name)
		},
	}
}

// data-*
func attributeDataWith() tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/data-*
				DataWith(name string, val any) Html%s
		
				`,
				t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) DataWith(name string, val any) Html%s {
					p.dataWith(name, val)
					return p
				}`,
				t.Name, t.Name)
		},
	}
}

func attributeId() tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// Id set element id attribute
				Id(val string) Html%s
		
				`,
				t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) Id(val string) Html%s {
					p.SetUIElementId(val)
					return p
				}`,
				t.Name, t.Name)
		},
	}
}

func attributeClass() tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// Class set the class attribute with name, and specify the class content with items
				// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/class
				Class(name string, items ...*Ref[StyleItem]) Html%s
		
				`,
				t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) Class(name string, items ...*Ref[StyleItem]) Html%s {
					p.setClass(name, items...)
					return p
				}`,
				t.Name, t.Name)
		},
	}
}

func attributeStyle() tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// Style set the style attribute with items
				// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/style
				Style(items ...*Ref[StyleItem]) Html%s
		
				`,
				t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) Style(items ...*Ref[StyleItem]) Html%s {
					p.setStyle(items...)
					return p
				}`,
				t.Name, t.Name)
		},
	}
}

// example: about div click
// define:
//
//	 // OnClick register listener handler for click event
//		OnClick(listener func(event Event), useCapture bool) HtmlDiv
//
// implement:
//
//	func (p *htmlDiv) OnClick(listener func(event Event), useCapture bool) HtmlDiv {
//		  p.registerEventHandler("click", listener, useCapture)
//		  return p
//	}
func getEventWriter(event tagEvent) tagWriter {
	return tagWriter{
		define: func(t tag) string {
			return fmt.Sprintf(`
				// %s
				On%s(listener func(event Event)) Html%s
		
				`,
				event.doc,
				event.name, t.Name)
		},
		implement: func(t tag) string {
			return fmt.Sprintf(`
				func (p *html%s) On%s(listener func(event Event)) Html%s {
					p.registerEventHandler("%s", listener)
					return p
				}`,
				t.Name, event.name, t.Name,
				strings.ToLower(event.name))
		},
	}
}
