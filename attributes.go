package aitch

func Async() Node {
	return newBooleanAttribute(attrAsync)
}

func AutoFocus() Node {
	return newBooleanAttribute(attrAutofocus)
}

func AutoPlay() Node {
	return newAttribute(attrAutoplay)
}

func Checked() Node {
	return newBooleanAttribute(attrChecked)
}

func Controls() Node {
	return newBooleanAttribute(attrControls)
}

func CrossOrigin(v ...any) Node {
	return newAttribute(attrCrossorigin, v...)
}

func DateTime(v ...any) Node {
	return newAttribute(attrDatetime, v...)
}

func Defer() Node {
	return newBooleanAttribute(attrDefer)
}

func Disabled() Node {
	return newBooleanAttribute(attrDisabled)
}

func Download(v ...any) Node {
	return newAttribute(attrDownload, v...)
}

func Draggable(v ...any) Node {
	return newAttribute(attrDraggable, v...)
}

func Loop() Node {
	return newBooleanAttribute(attrLoop)
}

func Multiple() Node {
	return newBooleanAttribute(attrMultiple)
}

func Muted() Node {
	return newBooleanAttribute(attrMuted)
}

func PlaysInline() Node {
	return newBooleanAttribute(attrPlaysinline)
}

func ReadOnly() Node {
	return newBooleanAttribute(attrReadonly)
}

func Required() Node {
	return newBooleanAttribute(attrRequired)
}

func Selected() Node {
	return newBooleanAttribute(attrSelected)
}

func Accept(v ...any) Node {
	return newAttribute(attrAccept, v...)
}

func Action(v ...any) Node {
	return newAttribute(attrAction, v...)
}

func Alt(v ...any) Node {
	return newAttribute(attrAlt, v...)
}

func Aria(name string, v ...any) Node {
	return newAttribute(append(attrAria, []byte(name)...), v...)
}

func As(v ...any) Node {
	return newAttribute(attrAs, v...)
}

func AutoComplete(v ...any) Node {
	return newAttribute(attrAutocomplete, v...)
}

func Charset(v ...any) Node {
	return newAttribute(attrCharset, v...)
}

func CiteAttr(v ...any) Node {
	return newAttribute(attrCite, v...)
}

func Class(v ...any) Node {
	return newDelimitedAttribute(attrClass, space, v...)
}

func Cols(v ...any) Node {
	return newAttribute(attrCols, v...)
}

func ColSpan(v ...any) Node {
	return newAttribute(attrColspan, v...)
}

func Content(v ...any) Node {
	return newAttribute(attrContent, v...)
}

func ContentEditable(v ...any) Node {
	return newAttribute(attrContenteditable, v...)
}

func Dir(v ...any) Node {
	return newAttribute(attrDir, v...)
}

// Data attribute prefixed with "data-".
func Data(name string, v ...any) Node {
	return newAttribute(append(attrData, []byte(name)...), v...)
}

func EncType(v ...any) Node {
	return newAttribute(attrEnctype, v...)
}

func For(v ...any) Node {
	return newAttribute(attrFor, v...)
}

func FormAttr(v ...any) Node {
	return newAttribute(attrForm, v...)
}

func NoValidate() Node {
	return newBooleanAttribute(attrNovalidate)
}

func FormNoValidate() Node {
	return newBooleanAttribute(attrFormnovalidate)
}

func Height(v ...any) Node {
	return newAttribute(attrHeight, v...)
}

func Hidden() Node {
	return newBooleanAttribute(attrHidden)
}

func Href(v ...any) Node {
	return newAttribute(attrHref, v...)
}

func Id(v ...any) Node {
	return newAttribute(attrId, v...)
}

func Integrity(v ...any) Node {
	return newAttribute(attrIntegrity, v...)
}

func LabelAttr(v ...any) Node {
	return newAttribute(attrLabel, v...)
}

func Lang(v ...any) Node {
	return newAttribute(attrLang, v...)
}

func List(v ...any) Node {
	return newAttribute(attrList, v...)
}

func Loading(v ...any) Node {
	return newAttribute(attrLoading, v...)
}

func Max(v ...any) Node {
	return newAttribute(attrMax, v...)
}

func MaxLength(v ...any) Node {
	return newAttribute(attrMaxlength, v...)
}

func Method(v ...any) Node {
	return newAttribute(attrMethod, v...)
}

func Min(v ...any) Node {
	return newAttribute(attrMin, v...)
}

func MinLength(v ...any) Node {
	return newAttribute(attrMinlength, v...)
}

func Name(v ...any) Node {
	return newAttribute(attrName, v...)
}

func Pattern(v ...any) Node {
	return newAttribute(attrPattern, v...)
}

func Placeholder(v ...any) Node {
	return newAttribute(attrPlaceholder, v...)
}

func Popover(v ...any) Node {
	return newAttribute(attrPopover, v...)
}

func PopoverTarget(v ...any) Node {
	return newAttribute(attrPopovertarget, v...)
}

func PopoverTargetAction(v ...any) Node {
	return newAttribute(attrPopovertargetaction, v...)
}

func Poster(v ...any) Node {
	return newAttribute(attrPoster, v...)
}

func Preload(v ...any) Node {
	return newAttribute(attrPreload, v...)
}

func ReferrerPolicy(v ...any) Node {
	return newAttribute(attrReferrerpolicy, v...)
}

func Rel(v ...any) Node {
	return newAttribute(attrRel, v...)
}

func Role(v ...any) Node {
	return newAttribute(attrRole, v...)
}

func Rows(v ...any) Node {
	return newAttribute(attrRows, v...)
}

func RowSpan(v ...any) Node {
	return newAttribute(attrRowspan, v...)
}

func SlotAttr(v ...any) Node {
	return newAttribute(attrSlot, v...)
}

func Src(v ...any) Node {
	return newAttribute(attrSrc, v...)
}

func SrcSet(v ...any) Node {
	return newAttribute(attrSrcset, v...)
}

func Step(v ...any) Node {
	return newAttribute(attrStep, v...)
}

func Style(v ...any) Node {
	return newDelimitedAttribute(attrStyle, stylesDelim, v...)
}

func TabIndex(v ...any) Node {
	return newAttribute(attrTabindex, v...)
}

func Target(v ...any) Node {
	return newAttribute(attrTarget, v...)
}

func Title(v ...any) Node {
	return newAttribute(attrTitle, v...)
}

func Type(v ...any) Node {
	return newAttribute(attrType, v...)
}

func Value(v ...any) Node {
	return newAttribute(attrValue, v...)
}

func Width(v ...any) Node {
	return newAttribute(attrWidth, v...)
}

func AllowFullScreen() Node {
	return newBooleanAttribute(attrAllowfullscreen)
}
func Default() Node {
	return newBooleanAttribute(attrDefault)
}

func Inert() Node {
	return newBooleanAttribute(attrInert)
}

func IsMap() Node {
	return newBooleanAttribute(attrIsmap)
}

func ItemScope() Node {
	return newBooleanAttribute(attrItemscope)
}

func NoModule() Node {
	return newBooleanAttribute(attrNomodule)
}

func Open() Node {
	return newBooleanAttribute(attrOpen)
}

func Reversed() Node {
	return newBooleanAttribute(attrReversed)
}

func Spellcheck() Node {
	return newBooleanAttribute(attrSpellcheck)
}

func Translate() Node {
	return newBooleanAttribute(attrTranslate)
}

func AriaAtomic(v ...any) Node {
	return newAttribute(attrAriaAtomic, v...)
}

func AriaBusy(v ...any) Node {
	return newAttribute(attrAriaBusy, v...)
}

func AriaChecked(v ...any) Node {
	return newAttribute(attrAriaChecked, v...)
}

func AriaControls(v ...any) Node {
	return newAttribute(attrAriaControls, v...)
}

func AriaDescribedBy(v ...any) Node {
	return newAttribute(attrAriaDescribedby, v...)
}

func AriaDisabled(v ...any) Node {
	return newAttribute(attrAriaDisabled, v...)
}

func AriaExpanded(v ...any) Node {
	return newAttribute(attrAriaExpanded, v...)
}

func AriaFlowTo(v ...any) Node {
	return newAttribute(attrAriaFlowto, v...)
}

func AriaHidden(v ...any) Node {
	return newAttribute(attrAriaHidden, v...)
}

func AriaInvalid(v ...any) Node {
	return newAttribute(attrAriaInvalid, v...)
}

func AriaLabel(v ...any) Node {
	return newAttribute(attrAriaLabel, v...)
}

func AriaLabelledBy(v ...any) Node {
	return newAttribute(attrAriaLabelledby, v...)
}

func AriaLive(v ...any) Node {
	return newAttribute(attrAriaLive, v...)
}

func AriaOwns(v ...any) Node {
	return newAttribute(attrAriaOwns, v...)
}

func AriaPlaceholder(v ...any) Node {
	return newAttribute(attrAriaPlaceholder, v...)
}

func AriaPosInSet(v ...any) Node {
	return newAttribute(attrAriaPosinset, v...)
}

func AriaPressed(v ...any) Node {
	return newAttribute(attrAriaPressed, v...)
}

func AriaReadonly(v ...any) Node {
	return newAttribute(attrAriaReadonly, v...)
}

func AriaRelevant(v ...any) Node {
	return newAttribute(attrAriaRelevant, v...)
}

func AriaRequired(v ...any) Node {
	return newAttribute(attrAriaRequired, v...)
}

func AriaSelected(v ...any) Node {
	return newAttribute(attrAriaSelected, v...)
}

func AriaSetSize(v ...any) Node {
	return newAttribute(attrAriaSetsize, v...)
}

func AriaValueMax(v ...any) Node {
	return newAttribute(attrAriaValuemax, v...)
}

func AriaValueMin(v ...any) Node {
	return newAttribute(attrAriaValuemin, v...)
}

func AriaValueNow(v ...any) Node {
	return newAttribute(attrAriaValuenow, v...)
}

func AriaValueText(v ...any) Node {
	return newAttribute(attrAriaValuetext, v...)
}

func OnAbort(v ...any) Node {
	return newAttribute(attrOnAbort, v...)
}

func OnBeforeUnload(v ...any) Node {
	return newAttribute(attrOnBeforeUnload, v...)
}

func OnBlur(v ...any) Node {
	return newAttribute(attrOnBlur, v...)
}

func OnCanPlay(v ...any) Node {
	return newAttribute(attrOnCanPlay, v...)
}

func OnCanPlayThrough(v ...any) Node {
	return newAttribute(attrOnCanPlayThrough, v...)
}

func OnChange(v ...any) Node {
	return newAttribute(attrOnChange, v...)
}

func OnClick(v ...any) Node {
	return newAttribute(attrOnClick, v...)
}

func OnContextMenu(v ...any) Node {
	return newAttribute(attrOnContextMenu, v...)
}

func OnCopy(v ...any) Node {
	return newAttribute(attrOnCopy, v...)
}

func OnCut(v ...any) Node {
	return newAttribute(attrOnCut, v...)
}

func OnDblClick(v ...any) Node {
	return newAttribute(attrOnDblClick, v...)
}

func OnDrag(v ...any) Node {
	return newAttribute(attrOnDrag, v...)
}

func OnDragEnd(v ...any) Node {
	return newAttribute(attrOnDragEnd, v...)
}

func OnDragEnter(v ...any) Node {
	return newAttribute(attrOnDragEnter, v...)
}

func OnDragLeave(v ...any) Node {
	return newAttribute(attrOnDragLeave, v...)
}

func OnDragOver(v ...any) Node {
	return newAttribute(attrOnDragOver, v...)
}

func OnDragStart(v ...any) Node {
	return newAttribute(attrOnDragStart, v...)
}

func OnDrop(v ...any) Node {
	return newAttribute(attrOnDrop, v...)
}

func OnDurationChange(v ...any) Node {
	return newAttribute(attrOnDurationChange, v...)
}

func OnEnded(v ...any) Node {
	return newAttribute(attrOnEnded, v...)
}

func OnError(v ...any) Node {
	return newAttribute(attrOnError, v...)
}

func OnFocus(v ...any) Node {
	return newAttribute(attrOnFocus, v...)
}

func OnHashChange(v ...any) Node {
	return newAttribute(attrOnHashChange, v...)
}

func OnInput(v ...any) Node {
	return newAttribute(attrOnInput, v...)
}

func OnInvalid(v ...any) Node {
	return newAttribute(attrOnInvalid, v...)
}

func OnKeyDown(v ...any) Node {
	return newAttribute(attrOnKeyDown, v...)
}

func OnKeyPress(v ...any) Node {
	return newAttribute(attrOnKeyPress, v...)
}

func OnKeyUp(v ...any) Node {
	return newAttribute(attrOnKeyUp, v...)
}

func OnLoad(v ...any) Node {
	return newAttribute(attrOnLoad, v...)
}

func OnLoadedData(v ...any) Node {
	return newAttribute(attrOnLoadedData, v...)
}

func OnLoadedMetadata(v ...any) Node {
	return newAttribute(attrOnLoadedMetadata, v...)
}

func OnMouseDown(v ...any) Node {
	return newAttribute(attrOnMouseDown, v...)
}

func OnMouseEnter(v ...any) Node {
	return newAttribute(attrOnMouseEnter, v...)
}

func OnMouseLeave(v ...any) Node {
	return newAttribute(attrOnMouseLeave, v...)
}

func OnMouseMove(v ...any) Node {
	return newAttribute(attrOnMouseMove, v...)
}

func OnMouseOut(v ...any) Node {
	return newAttribute(attrOnMouseOut, v...)
}

func OnMouseOver(v ...any) Node {
	return newAttribute(attrOnMouseOver, v...)
}

func OnMouseUp(v ...any) Node {
	return newAttribute(attrOnMouseUp, v...)
}

func OnPageHide(v ...any) Node {
	return newAttribute(attrOnPageHide, v...)
}

func OnPageShow(v ...any) Node {
	return newAttribute(attrOnPageShow, v...)
}

func OnPaste(v ...any) Node {
	return newAttribute(attrOnPaste, v...)
}

func OnPause(v ...any) Node {
	return newAttribute(attrOnPause, v...)
}

func OnPlay(v ...any) Node {
	return newAttribute(attrOnPlay, v...)
}

func OnPlaying(v ...any) Node {
	return newAttribute(attrOnPlaying, v...)
}

func OnPopState(v ...any) Node {
	return newAttribute(attrOnPopState, v...)
}

func OnProgress(v ...any) Node {
	return newAttribute(attrOnProgress, v...)
}

func OnRateChange(v ...any) Node {
	return newAttribute(attrOnRateChange, v...)
}

func OnReset(v ...any) Node {
	return newAttribute(attrOnReset, v...)
}

func OnResize(v ...any) Node {
	return newAttribute(attrOnResize, v...)
}

func OnScroll(v ...any) Node {
	return newAttribute(attrOnScroll, v...)
}

func OnSeeked(v ...any) Node {
	return newAttribute(attrOnSeeked, v...)
}

func OnSeeking(v ...any) Node {
	return newAttribute(attrOnSeeking, v...)
}

func OnSelect(v ...any) Node {
	return newAttribute(attrOnSelect, v...)
}

func OnStalled(v ...any) Node {
	return newAttribute(attrOnStalled, v...)
}

func OnSubmit(v ...any) Node {
	return newAttribute(attrOnSubmit, v...)
}

func OnSuspend(v ...any) Node {
	return newAttribute(attrOnSuspend, v...)
}

func OnTimeUpdate(v ...any) Node {
	return newAttribute(attrOnTimeUpdate, v...)
}

func OnToggle(v ...any) Node {
	return newAttribute(attrOnToggle, v...)
}

func OnUnload(v ...any) Node {
	return newAttribute(attrOnUnload, v...)
}

func OnVolumeChange(v ...any) Node {
	return newAttribute(attrOnVolumeChange, v...)
}

func OnWaiting(v ...any) Node {
	return newAttribute(attrOnWaiting, v...)
}

var (
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
