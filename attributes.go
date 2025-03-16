package aitch

func Async() Node {
	return newEmptyAttribute(attrAsync)
}

func AutoFocus() Node {
	return newEmptyAttribute(attrAutofocus)
}

func AutoPlay() Node {
	return newAttribute(attrAutoplay)
}

func Checked() Node {
	return newEmptyAttribute(attrChecked)
}

func Controls() Node {
	return newEmptyAttribute(attrControls)
}

func CrossOrigin(v ...any) Node {
	return newAttribute(attrCrossorigin, v...)
}

func DateTime(v ...any) Node {
	return newAttribute(attrDatetime, v...)
}

func Defer() Node {
	return newEmptyAttribute(attrDefer)
}

func Disabled() Node {
	return newEmptyAttribute(attrDisabled)
}

func Download(v ...any) Node {
	return newAttribute(attrDownload, v...)
}

func Draggable(v ...any) Node {
	return newAttribute(attrDraggable, v...)
}

func Loop() Node {
	return newEmptyAttribute(attrLoop)
}

func Multiple() Node {
	return newEmptyAttribute(attrMultiple)
}

func Muted() Node {
	return newEmptyAttribute(attrMuted)
}

func PlaysInline() Node {
	return newEmptyAttribute(attrPlaysinline)
}

func ReadOnly() Node {
	return newEmptyAttribute(attrReadonly)
}

func Required() Node {
	return newEmptyAttribute(attrRequired)
}

func Selected() Node {
	return newEmptyAttribute(attrSelected)
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
	return DelimitedAttribute(attrClass, space, v...)
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

// Data attributes automatically have their name prefixed with "data-".
func Data(name string, v ...any) Node {
	return newAttribute(append(attrData, []byte(name)...), v...)
}

func SlotAttr(v ...any) Node {
	return newAttribute(attrSlot, v...)
}

func For(v ...any) Node {
	return newAttribute(attrFor, v...)
}

func FormAttr(v ...any) Node {
	return newAttribute(attrForm, v...)
}

func FormNoValidate() Node {
	return newEmptyAttribute(attrNovalidate)
}

func Height(v ...any) Node {
	return newAttribute(attrHeight, v...)
}

func Hidden(v ...any) Node {
	return newAttribute(attrHidden, v...)
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
	return DelimitedAttribute(attrStyle, stylesDelim, v...)
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

func EncType(v ...any) Node {
	return newAttribute(attrEnctype, v...)
}

func Dir(v ...any) Node {
	return newAttribute(attrDir, v...)
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
)
