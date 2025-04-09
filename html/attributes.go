package html

import "github.com/go-andiamo/aitch"

// Async declares an HTML "async" boolean attribute
//
// only applies to Script <script> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Async() aitch.Node {
	return aitch.NewBooleanAttribute(attrAsync)
}

// AutoFocus declares an HTML "autofocus" boolean attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autofocus
func AutoFocus() aitch.Node {
	return aitch.NewBooleanAttribute(attrAutofocus)
}

// AutoPlay declares an HTML "autoplay" boolean attribute
//
// only applies to Audio <audio> and Video <video> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
// & https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func AutoPlay() aitch.Node {
	return aitch.NewBooleanAttribute(attrAutoplay)
}

// Checked declares an HTML "checked" boolean attribute
//
// only applies to Input <input> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Checked() aitch.Node {
	return aitch.NewBooleanAttribute(attrChecked)
}

// Controls declares an HTML "controls" boolean attribute
//
// only applies to Audio <audio> and Video <video> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
// & https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Controls() aitch.Node {
	return aitch.NewBooleanAttribute(attrControls)
}

// CrossOrigin declares an HTML "crossorigin" attribute
//
// only applies to Img <img>, Script <script>, Link <link>, Audio <audio>, Video <video>, IFrame <iframe> & svg.Use <use> elements
func CrossOrigin(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCrossorigin, value...)
}

// DateTime declares an HTML "datetime" attribute
//
// only applies to Time <time> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func DateTime(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDatetime, value...)
}

// Defer declares an HTML "defer" boolean attribute
//
// only applies to Script <script> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Defer() aitch.Node {
	return aitch.NewBooleanAttribute(attrDefer)
}

// Disabled declares an HTML "disabled" boolean attribute
//
// only applies to Button <button>, Input <input>, Select <select>, Textarea <textarea>, Option <option>, OptGroup <optgroup>, FieldSet <fieldset> & Form <form> elements
func Disabled() aitch.Node {
	return aitch.NewBooleanAttribute(attrDisabled)
}

// Download declares an HTML "download" attribute
//
// only applies to A <a> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Download(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDownload, value...)
}

// Draggable declares an HTML "draggable" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/draggable
func Draggable(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDraggable, value...)
}

// Loop declares an HTML "loop" boolean attribute
//
// only applies to Audio <audio> and Video <video> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
// & https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Loop() aitch.Node {
	return aitch.NewBooleanAttribute(attrLoop)
}

// Multiple declares an HTML "multiple" boolean attribute
//
// only applies to Input <input> and Select <select> elements
func Multiple() aitch.Node {
	return aitch.NewBooleanAttribute(attrMultiple)
}

// Muted declares an HTML "muted" boolean attribute
//
// only applies to Audio <audio> and Video <video> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
// & https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Muted() aitch.Node {
	return aitch.NewBooleanAttribute(attrMuted)
}

// PlaysInline declares an HTML "playsinline" boolean attribute
//
// only applies to Video <video> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func PlaysInline() aitch.Node {
	return aitch.NewBooleanAttribute(attrPlaysinline)
}

// ReadOnly declares an HTML "readonly" boolean attribute
//
// only applies to Input <input> and Textarea <textarea> elements
func ReadOnly() aitch.Node {
	return aitch.NewBooleanAttribute(attrReadonly)
}

// Required declares an HTML "required" boolean attribute
//
// only applies to Input <input>, Select <select> and Textarea <textarea> elements
func Required() aitch.Node {
	return aitch.NewBooleanAttribute(attrRequired)
}

// Selected declares an HTML "selected" boolean attribute
//
// only applies to Option <option> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Selected() aitch.Node {
	return aitch.NewBooleanAttribute(attrSelected)
}

// Accept declares an HTML "accept" attribute
//
// only applies to Input <input> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Accept(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAccept, value...)
}

// Action declares an HTML "action" attribute
//
// only applies to Form <form> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Action(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAction, value...)
}

// Alt declares an HTML "alt" attribute
//
// only applies to Img <img>, Area <area> & Input <input type="image"> elements
func Alt(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAlt, value...)
}

// Aria declares a custom HTML "aria-*" attribute using the provided name and value
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes
func Aria(name string, value ...any) aitch.Node {
	return aitch.NewAttribute(append(attrAria, []byte(name)...), value...)
}

// As declares an HTML "as" attribute
//
// only applies to Link <link> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func As(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAs, value...)
}

// AutoComplete declares an HTML "autocomplete" attribute
//
// only applies to Form <form>, Input <input>, Select <select> & Textarea <textarea> elements
func AutoComplete(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAutocomplete, value...)
}

// Charset declares an HTML "charset" attribute
//
// only applies to Meta <meta> & Script <script> elements
func Charset(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCharset, value...)
}

// CiteAttr declares an HTML "cite" attribute
//
// only applies to BlockQuote <blockquote>, Q <q>, Del <del> & Ins <ins> elements
func CiteAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCite, value...)
}

// Class declares an HTML "class" attribute
//
// Use multiple Class() declarations on a single element - the values are joined into one "class" attribute delimited with " " (space)
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/class
func Class(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrClass, space, value...)
}

// Cols declares an HTML "cols" attribute
//
// only applies to Textarea <textarea> element
func Cols(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCols, value...)
}

// ColSpan declares an HTML "colspan" attribute
//
// only applies to Td <td> & Th <th> elements
func ColSpan(value ...any) aitch.Node {
	return aitch.NewAttribute(attrColspan, value...)
}

// Content declares an HTML "content" attribute
//
// only applies to Meta <meta> element
func Content(value ...any) aitch.Node {
	return aitch.NewAttribute(attrContent, value...)
}

// ContentEditable declares an HTML "contenteditable" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/contenteditable
func ContentEditable(value ...any) aitch.Node {
	return aitch.NewAttribute(attrContenteditable, value...)
}

// Dir declares an HTML "dir" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/dir
func Dir(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDir, value...)
}

// Data declares an HTML data attribute prefixed with "data-".
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/data-*
func Data(name string, value ...any) aitch.Node {
	return aitch.NewAttribute(append(attrData, []byte(name)...), value...)
}

// EncType declares an HTML "enctype" attribute
//
// only applies to Form <form> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func EncType(value ...any) aitch.Node {
	return aitch.NewAttribute(attrEnctype, value...)
}

// For declares an HTML "for" attribute
//
// only applies to Label <label> & Output <output> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func For(value ...any) aitch.Node {
	return aitch.NewAttribute(attrFor, value...)
}

// FormAttr declares an HTML "form" attribute
//
// only applies to Input <input>, Button <button>, Select <select>, Textarea <textarea>, Output <output>, FieldSet <fieldset>, Label <label> & Object <object> elements
func FormAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrForm, value...)
}

// NoValidate declares an HTML "novalidate" boolean attribute
//
// only applies to Form <form> element
func NoValidate() aitch.Node {
	return aitch.NewBooleanAttribute(attrNovalidate)
}

// FormNoValidate declares an HTML "formnovalidate" boolean attribute
//
// only applies to Button <button> & Input <input> elements
func FormNoValidate() aitch.Node {
	return aitch.NewBooleanAttribute(attrFormnovalidate)
}

// Height declares an HTML "height" attribute
//
// only applies to Img <img>, Video <video>, Canvas <canvas>, IFrame <iframe>, Object <object>, Embed <embed> & Input <input type="image"> elements
func Height(value ...any) aitch.Node {
	return aitch.NewAttribute(attrHeight, value...)
}

// Hidden declares an HTML "hidden" boolean attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/hidden
func Hidden() aitch.Node {
	return aitch.NewBooleanAttribute(attrHidden)
}

// Href declares an HTML "href" attribute
//
// only applies to A <a>, Area <area>, Base <base> & Link <link> elements
func Href(value ...any) aitch.Node {
	return aitch.NewAttribute(attrHref, value...)
}

// Id declares an HTML "id" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id
func Id(value ...any) aitch.Node {
	return aitch.NewAttribute(attrId, value...)
}

// Integrity declares an HTML "integrity" attribute
//
// only applies to Link <link> & Script <script> elements
func Integrity(value ...any) aitch.Node {
	return aitch.NewAttribute(attrIntegrity, value...)
}

// Lang declares an HTML "lang" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/lang
func Lang(value ...any) aitch.Node {
	return aitch.NewAttribute(attrLang, value...)
}

// List declares an HTML "list" attribute
//
// only applies to Input <input> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func List(value ...any) aitch.Node {
	return aitch.NewAttribute(attrList, value...)
}

// Loading declares an HTML "loading" attribute
//
// only applies to Img <img> & IFrame <iframe> elements
func Loading(value ...any) aitch.Node {
	return aitch.NewAttribute(attrLoading, value...)
}

// Max declares an HTML "max" attribute
//
// only applies to Input <input> & Meter <meter> elements
func Max(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMax, value...)
}

// MaxLength declares an HTML "maxlength" attribute
//
// only applies to Input <input> & Textarea <textarea> elements
func MaxLength(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMaxlength, value...)
}

// Method declares an HTML "method" attribute
//
// only applies to Form <form> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Method(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMethod, value...)
}

// Min declares an HTML "min" attribute
//
// only applies to Input <input> & Meter <meter> elements
func Min(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMin, value...)
}

// MinLength declares an HTML "minlength" attribute
//
// only applies to Input <input> & Textarea <textarea> elements
func MinLength(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMinlength, value...)
}

// Name declares an HTML "name" attribute
//
// only applies to Input <input>, Textarea <textarea>, Select <select>, Button <button>, Form <form>, FieldSet <fieldset>, Object <object> & Output <output> elements
func Name(value ...any) aitch.Node {
	return aitch.NewAttribute(attrName, value...)
}

// Pattern declares an HTML "pattern" attribute
//
// only applies to Input <input> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Pattern(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPattern, value...)
}

// Placeholder declares an HTML "placeholder" attribute
//
// only applies to Input <input> & Textarea <textarea> elements
func Placeholder(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPlaceholder, value...)
}

// PopOver declares an HTML "popover" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/popover
func PopOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPopover, value...)
}

// PopOverTarget declares an HTML "popovertarget" attribute
func PopOverTarget(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPopovertarget, value...)
}

// PopOverTargetAction declares an HTML "popovertargetaction" attribute
func PopOverTargetAction(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPopovertargetaction, value...)
}

// Poster declares an HTML "poster" attribute
//
// only applies to Video <video> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Poster(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPoster, value...)
}

// Preload declares an HTML "preload" attribute
//
// only applies to Audio <audio> and Video <video> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
// & https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Preload(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPreload, value...)
}

// ReferrerPolicy declares an HTML "referrerpolicy" attribute
//
// only applies to A <a>, Img <img>, IFrame <iframe>, Script <script> & Link <link> elements
func ReferrerPolicy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrReferrerpolicy, value...)
}

// Rel declares an HTML "rel" attribute
//
// only applies to A <a>, Area <area>, Link <link> & Form <form> elements
func Rel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRel, value...)
}

// Role declares an HTML "role" attribute
func Role(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRole, value...)
}

// Rows declares an HTML "rows" attribute
//
// only applies to Textarea <textarea> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func Rows(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRows, value...)
}

// RowSpan declares an HTML "rowspan" attribute
//
// only applies to Td <td> & Th <th> elements
func RowSpan(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRowspan, value...)
}

// SlotAttr declares an HTML "slot" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/slot
func SlotAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSlot, value...)
}

// Src declares an HTML "src" attribute
//
// only applies to Img <img>, Audio <audio>, Video <video>, IFrame <iframe>, Script <script>, Source <source> & <frame> elements
func Src(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSrc, value...)
}

// SrcSet declares an HTML "srcset" attribute
//
// only applies to Img <img> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func SrcSet(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSrcset, value...)
}

// Step declares an HTML "step" attribute
//
// only applies to Input <input type="number"> & <input type="range"> elements
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Step(value ...any) aitch.Node {
	return aitch.NewAttribute(attrStep, value...)
}

// Style declares an HTML "style" attribute
//
// Use multiple Style() declarations on a single element - the values are joined into one "style" attribute delimited with "; "
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/style
func Style(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrStyle, stylesDelim, value...)
}

// TabIndex declares an HTML "tabindex" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/tabindex
func TabIndex(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTabindex, value...)
}

// Target declares an HTML "target" attribute
//
// only applies to A <a> & Form <form> elements
func Target(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTarget, value...)
}

// Title declares an HTML "title" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/title
func Title(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTitle, value...)
}

// Type declares an HTML "type" attribute
//
// only applies to Input <input>, Button <button>, Script <script>, Style <style>, Form <form> & Source <source> elements
func Type(value ...any) aitch.Node {
	return aitch.NewAttribute(attrType, value...)
}

// Value declares an HTML "value" attribute
//
// only applies to Input <input>, Option <option> & Button <button> elements
func Value(value ...any) aitch.Node {
	return aitch.NewAttribute(attrValue, value...)
}

// Width declares an HTML "width" attribute
//
// only applies to Img <img>, Video <video>, Canvas <canvas>, Table <table>, Input <input> & Textarea <textarea> elements
func Width(value ...any) aitch.Node {
	return aitch.NewAttribute(attrWidth, value...)
}

// AllowFullScreen declares an HTML "allowfullscreen" boolean attribute
//
// only applies to IFrame <iframe> element
func AllowFullScreen() aitch.Node {
	return aitch.NewBooleanAttribute(attrAllowfullscreen)
}

// Default declares an HTML "default" boolean attribute
//
// only applies to Option <option> & Track <track> elements
func Default() aitch.Node {
	return aitch.NewBooleanAttribute(attrDefault)
}

// Inert declares an HTML "inert" boolean attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inert
func Inert() aitch.Node {
	return aitch.NewBooleanAttribute(attrInert)
}

// IsMap declares an HTML "ismap" boolean attribute
//
// only applies to Img <img> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func IsMap() aitch.Node {
	return aitch.NewBooleanAttribute(attrIsmap)
}

// ItemScope declares an HTML "itemscope" boolean attribute
func ItemScope() aitch.Node {
	return aitch.NewBooleanAttribute(attrItemscope)
}

// NoModule declares an HTML "nomodule" boolean attribute
//
// only applies to Script <script> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func NoModule() aitch.Node {
	return aitch.NewBooleanAttribute(attrNomodule)
}

// Open declares an HTML "open" boolean attribute
//
// only applies to Details <details> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Open() aitch.Node {
	return aitch.NewBooleanAttribute(attrOpen)
}

// Reversed declares an HTML "reversed" boolean attribute
//
// only applies to Ol <ol> element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func Reversed() aitch.Node {
	return aitch.NewBooleanAttribute(attrReversed)
}

// Spellcheck declares an HTML "spellcheck" boolean attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/spellcheck
func Spellcheck() aitch.Node {
	return aitch.NewBooleanAttribute(attrSpellcheck)
}

// Translate declares an HTML "translate" boolean attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/translate
func Translate() aitch.Node {
	return aitch.NewBooleanAttribute(attrTranslate)
}

// AriaAtomic declares an HTML "aria-atomic" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-atomic
func AriaAtomic(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaAtomic, value...)
}

// AriaBusy declares an HTML "aria-busy" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-busy
func AriaBusy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaBusy, value...)
}

// AriaChecked declares an HTML "aria-checked" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-checked
func AriaChecked(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaChecked, value...)
}

// AriaControls declares an HTML "aria-controls" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-controls
func AriaControls(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaControls, value...)
}

// AriaDescribedBy declares an HTML "aria-describedby" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-describedby
func AriaDescribedBy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaDescribedby, value...)
}

// AriaDisabled declares an HTML "aria-disabled" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-disabled
func AriaDisabled(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaDisabled, value...)
}

// AriaExpanded declares an HTML "aria-expanded" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-expanded
func AriaExpanded(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaExpanded, value...)
}

// AriaFlowTo declares an HTML "aria-flowto" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-flowto
func AriaFlowTo(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaFlowto, value...)
}

// AriaHidden declares an HTML "aria-hidden" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-hidden
func AriaHidden(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaHidden, value...)
}

// AriaInvalid declares an HTML "aria-invalid" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-invalid
func AriaInvalid(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaInvalid, value...)
}

// AriaLabel declares an HTML "aria-label" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-label
func AriaLabel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLabel, value...)
}

// AriaLabelledBy declares an HTML "aria-labelledby" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-labelledby
func AriaLabelledBy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLabelledby, value...)
}

// AriaLive declares an HTML "aria-live" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-live
func AriaLive(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLive, value...)
}

// AriaOwns declares an HTML "aria-owns" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-owns
func AriaOwns(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaOwns, value...)
}

// AriaPlaceholder declares an HTML "aria-placeholder" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-placeholder
func AriaPlaceholder(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPlaceholder, value...)
}

// AriaPosInSet declares an HTML "aria-posinset" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-posinset
func AriaPosInSet(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPosinset, value...)
}

// AriaPressed declares an HTML "aria-pressed" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-pressed
func AriaPressed(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPressed, value...)
}

// AriaReadonly declares an HTML "aria-readonly" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-readonly
func AriaReadonly(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaReadonly, value...)
}

// AriaRelevant declares an HTML "aria-relevant" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-relevant
func AriaRelevant(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaRelevant, value...)
}

// AriaRequired declares an HTML "aria-required" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-required
func AriaRequired(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaRequired, value...)
}

// AriaSelected declares an HTML "aria-selected" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-selected
func AriaSelected(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaSelected, value...)
}

// AriaSetSize declares an HTML "aria-setsize" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-setsize
func AriaSetSize(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaSetsize, value...)
}

// AriaValueMax declares an HTML "aria-valuemax" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-valuemax
func AriaValueMax(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuemax, value...)
}

// AriaValueMin declares an HTML "aria-valuemin" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-valuemin
func AriaValueMin(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuemin, value...)
}

// AriaValueNow declares an HTML "aria-valuenow" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-valuenow
func AriaValueNow(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuenow, value...)
}

// AriaValueText declares an HTML "aria-valuetext" attribute
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Attributes/aria-valuetext
func AriaValueText(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuetext, value...)
}

// OnAbort declares an HTML "onabort" event attribute
func OnAbort(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAbort, value...)
}

// OnBeforeUnload declares an HTML "onbeforeunload" event attribute
func OnBeforeUnload(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeUnload, value...)
}

// OnBlur declares an HTML "onblur" event attribute
func OnBlur(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBlur, value...)
}

// OnCanPlay declares an HTML "oncanplay" event attribute
func OnCanPlay(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCanPlay, value...)
}

// OnCanPlayThrough declares an HTML "oncanplaythrough" event attribute
func OnCanPlayThrough(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCanPlayThrough, value...)
}

// OnChange declares an HTML "onchange" event attribute
func OnChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnChange, value...)
}

// OnClick declares an HTML "onclick" event attribute
func OnClick(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnClick, value...)
}

// OnContextMenu declares an HTML "oncontextmenu" event attribute
func OnContextMenu(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnContextMenu, value...)
}

// OnCopy declares an HTML "oncopy" event attribute
func OnCopy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCopy, value...)
}

// OnCut declares an HTML "oncut" event attribute
func OnCut(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCut, value...)
}

// OnDblClick declares an HTML "ondblclick" event attribute
func OnDblClick(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDblClick, value...)
}

// OnDrag declares an HTML "ondrag" event attribute
func OnDrag(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDrag, value...)
}

// OnDragEnd declares an HTML "ondragend" event attribute
func OnDragEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragEnd, value...)
}

// OnDragEnter declares an HTML "ondragenter" event attribute
func OnDragEnter(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragEnter, value...)
}

// OnDragLeave declares an HTML "ondragleave" event attribute
func OnDragLeave(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragLeave, value...)
}

// OnDragOver declares an HTML "ondragover" event attribute
func OnDragOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragOver, value...)
}

// OnDragStart declares an HTML "ondragstart" event attribute
func OnDragStart(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragStart, value...)
}

// OnDrop declares an HTML "ondrop" event attribute
func OnDrop(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDrop, value...)
}

// OnDurationChange declares an HTML "ondurationchange" event attribute
func OnDurationChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDurationChange, value...)
}

// OnEnded declares an HTML "onended" event attribute
func OnEnded(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnEnded, value...)
}

// OnError declares an HTML "onerror" event attribute
func OnError(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnError, value...)
}

// OnFocus declares an HTML "onfocus" event attribute
func OnFocus(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnFocus, value...)
}

// OnHashChange declares an HTML "onhashchange" event attribute
func OnHashChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnHashChange, value...)
}

// OnInput declares an HTML "oninput" event attribute
func OnInput(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnInput, value...)
}

// OnInvalid declares an HTML "oninvalid" event attribute
func OnInvalid(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnInvalid, value...)
}

// OnKeyDown declares an HTML "onkeydown" event attribute
func OnKeyDown(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyDown, value...)
}

// OnKeyPress declares an HTML "onkeypress" event attribute
func OnKeyPress(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyPress, value...)
}

// OnKeyUp declares an HTML "onkeyup" event attribute
func OnKeyUp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyUp, value...)
}

// OnLoad declares an HTML "onload" event attribute
func OnLoad(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoad, value...)
}

// OnLoadedData declares an HTML "onloadeddata" event attribute
func OnLoadedData(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoadedData, value...)
}

// OnLoadedMetadata declares an HTML "onloadedmetadata" event attribute
func OnLoadedMetadata(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoadedMetadata, value...)
}

// OnMouseDown declares an HTML "onmousedown" event attribute
func OnMouseDown(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseDown, value...)
}

// OnMouseEnter declares an HTML "onmouseenter" event attribute
func OnMouseEnter(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseEnter, value...)
}

// OnMouseLeave declares an HTML "onmouseleave" event attribute
func OnMouseLeave(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseLeave, value...)
}

// OnMouseMove declares an HTML "onmousemove" event attribute
func OnMouseMove(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseMove, value...)
}

// OnMouseOut declares an HTML "onmouseout" event attribute
func OnMouseOut(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseOut, value...)
}

// OnMouseOver declares an HTML "onmouseover" event attribute
func OnMouseOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseOver, value...)
}

// OnMouseUp declares an HTML "onmouseup" event attribute
func OnMouseUp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseUp, value...)
}

// OnPageHide declares an HTML "onpagehide" event attribute
func OnPageHide(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPageHide, value...)
}

// OnPageShow declares an HTML "onpageshow" event attribute
func OnPageShow(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPageShow, value...)
}

// OnPaste declares an HTML "onpaste" event attribute
func OnPaste(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPaste, value...)
}

// OnPause declares an HTML "onpause" event attribute
func OnPause(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPause, value...)
}

// OnPlay declares an HTML "onplay" event attribute
func OnPlay(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPlay, value...)
}

// OnPlaying declares an HTML "onplaying" event attribute
func OnPlaying(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPlaying, value...)
}

// OnPopState declares an HTML "onpopstate" event attribute
func OnPopState(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPopState, value...)
}

// OnProgress declares an HTML "onprogress" event attribute
func OnProgress(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnProgress, value...)
}

// OnRateChange declares an HTML "onratechange" event attribute
func OnRateChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnRateChange, value...)
}

// OnReset declares an HTML "onreset" event attribute
func OnReset(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnReset, value...)
}

// OnResize declares an HTML "onresize" event attribute
func OnResize(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnResize, value...)
}

// OnScroll declares an HTML "onscroll" event attribute
func OnScroll(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnScroll, value...)
}

// OnSeeked declares an HTML "onseeked" event attribute
func OnSeeked(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSeeked, value...)
}

// OnSeeking declares an HTML "onseeking" event attribute
func OnSeeking(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSeeking, value...)
}

// OnSelect declares an HTML "onselect" event attribute
func OnSelect(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSelect, value...)
}

// OnStalled declares an HTML "onstalled" event attribute
func OnStalled(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnStalled, value...)
}

// OnSubmit declares an HTML "onsubmit" event attribute
func OnSubmit(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSubmit, value...)
}

// OnSuspend declares an HTML "onsuspend" event attribute
func OnSuspend(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSuspend, value...)
}

// OnTimeUpdate declares an HTML "ontimeupdate" event attribute
func OnTimeUpdate(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTimeUpdate, value...)
}

// OnToggle declares an HTML "ontoggle" event attribute
func OnToggle(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnToggle, value...)
}

// OnUnload declares an HTML "onunload" event attribute
func OnUnload(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnUnload, value...)
}

// OnVolumeChange declares an HTML "onvolumechange" event attribute
func OnVolumeChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnVolumeChange, value...)
}

// OnWaiting declares an HTML "onwaiting" event attribute
func OnWaiting(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnWaiting, value...)
}

// OnPointerDown declares an HTML "onpointerdown" event attribute
func OnPointerDown(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerDown, value...)
}

// OnPointerUp declares an HTML "onpointerup" event attribute
func OnPointerUp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerUp, value...)
}

// OnPointerMove declares an HTML "onpointermove" event attribute
func OnPointerMove(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerMove, value...)
}

// OnPointerEnter declares an HTML "onpointerenter" event attribute
func OnPointerEnter(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerEnter, value...)
}

// OnPointerLeave declares an HTML "onpointerleave" event attribute
func OnPointerLeave(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerLeave, value...)
}

// OnPointerOver declares an HTML "onpointerover" event attribute
func OnPointerOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerOver, value...)
}

// OnPointerOut declares an HTML "onpointerout" event attribute
func OnPointerOut(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerOut, value...)
}

// OnPointerCancel declares an HTML "onpointercancel" event attribute
func OnPointerCancel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerCancel, value...)
}

// OnTouchStart declares an HTML "ontouchstart" event attribute
func OnTouchStart(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchStart, value...)
}

// OnTouchMove declares an HTML "ontouchmove" event attribute
func OnTouchMove(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchMove, value...)
}

// OnTouchEnd declares an HTML "ontouchend" event attribute
func OnTouchEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchEnd, value...)
}

// OnTouchCancel declares an HTML "ontouchcancel" event attribute
func OnTouchCancel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchCancel, value...)
}

// OnAnimationStart declares an HTML "onanimationstart" event attribute
func OnAnimationStart(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAnimationStart, value...)
}

// OnAnimationEnd declares an HTML "onanimationend" event attribute
func OnAnimationEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAnimationEnd, value...)
}

// OnAnimationIteration declares an HTML "onanimationiteration" event attribute
func OnAnimationIteration(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAnimationIteration, value...)
}

// OnTransitionEnd declares an HTML "ontransitionend" event attribute
func OnTransitionEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTransitionEnd, value...)
}

// OnBeforeInput declares an HTML "onbeforeinput" event attribute
func OnBeforeInput(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeInput, value...)
}

// OnSlotChange declares an HTML "onslotchange" event attribute
func OnSlotChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSlotChange, value...)
}

// OnFormData declares an HTML "onformdata" event attribute
func OnFormData(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnFormData, value...)
}

// AccessKey declares an HTML "accesskey" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/accesskey
func AccessKey(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAccessKey, value...)
}

// AutoCapitalize declares an HTML "autocapitalize" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autocapitalize
func AutoCapitalize(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAutoCapitalize, value...)
}

// AutoCorrect declares an HTML "autocorrect" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autocorrect
func AutoCorrect(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAutoCorrect, value...)
}

// EnterKeyHint declares an HTML "enterkeyhint" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/enterkeyhint
func EnterKeyHint(value ...any) aitch.Node {
	return aitch.NewAttribute(attrEnterKeyHint, value...)
}

// ExportParts declares an HTML "exportparts" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/exportparts
func ExportParts(value ...any) aitch.Node {
	return aitch.NewAttribute(attrExportParts, value...)
}

// InputMode declares an HTML "inputmode" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/inputmode
func InputMode(value ...any) aitch.Node {
	return aitch.NewAttribute(attrInputMode, value...)
}

// Is declares an HTML "is" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/is
func Is(value ...any) aitch.Node {
	return aitch.NewAttribute(attrIs, value...)
}

// ItemId declares an HTML "itemid" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemid
func ItemId(value ...any) aitch.Node {
	return aitch.NewAttribute(attrItemId, value...)
}

// ItemProp declares an HTML "itemprop" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemprop
func ItemProp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrItemProp, value...)
}

// ItemRef declares an HTML "itemref" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemref
func ItemRef(value ...any) aitch.Node {
	return aitch.NewAttribute(attrItemRef, value...)
}

// ItemType declares an HTML "itemtype" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/itemtype
func ItemType(value ...any) aitch.Node {
	return aitch.NewAttribute(attrItemType, value...)
}

// Nonce declares an HTML "nonce" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/nonce
func Nonce(value ...any) aitch.Node {
	return aitch.NewAttribute(attrNonce, value...)
}

// Part declares an HTML "part" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/part
func Part(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPart, value...)
}

// WritingSuggestions declares an HTML "writingsuggestions" attribute
//
// global attribute - applies to any element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/writingsuggestions
func WritingSuggestions(value ...any) aitch.Node {
	return aitch.NewAttribute(attrWritingSuggestions, value...)
}

var (
	space                    = []byte{' '}
	stylesDelim              = []byte{';', ' '}
	attrAccept               = []byte("accept")
	attrAction               = []byte("action")
	attrAlt                  = []byte("alt")
	attrAria                 = []byte("aria-")
	attrAs                   = []byte("as")
	attrAsync                = []byte("async")
	attrAutocomplete         = []byte("autocomplete")
	attrAutofocus            = []byte("autofocus")
	attrAutoplay             = []byte("autoplay")
	attrCharset              = []byte("charset")
	attrChecked              = []byte("checked")
	attrCite                 = []byte("cite")
	attrClass                = []byte("class")
	attrCols                 = []byte("cols")
	attrColspan              = []byte("colspan")
	attrContent              = []byte("content")
	attrContenteditable      = []byte("contenteditable")
	attrControls             = []byte("controls")
	attrCrossorigin          = []byte("crossorigin")
	attrData                 = []byte("data-")
	attrDatetime             = []byte("datetime")
	attrDefer                = []byte("defer")
	attrDir                  = []byte("dir")
	attrDisabled             = []byte("disabled")
	attrDownload             = []byte("download")
	attrDraggable            = []byte("draggable")
	attrEnctype              = []byte("enctype")
	attrFor                  = []byte("for")
	attrForm                 = []byte("form")
	attrHeight               = []byte("height")
	attrHidden               = []byte("hidden")
	attrHref                 = []byte("href")
	attrId                   = []byte("id")
	attrIntegrity            = []byte("integrity")
	attrLang                 = []byte("lang")
	attrList                 = []byte("list")
	attrLoading              = []byte("loading")
	attrLoop                 = []byte("loop")
	attrMax                  = []byte("max")
	attrMaxlength            = []byte("maxlength")
	attrMethod               = []byte("method")
	attrMin                  = []byte("min")
	attrMinlength            = []byte("minlength")
	attrMultiple             = []byte("multiple")
	attrMuted                = []byte("muted")
	attrName                 = []byte("name")
	attrNovalidate           = []byte("novalidate")
	attrFormnovalidate       = []byte("formnovalidate")
	attrPattern              = []byte("pattern")
	attrPlaceholder          = []byte("placeholder")
	attrPlaysinline          = []byte("playsinline")
	attrPopover              = []byte("popover")
	attrPopovertarget        = []byte("popovertarget")
	attrPopovertargetaction  = []byte("popovertargetaction")
	attrPoster               = []byte("poster")
	attrPreload              = []byte("preload")
	attrReadonly             = []byte("readonly")
	attrReferrerpolicy       = []byte("referrerpolicy")
	attrRel                  = []byte("rel")
	attrRequired             = []byte("required")
	attrRole                 = []byte("role")
	attrRows                 = []byte("rows")
	attrRowspan              = []byte("rowspan")
	attrSelected             = []byte("selected")
	attrSlot                 = []byte("slot")
	attrSrc                  = []byte("src")
	attrSrcset               = []byte("srcset")
	attrStep                 = []byte("step")
	attrStyle                = []byte("style")
	attrTabindex             = []byte("tabindex")
	attrTarget               = []byte("target")
	attrTitle                = []byte("title")
	attrType                 = []byte("type")
	attrValue                = []byte("value")
	attrWidth                = []byte("width")
	attrAllowfullscreen      = []byte("allowfullscreen")
	attrDefault              = []byte("default")
	attrInert                = []byte("inert")
	attrIsmap                = []byte("ismap")
	attrItemscope            = []byte("itemscope")
	attrNomodule             = []byte("nomodule")
	attrOpen                 = []byte("open")
	attrReversed             = []byte("reversed")
	attrSpellcheck           = []byte("spellcheck")
	attrTranslate            = []byte("translate")
	attrAriaAtomic           = []byte("aria-atomic")
	attrAriaBusy             = []byte("aria-busy")
	attrAriaChecked          = []byte("aria-checked")
	attrAriaControls         = []byte("aria-controls")
	attrAriaDescribedby      = []byte("aria-describedby")
	attrAriaDisabled         = []byte("aria-disabled")
	attrAriaExpanded         = []byte("aria-expanded")
	attrAriaFlowto           = []byte("aria-flowto")
	attrAriaHidden           = []byte("aria-hidden")
	attrAriaInvalid          = []byte("aria-invalid")
	attrAriaLabel            = []byte("aria-label")
	attrAriaLabelledby       = []byte("aria-labelledby")
	attrAriaLive             = []byte("aria-live")
	attrAriaOwns             = []byte("aria-owns")
	attrAriaPlaceholder      = []byte("aria-placeholder")
	attrAriaPosinset         = []byte("aria-posinset")
	attrAriaPressed          = []byte("aria-pressed")
	attrAriaReadonly         = []byte("aria-readonly")
	attrAriaRelevant         = []byte("aria-relevant")
	attrAriaRequired         = []byte("aria-required")
	attrAriaSelected         = []byte("aria-selected")
	attrAriaSetsize          = []byte("aria-setsize")
	attrAriaValuemax         = []byte("aria-valuemax")
	attrAriaValuemin         = []byte("aria-valuemin")
	attrAriaValuenow         = []byte("aria-valuenow")
	attrAriaValuetext        = []byte("aria-valuetext")
	attrOnAbort              = []byte("onabort")
	attrOnBeforeUnload       = []byte("onbeforeunload")
	attrOnBlur               = []byte("onblur")
	attrOnCanPlay            = []byte("oncanplay")
	attrOnCanPlayThrough     = []byte("oncanplaythrough")
	attrOnChange             = []byte("onchange")
	attrOnClick              = []byte("onclick")
	attrOnContextMenu        = []byte("oncontextmenu")
	attrOnCopy               = []byte("oncopy")
	attrOnCut                = []byte("oncut")
	attrOnDblClick           = []byte("ondblclick")
	attrOnDrag               = []byte("ondrag")
	attrOnDragEnd            = []byte("ondragend")
	attrOnDragEnter          = []byte("ondragenter")
	attrOnDragLeave          = []byte("ondragleave")
	attrOnDragOver           = []byte("ondragover")
	attrOnDragStart          = []byte("ondragstart")
	attrOnDrop               = []byte("ondrop")
	attrOnDurationChange     = []byte("ondurationchange")
	attrOnEnded              = []byte("onended")
	attrOnError              = []byte("onerror")
	attrOnFocus              = []byte("onfocus")
	attrOnHashChange         = []byte("onhashchange")
	attrOnInput              = []byte("oninput")
	attrOnInvalid            = []byte("oninvalid")
	attrOnKeyDown            = []byte("onkeydown")
	attrOnKeyPress           = []byte("onkeypress")
	attrOnKeyUp              = []byte("onkeyup")
	attrOnLoad               = []byte("onload")
	attrOnLoadedData         = []byte("onloadeddata")
	attrOnLoadedMetadata     = []byte("onloadedmetadata")
	attrOnMouseDown          = []byte("onmousedown")
	attrOnMouseEnter         = []byte("onmouseenter")
	attrOnMouseLeave         = []byte("onmouseleave")
	attrOnMouseMove          = []byte("onmousemove")
	attrOnMouseOut           = []byte("onmouseout")
	attrOnMouseOver          = []byte("onmouseover")
	attrOnMouseUp            = []byte("onmouseup")
	attrOnPageHide           = []byte("onpagehide")
	attrOnPageShow           = []byte("onpageshow")
	attrOnPaste              = []byte("onpaste")
	attrOnPause              = []byte("onpause")
	attrOnPlay               = []byte("onplay")
	attrOnPlaying            = []byte("onplaying")
	attrOnPopState           = []byte("onpopstate")
	attrOnProgress           = []byte("onprogress")
	attrOnRateChange         = []byte("onratechange")
	attrOnReset              = []byte("onreset")
	attrOnResize             = []byte("onresize")
	attrOnScroll             = []byte("onscroll")
	attrOnSeeked             = []byte("onseeked")
	attrOnSeeking            = []byte("onseeking")
	attrOnSelect             = []byte("onselect")
	attrOnStalled            = []byte("onstalled")
	attrOnSubmit             = []byte("onsubmit")
	attrOnSuspend            = []byte("onsuspend")
	attrOnTimeUpdate         = []byte("ontimeupdate")
	attrOnToggle             = []byte("ontoggle")
	attrOnUnload             = []byte("onunload")
	attrOnVolumeChange       = []byte("onvolumechange")
	attrOnWaiting            = []byte("onwaiting")
	attrOnPointerDown        = []byte("onpointerdown")
	attrOnPointerUp          = []byte("onpointerup")
	attrOnPointerMove        = []byte("onpointermove")
	attrOnPointerEnter       = []byte("onpointerenter")
	attrOnPointerLeave       = []byte("onpointerleave")
	attrOnPointerOver        = []byte("onpointerover")
	attrOnPointerOut         = []byte("onpointerout")
	attrOnPointerCancel      = []byte("onpointercancel")
	attrOnTouchStart         = []byte("ontouchstart")
	attrOnTouchMove          = []byte("ontouchmove")
	attrOnTouchEnd           = []byte("ontouchend")
	attrOnTouchCancel        = []byte("ontouchcancel")
	attrOnAnimationStart     = []byte("onanimationstart")
	attrOnAnimationEnd       = []byte("onanimationend")
	attrOnAnimationIteration = []byte("onanimationiteration")
	attrOnTransitionEnd      = []byte("ontransitionend")
	attrOnBeforeInput        = []byte("onbeforeinput")
	attrOnSlotChange         = []byte("onslotchange")
	attrOnFormData           = []byte("onformdata")
	attrAccessKey            = []byte("accesskey")
	attrAutoCapitalize       = []byte("autocapitalize")
	attrAutoCorrect          = []byte("autocorrect")
	attrEnterKeyHint         = []byte("enterkeyhint")
	attrExportParts          = []byte("exportparts")
	attrInputMode            = []byte("inputmode")
	attrIs                   = []byte("is")
	attrItemId               = []byte("itemid")
	attrItemProp             = []byte("itemprop")
	attrItemRef              = []byte("itemref")
	attrItemType             = []byte("itemtype")
	attrNonce                = []byte("nonce")
	attrPart                 = []byte("part")
	attrWritingSuggestions   = []byte("writingsuggestions")
)
