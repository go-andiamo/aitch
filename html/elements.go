package html

import (
	"github.com/go-andiamo/aitch"
)

func A(contents ...any) aitch.Node {
	return aitch.NewElement(tagA, contents...)
}

func Abbr(contents ...any) aitch.Node {
	return aitch.NewElement(tagAbbr, contents...)
}

func Address(contents ...any) aitch.Node {
	return aitch.NewElement(tagAddress, contents...)
}

func Area(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagArea, contents...)
}

func Article(contents ...any) aitch.Node {
	return aitch.NewElement(tagArticle, contents...)
}

func Aside(contents ...any) aitch.Node {
	return aitch.NewElement(tagAside, contents...)
}

func Audio(contents ...any) aitch.Node {
	return aitch.NewElement(tagAudio, contents...)
}

func B(contents ...any) aitch.Node {
	return aitch.NewElement(tagB, contents...)
}

func Base(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagBase, contents...)
}

func BlockQuote(contents ...any) aitch.Node {
	return aitch.NewElement(tagBlockquote, contents...)
}

func Body(contents ...any) aitch.Node {
	return aitch.NewElement(tagBody, contents...)
}

func Br(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagBr, contents...)
}

func Button(contents ...any) aitch.Node {
	return aitch.NewElement(tagButton, contents...)
}

func Canvas(contents ...any) aitch.Node {
	return aitch.NewElement(tagCanvas, contents...)
}

func Caption(contents ...any) aitch.Node {
	return aitch.NewElement(tagCaption, contents...)
}

func Cite(contents ...any) aitch.Node {
	return aitch.NewElement(tagCite, contents...)
}

func Code(contents ...any) aitch.Node {
	return aitch.NewElement(tagCode, contents...)
}

func Col(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagCol, contents...)
}

func ColGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagColgroup, contents...)
}

func DataElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagData, contents...)
}

func DataList(contents ...any) aitch.Node {
	return aitch.NewElement(tagDatalist, contents...)
}

func Dd(contents ...any) aitch.Node {
	return aitch.NewElement(tagDd, contents...)
}

func Del(contents ...any) aitch.Node {
	return aitch.NewElement(tagDel, contents...)
}

func Dfn(contents ...any) aitch.Node {
	return aitch.NewElement(tagDfn, contents...)
}

func Dt(contents ...any) aitch.Node {
	return aitch.NewElement(tagDt, contents...)
}

func Details(contents ...any) aitch.Node {
	return aitch.NewElement(tagDetails, contents...)
}

func Dialog(contents ...any) aitch.Node {
	return aitch.NewElement(tagDialog, contents...)
}

func Div(contents ...any) aitch.Node {
	return aitch.NewElement(tagDiv, contents...)
}

func Dl(contents ...any) aitch.Node {
	return aitch.NewElement(tagDl, contents...)
}

func Em(contents ...any) aitch.Node {
	return aitch.NewElement(tagEm, contents...)
}

func Embed(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagEmbed, contents...)
}

func Form(contents ...any) aitch.Node {
	return aitch.NewElement(tagForm, contents...)
}

func FieldSet(contents ...any) aitch.Node {
	return aitch.NewElement(tagFieldset, contents...)
}

func FigCaption(contents ...any) aitch.Node {
	return aitch.NewElement(tagFigcaption, contents...)
}

func Figure(contents ...any) aitch.Node {
	return aitch.NewElement(tagFigure, contents...)
}

func Footer(contents ...any) aitch.Node {
	return aitch.NewElement(tagFooter, contents...)
}

func H1(contents ...any) aitch.Node {
	return aitch.NewElement(tagH1, contents...)
}

func H2(contents ...any) aitch.Node {
	return aitch.NewElement(tagH2, contents...)
}

func H3(contents ...any) aitch.Node {
	return aitch.NewElement(tagH3, contents...)
}

func H4(contents ...any) aitch.Node {
	return aitch.NewElement(tagH4, contents...)
}

func H5(contents ...any) aitch.Node {
	return aitch.NewElement(tagH5, contents...)
}

func H6(contents ...any) aitch.Node {
	return aitch.NewElement(tagH6, contents...)
}

func Head(contents ...any) aitch.Node {
	return aitch.NewElement(tagHead, contents...)
}

func Header(contents ...any) aitch.Node {
	return aitch.NewElement(tagHeader, contents...)
}

func HGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagHgroup, contents...)
}

func Hr(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagHr, contents...)
}

func Html(contents ...any) aitch.Node {
	return aitch.NewElement(tagHtml, contents...)
}

func I(contents ...any) aitch.Node {
	return aitch.NewElement(tagI, contents...)
}

func IFrame(contents ...any) aitch.Node {
	return aitch.NewElement(tagIframe, contents...)
}

func Img(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagImg, contents...)
}

func Input(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagInput, contents...)
}

func Ins(contents ...any) aitch.Node {
	return aitch.NewElement(tagIns, contents...)
}

func Kbd(contents ...any) aitch.Node {
	return aitch.NewElement(tagKbd, contents...)
}

func Label(contents ...any) aitch.Node {
	return aitch.NewElement(tagLabel, contents...)
}

func Legend(contents ...any) aitch.Node {
	return aitch.NewElement(tagLegend, contents...)
}

func Li(contents ...any) aitch.Node {
	return aitch.NewElement(tagLi, contents...)
}

func Link(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagLink, contents...)
}

func Main(contents ...any) aitch.Node {
	return aitch.NewElement(tagMain, contents...)
}

func Mark(contents ...any) aitch.Node {
	return aitch.NewElement(tagMark, contents...)
}

func Menu(contents ...any) aitch.Node {
	return aitch.NewElement(tagMenu, contents...)
}

func Meta(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagMeta, contents...)
}

func Meter(contents ...any) aitch.Node {
	return aitch.NewElement(tagMeter, contents...)
}

func Nav(contents ...any) aitch.Node {
	return aitch.NewElement(tagNav, contents...)
}

func NoScript(contents ...any) aitch.Node {
	return aitch.NewElement(tagNoscript, contents...)
}

func Object(contents ...any) aitch.Node {
	return aitch.NewElement(tagObject, contents...)
}

func Ol(contents ...any) aitch.Node {
	return aitch.NewElement(tagOl, contents...)
}

func OptGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagOptgroup, contents...)
}

func Option(contents ...any) aitch.Node {
	return aitch.NewElement(tagOption, contents...)
}

func P(contents ...any) aitch.Node {
	return aitch.NewElement(tagP, contents...)
}

func Para(contents ...any) aitch.Node {
	return aitch.NewElement(tagP, contents...)
}

func Param(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagParam, contents...)
}

func Picture(contents ...any) aitch.Node {
	return aitch.NewElement(tagPicture, contents...)
}

func Pre(contents ...any) aitch.Node {
	return aitch.NewElement(tagPre, contents...)
}

func Progress(contents ...any) aitch.Node {
	return aitch.NewElement(tagProgress, contents...)
}

func Q(contents ...any) aitch.Node {
	return aitch.NewElement(tagQ, contents...)
}

func S(contents ...any) aitch.Node {
	return aitch.NewElement(tagS, contents...)
}

func Samp(contents ...any) aitch.Node {
	return aitch.NewElement(tagSamp, contents...)
}

func Script(contents ...any) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

func Section(contents ...any) aitch.Node {
	return aitch.NewElement(tagSection, contents...)
}

func Select(contents ...any) aitch.Node {
	return aitch.NewElement(tagSelect, contents...)
}

func Slot(contents ...any) aitch.Node {
	return aitch.NewElement(tagSlot, contents...)
}

func Small(contents ...any) aitch.Node {
	return aitch.NewElement(tagSmall, contents...)
}

func Source(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagSource, contents...)
}

func Span(contents ...any) aitch.Node {
	return aitch.NewElement(tagSpan, contents...)
}

func Strong(contents ...any) aitch.Node {
	return aitch.NewElement(tagStrong, contents...)
}

func StyleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

func Sub(contents ...any) aitch.Node {
	return aitch.NewElement(tagSub, contents...)
}

func Sup(contents ...any) aitch.Node {
	return aitch.NewElement(tagSup, contents...)
}

func Summary(contents ...any) aitch.Node {
	return aitch.NewElement(tagSummary, contents...)
}

func Table(contents ...any) aitch.Node {
	return aitch.NewElement(tagTable, contents...)
}

func TBody(contents ...any) aitch.Node {
	return aitch.NewElement(tagTbody, contents...)
}

func Td(contents ...any) aitch.Node {
	return aitch.NewElement(tagTd, contents...)
}

func Template(contents ...any) aitch.Node {
	return aitch.NewElement(tagTemplate, contents...)
}

func Textarea(contents ...any) aitch.Node {
	return aitch.NewElement(tagTextarea, contents...)
}

func TFoot(contents ...any) aitch.Node {
	return aitch.NewElement(tagTfoot, contents...)
}

func Th(contents ...any) aitch.Node {
	return aitch.NewElement(tagTh, contents...)
}

func THead(contents ...any) aitch.Node {
	return aitch.NewElement(tagThead, contents...)
}

func Time(contents ...any) aitch.Node {
	return aitch.NewElement(tagTime, contents...)
}

func TitleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

func Tr(contents ...any) aitch.Node {
	return aitch.NewElement(tagTr, contents...)
}

func Ul(contents ...any) aitch.Node {
	return aitch.NewElement(tagUl, contents...)
}

func U(contents ...any) aitch.Node {
	return aitch.NewElement(tagU, contents...)
}

func Var(contents ...any) aitch.Node {
	return aitch.NewElement(tagVar, contents...)
}

func Video(contents ...any) aitch.Node {
	return aitch.NewElement(tagVideo, contents...)
}

func Wbr(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagWbr, contents...)
}

var (
	tagA          = []byte("a")
	tagAbbr       = []byte("abbr")
	tagAddress    = []byte("address")
	tagArea       = []byte("area")
	tagArticle    = []byte("article")
	tagAside      = []byte("aside")
	tagAudio      = []byte("audio")
	tagB          = []byte("b")
	tagBase       = []byte("base")
	tagBlockquote = []byte("blockquote")
	tagBody       = []byte("body")
	tagBr         = []byte("br")
	tagButton     = []byte("button")
	tagCanvas     = []byte("canvas")
	tagCaption    = []byte("caption")
	tagCite       = []byte("cite")
	tagCode       = []byte("code")
	tagCol        = []byte("col")
	tagColgroup   = []byte("colgroup")
	tagData       = []byte("data")
	tagDatalist   = []byte("datalist")
	tagDd         = []byte("dd")
	tagDel        = []byte("del")
	tagDetails    = []byte("details")
	tagDfn        = []byte("dfn")
	tagDialog     = []byte("dialog")
	tagDiv        = []byte("div")
	tagDl         = []byte("dl")
	tagDt         = []byte("dt")
	tagEm         = []byte("em")
	tagEmbed      = []byte("embed")
	tagFieldset   = []byte("fieldset")
	tagFigcaption = []byte("figcaption")
	tagFigure     = []byte("figure")
	tagFooter     = []byte("footer")
	tagForm       = []byte("form")
	tagH1         = []byte("h1")
	tagH2         = []byte("h2")
	tagH3         = []byte("h3")
	tagH4         = []byte("h4")
	tagH5         = []byte("h5")
	tagH6         = []byte("h6")
	tagHead       = []byte("head")
	tagHeader     = []byte("header")
	tagHgroup     = []byte("hgroup")
	tagHr         = []byte("hr")
	tagHtml       = []byte("html")
	tagI          = []byte("i")
	tagIframe     = []byte("iframe")
	tagImg        = []byte("img")
	tagInput      = []byte("input")
	tagIns        = []byte("ins")
	tagKbd        = []byte("kbd")
	tagLabel      = []byte("label")
	tagLegend     = []byte("legend")
	tagLi         = []byte("li")
	tagLink       = []byte("link")
	tagMain       = []byte("main")
	tagMark       = []byte("mark")
	tagMenu       = []byte("menu")
	tagMeta       = []byte("meta")
	tagMeter      = []byte("meter")
	tagNav        = []byte("nav")
	tagNoscript   = []byte("noscript")
	tagObject     = []byte("object")
	tagOl         = []byte("ol")
	tagOptgroup   = []byte("optgroup")
	tagOption     = []byte("option")
	tagP          = []byte("p")
	tagParam      = []byte("param")
	tagPicture    = []byte("picture")
	tagPre        = []byte("pre")
	tagProgress   = []byte("progress")
	tagQ          = []byte("q")
	tagS          = []byte("s")
	tagSamp       = []byte("samp")
	tagScript     = []byte("script")
	tagSection    = []byte("section")
	tagSelect     = []byte("select")
	tagSlot       = []byte("slot")
	tagSmall      = []byte("small")
	tagSource     = []byte("source")
	tagSpan       = []byte("span")
	tagStrong     = []byte("strong")
	tagStyle      = []byte("style")
	tagSub        = []byte("sub")
	tagSummary    = []byte("summary")
	tagSup        = []byte("sup")
	tagTable      = []byte("table")
	tagTbody      = []byte("tbody")
	tagTd         = []byte("td")
	tagTemplate   = []byte("template")
	tagTextarea   = []byte("textarea")
	tagTfoot      = []byte("tfoot")
	tagTh         = []byte("th")
	tagThead      = []byte("thead")
	tagTime       = []byte("time")
	tagTitle      = []byte("title")
	tagTr         = []byte("tr")
	tagU          = []byte("u")
	tagUl         = []byte("ul")
	tagVar        = []byte("var")
	tagVideo      = []byte("video")
	tagWbr        = []byte("wbr")
)
