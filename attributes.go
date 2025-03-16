package aitch

func Async() Node {
	return newEmptyAttribute("async")
}

func AutoFocus() Node {
	return newEmptyAttribute("autofocus")
}

func AutoPlay() Node {
	return newAttribute("autoplay")
}

func Checked() Node {
	return newEmptyAttribute("checked")
}

func Controls() Node {
	return newEmptyAttribute("controls")
}

func CrossOrigin(v ...any) Node {
	return newAttribute("crossorigin", v...)
}

func DateTime(v ...any) Node {
	return newAttribute("datetime", v...)
}

func Defer() Node {
	return newEmptyAttribute("defer")
}

func Disabled() Node {
	return newEmptyAttribute("disabled")
}

func Download(v ...any) Node {
	return newAttribute("download", v...)
}

func Draggable(v ...any) Node {
	return newAttribute("draggable", v...)
}

func Loop() Node {
	return newEmptyAttribute("loop")
}

func Multiple() Node {
	return newEmptyAttribute("multiple")
}

func Muted() Node {
	return newEmptyAttribute("muted")
}

func PlaysInline() Node {
	return newEmptyAttribute("playsinline")
}

func ReadOnly() Node {
	return newEmptyAttribute("readonly")
}

func Required() Node {
	return newEmptyAttribute("required")
}

func Selected() Node {
	return newEmptyAttribute("selected")
}

func Accept(v ...any) Node {
	return newAttribute("accept", v...)
}

func Action(v ...any) Node {
	return newAttribute("action", v...)
}

func Alt(v ...any) Node {
	return newAttribute("alt", v...)
}

func Aria(name string, v ...any) Node {
	return newAttribute("aria-"+name, v...)
}

func As(v ...any) Node {
	return newAttribute("as", v...)
}

func AutoComplete(v ...any) Node {
	return newAttribute("autocomplete", v...)
}

func Charset(v ...any) Node {
	return newAttribute("charset", v...)
}

func CiteAttr(v ...any) Node {
	return newAttribute("cite", v...)
}

func Class(v ...any) Node {
	return DelimitedAttribute("class", space, v...)
}

func Cols(v ...any) Node {
	return newAttribute("cols", v...)
}

func ColSpan(v ...any) Node {
	return newAttribute("colspan", v...)
}

func Content(v ...any) Node {
	return newAttribute("content", v...)
}

// Data attributes automatically have their name prefixed with "data-".
func Data(name string, v ...any) Node {
	return newAttribute("data-"+name, v...)
}

func SlotAttr(v ...any) Node {
	return newAttribute("slot", v...)
}

func For(v ...any) Node {
	return newAttribute("for", v...)
}

func FormAction(v ...any) Node {
	return newAttribute("action", v...)
}

func FormAttr(v ...any) Node {
	return newAttribute("form", v...)
}

func FormEncType(v ...any) Node {
	return newAttribute("enctype", v...)
}

func FormMethod(v ...any) Node {
	return newAttribute("method", v...)
}

func FormNoValidate() Node {
	return newEmptyAttribute("novalidate")
}

func FormTarget(v ...any) Node {
	return newAttribute("target", v...)
}

func Height(v ...any) Node {
	return newAttribute("height", v...)
}

func Hidden(v ...any) Node {
	return newAttribute("hidden", v...)
}

func Href(v ...any) Node {
	return newAttribute("href", v...)
}

func Id(v ...any) Node {
	return newAttribute("id", v...)
}

func Integrity(v ...any) Node {
	return newAttribute("integrity", v...)
}

func LabelAttr(v ...any) Node {
	return newAttribute("label", v...)
}

func Lang(v ...any) Node {
	return newAttribute("lang", v...)
}

func List(v ...any) Node {
	return newAttribute("list", v...)
}

func Loading(v ...any) Node {
	return newAttribute("loading", v...)
}

func Max(v ...any) Node {
	return newAttribute("max", v...)
}

func MaxLength(v ...any) Node {
	return newAttribute("maxlength", v...)
}

func Method(v ...any) Node {
	return newAttribute("method", v...)
}

func Min(v ...any) Node {
	return newAttribute("min", v...)
}

func MinLength(v ...any) Node {
	return newAttribute("minlength", v...)
}

func Name(v ...any) Node {
	return newAttribute("name", v...)
}

func Pattern(v ...any) Node {
	return newAttribute("pattern", v...)
}

func Placeholder(v ...any) Node {
	return newAttribute("placeholder", v...)
}

func Popover(v ...any) Node {
	return newAttribute("popover", v...)
}

func PopoverTarget(v ...any) Node {
	return newAttribute("popovertarget", v...)
}

func PopoverTargetAction(v ...any) Node {
	return newAttribute("popovertargetaction", v...)
}

func Poster(v ...any) Node {
	return newAttribute("poster", v...)
}

func Preload(v ...any) Node {
	return newAttribute("preload", v...)
}

func ReferrerPolicy(v ...any) Node {
	return newAttribute("referrerpolicy", v...)
}

func Rel(v ...any) Node {
	return newAttribute("rel", v...)
}

func Role(v ...any) Node {
	return newAttribute("role", v...)
}

func Rows(v ...any) Node {
	return newAttribute("rows", v...)
}

func RowSpan(v ...any) Node {
	return newAttribute("rowspan", v...)
}

func Src(v ...any) Node {
	return newAttribute("src", v...)
}

func SrcSet(v ...any) Node {
	return newAttribute("srcset", v...)
}

func Step(v ...any) Node {
	return newAttribute("step", v...)
}

func Style(v ...any) Node {
	return DelimitedAttribute("style", stylesDelim, v...)
}

func TabIndex(v ...any) Node {
	return newAttribute("tabindex", v...)
}

func Target(v ...any) Node {
	return newAttribute("target", v...)
}

func Title(v ...any) Node {
	return newAttribute("title", v...)
}

func Type(v ...any) Node {
	return newAttribute("type", v...)
}

func Value(v ...any) Node {
	return newAttribute("value", v...)
}

func Width(v ...any) Node {
	return newAttribute("width", v...)
}

func EncType(v ...any) Node {
	return newAttribute("enctype", v...)
}

func Dir(v ...any) Node {
	return newAttribute("dir", v...)
}
