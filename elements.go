package aitch

func A(contents ...Node) Node {
	return newElement(tagA, contents...)
}

func Abbr(contents ...Node) Node {
	return newElement(tagAbbr, contents...)
}

func Address(contents ...Node) Node {
	return newElement(tagAddress, contents...)
}

func Area(contents ...Node) Node {
	return newVoidElement(tagArea, contents...)
}

func Article(contents ...Node) Node {
	return newElement(tagArticle, contents...)
}

func Aside(contents ...Node) Node {
	return newElement(tagAside, contents...)
}

func Audio(contents ...Node) Node {
	return newElement(tagAudio, contents...)
}

func B(contents ...Node) Node {
	return newElement(tagB, contents...)
}

func Base(contents ...Node) Node {
	return newVoidElement(tagBase, contents...)
}

func BlockQuote(contents ...Node) Node {
	return newElement(tagBlockquote, contents...)
}

func Body(contents ...Node) Node {
	return newElement(tagBody, contents...)
}

func Br(contents ...Node) Node {
	return newVoidElement(tagBr, contents...)
}

func Button(contents ...Node) Node {
	return newElement(tagButton, contents...)
}

func Canvas(contents ...Node) Node {
	return newElement(tagCanvas, contents...)
}

func Caption(contents ...Node) Node {
	return newElement(tagCaption, contents...)
}

func Cite(contents ...Node) Node {
	return newElement(tagCite, contents...)
}

func Code(contents ...Node) Node {
	return newElement(tagCode, contents...)
}

func Col(contents ...Node) Node {
	return newVoidElement(tagCol, contents...)
}

func ColGroup(contents ...Node) Node {
	return newElement(tagColgroup, contents...)
}

func DataElement(contents ...Node) Node {
	return newElement(tagData, contents...)
}

func DataList(contents ...Node) Node {
	return newElement(tagDatalist, contents...)
}

func Dd(contents ...Node) Node {
	return newElement(tagDd, contents...)
}

func Del(contents ...Node) Node {
	return newElement(tagDel, contents...)
}

func Dfn(contents ...Node) Node {
	return newElement(tagDfn, contents...)
}

func Dt(contents ...Node) Node {
	return newElement(tagDt, contents...)
}

func Details(contents ...Node) Node {
	return newElement(tagDetails, contents...)
}

func Dialog(contents ...Node) Node {
	return newElement(tagDialog, contents...)
}

func Div(contents ...Node) Node {
	return newElement(tagDiv, contents...)
}

func Dl(contents ...Node) Node {
	return newElement(tagDl, contents...)
}

func Em(contents ...Node) Node {
	return newElement(tagEm, contents...)
}

func Embed(contents ...Node) Node {
	return newVoidElement(tagEmbed, contents...)
}

func Form(contents ...Node) Node {
	return newElement(tagForm, contents...)
}

func FieldSet(contents ...Node) Node {
	return newElement(tagFieldset, contents...)
}

func FigCaption(contents ...Node) Node {
	return newElement(tagFigcaption, contents...)
}

func Figure(contents ...Node) Node {
	return newElement(tagFigure, contents...)
}

func Footer(contents ...Node) Node {
	return newElement(tagFooter, contents...)
}

func H1(contents ...Node) Node {
	return newElement(tagH1, contents...)
}

func H2(contents ...Node) Node {
	return newElement(tagH2, contents...)
}

func H3(contents ...Node) Node {
	return newElement(tagH3, contents...)
}

func H4(contents ...Node) Node {
	return newElement(tagH4, contents...)
}

func H5(contents ...Node) Node {
	return newElement(tagH5, contents...)
}

func H6(contents ...Node) Node {
	return newElement(tagH6, contents...)
}

func Head(contents ...Node) Node {
	return newElement(tagHead, contents...)
}

func Header(contents ...Node) Node {
	return newElement(tagHeader, contents...)
}

func HGroup(contents ...Node) Node {
	return newElement(tagHgroup, contents...)
}

func Hr(contents ...Node) Node {
	return newVoidElement(tagHr, contents...)
}

func Html(contents ...Node) Node {
	return newElement(tagHtml, contents...)
}

func I(contents ...Node) Node {
	return newElement(tagI, contents...)
}

func IFrame(contents ...Node) Node {
	return newElement(tagIframe, contents...)
}

func Img(contents ...Node) Node {
	return newVoidElement(tagImg, contents...)
}

func Input(contents ...Node) Node {
	return newVoidElement(tagInput, contents...)
}

func Ins(contents ...Node) Node {
	return newElement(tagIns, contents...)
}

func Kbd(contents ...Node) Node {
	return newElement(tagKbd, contents...)
}

func Label(contents ...Node) Node {
	return newElement(tagLabel, contents...)
}

func Legend(contents ...Node) Node {
	return newElement(tagLegend, contents...)
}

func Li(contents ...Node) Node {
	return newElement(tagLi, contents...)
}

func Link(contents ...Node) Node {
	return newVoidElement(tagLink, contents...)
}

func Main(contents ...Node) Node {
	return newElement(tagMain, contents...)
}

func Mark(contents ...Node) Node {
	return newElement(tagMark, contents...)
}

func Menu(contents ...Node) Node {
	return newElement(tagMenu, contents...)
}

func Meta(contents ...Node) Node {
	return newVoidElement(tagMeta, contents...)
}

func Meter(contents ...Node) Node {
	return newElement(tagMeter, contents...)
}

func Nav(contents ...Node) Node {
	return newElement(tagNav, contents...)
}

func NoScript(contents ...Node) Node {
	return newElement(tagNoscript, contents...)
}

func Object(contents ...Node) Node {
	return newElement(tagObject, contents...)
}

func Ol(contents ...Node) Node {
	return newElement(tagOl, contents...)
}

func OptGroup(contents ...Node) Node {
	return newElement(tagOptgroup, contents...)
}

func Option(contents ...Node) Node {
	return newElement(tagOption, contents...)
}

func P(contents ...Node) Node {
	return newElement(tagP, contents...)
}

func Para(contents ...Node) Node {
	return newElement(tagP, contents...)
}

func Param(contents ...Node) Node {
	return newVoidElement(tagParam, contents...)
}

func Picture(contents ...Node) Node {
	return newElement(tagPicture, contents...)
}

func Pre(contents ...Node) Node {
	return newElement(tagPre, contents...)
}

func Progress(contents ...Node) Node {
	return newElement(tagProgress, contents...)
}

func Q(contents ...Node) Node {
	return newElement(tagQ, contents...)
}

func S(contents ...Node) Node {
	return newElement(tagS, contents...)
}

func Samp(contents ...Node) Node {
	return newElement(tagSamp, contents...)
}

func Script(contents ...Node) Node {
	return newElement(tagScript, contents...)
}

func Section(contents ...Node) Node {
	return newElement(tagSection, contents...)
}

func Select(contents ...Node) Node {
	return newElement(tagSelect, contents...)
}

func Slot(contents ...Node) Node {
	return newElement(tagSlot, contents...)
}

func Small(contents ...Node) Node {
	return newElement(tagSmall, contents...)
}

func Source(contents ...Node) Node {
	return newVoidElement(tagSource, contents...)
}

func Span(contents ...Node) Node {
	return newElement(tagSpan, contents...)
}

func Strong(contents ...Node) Node {
	return newElement(tagStrong, contents...)
}

func StyleElement(contents ...Node) Node {
	return newElement(tagStyle, contents...)
}

func Sub(contents ...Node) Node {
	return newElement(tagSub, contents...)
}

func Sup(contents ...Node) Node {
	return newElement(tagSup, contents...)
}

func Summary(contents ...Node) Node {
	return newElement(tagSummary, contents...)
}

func Svg(contents ...Node) Node {
	return newElement(tagSvg, contents...)
}

func Table(contents ...Node) Node {
	return newElement(tagTable, contents...)
}

func TBody(contents ...Node) Node {
	return newElement(tagTbody, contents...)
}

func Td(contents ...Node) Node {
	return newElement(tagTd, contents...)
}

func Template(contents ...Node) Node {
	return newElement(tagTemplate, contents...)
}

func Textarea(contents ...Node) Node {
	return newElement(tagTextarea, contents...)
}

func TFoot(contents ...Node) Node {
	return newElement(tagTfoot, contents...)
}

func Th(contents ...Node) Node {
	return newElement(tagTh, contents...)
}

func THead(contents ...Node) Node {
	return newElement(tagThead, contents...)
}

func Time(contents ...Node) Node {
	return newElement(tagTime, contents...)
}

func TitleElement(contents ...Node) Node {
	return newElement(tagTitle, contents...)
}

func Tr(contents ...Node) Node {
	return newElement(tagTr, contents...)
}

func Ul(contents ...Node) Node {
	return newElement(tagUl, contents...)
}

func U(contents ...Node) Node {
	return newElement(tagU, contents...)
}

func Var(contents ...Node) Node {
	return newElement(tagVar, contents...)
}

func Video(contents ...Node) Node {
	return newElement(tagVideo, contents...)
}

func Wbr(contents ...Node) Node {
	return newVoidElement(tagWbr, contents...)
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
	tagSvg        = []byte("svg")
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
