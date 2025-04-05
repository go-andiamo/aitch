package html

import (
	"github.com/go-andiamo/aitch"
)

// A declares a <a> HTML element
func A(contents ...any) aitch.Node {
	return aitch.NewElement(tagA, contents...)
}

// Abbr declares a <abbr> HTML element
func Abbr(contents ...any) aitch.Node {
	return aitch.NewElement(tagAbbr, contents...)
}

// Address declares a <address> HTML element
func Address(contents ...any) aitch.Node {
	return aitch.NewElement(tagAddress, contents...)
}

// Area declares a <area> HTML element
//
// As <area> is a void element - any non-attribute contents are ignored
func Area(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagArea, contents...)
}

// Article declares a <article> HTML element
func Article(contents ...any) aitch.Node {
	return aitch.NewElement(tagArticle, contents...)
}

// Aside declares a <aside> HTML element
func Aside(contents ...any) aitch.Node {
	return aitch.NewElement(tagAside, contents...)
}

// Audio declares a <audio> HTML element
func Audio(contents ...any) aitch.Node {
	return aitch.NewElement(tagAudio, contents...)
}

// B declares a <b> HTML element
func B(contents ...any) aitch.Node {
	return aitch.NewElement(tagB, contents...)
}

// Base declares a <base> HTML element
//
// As <base> is a void element - any non-attribute contents are ignored
func Base(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagBase, contents...)
}

// BlockQuote declares a <blockquote> HTML element
func BlockQuote(contents ...any) aitch.Node {
	return aitch.NewElement(tagBlockquote, contents...)
}

// Body declares a <body> HTML element
func Body(contents ...any) aitch.Node {
	return aitch.NewElement(tagBody, contents...)
}

// Br declares a <br> HTML element
//
// As <br> is a void element - any non-attribute contents are ignored
func Br(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagBr, contents...)
}

// Button declares a <button> HTML element
func Button(contents ...any) aitch.Node {
	return aitch.NewElement(tagButton, contents...)
}

// Canvas declares a <canvas> HTML element
func Canvas(contents ...any) aitch.Node {
	return aitch.NewElement(tagCanvas, contents...)
}

// Caption declares a <caption> HTML element
func Caption(contents ...any) aitch.Node {
	return aitch.NewElement(tagCaption, contents...)
}

// Cite declares a <cite> HTML element
func Cite(contents ...any) aitch.Node {
	return aitch.NewElement(tagCite, contents...)
}

// Code declares a <code> HTML element
func Code(contents ...any) aitch.Node {
	return aitch.NewElement(tagCode, contents...)
}

// Col declares a <col> HTML element
//
// As <col> is a void element - any non-attribute contents are ignored
func Col(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagCol, contents...)
}

// ColGroup declares a <colgroup> HTML element
func ColGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagColgroup, contents...)
}

// DataElement declares a <data> HTML element
func DataElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagData, contents...)
}

// DataList declares a <datalist> HTML element
func DataList(contents ...any) aitch.Node {
	return aitch.NewElement(tagDatalist, contents...)
}

// Dd declares a <dd> HTML element
func Dd(contents ...any) aitch.Node {
	return aitch.NewElement(tagDd, contents...)
}

// Del declares a <del> HTML element
func Del(contents ...any) aitch.Node {
	return aitch.NewElement(tagDel, contents...)
}

// Dfn declares a <dfn> HTML element
func Dfn(contents ...any) aitch.Node {
	return aitch.NewElement(tagDfn, contents...)
}

// Dt declares a <dt> HTML element
func Dt(contents ...any) aitch.Node {
	return aitch.NewElement(tagDt, contents...)
}

// Details declares a <details> HTML element
func Details(contents ...any) aitch.Node {
	return aitch.NewElement(tagDetails, contents...)
}

// Dialog declares a <dialog> HTML element
func Dialog(contents ...any) aitch.Node {
	return aitch.NewElement(tagDialog, contents...)
}

// Div declares a <div> HTML element
func Div(contents ...any) aitch.Node {
	return aitch.NewElement(tagDiv, contents...)
}

// Dl declares a <dl> HTML element
func Dl(contents ...any) aitch.Node {
	return aitch.NewElement(tagDl, contents...)
}

// Em declares a <em> HTML element
func Em(contents ...any) aitch.Node {
	return aitch.NewElement(tagEm, contents...)
}

// Embed declares a <embed> HTML element
//
// As <embed> is a void element - any non-attribute contents are ignored
func Embed(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagEmbed, contents...)
}

// Form declares a <form> HTML element
func Form(contents ...any) aitch.Node {
	return aitch.NewElement(tagForm, contents...)
}

// FieldSet declares a <fieldset> HTML element
func FieldSet(contents ...any) aitch.Node {
	return aitch.NewElement(tagFieldset, contents...)
}

// FigCaption declares a <figcaption> HTML element
func FigCaption(contents ...any) aitch.Node {
	return aitch.NewElement(tagFigcaption, contents...)
}

// Figure declares a <figure> HTML element
func Figure(contents ...any) aitch.Node {
	return aitch.NewElement(tagFigure, contents...)
}

// Footer declares a <footer> HTML element
func Footer(contents ...any) aitch.Node {
	return aitch.NewElement(tagFooter, contents...)
}

// H1 declares a <h1> HTML element
func H1(contents ...any) aitch.Node {
	return aitch.NewElement(tagH1, contents...)
}

// H2 declares a <h2> HTML element
func H2(contents ...any) aitch.Node {
	return aitch.NewElement(tagH2, contents...)
}

// H3 declares a <h3> HTML element
func H3(contents ...any) aitch.Node {
	return aitch.NewElement(tagH3, contents...)
}

// H4 declares a <h4> HTML element
func H4(contents ...any) aitch.Node {
	return aitch.NewElement(tagH4, contents...)
}

// H5 declares a <h5> HTML element
func H5(contents ...any) aitch.Node {
	return aitch.NewElement(tagH5, contents...)
}

// H6 declares a <h6> HTML element
func H6(contents ...any) aitch.Node {
	return aitch.NewElement(tagH6, contents...)
}

// Head declares a <head> HTML element
func Head(contents ...any) aitch.Node {
	return aitch.NewElement(tagHead, contents...)
}

// Header declares a <header> HTML element
func Header(contents ...any) aitch.Node {
	return aitch.NewElement(tagHeader, contents...)
}

// HGroup declares a <hgroup> HTML element
func HGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagHgroup, contents...)
}

// Hr declares a <hr> HTML element
//
// As <hr> is a void element - any non-attribute contents are ignored
func Hr(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagHr, contents...)
}

// Html declares a <html> HTML element
func Html(contents ...any) aitch.Node {
	return aitch.NewElement(tagHtml, contents...)
}

// I declares a <i> HTML element
func I(contents ...any) aitch.Node {
	return aitch.NewElement(tagI, contents...)
}

// IFrame declares a <iframe> HTML element
func IFrame(contents ...any) aitch.Node {
	return aitch.NewElement(tagIframe, contents...)
}

// Img declares a <img> HTML element
//
// As <img> is a void element - any non-attribute contents are ignored
func Img(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagImg, contents...)
}

// Input declares a <input> HTML element
//
// As <input> is a void element - any non-attribute contents are ignored
func Input(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagInput, contents...)
}

// Ins declares a <ins> HTML element
func Ins(contents ...any) aitch.Node {
	return aitch.NewElement(tagIns, contents...)
}

// Kbd declares a <kbd> HTML element
func Kbd(contents ...any) aitch.Node {
	return aitch.NewElement(tagKbd, contents...)
}

// Label declares a <label> HTML element
func Label(contents ...any) aitch.Node {
	return aitch.NewElement(tagLabel, contents...)
}

// Legend declares a <legend> HTML element
func Legend(contents ...any) aitch.Node {
	return aitch.NewElement(tagLegend, contents...)
}

// Li declares a <li> HTML element
func Li(contents ...any) aitch.Node {
	return aitch.NewElement(tagLi, contents...)
}

// Link declares a <link> HTML element
//
// As <link> is a void element - any non-attribute contents are ignored
func Link(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagLink, contents...)
}

// Main declares a <main> HTML element
func Main(contents ...any) aitch.Node {
	return aitch.NewElement(tagMain, contents...)
}

// Mark declares a <mark> HTML element
func Mark(contents ...any) aitch.Node {
	return aitch.NewElement(tagMark, contents...)
}

// Menu declares a <menu> HTML element
func Menu(contents ...any) aitch.Node {
	return aitch.NewElement(tagMenu, contents...)
}

// Meta declares a <meta> HTML element
//
// As <meta> is a void element - any non-attribute contents are ignored
func Meta(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagMeta, contents...)
}

// Meter declares a <meter> HTML element
func Meter(contents ...any) aitch.Node {
	return aitch.NewElement(tagMeter, contents...)
}

// Nav declares a <nav> HTML element
func Nav(contents ...any) aitch.Node {
	return aitch.NewElement(tagNav, contents...)
}

// NoScript declares a <noscript> HTML element
func NoScript(contents ...any) aitch.Node {
	return aitch.NewElement(tagNoscript, contents...)
}

// Object declares a <object> HTML element
func Object(contents ...any) aitch.Node {
	return aitch.NewElement(tagObject, contents...)
}

// Ol declares a <ol> HTML element
func Ol(contents ...any) aitch.Node {
	return aitch.NewElement(tagOl, contents...)
}

// OptGroup declares a <optgroup> HTML element
func OptGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagOptgroup, contents...)
}

// Option declares a <option> HTML element
func Option(contents ...any) aitch.Node {
	return aitch.NewElement(tagOption, contents...)
}

// P declares a <p> HTML element
func P(contents ...any) aitch.Node {
	return aitch.NewElement(tagP, contents...)
}

// Para declares a <para> HTML element
func Para(contents ...any) aitch.Node {
	return aitch.NewElement(tagP, contents...)
}

// Param declares a <param> HTML element
//
// As <param> is a void element - any non-attribute contents are ignored
func Param(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagParam, contents...)
}

// Picture declares a <picture> HTML element
func Picture(contents ...any) aitch.Node {
	return aitch.NewElement(tagPicture, contents...)
}

// Pre declares a <pre> HTML element
func Pre(contents ...any) aitch.Node {
	return aitch.NewElement(tagPre, contents...)
}

// Progress declares a <progress> HTML element
func Progress(contents ...any) aitch.Node {
	return aitch.NewElement(tagProgress, contents...)
}

// Q declares a <q> HTML element
func Q(contents ...any) aitch.Node {
	return aitch.NewElement(tagQ, contents...)
}

// S declares a <s> HTML element
func S(contents ...any) aitch.Node {
	return aitch.NewElement(tagS, contents...)
}

// Samp declares a <samp> HTML element
func Samp(contents ...any) aitch.Node {
	return aitch.NewElement(tagSamp, contents...)
}

// Script declares a <script> HTML element
func Script(contents ...any) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

// Section declares a <section> HTML element
func Section(contents ...any) aitch.Node {
	return aitch.NewElement(tagSection, contents...)
}

// Select declares a <select> HTML element
func Select(contents ...any) aitch.Node {
	return aitch.NewElement(tagSelect, contents...)
}

// Slot declares a <slot> HTML element
func Slot(contents ...any) aitch.Node {
	return aitch.NewElement(tagSlot, contents...)
}

// Small declares a <small> HTML element
func Small(contents ...any) aitch.Node {
	return aitch.NewElement(tagSmall, contents...)
}

// Source declares a <source> HTML element
//
// As <source> is a void element - any non-attribute contents are ignored
func Source(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagSource, contents...)
}

// Span declares a <span> HTML element
func Span(contents ...any) aitch.Node {
	return aitch.NewElement(tagSpan, contents...)
}

// Strong declares a <strong> HTML element
func Strong(contents ...any) aitch.Node {
	return aitch.NewElement(tagStrong, contents...)
}

// StyleElement declares a <style> HTML element
func StyleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

// Sub declares a <sub> HTML element
func Sub(contents ...any) aitch.Node {
	return aitch.NewElement(tagSub, contents...)
}

// Sup declares a <sup> HTML element
func Sup(contents ...any) aitch.Node {
	return aitch.NewElement(tagSup, contents...)
}

// Summary declares a <summary> HTML element
func Summary(contents ...any) aitch.Node {
	return aitch.NewElement(tagSummary, contents...)
}

// Table declares a <table> HTML element
func Table(contents ...any) aitch.Node {
	return aitch.NewElement(tagTable, contents...)
}

// TBody declares a <tbody> HTML element
func TBody(contents ...any) aitch.Node {
	return aitch.NewElement(tagTbody, contents...)
}

// Td declares a <td> HTML element
func Td(contents ...any) aitch.Node {
	return aitch.NewElement(tagTd, contents...)
}

// Template declares a <template> HTML element
func Template(contents ...any) aitch.Node {
	return aitch.NewElement(tagTemplate, contents...)
}

// Textarea declares a <textarea> HTML element
func Textarea(contents ...any) aitch.Node {
	return aitch.NewElement(tagTextarea, contents...)
}

// TFoot declares a <tfoot> HTML element
func TFoot(contents ...any) aitch.Node {
	return aitch.NewElement(tagTfoot, contents...)
}

// Th declares a <th> HTML element
func Th(contents ...any) aitch.Node {
	return aitch.NewElement(tagTh, contents...)
}

// THead declares a <thead> HTML element
func THead(contents ...any) aitch.Node {
	return aitch.NewElement(tagThead, contents...)
}

// Time declares a <time> HTML element
func Time(contents ...any) aitch.Node {
	return aitch.NewElement(tagTime, contents...)
}

// TitleElement declares a <title> HTML element
func TitleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

// Tr declares a <tr> HTML element
func Tr(contents ...any) aitch.Node {
	return aitch.NewElement(tagTr, contents...)
}

// Ul declares a <ul> HTML element
func Ul(contents ...any) aitch.Node {
	return aitch.NewElement(tagUl, contents...)
}

// U declares a <u> HTML element
func U(contents ...any) aitch.Node {
	return aitch.NewElement(tagU, contents...)
}

// Var declares a <var> HTML element
func Var(contents ...any) aitch.Node {
	return aitch.NewElement(tagVar, contents...)
}

// Video declares a <video> HTML element
func Video(contents ...any) aitch.Node {
	return aitch.NewElement(tagVideo, contents...)
}

// Wbr declares a <wbr> HTML element
//
// As <wbr> is a void element - any non-attribute contents are ignored
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
