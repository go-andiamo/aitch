package html

import "github.com/go-andiamo/aitch"

func Async() aitch.Node {
	return aitch.NewBooleanAttribute(attrAsync)
}

func AutoFocus() aitch.Node {
	return aitch.NewBooleanAttribute(attrAutofocus)
}

func AutoPlay() aitch.Node {
	return aitch.NewBooleanAttribute(attrAutoplay)
}

func Checked() aitch.Node {
	return aitch.NewBooleanAttribute(attrChecked)
}

func Controls() aitch.Node {
	return aitch.NewBooleanAttribute(attrControls)
}

func CrossOrigin(v ...any) aitch.Node {
	return aitch.NewAttribute(attrCrossorigin, v...)
}

func DateTime(v ...any) aitch.Node {
	return aitch.NewAttribute(attrDatetime, v...)
}

func Defer() aitch.Node {
	return aitch.NewBooleanAttribute(attrDefer)
}

func Disabled() aitch.Node {
	return aitch.NewBooleanAttribute(attrDisabled)
}

func Download(v ...any) aitch.Node {
	return aitch.NewAttribute(attrDownload, v...)
}

func Draggable(v ...any) aitch.Node {
	return aitch.NewAttribute(attrDraggable, v...)
}

func Loop() aitch.Node {
	return aitch.NewBooleanAttribute(attrLoop)
}

func Multiple() aitch.Node {
	return aitch.NewBooleanAttribute(attrMultiple)
}

func Muted() aitch.Node {
	return aitch.NewBooleanAttribute(attrMuted)
}

func PlaysInline() aitch.Node {
	return aitch.NewBooleanAttribute(attrPlaysinline)
}

func ReadOnly() aitch.Node {
	return aitch.NewBooleanAttribute(attrReadonly)
}

func Required() aitch.Node {
	return aitch.NewBooleanAttribute(attrRequired)
}

func Selected() aitch.Node {
	return aitch.NewBooleanAttribute(attrSelected)
}

func Accept(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAccept, v...)
}

func Action(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAction, v...)
}

func Alt(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAlt, v...)
}

func Aria(name string, v ...any) aitch.Node {
	return aitch.NewAttribute(append(attrAria, []byte(name)...), v...)
}

func As(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAs, v...)
}

func AutoComplete(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAutocomplete, v...)
}

func Charset(v ...any) aitch.Node {
	return aitch.NewAttribute(attrCharset, v...)
}

func CiteAttr(v ...any) aitch.Node {
	return aitch.NewAttribute(attrCite, v...)
}

func Class(v ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrClass, space, v...)
}

func Cols(v ...any) aitch.Node {
	return aitch.NewAttribute(attrCols, v...)
}

func ColSpan(v ...any) aitch.Node {
	return aitch.NewAttribute(attrColspan, v...)
}

func Content(v ...any) aitch.Node {
	return aitch.NewAttribute(attrContent, v...)
}

func ContentEditable(v ...any) aitch.Node {
	return aitch.NewAttribute(attrContenteditable, v...)
}

func Dir(v ...any) aitch.Node {
	return aitch.NewAttribute(attrDir, v...)
}

// Data attribute prefixed with "data-".
func Data(name string, v ...any) aitch.Node {
	return aitch.NewAttribute(append(attrData, []byte(name)...), v...)
}

func EncType(v ...any) aitch.Node {
	return aitch.NewAttribute(attrEnctype, v...)
}

func For(v ...any) aitch.Node {
	return aitch.NewAttribute(attrFor, v...)
}

func FormAttr(v ...any) aitch.Node {
	return aitch.NewAttribute(attrForm, v...)
}

func NoValidate() aitch.Node {
	return aitch.NewBooleanAttribute(attrNovalidate)
}

func FormNoValidate() aitch.Node {
	return aitch.NewBooleanAttribute(attrFormnovalidate)
}

func Height(v ...any) aitch.Node {
	return aitch.NewAttribute(attrHeight, v...)
}

func Hidden() aitch.Node {
	return aitch.NewBooleanAttribute(attrHidden)
}

func Href(v ...any) aitch.Node {
	return aitch.NewAttribute(attrHref, v...)
}

func Id(v ...any) aitch.Node {
	return aitch.NewAttribute(attrId, v...)
}

func Integrity(v ...any) aitch.Node {
	return aitch.NewAttribute(attrIntegrity, v...)
}

func LabelAttr(v ...any) aitch.Node {
	return aitch.NewAttribute(attrLabel, v...)
}

func Lang(v ...any) aitch.Node {
	return aitch.NewAttribute(attrLang, v...)
}

func List(v ...any) aitch.Node {
	return aitch.NewAttribute(attrList, v...)
}

func Loading(v ...any) aitch.Node {
	return aitch.NewAttribute(attrLoading, v...)
}

func Max(v ...any) aitch.Node {
	return aitch.NewAttribute(attrMax, v...)
}

func MaxLength(v ...any) aitch.Node {
	return aitch.NewAttribute(attrMaxlength, v...)
}

func Method(v ...any) aitch.Node {
	return aitch.NewAttribute(attrMethod, v...)
}

func Min(v ...any) aitch.Node {
	return aitch.NewAttribute(attrMin, v...)
}

func MinLength(v ...any) aitch.Node {
	return aitch.NewAttribute(attrMinlength, v...)
}

func Name(v ...any) aitch.Node {
	return aitch.NewAttribute(attrName, v...)
}

func Pattern(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPattern, v...)
}

func Placeholder(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPlaceholder, v...)
}

func PopOver(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPopover, v...)
}

func PopOverTarget(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPopovertarget, v...)
}

func PopOverTargetAction(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPopovertargetaction, v...)
}

func Poster(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPoster, v...)
}

func Preload(v ...any) aitch.Node {
	return aitch.NewAttribute(attrPreload, v...)
}

func ReferrerPolicy(v ...any) aitch.Node {
	return aitch.NewAttribute(attrReferrerpolicy, v...)
}

func Rel(v ...any) aitch.Node {
	return aitch.NewAttribute(attrRel, v...)
}

func Role(v ...any) aitch.Node {
	return aitch.NewAttribute(attrRole, v...)
}

func Rows(v ...any) aitch.Node {
	return aitch.NewAttribute(attrRows, v...)
}

func RowSpan(v ...any) aitch.Node {
	return aitch.NewAttribute(attrRowspan, v...)
}

func SlotAttr(v ...any) aitch.Node {
	return aitch.NewAttribute(attrSlot, v...)
}

func Src(v ...any) aitch.Node {
	return aitch.NewAttribute(attrSrc, v...)
}

func SrcSet(v ...any) aitch.Node {
	return aitch.NewAttribute(attrSrcset, v...)
}

func Step(v ...any) aitch.Node {
	return aitch.NewAttribute(attrStep, v...)
}

func Style(v ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrStyle, stylesDelim, v...)
}

func TabIndex(v ...any) aitch.Node {
	return aitch.NewAttribute(attrTabindex, v...)
}

func Target(v ...any) aitch.Node {
	return aitch.NewAttribute(attrTarget, v...)
}

func Title(v ...any) aitch.Node {
	return aitch.NewAttribute(attrTitle, v...)
}

func Type(v ...any) aitch.Node {
	return aitch.NewAttribute(attrType, v...)
}

func Value(v ...any) aitch.Node {
	return aitch.NewAttribute(attrValue, v...)
}

func Width(v ...any) aitch.Node {
	return aitch.NewAttribute(attrWidth, v...)
}

func AllowFullScreen() aitch.Node {
	return aitch.NewBooleanAttribute(attrAllowfullscreen)
}

func Default() aitch.Node {
	return aitch.NewBooleanAttribute(attrDefault)
}

func Inert() aitch.Node {
	return aitch.NewBooleanAttribute(attrInert)
}

func IsMap() aitch.Node {
	return aitch.NewBooleanAttribute(attrIsmap)
}

func ItemScope() aitch.Node {
	return aitch.NewBooleanAttribute(attrItemscope)
}

func NoModule() aitch.Node {
	return aitch.NewBooleanAttribute(attrNomodule)
}

func Open() aitch.Node {
	return aitch.NewBooleanAttribute(attrOpen)
}

func Reversed() aitch.Node {
	return aitch.NewBooleanAttribute(attrReversed)
}

func Spellcheck() aitch.Node {
	return aitch.NewBooleanAttribute(attrSpellcheck)
}

func Translate() aitch.Node {
	return aitch.NewBooleanAttribute(attrTranslate)
}

func AriaAtomic(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaAtomic, v...)
}

func AriaBusy(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaBusy, v...)
}

func AriaChecked(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaChecked, v...)
}

func AriaControls(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaControls, v...)
}

func AriaDescribedBy(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaDescribedby, v...)
}

func AriaDisabled(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaDisabled, v...)
}

func AriaExpanded(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaExpanded, v...)
}

func AriaFlowTo(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaFlowto, v...)
}

func AriaHidden(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaHidden, v...)
}

func AriaInvalid(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaInvalid, v...)
}

func AriaLabel(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLabel, v...)
}

func AriaLabelledBy(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLabelledby, v...)
}

func AriaLive(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaLive, v...)
}

func AriaOwns(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaOwns, v...)
}

func AriaPlaceholder(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPlaceholder, v...)
}

func AriaPosInSet(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPosinset, v...)
}

func AriaPressed(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaPressed, v...)
}

func AriaReadonly(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaReadonly, v...)
}

func AriaRelevant(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaRelevant, v...)
}

func AriaRequired(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaRequired, v...)
}

func AriaSelected(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaSelected, v...)
}

func AriaSetSize(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaSetsize, v...)
}

func AriaValueMax(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuemax, v...)
}

func AriaValueMin(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuemin, v...)
}

func AriaValueNow(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuenow, v...)
}

func AriaValueText(v ...any) aitch.Node {
	return aitch.NewAttribute(attrAriaValuetext, v...)
}

func OnAbort(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnAbort, v...)
}

func OnBeforeUnload(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBeforeUnload, v...)
}

func OnBlur(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnBlur, v...)
}

func OnCanPlay(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCanPlay, v...)
}

func OnCanPlayThrough(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCanPlayThrough, v...)
}

func OnChange(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnChange, v...)
}

func OnClick(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnClick, v...)
}

func OnContextMenu(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnContextMenu, v...)
}

func OnCopy(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCopy, v...)
}

func OnCut(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnCut, v...)
}

func OnDblClick(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDblClick, v...)
}

func OnDrag(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDrag, v...)
}

func OnDragEnd(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragEnd, v...)
}

func OnDragEnter(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragEnter, v...)
}

func OnDragLeave(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragLeave, v...)
}

func OnDragOver(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragOver, v...)
}

func OnDragStart(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDragStart, v...)
}

func OnDrop(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDrop, v...)
}

func OnDurationChange(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnDurationChange, v...)
}

func OnEnded(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnEnded, v...)
}

func OnError(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnError, v...)
}

func OnFocus(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnFocus, v...)
}

func OnHashChange(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnHashChange, v...)
}

func OnInput(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnInput, v...)
}

func OnInvalid(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnInvalid, v...)
}

func OnKeyDown(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyDown, v...)
}

func OnKeyPress(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyPress, v...)
}

func OnKeyUp(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnKeyUp, v...)
}

func OnLoad(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoad, v...)
}

func OnLoadedData(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoadedData, v...)
}

func OnLoadedMetadata(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnLoadedMetadata, v...)
}

func OnMouseDown(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseDown, v...)
}

func OnMouseEnter(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseEnter, v...)
}

func OnMouseLeave(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseLeave, v...)
}

func OnMouseMove(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseMove, v...)
}

func OnMouseOut(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseOut, v...)
}

func OnMouseOver(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseOver, v...)
}

func OnMouseUp(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnMouseUp, v...)
}

func OnPageHide(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPageHide, v...)
}

func OnPageShow(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPageShow, v...)
}

func OnPaste(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPaste, v...)
}

func OnPause(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPause, v...)
}

func OnPlay(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPlay, v...)
}

func OnPlaying(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPlaying, v...)
}

func OnPopState(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnPopState, v...)
}

func OnProgress(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnProgress, v...)
}

func OnRateChange(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnRateChange, v...)
}

func OnReset(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnReset, v...)
}

func OnResize(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnResize, v...)
}

func OnScroll(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnScroll, v...)
}

func OnSeeked(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSeeked, v...)
}

func OnSeeking(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSeeking, v...)
}

func OnSelect(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSelect, v...)
}

func OnStalled(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnStalled, v...)
}

func OnSubmit(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSubmit, v...)
}

func OnSuspend(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnSuspend, v...)
}

func OnTimeUpdate(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnTimeUpdate, v...)
}

func OnToggle(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnToggle, v...)
}

func OnUnload(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnUnload, v...)
}

func OnVolumeChange(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnVolumeChange, v...)
}

func OnWaiting(v ...any) aitch.Node {
	return aitch.NewAttribute(attrOnWaiting, v...)
}

var (
	space                   = []byte{' '}
	stylesDelim             = []byte{';', ' '}
	attrAccept              = []byte("accept")
	attrAction              = []byte("action")
	attrAlt                 = []byte("alt")
	attrAria                = []byte("aria-")
	attrAs                  = []byte("as")
	attrAsync               = []byte("async")
	attrAutocomplete        = []byte("autocomplete")
	attrAutofocus           = []byte("autofocus")
	attrAutoplay            = []byte("autoplay")
	attrCharset             = []byte("charset")
	attrChecked             = []byte("checked")
	attrCite                = []byte("cite")
	attrClass               = []byte("class")
	attrCols                = []byte("cols")
	attrColspan             = []byte("colspan")
	attrContent             = []byte("content")
	attrContenteditable     = []byte("contenteditable")
	attrControls            = []byte("controls")
	attrCrossorigin         = []byte("crossorigin")
	attrData                = []byte("data-")
	attrDatetime            = []byte("datetime")
	attrDefer               = []byte("defer")
	attrDir                 = []byte("dir")
	attrDisabled            = []byte("disabled")
	attrDownload            = []byte("download")
	attrDraggable           = []byte("draggable")
	attrEnctype             = []byte("enctype")
	attrFor                 = []byte("for")
	attrForm                = []byte("form")
	attrHeight              = []byte("height")
	attrHidden              = []byte("hidden")
	attrHref                = []byte("href")
	attrId                  = []byte("id")
	attrIntegrity           = []byte("integrity")
	attrLabel               = []byte("label")
	attrLang                = []byte("lang")
	attrList                = []byte("list")
	attrLoading             = []byte("loading")
	attrLoop                = []byte("loop")
	attrMax                 = []byte("max")
	attrMaxlength           = []byte("maxlength")
	attrMethod              = []byte("method")
	attrMin                 = []byte("min")
	attrMinlength           = []byte("minlength")
	attrMultiple            = []byte("multiple")
	attrMuted               = []byte("muted")
	attrName                = []byte("name")
	attrNovalidate          = []byte("novalidate")
	attrFormnovalidate      = []byte("formnovalidate")
	attrPattern             = []byte("pattern")
	attrPlaceholder         = []byte("placeholder")
	attrPlaysinline         = []byte("playsinline")
	attrPopover             = []byte("popover")
	attrPopovertarget       = []byte("popovertarget")
	attrPopovertargetaction = []byte("popovertargetaction")
	attrPoster              = []byte("poster")
	attrPreload             = []byte("preload")
	attrReadonly            = []byte("readonly")
	attrReferrerpolicy      = []byte("referrerpolicy")
	attrRel                 = []byte("rel")
	attrRequired            = []byte("required")
	attrRole                = []byte("role")
	attrRows                = []byte("rows")
	attrRowspan             = []byte("rowspan")
	attrSelected            = []byte("selected")
	attrSlot                = []byte("slot")
	attrSrc                 = []byte("src")
	attrSrcset              = []byte("srcset")
	attrStep                = []byte("step")
	attrStyle               = []byte("style")
	attrTabindex            = []byte("tabindex")
	attrTarget              = []byte("target")
	attrTitle               = []byte("title")
	attrType                = []byte("type")
	attrValue               = []byte("value")
	attrWidth               = []byte("width")
	attrAllowfullscreen     = []byte("allowfullscreen")
	attrDefault             = []byte("default")
	attrInert               = []byte("inert")
	attrIsmap               = []byte("ismap")
	attrItemscope           = []byte("itemscope")
	attrNomodule            = []byte("nomodule")
	attrOpen                = []byte("open")
	attrReversed            = []byte("reversed")
	attrSpellcheck          = []byte("spellcheck")
	attrTranslate           = []byte("translate")
	attrAriaAtomic          = []byte("aria-atomic")
	attrAriaBusy            = []byte("aria-busy")
	attrAriaChecked         = []byte("aria-checked")
	attrAriaControls        = []byte("aria-controls")
	attrAriaDescribedby     = []byte("aria-describedby")
	attrAriaDisabled        = []byte("aria-disabled")
	attrAriaExpanded        = []byte("aria-expanded")
	attrAriaFlowto          = []byte("aria-flowto")
	attrAriaHidden          = []byte("aria-hidden")
	attrAriaInvalid         = []byte("aria-invalid")
	attrAriaLabel           = []byte("aria-label")
	attrAriaLabelledby      = []byte("aria-labelledby")
	attrAriaLive            = []byte("aria-live")
	attrAriaOwns            = []byte("aria-owns")
	attrAriaPlaceholder     = []byte("aria-placeholder")
	attrAriaPosinset        = []byte("aria-posinset")
	attrAriaPressed         = []byte("aria-pressed")
	attrAriaReadonly        = []byte("aria-readonly")
	attrAriaRelevant        = []byte("aria-relevant")
	attrAriaRequired        = []byte("aria-required")
	attrAriaSelected        = []byte("aria-selected")
	attrAriaSetsize         = []byte("aria-setsize")
	attrAriaValuemax        = []byte("aria-valuemax")
	attrAriaValuemin        = []byte("aria-valuemin")
	attrAriaValuenow        = []byte("aria-valuenow")
	attrAriaValuetext       = []byte("aria-valuetext")
	attrOnAbort             = []byte("onabort")
	attrOnBeforeUnload      = []byte("onbeforeunload")
	attrOnBlur              = []byte("onblur")
	attrOnCanPlay           = []byte("oncanplay")
	attrOnCanPlayThrough    = []byte("oncanplaythrough")
	attrOnChange            = []byte("onchange")
	attrOnClick             = []byte("onclick")
	attrOnContextMenu       = []byte("oncontextmenu")
	attrOnCopy              = []byte("oncopy")
	attrOnCut               = []byte("oncut")
	attrOnDblClick          = []byte("ondblclick")
	attrOnDrag              = []byte("ondrag")
	attrOnDragEnd           = []byte("ondragend")
	attrOnDragEnter         = []byte("ondragenter")
	attrOnDragLeave         = []byte("ondragleave")
	attrOnDragOver          = []byte("ondragover")
	attrOnDragStart         = []byte("ondragstart")
	attrOnDrop              = []byte("ondrop")
	attrOnDurationChange    = []byte("ondurationchange")
	attrOnEnded             = []byte("onended")
	attrOnError             = []byte("onerror")
	attrOnFocus             = []byte("onfocus")
	attrOnHashChange        = []byte("onhashchange")
	attrOnInput             = []byte("oninput")
	attrOnInvalid           = []byte("oninvalid")
	attrOnKeyDown           = []byte("onkeydown")
	attrOnKeyPress          = []byte("onkeypress")
	attrOnKeyUp             = []byte("onkeyup")
	attrOnLoad              = []byte("onload")
	attrOnLoadedData        = []byte("onloadeddata")
	attrOnLoadedMetadata    = []byte("onloadedmetadata")
	attrOnMouseDown         = []byte("onmousedown")
	attrOnMouseEnter        = []byte("onmouseenter")
	attrOnMouseLeave        = []byte("onmouseleave")
	attrOnMouseMove         = []byte("onmousemove")
	attrOnMouseOut          = []byte("onmouseout")
	attrOnMouseOver         = []byte("onmouseover")
	attrOnMouseUp           = []byte("onmouseup")
	attrOnPageHide          = []byte("onpagehide")
	attrOnPageShow          = []byte("onpageshow")
	attrOnPaste             = []byte("onpaste")
	attrOnPause             = []byte("onpause")
	attrOnPlay              = []byte("onplay")
	attrOnPlaying           = []byte("onplaying")
	attrOnPopState          = []byte("onpopstate")
	attrOnProgress          = []byte("onprogress")
	attrOnRateChange        = []byte("onratechange")
	attrOnReset             = []byte("onreset")
	attrOnResize            = []byte("onresize")
	attrOnScroll            = []byte("onscroll")
	attrOnSeeked            = []byte("onseeked")
	attrOnSeeking           = []byte("onseeking")
	attrOnSelect            = []byte("onselect")
	attrOnStalled           = []byte("onstalled")
	attrOnSubmit            = []byte("onsubmit")
	attrOnSuspend           = []byte("onsuspend")
	attrOnTimeUpdate        = []byte("ontimeupdate")
	attrOnToggle            = []byte("ontoggle")
	attrOnUnload            = []byte("onunload")
	attrOnVolumeChange      = []byte("onvolumechange")
	attrOnWaiting           = []byte("onwaiting")
)
