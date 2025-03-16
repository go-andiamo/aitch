package aitch

func A(contents ...Node) Node {
	return newElement(false, tagA, contents...)
}

func Abbr(contents ...Node) Node {
	return newElement(false, tagAbbr, contents...)
}

func Address(contents ...Node) Node {
	return newElement(false, tagAddress, contents...)
}

func Area(contents ...Node) Node {
	return newElement(true, tagArea, contents...)
}

func Article(contents ...Node) Node {
	return newElement(false, tagArticle, contents...)
}

func Aside(contents ...Node) Node {
	return newElement(false, tagAside, contents...)
}

func Audio(contents ...Node) Node {
	return newElement(false, tagAudio, contents...)
}

func B(contents ...Node) Node {
	return newElement(false, tagB, contents...)
}

func Base(contents ...Node) Node {
	return newElement(true, tagBase, contents...)
}

func BlockQuote(contents ...Node) Node {
	return newElement(false, tagBlockquote, contents...)
}

func Body(contents ...Node) Node {
	return newElement(false, tagBody, contents...)
}

func Br(contents ...Node) Node {
	return newElement(true, tagBr, contents...)
}

func Button(contents ...Node) Node {
	return newElement(false, tagButton, contents...)
}

func Canvas(contents ...Node) Node {
	return newElement(false, tagCanvas, contents...)
}

func Caption(contents ...Node) Node {
	return newElement(false, tagCaption, contents...)
}

func Cite(contents ...Node) Node {
	return newElement(false, tagCite, contents...)
}

func Code(contents ...Node) Node {
	return newElement(false, tagCode, contents...)
}

func Col(contents ...Node) Node {
	return newElement(true, tagCol, contents...)
}

func ColGroup(contents ...Node) Node {
	return newElement(false, tagColgroup, contents...)
}

func DataElement(contents ...Node) Node {
	return newElement(false, tagData, contents...)
}

func DataList(contents ...Node) Node {
	return newElement(false, tagDatalist, contents...)
}

func Dd(contents ...Node) Node {
	return newElement(false, tagDd, contents...)
}

func Del(contents ...Node) Node {
	return newElement(false, tagDel, contents...)
}

func Dfn(contents ...Node) Node {
	return newElement(false, tagDfn, contents...)
}

func Dt(contents ...Node) Node {
	return newElement(false, tagDt, contents...)
}

func Details(contents ...Node) Node {
	return newElement(false, tagDetails, contents...)
}

func Dialog(contents ...Node) Node {
	return newElement(false, tagDialog, contents...)
}

func Div(contents ...Node) Node {
	return newElement(false, tagDiv, contents...)
}

func Dl(contents ...Node) Node {
	return newElement(false, tagDl, contents...)
}

func Em(contents ...Node) Node {
	return newElement(false, tagEm, contents...)
}

func Embed(contents ...Node) Node {
	return newElement(true, tagEmbed, contents...)
}

func Form(contents ...Node) Node {
	return newElement(false, tagForm, contents...)
}

func FieldSet(contents ...Node) Node {
	return newElement(false, tagFieldset, contents...)
}

func FigCaption(contents ...Node) Node {
	return newElement(false, tagFigcaption, contents...)
}

func Figure(contents ...Node) Node {
	return newElement(false, tagFigure, contents...)
}

func Footer(contents ...Node) Node {
	return newElement(false, tagFooter, contents...)
}

func H1(contents ...Node) Node {
	return newElement(false, tagH1, contents...)
}

func H2(contents ...Node) Node {
	return newElement(false, tagH2, contents...)
}

func H3(contents ...Node) Node {
	return newElement(false, tagH3, contents...)
}

func H4(contents ...Node) Node {
	return newElement(false, tagH4, contents...)
}

func H5(contents ...Node) Node {
	return newElement(false, tagH5, contents...)
}

func H6(contents ...Node) Node {
	return newElement(false, tagH6, contents...)
}

func Head(contents ...Node) Node {
	return newElement(false, tagHead, contents...)
}

func Header(contents ...Node) Node {
	return newElement(false, tagHeader, contents...)
}

func HGroup(contents ...Node) Node {
	return newElement(false, tagHgroup, contents...)
}

func Hr(contents ...Node) Node {
	return newElement(true, tagHr, contents...)
}

func Html(contents ...Node) Node {
	return newElement(false, tagHtml, contents...)
}

func I(contents ...Node) Node {
	return newElement(false, tagI, contents...)
}

func IFrame(contents ...Node) Node {
	return newElement(false, tagIframe, contents...)
}

func Img(contents ...Node) Node {
	return newElement(true, tagImg, contents...)
}

func Input(contents ...Node) Node {
	return newElement(true, tagInput, contents...)
}

func Ins(contents ...Node) Node {
	return newElement(false, tagIns, contents...)
}

func Kbd(contents ...Node) Node {
	return newElement(false, tagKbd, contents...)
}

func Label(contents ...Node) Node {
	return newElement(false, tagLabel, contents...)
}

func Legend(contents ...Node) Node {
	return newElement(false, tagLegend, contents...)
}

func Li(contents ...Node) Node {
	return newElement(false, tagLi, contents...)
}

func Link(contents ...Node) Node {
	return newElement(true, tagLink, contents...)
}

func Main(contents ...Node) Node {
	return newElement(false, tagMain, contents...)
}

func Mark(contents ...Node) Node {
	return newElement(false, tagMark, contents...)
}

func Menu(contents ...Node) Node {
	return newElement(false, tagMenu, contents...)
}

func Meta(contents ...Node) Node {
	return newElement(true, tagMeta, contents...)
}

func Meter(contents ...Node) Node {
	return newElement(false, tagMeter, contents...)
}

func Nav(contents ...Node) Node {
	return newElement(false, tagNav, contents...)
}

func NoScript(contents ...Node) Node {
	return newElement(false, tagNoscript, contents...)
}

func Object(contents ...Node) Node {
	return newElement(false, tagObject, contents...)
}

func Ol(contents ...Node) Node {
	return newElement(false, tagOl, contents...)
}

func OptGroup(contents ...Node) Node {
	return newElement(false, tagOptgroup, contents...)
}

func Option(contents ...Node) Node {
	return newElement(false, tagOption, contents...)
}

func P(contents ...Node) Node {
	return newElement(false, tagP, contents...)
}

func Para(contents ...Node) Node {
	return newElement(false, tagP, contents...)
}

func Param(contents ...Node) Node {
	return newElement(true, tagParam, contents...)
}

func Picture(contents ...Node) Node {
	return newElement(false, tagPicture, contents...)
}

func Pre(contents ...Node) Node {
	return newElement(false, tagPre, contents...)
}

func Progress(contents ...Node) Node {
	return newElement(false, tagProgress, contents...)
}

func Q(contents ...Node) Node {
	return newElement(false, tagQ, contents...)
}

func S(contents ...Node) Node {
	return newElement(false, tagS, contents...)
}

func Samp(contents ...Node) Node {
	return newElement(false, tagSamp, contents...)
}

func Script(contents ...Node) Node {
	return newElement(false, tagScript, contents...)
}

func Section(contents ...Node) Node {
	return newElement(false, tagSection, contents...)
}

func Select(contents ...Node) Node {
	return newElement(false, tagSelect, contents...)
}

func Slot(contents ...Node) Node {
	return newElement(false, tagSlot, contents...)
}

func Small(contents ...Node) Node {
	return newElement(false, tagSmall, contents...)
}

func Source(contents ...Node) Node {
	return newElement(true, tagSource, contents...)
}

func Span(contents ...Node) Node {
	return newElement(false, tagSpan, contents...)
}

func Strong(contents ...Node) Node {
	return newElement(false, tagStrong, contents...)
}

func StyleElement(contents ...Node) Node {
	return newElement(false, tagStyle, contents...)
}

func Sub(contents ...Node) Node {
	return newElement(false, tagSub, contents...)
}

func Sup(contents ...Node) Node {
	return newElement(false, tagSup, contents...)
}

func Summary(contents ...Node) Node {
	return newElement(false, tagSummary, contents...)
}

func Svg(contents ...Node) Node {
	return newElement(false, tagSvg, contents...)
}

func Table(contents ...Node) Node {
	return newElement(false, tagTable, contents...)
}

func TBody(contents ...Node) Node {
	return newElement(false, tagTbody, contents...)
}

func Td(contents ...Node) Node {
	return newElement(false, tagTd, contents...)
}

func Template(contents ...Node) Node {
	return newElement(false, tagTemplate, contents...)
}

func Textarea(contents ...Node) Node {
	return newElement(false, tagTextarea, contents...)
}

func TFoot(contents ...Node) Node {
	return newElement(false, tagTfoot, contents...)
}

func Th(contents ...Node) Node {
	return newElement(false, tagTh, contents...)
}

func THead(contents ...Node) Node {
	return newElement(false, tagThead, contents...)
}

func Time(contents ...Node) Node {
	return newElement(false, tagTime, contents...)
}

func TitleElement(contents ...Node) Node {
	return newElement(false, tagTitle, contents...)
}

func Tr(contents ...Node) Node {
	return newElement(false, tagTr, contents...)
}

func Ul(contents ...Node) Node {
	return newElement(false, tagUl, contents...)
}

func U(contents ...Node) Node {
	return newElement(false, tagU, contents...)
}

func Var(contents ...Node) Node {
	return newElement(false, tagVar, contents...)
}

func Video(contents ...Node) Node {
	return newElement(false, tagVideo, contents...)
}

func Wbr(contents ...Node) Node {
	return newElement(true, tagWbr, contents...)
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
