package html

import (
	"golang.org/x/net/html"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// GenerateCode reads HTML and writes fluent HTML code
func GenerateCode(r io.Reader, w io.Writer) error {
	if w == nil {
		w = os.Stdout
	}
	elemMap, attrMap, boolAttrs := mappings()
	tokenizer := html.NewTokenizer(r)
	indent := []byte{}
	eof := false
	var err error
	nl := []byte{'\n'}
	tab := byte('\t')
	funcName := func(fn any) string {
		v := reflect.ValueOf(fn)
		result := "Unknown"
		if v.Kind() == reflect.Func {
			if fpp := runtime.FuncForPC(v.Pointer()); fpp != nil {
				name := fpp.Name()
				if ls := strings.LastIndex(name, "/"); ls > 0 {
					result = name[ls+1:]
				}
			}
		}
		return result
	}
	write := func(s string) {
		_, err = w.Write(indent)
		_, err = w.Write([]byte(s))
		_, err = w.Write(nl)
	}
	incIndent := func() {
		indent = append(indent, tab)
	}
	decIndent := func() {
		if len(indent) > 0 {
			indent = indent[:len(indent)-1]
		}
	}
	escape := func(s string) string {
		return strings.Replace(
			strings.Replace(
				strings.Replace(strings.Replace(s, "\\", `\\`, -1), "\t", `\t`, -1), "\n", `\n`, -1),
			`"`, `\"`, -1)
	}
	writeElementStart := func(name string, attrs []html.Attribute, selfClosing bool) {
		start := ""
		if elem, ok := elemMap[strings.ToLower(string(name))]; ok {
			start = funcName(elem) + "("
		} else {
			start = `aitch.Element("` + string(name) + `", `
		}
		_, err = w.Write(indent)
		_, err = w.Write([]byte(start))
		if selfClosing && len(attrs) == 0 {
			_, err = w.Write([]byte{')', ',', '\n'})
			return
		}
		for i, a := range attrs {
			if i > 0 {
				_, err = w.Write([]byte{',', ' '})
			}
			if attr, ok := attrMap[strings.ToLower(a.Key)]; ok {
				if boolAttrs[strings.ToLower(a.Key)] {
					_, err = w.Write([]byte(funcName(attr)))
					_, err = w.Write([]byte{'(', ')'})
				} else {
					_, err = w.Write([]byte(funcName(attr)))
					_, err = w.Write([]byte{'(', '"'})
					_, err = w.Write([]byte(escape(a.Val)))
					_, err = w.Write([]byte{'"', ')'})
				}
			} else {
				_, err = w.Write([]byte(`aitch.Attribute("`))
				_, err = w.Write([]byte(a.Key))
				_, err = w.Write([]byte{'"', ' ', ',', '"'})
				_, err = w.Write([]byte(escape(a.Val)))
				_, err = w.Write([]byte{'"', ')'})
			}
		}
		if selfClosing {
			_, err = w.Write([]byte{')', ','})
		} else if len(attrs) > 0 {
			_, err = w.Write([]byte{','})
		}
		_, err = w.Write([]byte{'\n'})
		if !selfClosing {
			incIndent()
		}
	}
	write("var template = aitch.Collection(")
	incIndent()
	for err == nil && !eof {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			if err = tokenizer.Err(); err == io.EOF {
				eof = true
				err = nil
			}
		case html.TextToken:
			token := tokenizer.Token()
			if strings.Trim(token.Data, " \n\t") != "" {
				write(`html.Text("` + escape(token.Data) + `"),`)
			}
		case html.StartTagToken:
			token := tokenizer.Token()
			writeElementStart(token.Data, token.Attr, false)
		case html.SelfClosingTagToken:
			token := tokenizer.Token()
			writeElementStart(token.Data, token.Attr, true)
		case html.EndTagToken:
			decIndent()
			write("),")
		case html.CommentToken:
			token := tokenizer.Token()
			write(`html.Comment("` + escape(token.Data) + `"),`)
		case html.DoctypeToken:
			token := tokenizer.Token()
			write(`html.PI("DOCTYPE", "` + escape(token.Data) + `"),`)
		}
	}
	if err == nil {
		decIndent()
		write(")\n")
	}
	return err
}

func mappings() (elems map[string]any, attrs map[string]any, boolAttrs map[string]bool) {
	elems = map[string]any{
		"a":          A,
		"abbr":       Abbr,
		"address":    Address,
		"area":       Area,
		"article":    Article,
		"aside":      Aside,
		"audio":      Audio,
		"b":          B,
		"base":       Base,
		"blockquote": BlockQuote,
		"body":       Body,
		"br":         Br,
		"button":     Button,
		"canvas":     Canvas,
		"caption":    Caption,
		"cite":       Cite,
		"code":       Code,
		"col":        Col,
		"colgroup":   ColGroup,
		"data":       DataElement,
		"datalist":   DataList,
		"dd":         Dd,
		"del":        Del,
		"details":    Details,
		"dfn":        Dfn,
		"dialog":     Dialog,
		"div":        Div,
		"dl":         Dl,
		"dt":         Dt,
		"em":         Em,
		"embed":      Embed,
		"fieldset":   FieldSet,
		"figcaption": FigCaption,
		"figure":     Figure,
		"footer":     Footer,
		"form":       Form,
		"h1":         H1,
		"h2":         H2,
		"h3":         H3,
		"h4":         H4,
		"h5":         H5,
		"h6":         H6,
		"head":       Head,
		"header":     Header,
		"hgroup":     HGroup,
		"hr":         Hr,
		"html":       Html,
		"i":          I,
		"iframe":     IFrame,
		"img":        Img,
		"input":      Input,
		"ins":        Ins,
		"kbd":        Kbd,
		"label":      Label,
		"legend":     Legend,
		"li":         Li,
		"link":       Link,
		"main":       Main,
		"mark":       Mark,
		"menu":       Menu,
		"meta":       Meta,
		"meter":      Meter,
		"nav":        Nav,
		"noscript":   NoScript,
		"object":     Object,
		"ol":         Ol,
		"optgroup":   OptGroup,
		"option":     Option,
		"p":          P,
		"param":      Param,
		"picture":    Picture,
		"pre":        Pre,
		"progress":   Progress,
		"q":          Q,
		"s":          S,
		"samp":       Samp,
		"script":     Script,
		"section":    Section,
		"select":     Select,
		"slot":       Slot,
		"small":      Small,
		"source":     Source,
		"span":       Span,
		"strong":     Strong,
		"style":      StyleElement,
		"sub":        Sub,
		"summary":    Summary,
		"sup":        Sup,
		"table":      Table,
		"tbody":      TBody,
		"td":         Td,
		"template":   Template,
		"textarea":   Textarea,
		"tfoot":      TFoot,
		"th":         Th,
		"thead":      THead,
		"time":       Time,
		"title":      TitleElement,
		"tr":         Tr,
		"u":          U,
		"ul":         Ul,
		"var":        Var,
		"video":      Video,
		"wbr":        Wbr,
	}
	attrs = map[string]any{
		"accept":              Accept,
		"action":              Action,
		"allowfullscreen":     AllowFullScreen,
		"alt":                 Alt,
		"aria-":               Aria,
		"aria-atomic":         AriaAtomic,
		"aria-busy":           AriaBusy,
		"aria-checked":        AriaChecked,
		"aria-controls":       AriaControls,
		"aria-describedby":    AriaDescribedBy,
		"aria-disabled":       AriaDisabled,
		"aria-expanded":       AriaExpanded,
		"aria-flowto":         AriaFlowTo,
		"aria-hidden":         AriaHidden,
		"aria-invalid":        AriaInvalid,
		"aria-label":          AriaLabel,
		"aria-labelledby":     AriaLabelledBy,
		"aria-live":           AriaLive,
		"aria-owns":           AriaOwns,
		"aria-placeholder":    AriaPlaceholder,
		"aria-posinset":       AriaPosInSet,
		"aria-pressed":        AriaPressed,
		"aria-readonly":       AriaReadonly,
		"aria-relevant":       AriaRelevant,
		"aria-required":       AriaRequired,
		"aria-selected":       AriaSelected,
		"aria-setsize":        AriaSetSize,
		"aria-valuemax":       AriaValueMax,
		"aria-valuemin":       AriaValueMin,
		"aria-valuenow":       AriaValueNow,
		"aria-valuetext":      AriaValueText,
		"as":                  As,
		"async":               Async,
		"autocomplete":        AutoComplete,
		"autofocus":           AutoFocus,
		"autoplay":            AutoPlay,
		"charset":             Charset,
		"checked":             Checked,
		"cite":                CiteAttr,
		"class":               Class,
		"cols":                Cols,
		"colspan":             ColSpan,
		"content":             Content,
		"contenteditable":     ContentEditable,
		"controls":            Controls,
		"crossorigin":         CrossOrigin,
		"data-":               Data,
		"datetime":            DateTime,
		"default":             Default,
		"defer":               Defer,
		"dir":                 Dir,
		"disabled":            Disabled,
		"download":            Download,
		"draggable":           Draggable,
		"enctype":             EncType,
		"for":                 For,
		"form":                FormAttr,
		"formnovalidate":      FormNoValidate,
		"height":              Height,
		"hidden":              Hidden,
		"href":                Href,
		"id":                  Id,
		"inert":               Inert,
		"integrity":           Integrity,
		"ismap":               IsMap,
		"itemscope":           ItemScope,
		"lang":                Lang,
		"list":                List,
		"loading":             Loading,
		"loop":                Loop,
		"max":                 Max,
		"maxlength":           MaxLength,
		"method":              Method,
		"min":                 Min,
		"minlength":           MinLength,
		"multiple":            Multiple,
		"muted":               Muted,
		"name":                Name,
		"nomodule":            NoModule,
		"novalidate":          NoValidate,
		"onabort":             OnAbort,
		"onbeforeunload":      OnBeforeUnload,
		"onblur":              OnBlur,
		"oncanplay":           OnCanPlay,
		"oncanplaythrough":    OnCanPlayThrough,
		"onchange":            OnChange,
		"onclick":             OnClick,
		"oncontextmenu":       OnContextMenu,
		"oncopy":              OnCopy,
		"oncut":               OnCut,
		"ondblclick":          OnDblClick,
		"ondrag":              OnDrag,
		"ondragend":           OnDragEnd,
		"ondragenter":         OnDragEnter,
		"ondragleave":         OnDragLeave,
		"ondragover":          OnDragOver,
		"ondragstart":         OnDragStart,
		"ondrop":              OnDrop,
		"ondurationchange":    OnDurationChange,
		"onended":             OnEnded,
		"onerror":             OnError,
		"onfocus":             OnFocus,
		"onhashchange":        OnHashChange,
		"oninput":             OnInput,
		"oninvalid":           OnInvalid,
		"onkeydown":           OnKeyDown,
		"onkeypress":          OnKeyPress,
		"onkeyup":             OnKeyUp,
		"onload":              OnLoad,
		"onloadeddata":        OnLoadedData,
		"onloadedmetadata":    OnLoadedMetadata,
		"onmousedown":         OnMouseDown,
		"onmouseenter":        OnMouseEnter,
		"onmouseleave":        OnMouseLeave,
		"onmousemove":         OnMouseMove,
		"onmouseout":          OnMouseOut,
		"onmouseover":         OnMouseOver,
		"onmouseup":           OnMouseUp,
		"onpagehide":          OnPageHide,
		"onpageshow":          OnPageShow,
		"onpaste":             OnPaste,
		"onpause":             OnPause,
		"onplay":              OnPlay,
		"onplaying":           OnPlaying,
		"onpopstate":          OnPopState,
		"onprogress":          OnProgress,
		"onratechange":        OnRateChange,
		"onreset":             OnReset,
		"onresize":            OnResize,
		"onscroll":            OnScroll,
		"onseeked":            OnSeeked,
		"onseeking":           OnSeeking,
		"onselect":            OnSelect,
		"onstalled":           OnStalled,
		"onsubmit":            OnSubmit,
		"onsuspend":           OnSuspend,
		"ontimeupdate":        OnTimeUpdate,
		"ontoggle":            OnToggle,
		"onunload":            OnUnload,
		"onvolumechange":      OnVolumeChange,
		"onwaiting":           OnWaiting,
		"open":                Open,
		"pattern":             Pattern,
		"placeholder":         Placeholder,
		"playsinline":         PlaysInline,
		"popover":             PopOver,
		"popovertarget":       PopOverTarget,
		"popovertargetaction": PopOverTargetAction,
		"poster":              Poster,
		"preload":             Preload,
		"readonly":            ReadOnly,
		"referrerpolicy":      ReferrerPolicy,
		"rel":                 Rel,
		"required":            Required,
		"reversed":            Reversed,
		"role":                Role,
		"rows":                Rows,
		"rowspan":             RowSpan,
		"selected":            Selected,
		"slot":                SlotAttr,
		"spellcheck":          Spellcheck,
		"src":                 Src,
		"srcset":              SrcSet,
		"step":                Step,
		"style":               Style,
		"tabindex":            TabIndex,
		"target":              Target,
		"title":               Title,
		"translate":           Translate,
		"type":                Type,
		"value":               Value,
		"width":               Width,
	}
	boolAttrs = map[string]bool{
		"allowfullscreen": true,
		"async":           true,
		"autofocus":       true,
		"autoplay":        true,
		"checked":         true,
		"controls":        true,
		"default":         true,
		"defer":           true,
		"disabled":        true,
		"formnovalidate":  true,
		"hidden":          true,
		"inert":           true,
		"ismap":           true,
		"itemscope":       true,
		"loop":            true,
		"multiple":        true,
		"muted":           true,
		"nomodule":        true,
		"novalidate":      true,
		"open":            true,
		"playsinline":     true,
		"readonly":        true,
		"required":        true,
		"reversed":        true,
		"selected":        true,
		"spellcheck":      true,
		"translate":       true,
	}
	return
}
