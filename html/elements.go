package html

import (
	"github.com/go-andiamo/aitch"
)

// A declares an <a> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func A(contents ...any) aitch.Node {
	return aitch.NewElement(tagA, contents...)
}

// Abbr declares an <abbr> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbr(contents ...any) aitch.Node {
	return aitch.NewElement(tagAbbr, contents...)
}

// Address declares an <address> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(contents ...any) aitch.Node {
	return aitch.NewElement(tagAddress, contents...)
}

// Area declares an <area> HTML element
//
// As <area> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagArea, contents...)
}

// Article declares an <article> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(contents ...any) aitch.Node {
	return aitch.NewElement(tagArticle, contents...)
}

// Aside declares an <aside> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(contents ...any) aitch.Node {
	return aitch.NewElement(tagAside, contents...)
}

// Audio declares an <audio> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(contents ...any) aitch.Node {
	return aitch.NewElement(tagAudio, contents...)
}

// B declares a <b> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func B(contents ...any) aitch.Node {
	return aitch.NewElement(tagB, contents...)
}

// Base declares a <base> HTML element
//
// As <base> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagBase, contents...)
}

// BlockQuote declares a <blockquote> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(contents ...any) aitch.Node {
	return aitch.NewElement(tagBlockquote, contents...)
}

// Body declares a <body> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func Body(contents ...any) aitch.Node {
	return aitch.NewElement(tagBody, contents...)
}

// Br declares a <br> HTML element
//
// As <br> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Br(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagBr, contents...)
}

// Button declares a <button> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(contents ...any) aitch.Node {
	return aitch.NewElement(tagButton, contents...)
}

// Canvas declares a <canvas> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(contents ...any) aitch.Node {
	return aitch.NewElement(tagCanvas, contents...)
}

// Caption declares a <caption> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(contents ...any) aitch.Node {
	return aitch.NewElement(tagCaption, contents...)
}

// Cite declares a <cite> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Cite(contents ...any) aitch.Node {
	return aitch.NewElement(tagCite, contents...)
}

// Code declares a <code> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(contents ...any) aitch.Node {
	return aitch.NewElement(tagCode, contents...)
}

// Col declares a <col> HTML element
//
// As <col> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Col(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagCol, contents...)
}

// ColGroup declares a <colgroup> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagColgroup, contents...)
}

// DataElement declares a <data> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func DataElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagData, contents...)
}

// DataList declares a <datalist> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(contents ...any) aitch.Node {
	return aitch.NewElement(tagDatalist, contents...)
}

// Dd declares a <dd> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Dd(contents ...any) aitch.Node {
	return aitch.NewElement(tagDd, contents...)
}

// Del declares a <del> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func Del(contents ...any) aitch.Node {
	return aitch.NewElement(tagDel, contents...)
}

// Dfn declares a <dfn> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Dfn(contents ...any) aitch.Node {
	return aitch.NewElement(tagDfn, contents...)
}

// Dt declares a <dt> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func Dt(contents ...any) aitch.Node {
	return aitch.NewElement(tagDt, contents...)
}

// Details declares a <details> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(contents ...any) aitch.Node {
	return aitch.NewElement(tagDetails, contents...)
}

// Dialog declares a <dialog> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(contents ...any) aitch.Node {
	return aitch.NewElement(tagDialog, contents...)
}

// Div declares a <div> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(contents ...any) aitch.Node {
	return aitch.NewElement(tagDiv, contents...)
}

// Dl declares a <dl> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func Dl(contents ...any) aitch.Node {
	return aitch.NewElement(tagDl, contents...)
}

// Em declares an <em> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Em(contents ...any) aitch.Node {
	return aitch.NewElement(tagEm, contents...)
}

// Embed declares an <embed> HTML element
//
// As <embed> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagEmbed, contents...)
}

// Form declares a <form> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(contents ...any) aitch.Node {
	return aitch.NewElement(tagForm, contents...)
}

// FieldSet declares a <fieldset> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(contents ...any) aitch.Node {
	return aitch.NewElement(tagFieldset, contents...)
}

// FigCaption declares a <figcaption> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigCaption(contents ...any) aitch.Node {
	return aitch.NewElement(tagFigcaption, contents...)
}

// Figure declares a <figure> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(contents ...any) aitch.Node {
	return aitch.NewElement(tagFigure, contents...)
}

// Footer declares a <footer> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(contents ...any) aitch.Node {
	return aitch.NewElement(tagFooter, contents...)
}

// H1 declares a <h1> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H1(contents ...any) aitch.Node {
	return aitch.NewElement(tagH1, contents...)
}

// H2 declares a <h2> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H2(contents ...any) aitch.Node {
	return aitch.NewElement(tagH2, contents...)
}

// H3 declares a <h3> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H3(contents ...any) aitch.Node {
	return aitch.NewElement(tagH3, contents...)
}

// H4 declares a <h4> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H4(contents ...any) aitch.Node {
	return aitch.NewElement(tagH4, contents...)
}

// H5 declares a <h5> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H5(contents ...any) aitch.Node {
	return aitch.NewElement(tagH5, contents...)
}

// H6 declares a <h6> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func H6(contents ...any) aitch.Node {
	return aitch.NewElement(tagH6, contents...)
}

// Head declares a <head> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
func Head(contents ...any) aitch.Node {
	return aitch.NewElement(tagHead, contents...)
}

// Header declares a <header> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(contents ...any) aitch.Node {
	return aitch.NewElement(tagHeader, contents...)
}

// HGroup declares a <hgroup> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagHgroup, contents...)
}

// Hr declares a <hr> HTML element
//
// As <hr> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func Hr(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagHr, contents...)
}

// Html declares a <html> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html
func Html(contents ...any) aitch.Node {
	return aitch.NewElement(tagHtml, contents...)
}

// I declares an <i> HTML element (idiomatic text)
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func I(contents ...any) aitch.Node {
	return aitch.NewElement(tagI, contents...)
}

// IFrame declares an <iframe> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func IFrame(contents ...any) aitch.Node {
	return aitch.NewElement(tagIframe, contents...)
}

// Img declares an <img> HTML element
//
// As <img> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Img(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagImg, contents...)
}

// Input declares an <input> HTML element
//
// As <input> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagInput, contents...)
}

// Ins declares an <ins> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func Ins(contents ...any) aitch.Node {
	return aitch.NewElement(tagIns, contents...)
}

// Kbd declares a <kbd> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func Kbd(contents ...any) aitch.Node {
	return aitch.NewElement(tagKbd, contents...)
}

// Label declares a <label> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(contents ...any) aitch.Node {
	return aitch.NewElement(tagLabel, contents...)
}

// Legend declares a <legend> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(contents ...any) aitch.Node {
	return aitch.NewElement(tagLegend, contents...)
}

// Li declares a <li> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func Li(contents ...any) aitch.Node {
	return aitch.NewElement(tagLi, contents...)
}

// Link declares a <link> HTML element
//
// As <link> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagLink, contents...)
}

// Main declares a <main> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(contents ...any) aitch.Node {
	return aitch.NewElement(tagMain, contents...)
}

// Mark declares a <mark> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(contents ...any) aitch.Node {
	return aitch.NewElement(tagMark, contents...)
}

// Menu declares a <menu> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(contents ...any) aitch.Node {
	return aitch.NewElement(tagMenu, contents...)
}

// Meta declares a <meta> HTML element
//
// As <meta> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagMeta, contents...)
}

// Meter declares a <meter> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(contents ...any) aitch.Node {
	return aitch.NewElement(tagMeter, contents...)
}

// Nav declares a <nav> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Nav(contents ...any) aitch.Node {
	return aitch.NewElement(tagNav, contents...)
}

// NoScript declares a <noscript> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(contents ...any) aitch.Node {
	return aitch.NewElement(tagNoscript, contents...)
}

// Object declares an <object> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(contents ...any) aitch.Node {
	return aitch.NewElement(tagObject, contents...)
}

// Ol declares an <ol> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func Ol(contents ...any) aitch.Node {
	return aitch.NewElement(tagOl, contents...)
}

// OptGroup declares an <optgroup> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptGroup(contents ...any) aitch.Node {
	return aitch.NewElement(tagOptgroup, contents...)
}

// Option declares an <option> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(contents ...any) aitch.Node {
	return aitch.NewElement(tagOption, contents...)
}

// P declares a <p> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func P(contents ...any) aitch.Node {
	return aitch.NewElement(tagP, contents...)
}

// Para declares a <p> HTML element (same as P)
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Para(contents ...any) aitch.Node {
	return aitch.NewElement(tagP, contents...)
}

// Param declares a <param> HTML element
//
// As <param> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Param(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagParam, contents...)
}

// Picture declares a <picture> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(contents ...any) aitch.Node {
	return aitch.NewElement(tagPicture, contents...)
}

// Pre declares a <pre> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Pre(contents ...any) aitch.Node {
	return aitch.NewElement(tagPre, contents...)
}

// Progress declares a <progress> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(contents ...any) aitch.Node {
	return aitch.NewElement(tagProgress, contents...)
}

// Q declares a <q> HTML element (inline quotation)
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Q(contents ...any) aitch.Node {
	return aitch.NewElement(tagQ, contents...)
}

// S declares an <s> HTML element (strikethrough)
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func S(contents ...any) aitch.Node {
	return aitch.NewElement(tagS, contents...)
}

// Samp declares a <samp> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Samp(contents ...any) aitch.Node {
	return aitch.NewElement(tagSamp, contents...)
}

// Script declares a <script> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(contents ...any) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

// Section declares a <section> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(contents ...any) aitch.Node {
	return aitch.NewElement(tagSection, contents...)
}

// Select declares a <select> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(contents ...any) aitch.Node {
	return aitch.NewElement(tagSelect, contents...)
}

// Slot declares a <slot> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot
func Slot(contents ...any) aitch.Node {
	return aitch.NewElement(tagSlot, contents...)
}

// Small declares a <small> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(contents ...any) aitch.Node {
	return aitch.NewElement(tagSmall, contents...)
}

// Source declares a <source> HTML element
//
// As <source> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagSource, contents...)
}

// Span declares a <span> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(contents ...any) aitch.Node {
	return aitch.NewElement(tagSpan, contents...)
}

// Strong declares a <strong> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(contents ...any) aitch.Node {
	return aitch.NewElement(tagStrong, contents...)
}

// StyleElement declares a <style> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func StyleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

// Sub declares a <sub> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Sub(contents ...any) aitch.Node {
	return aitch.NewElement(tagSub, contents...)
}

// Sup declares a <sup> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Sup(contents ...any) aitch.Node {
	return aitch.NewElement(tagSup, contents...)
}

// Summary declares a <summary> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(contents ...any) aitch.Node {
	return aitch.NewElement(tagSummary, contents...)
}

// Table declares a <table> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(contents ...any) aitch.Node {
	return aitch.NewElement(tagTable, contents...)
}

// TBody declares a <tbody> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TBody(contents ...any) aitch.Node {
	return aitch.NewElement(tagTbody, contents...)
}

// Td declares a <td> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func Td(contents ...any) aitch.Node {
	return aitch.NewElement(tagTd, contents...)
}

// Template declares a <template> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(contents ...any) aitch.Node {
	return aitch.NewElement(tagTemplate, contents...)
}

// Textarea declares a <textarea> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func Textarea(contents ...any) aitch.Node {
	return aitch.NewElement(tagTextarea, contents...)
}

// TFoot declares a <tfoot> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TFoot(contents ...any) aitch.Node {
	return aitch.NewElement(tagTfoot, contents...)
}

// Th declares a <th> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func Th(contents ...any) aitch.Node {
	return aitch.NewElement(tagTh, contents...)
}

// THead declares a <thead> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func THead(contents ...any) aitch.Node {
	return aitch.NewElement(tagThead, contents...)
}

// Time declares a <time> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(contents ...any) aitch.Node {
	return aitch.NewElement(tagTime, contents...)
}

// TitleElement declares a <title> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func TitleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

// Tr declares a <tr> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func Tr(contents ...any) aitch.Node {
	return aitch.NewElement(tagTr, contents...)
}

// Ul declares a <ul> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func Ul(contents ...any) aitch.Node {
	return aitch.NewElement(tagUl, contents...)
}

// U declares a <u> HTML element (underline)
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func U(contents ...any) aitch.Node {
	return aitch.NewElement(tagU, contents...)
}

// Var declares a <var> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Var(contents ...any) aitch.Node {
	return aitch.NewElement(tagVar, contents...)
}

// Video declares a <video> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(contents ...any) aitch.Node {
	return aitch.NewElement(tagVideo, contents...)
}

// Wbr declares a <wbr> HTML element
//
// As <wbr> is a void element - any non-attribute contents are ignored
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func Wbr(contents ...any) aitch.Node {
	return aitch.NewVoidElement(tagWbr, contents...)
}

// Bdi declares a <bdi> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func Bdi(contents ...any) aitch.Node {
	return aitch.NewElement(tagBdi, contents...)
}

// Bdo declares a <bdo> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func Bdo(contents ...any) aitch.Node {
	return aitch.NewElement(tagBdo, contents...)
}

// Map declares a <map> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(contents ...any) aitch.Node {
	return aitch.NewElement(tagMap, contents...)
}

// Output declares a <output> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(contents ...any) aitch.Node {
	return aitch.NewElement(tagOutput, contents...)
}

// Rp declares a <rp> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func Rp(contents ...any) aitch.Node {
	return aitch.NewElement(tagRp, contents...)
}

// Rt declares a <rt> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func Rt(contents ...any) aitch.Node {
	return aitch.NewElement(tagRt, contents...)
}

// Ruby declares a <ruby> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(contents ...any) aitch.Node {
	return aitch.NewElement(tagRuby, contents...)
}

// Search declares a <search> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/search
func Search(contents ...any) aitch.Node {
	return aitch.NewElement(tagSearch, contents...)
}

// Track declares a <track> HTML element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(contents ...any) aitch.Node {
	return aitch.NewElement(tagTrack, contents...)
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
	tagBdi        = []byte("bdi")
	tagBdo        = []byte("bdo")
	tagMap        = []byte("map")
	tagOutput     = []byte("output")
	tagRp         = []byte("rp")
	tagRt         = []byte("rt")
	tagRuby       = []byte("ruby")
	tagSearch     = []byte("search")
	tagTrack      = []byte("track")
)
