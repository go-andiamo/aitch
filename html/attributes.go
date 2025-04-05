package html

import "github.com/go-andiamo/aitch"

// Async declares a HTML "async" boolean attribute
func Async() aitch.Node {
	return aitch.NewBooleanAttribute(attrAsync)
}

// AutoFocus declares a HTML "autofocus" boolean attribute
func AutoFocus() aitch.Node {
	return aitch.NewBooleanAttribute(attrAutofocus)
}

// AutoPlay declares a HTML "autoplay" boolean attribute
func AutoPlay() aitch.Node {
	return aitch.NewBooleanAttribute(attrAutoplay)
}

// Checked declares a HTML "checked" boolean attribute
func Checked() aitch.Node {
	return aitch.NewBooleanAttribute(attrChecked)
}

// Controls declares a HTML "controls" boolean attribute
func Controls() aitch.Node {
	return aitch.NewBooleanAttribute(attrControls)
}

// CrossOrigin declares a HTML "crossorigin" attribute
func CrossOrigin(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCrossorigin, value...)
}

// DateTime declares a HTML "datetime" attribute
func DateTime(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDatetime, value...)
}

// Defer declares a HTML "defer" boolean attribute
func Defer() aitch.Node {
	return aitch.NewBooleanAttribute(attrDefer)
}

// Disabled declares a HTML "disabled" boolean attribute
func Disabled() aitch.Node {
	return aitch.NewBooleanAttribute(attrDisabled)
}

// Download declares a HTML "download" attribute
func Download(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDownload, value...)
}

// Draggable declares a HTML "draggable" attribute
func Draggable(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDraggable, value...)
}

// Loop declares a HTML "loop" boolean attribute
func Loop() aitch.Node {
	return aitch.NewBooleanAttribute(attrLoop)
}

// Multiple declares a HTML "multiple" boolean attribute
func Multiple() aitch.Node {
	return aitch.NewBooleanAttribute(attrMultiple)
}

// Muted declares a HTML "muted" boolean attribute
func Muted() aitch.Node {
	return aitch.NewBooleanAttribute(attrMuted)
}

// PlaysInline declares a HTML "playsinline" boolean attribute
func PlaysInline() aitch.Node {
	return aitch.NewBooleanAttribute(attrPlaysinline)
}

// ReadOnly declares a HTML "readonly" boolean attribute
func ReadOnly() aitch.Node {
	return aitch.NewBooleanAttribute(attrReadonly)
}

// Required declares a HTML "required" boolean attribute
func Required() aitch.Node {
	return aitch.NewBooleanAttribute(attrRequired)
}

// Selected declares a HTML "selected" boolean attribute
func Selected() aitch.Node {
	return aitch.NewBooleanAttribute(attrSelected)
}

// Accept declares a HTML "accept" attribute
func Accept(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAccept, value...)
}

// Action declares a HTML "action" attribute
func Action(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAction, value...)
}

// Alt declares a HTML "alt" attribute
func Alt(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAlt, value...)
}

// Aria declares a custom HTML "aria-*" attribute using the provided name and value
func Aria(name string, value ...any) aitch.Node {
	return aitch.NewAttribute(append(attrAria, []byte(name)...), value...)
}

// As declares a HTML "as" attribute
func As(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAs, value...)
}

// AutoComplete declares a HTML "autocomplete" attribute
func AutoComplete(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAutocomplete, value...)
}

// Charset declares a HTML "charset" attribute
func Charset(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCharset, value...)
}

// CiteAttr declares a HTML "cite" attribute
func CiteAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCite, value...)
}

// Class declares a HTML "class" attribute
func Class(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrClass, space, value...)
}

// Cols declares a HTML "cols" attribute
func Cols(value ...any) aitch.Node {
	return aitch.NewAttribute(attrCols, value...)
}

// ColSpan declares a HTML "colspan" attribute
func ColSpan(value ...any) aitch.Node {
	return aitch.NewAttribute(attrColspan, value...)
}

// Content declares a HTML "content" attribute
func Content(value ...any) aitch.Node {
	return aitch.NewAttribute(attrContent, value...)
}

// ContentEditable declares a HTML "contenteditable" attribute
func ContentEditable(value ...any) aitch.Node {
	return aitch.NewAttribute(attrContenteditable, value...)
}

// Dir declares a HTML "dir" attribute
func Dir(value ...any) aitch.Node {
	return aitch.NewAttribute(attrDir, value...)
}

// Data declares a HTML data attribute prefixed with "data-".
func Data(name string, value ...any) aitch.Node {
	return aitch.NewAttribute(append(attrData, []byte(name)...), value...)
}

// EncType declares a HTML "enctype" attribute
func EncType(value ...any) aitch.Node {
	return aitch.NewAttribute(attrEnctype, value...)
}

// For declares a HTML "for" attribute
func For(value ...any) aitch.Node {
	return aitch.NewAttribute(attrFor, value...)
}

// FormAttr declares a HTML "form" attribute
func FormAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrForm, value...)
}

// NoValidate declares a HTML "novalidate" boolean attribute
func NoValidate() aitch.Node {
	return aitch.NewBooleanAttribute(attrNovalidate)
}

// FormNoValidate declares a HTML "formnovalidate" boolean attribute
func FormNoValidate() aitch.Node {
	return aitch.NewBooleanAttribute(attrFormnovalidate)
}

// Height declares a HTML "height" attribute
func Height(value ...any) aitch.Node {
	return aitch.NewAttribute(attrHeight, value...)
}

// Hidden declares a HTML "hidden" boolean attribute
func Hidden() aitch.Node {
	return aitch.NewBooleanAttribute(attrHidden)
}

// Href declares a HTML "href" attribute
func Href(value ...any) aitch.Node {
	return aitch.NewAttribute(attrHref, value...)
}

// Id declares a HTML "id" attribute
func Id(value ...any) aitch.Node {
	return aitch.NewAttribute(attrId, value...)
}

// Integrity declares a HTML "integrity" attribute
func Integrity(value ...any) aitch.Node {
	return aitch.NewAttribute(attrIntegrity, value...)
}

// LabelAttr declares a HTML "label" attribute
func LabelAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrLabel, value...)
}

// Lang declares a HTML "lang" attribute
func Lang(value ...any) aitch.Node {
	return aitch.NewAttribute(attrLang, value...)
}

// List declares a HTML "list" attribute
func List(value ...any) aitch.Node {
	return aitch.NewAttribute(attrList, value...)
}

// Loading declares a HTML "loading" attribute
func Loading(value ...any) aitch.Node {
	return aitch.NewAttribute(attrLoading, value...)
}

// Max declares a HTML "max" attribute
func Max(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMax, value...)
}

// MaxLength declares a HTML "maxlength" attribute
func MaxLength(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMaxlength, value...)
}

// Method declares a HTML "method" attribute
func Method(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMethod, value...)
}

// Min declares a HTML "min" attribute
func Min(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMin, value...)
}

// MinLength declares a HTML "minlength" attribute
func MinLength(value ...any) aitch.Node {
	return aitch.NewAttribute(attrMinlength, value...)
}

// Name declares a HTML "name" attribute
func Name(value ...any) aitch.Node {
	return aitch.NewAttribute(attrName, value...)
}

// Pattern declares a HTML "pattern" attribute
func Pattern(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPattern, value...)
}

// Placeholder declares a HTML "placeholder" attribute
func Placeholder(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPlaceholder, value...)
}

// PopOver declares a HTML "popover" attribute
func PopOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPopover, value...)
}

// PopOverTarget declares a HTML "popovertarget" attribute
func PopOverTarget(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPopovertarget, value...)
}

// PopOverTargetAction declares a HTML "popovertargetaction" attribute
func PopOverTargetAction(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPopovertargetaction, value...)
}

// Poster declares a HTML "poster" attribute
func Poster(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPoster, value...)
}

// Preload declares a HTML "preload" attribute
func Preload(value ...any) aitch.Node {
	return aitch.NewAttribute(attrPreload, value...)
}

// ReferrerPolicy declares a HTML "referrerpolicy" attribute
func ReferrerPolicy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrReferrerpolicy, value...)
}

// Rel declares a HTML "rel" attribute
func Rel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRel, value...)
}

// Role declares a HTML "role" attribute
func Role(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRole, value...)
}

// Rows declares a HTML "rows" attribute
func Rows(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRows, value...)
}

// RowSpan declares a HTML "rowspan" attribute
func RowSpan(value ...any) aitch.Node {
	return aitch.NewAttribute(attrRowspan, value...)
}

// SlotAttr declares a HTML "slot" attribute
func SlotAttr(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSlot, value...)
}

// Src declares a HTML "src" attribute
func Src(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSrc, value...)
}

// SrcSet declares a HTML "srcset" attribute
func SrcSet(value ...any) aitch.Node {
	return aitch.NewAttribute(attrSrcset, value...)
}

// Step declares a HTML "step" attribute
func Step(value ...any) aitch.Node {
	return aitch.NewAttribute(attrStep, value...)
}

// Style declares a HTML "style" attribute
func Style(value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrStyle, stylesDelim, value...)
}

// TabIndex declares a HTML "tabindex" attribute
func TabIndex(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTabindex, value...)
}

// Target declares a HTML "target" attribute
func Target(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTarget, value...)
}

// Title declares a HTML "title" attribute
func Title(value ...any) aitch.Node {
	return aitch.NewAttribute(attrTitle, value...)
}

// Type declares a HTML "type" attribute
func Type(value ...any) aitch.Node {
	return aitch.NewAttribute(attrType, value...)
}

// Value declares a HTML "value" attribute
func Value(value ...any) aitch.Node {
	return aitch.NewAttribute(attrValue, value...)
}

// Width declares a HTML "width" attribute
func Width(value ...any) aitch.Node {
	return aitch.NewAttribute(attrWidth, value...)
}

// AllowFullScreen declares a HTML "allowfullscreen" boolean attribute
func AllowFullScreen() aitch.Node {
	return aitch.NewBooleanAttribute(attrAllowfullscreen)
}

// Default declares a HTML "default" boolean attribute
func Default() aitch.Node {
	return aitch.NewBooleanAttribute(attrDefault)
}

// Inert declares a HTML "inert" boolean attribute
func Inert() aitch.Node {
	return aitch.NewBooleanAttribute(attrInert)
}

// IsMap declares a HTML "ismap" boolean attribute
func IsMap() aitch.Node {
	return aitch.NewBooleanAttribute(attrIsmap)
}

// ItemScope declares a HTML "itemscope" boolean attribute
func ItemScope() aitch.Node {
	return aitch.NewBooleanAttribute(attrItemscope)
}

// NoModule declares a HTML "nomodule" boolean attribute
func NoModule() aitch.Node {
	return aitch.NewBooleanAttribute(attrNomodule)
}

// Open declares a HTML "open" boolean attribute
func Open() aitch.Node {
	return aitch.NewBooleanAttribute(attrOpen)
}

// Reversed declares a HTML "reversed" boolean attribute
func Reversed() aitch.Node {
	return aitch.NewBooleanAttribute(attrReversed)
}

// Spellcheck declares a HTML "spellcheck" boolean attribute
func Spellcheck() aitch.Node {
	return aitch.NewBooleanAttribute(attrSpellcheck)
}

// Translate declares a HTML "translate" boolean attribute
func Translate() aitch.Node {
	return aitch.NewBooleanAttribute(attrTranslate)
}

// AriaAtomic declares a HTML "aria-atomic" attribute
func AriaAtomic(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaAtomic, value...)
}

// AriaBusy declares a HTML "aria-busy" attribute
func AriaBusy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaBusy, value...)
}

// AriaChecked declares a HTML "aria-checked" attribute
func AriaChecked(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaChecked, value...)
}

// AriaControls declares a HTML "aria-controls" attribute
func AriaControls(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaControls, value...)
}

// AriaDescribedBy declares a HTML "aria-describedby" attribute
func AriaDescribedBy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaDescribedby, value...)
}

// AriaDisabled declares a HTML "aria-disabled" attribute
func AriaDisabled(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaDisabled, value...)
}

// AriaExpanded declares a HTML "aria-expanded" attribute
func AriaExpanded(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaExpanded, value...)
}

// AriaFlowTo declares a HTML "aria-flowto" attribute
func AriaFlowTo(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaFlowto, value...)
}

// AriaHidden declares a HTML "aria-hidden" attribute
func AriaHidden(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaHidden, value...)
}

// AriaInvalid declares a HTML "aria-invalid" attribute
func AriaInvalid(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaInvalid, value...)
}

// AriaLabel declares a HTML "aria-label" attribute
func AriaLabel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLabel, value...)
}

// AriaLabelledBy declares a HTML "aria-labelledby" attribute
func AriaLabelledBy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLabelledby, value...)
}

// AriaLive declares a HTML "aria-live" attribute
func AriaLive(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLive, value...)
}

// AriaOwns declares a HTML "aria-owns" attribute
func AriaOwns(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaOwns, value...)
}

// AriaPlaceholder declares a HTML "aria-placeholder" attribute
func AriaPlaceholder(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPlaceholder, value...)
}

// AriaPosInSet declares a HTML "aria-posinset" attribute
func AriaPosInSet(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPosinset, value...)
}

// AriaPressed declares a HTML "aria-pressed" attribute
func AriaPressed(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPressed, value...)
}

// AriaReadonly declares a HTML "aria-readonly" attribute
func AriaReadonly(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaReadonly, value...)
}

// AriaRelevant declares a HTML "aria-relevant" attribute
func AriaRelevant(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaRelevant, value...)
}

// AriaRequired declares a HTML "aria-required" attribute
func AriaRequired(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaRequired, value...)
}

// AriaSelected declares a HTML "aria-selected" attribute
func AriaSelected(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaSelected, value...)
}

// AriaSetSize declares a HTML "aria-setsize" attribute
func AriaSetSize(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaSetsize, value...)
}

// AriaValueMax declares a HTML "aria-valuemax" attribute
func AriaValueMax(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuemax, value...)
}

// AriaValueMin declares a HTML "aria-valuemin" attribute
func AriaValueMin(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuemin, value...)
}

// AriaValueNow declares a HTML "aria-valuenow" attribute
func AriaValueNow(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuenow, value...)
}

// AriaValueText declares a HTML "aria-valuetext" attribute
func AriaValueText(value ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuetext, value...)
}

// OnAbort declares a HTML "onabort" event attribute
func OnAbort(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAbort, value...)
}

// OnBeforeUnload declares a HTML "onbeforeunload" event attribute
func OnBeforeUnload(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeUnload, value...)
}

// OnBlur declares a HTML "onblur" event attribute
func OnBlur(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBlur, value...)
}

// OnCanPlay declares a HTML "oncanplay" event attribute
func OnCanPlay(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCanPlay, value...)
}

// OnCanPlayThrough declares a HTML "oncanplaythrough" event attribute
func OnCanPlayThrough(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCanPlayThrough, value...)
}

// OnChange declares a HTML "onchange" event attribute
func OnChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnChange, value...)
}

// OnClick declares a HTML "onclick" event attribute
func OnClick(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnClick, value...)
}

// OnContextMenu declares a HTML "oncontextmenu" event attribute
func OnContextMenu(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnContextMenu, value...)
}

// OnCopy declares a HTML "oncopy" event attribute
func OnCopy(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCopy, value...)
}

// OnCut declares a HTML "oncut" event attribute
func OnCut(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCut, value...)
}

// OnDblClick declares a HTML "ondblclick" event attribute
func OnDblClick(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDblClick, value...)
}

// OnDrag declares a HTML "ondrag" event attribute
func OnDrag(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDrag, value...)
}

// OnDragEnd declares a HTML "ondragend" event attribute
func OnDragEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragEnd, value...)
}

// OnDragEnter declares a HTML "ondragenter" event attribute
func OnDragEnter(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragEnter, value...)
}

// OnDragLeave declares a HTML "ondragleave" event attribute
func OnDragLeave(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragLeave, value...)
}

// OnDragOver declares a HTML "ondragover" event attribute
func OnDragOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragOver, value...)
}

// OnDragStart declares a HTML "ondragstart" event attribute
func OnDragStart(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragStart, value...)
}

// OnDrop declares a HTML "ondrop" event attribute
func OnDrop(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDrop, value...)
}

// OnDurationChange declares a HTML "ondurationchange" event attribute
func OnDurationChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDurationChange, value...)
}

// OnEnded declares a HTML "onended" event attribute
func OnEnded(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnEnded, value...)
}

// OnError declares a HTML "onerror" event attribute
func OnError(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnError, value...)
}

// OnFocus declares a HTML "onfocus" event attribute
func OnFocus(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnFocus, value...)
}

// OnHashChange declares a HTML "onhashchange" event attribute
func OnHashChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnHashChange, value...)
}

// OnInput declares a HTML "oninput" event attribute
func OnInput(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnInput, value...)
}

// OnInvalid declares a HTML "oninvalid" event attribute
func OnInvalid(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnInvalid, value...)
}

// OnKeyDown declares a HTML "onkeydown" event attribute
func OnKeyDown(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyDown, value...)
}

// OnKeyPress declares a HTML "onkeypress" event attribute
func OnKeyPress(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyPress, value...)
}

// OnKeyUp declares a HTML "onkeyup" event attribute
func OnKeyUp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyUp, value...)
}

// OnLoad declares a HTML "onload" event attribute
func OnLoad(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoad, value...)
}

// OnLoadedData declares a HTML "onloadeddata" event attribute
func OnLoadedData(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoadedData, value...)
}

// OnLoadedMetadata declares a HTML "onloadedmetadata" event attribute
func OnLoadedMetadata(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoadedMetadata, value...)
}

// OnMouseDown declares a HTML "onmousedown" event attribute
func OnMouseDown(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseDown, value...)
}

// OnMouseEnter declares a HTML "onmouseenter" event attribute
func OnMouseEnter(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseEnter, value...)
}

// OnMouseLeave declares a HTML "onmouseleave" event attribute
func OnMouseLeave(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseLeave, value...)
}

// OnMouseMove declares a HTML "onmousemove" event attribute
func OnMouseMove(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseMove, value...)
}

// OnMouseOut declares a HTML "onmouseout" event attribute
func OnMouseOut(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseOut, value...)
}

// OnMouseOver declares a HTML "onmouseover" event attribute
func OnMouseOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseOver, value...)
}

// OnMouseUp declares a HTML "onmouseup" event attribute
func OnMouseUp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseUp, value...)
}

// OnPageHide declares a HTML "onpagehide" event attribute
func OnPageHide(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPageHide, value...)
}

// OnPageShow declares a HTML "onpageshow" event attribute
func OnPageShow(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPageShow, value...)
}

// OnPaste declares a HTML "onpaste" event attribute
func OnPaste(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPaste, value...)
}

// OnPause declares a HTML "onpause" event attribute
func OnPause(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPause, value...)
}

// OnPlay declares a HTML "onplay" event attribute
func OnPlay(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPlay, value...)
}

// OnPlaying declares a HTML "onplaying" event attribute
func OnPlaying(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPlaying, value...)
}

// OnPopState declares a HTML "onpopstate" event attribute
func OnPopState(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPopState, value...)
}

// OnProgress declares a HTML "onprogress" event attribute
func OnProgress(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnProgress, value...)
}

// OnRateChange declares a HTML "onratechange" event attribute
func OnRateChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnRateChange, value...)
}

// OnReset declares a HTML "onreset" event attribute
func OnReset(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnReset, value...)
}

// OnResize declares a HTML "onresize" event attribute
func OnResize(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnResize, value...)
}

// OnScroll declares a HTML "onscroll" event attribute
func OnScroll(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnScroll, value...)
}

// OnSeeked declares a HTML "onseeked" event attribute
func OnSeeked(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSeeked, value...)
}

// OnSeeking declares a HTML "onseeking" event attribute
func OnSeeking(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSeeking, value...)
}

// OnSelect declares a HTML "onselect" event attribute
func OnSelect(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSelect, value...)
}

// OnStalled declares a HTML "onstalled" event attribute
func OnStalled(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnStalled, value...)
}

// OnSubmit declares a HTML "onsubmit" event attribute
func OnSubmit(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSubmit, value...)
}

// OnSuspend declares a HTML "onsuspend" event attribute
func OnSuspend(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSuspend, value...)
}

// OnTimeUpdate declares a HTML "ontimeupdate" event attribute
func OnTimeUpdate(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTimeUpdate, value...)
}

// OnToggle declares a HTML "ontoggle" event attribute
func OnToggle(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnToggle, value...)
}

// OnUnload declares a HTML "onunload" event attribute
func OnUnload(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnUnload, value...)
}

// OnVolumeChange declares a HTML "onvolumechange" event attribute
func OnVolumeChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnVolumeChange, value...)
}

// OnWaiting declares a HTML "onwaiting" event attribute
func OnWaiting(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnWaiting, value...)
}

// OnPointerDown declares a HTML "onpointerdown" event attribute
func OnPointerDown(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerDown, value...)
}

// OnPointerUp declares a HTML "onpointerup" event attribute
func OnPointerUp(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerUp, value...)
}

// OnPointerMove declares a HTML "onpointermove" event attribute
func OnPointerMove(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerMove, value...)
}

// OnPointerEnter declares a HTML "onpointerenter" event attribute
func OnPointerEnter(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerEnter, value...)
}

// OnPointerLeave declares a HTML "onpointerleave" event attribute
func OnPointerLeave(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerLeave, value...)
}

// OnPointerOver declares a HTML "onpointerover" event attribute
func OnPointerOver(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerOver, value...)
}

// OnPointerOut declares a HTML "onpointerout" event attribute
func OnPointerOut(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerOut, value...)
}

// OnPointerCancel declares a HTML "onpointercancel" event attribute
func OnPointerCancel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPointerCancel, value...)
}

// OnTouchStart declares a HTML "ontouchstart" event attribute
func OnTouchStart(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchStart, value...)
}

// OnTouchMove declares a HTML "ontouchmove" event attribute
func OnTouchMove(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchMove, value...)
}

// OnTouchEnd declares a HTML "ontouchend" event attribute
func OnTouchEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchEnd, value...)
}

// OnTouchCancel declares a HTML "ontouchcancel" event attribute
func OnTouchCancel(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTouchCancel, value...)
}

// OnAnimationStart declares a HTML "onanimationstart" event attribute
func OnAnimationStart(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAnimationStart, value...)
}

// OnAnimationEnd declares a HTML "onanimationend" event attribute
func OnAnimationEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAnimationEnd, value...)
}

// OnAnimationIteration declares a HTML "onanimationiteration" event attribute
func OnAnimationIteration(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAnimationIteration, value...)
}

// OnTransitionEnd declares a HTML "ontransitionend" event attribute
func OnTransitionEnd(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTransitionEnd, value...)
}

// OnBeforeInput declares a HTML "onbeforeinput" event attribute
func OnBeforeInput(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeInput, value...)
}

// OnSlotChange declares a HTML "onslotchange" event attribute
func OnSlotChange(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSlotChange, value...)
}

// OnFormData declares a HTML "onformdata" event attribute
func OnFormData(value ...any) aitch.Node {
	return aitch.NewAttribute(attrOnFormData, value...)
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
	attrLabel                = []byte("label")
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
)
