package html

import (
	"bytes"
	"github.com/go-andiamo/aitch"
	"github.com/stretchr/testify/require"
	"testing"
)

func testRender(node aitch.Node, t *testing.T) string {
	var w bytes.Buffer
	err := node.Render(&w, nil)
	require.NoError(t, err)
	return w.String()
}

func TestA(t *testing.T) {
	el := A(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "a", el.Name())
	require.Equal(t, `<a id="1">foo</a>`, testRender(el, t))
}

func TestAbbr(t *testing.T) {
	el := Abbr(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "abbr", el.Name())
	require.Equal(t, `<abbr id="1">foo</abbr>`, testRender(el, t))
}

func TestB(t *testing.T) {
	el := B(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "b", el.Name())
	require.Equal(t, `<b id="1">foo</b>`, testRender(el, t))
}

func TestCaption(t *testing.T) {
	el := Caption(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "caption", el.Name())
	require.Equal(t, `<caption id="1">foo</caption>`, testRender(el, t))
}

func TestDd(t *testing.T) {
	el := Dd(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "dd", el.Name())
	require.Equal(t, `<dd id="1">foo</dd>`, testRender(el, t))
}

func TestDel(t *testing.T) {
	el := Del(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "del", el.Name())
	require.Equal(t, `<del id="1">foo</del>`, testRender(el, t))
}

func TestDfn(t *testing.T) {
	el := Dfn(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "dfn", el.Name())
	require.Equal(t, `<dfn id="1">foo</dfn>`, testRender(el, t))
}

func TestDt(t *testing.T) {
	el := Dt(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "dt", el.Name())
	require.Equal(t, `<dt id="1">foo</dt>`, testRender(el, t))
}

func TestEm(t *testing.T) {
	el := Em(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "em", el.Name())
	require.Equal(t, `<em id="1">foo</em>`, testRender(el, t))
}

func TestFigCaption(t *testing.T) {
	el := FigCaption(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "figcaption", el.Name())
	require.Equal(t, `<figcaption id="1">foo</figcaption>`, testRender(el, t))
}

func TestH1(t *testing.T) {
	el := H1(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "h1", el.Name())
	require.Equal(t, `<h1 id="1">foo</h1>`, testRender(el, t))
}

func TestH2(t *testing.T) {
	el := H2(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "h2", el.Name())
	require.Equal(t, `<h2 id="1">foo</h2>`, testRender(el, t))
}

func TestH3(t *testing.T) {
	el := H3(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "h3", el.Name())
	require.Equal(t, `<h3 id="1">foo</h3>`, testRender(el, t))
}

func TestH4(t *testing.T) {
	el := H4(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "h4", el.Name())
	require.Equal(t, `<h4 id="1">foo</h4>`, testRender(el, t))
}

func TestH5(t *testing.T) {
	el := H5(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "h5", el.Name())
	require.Equal(t, `<h5 id="1">foo</h5>`, testRender(el, t))
}

func TestH6(t *testing.T) {
	el := H6(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "h6", el.Name())
	require.Equal(t, `<h6 id="1">foo</h6>`, testRender(el, t))
}

func TestI(t *testing.T) {
	el := I(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "i", el.Name())
	require.Equal(t, `<i id="1">foo</i>`, testRender(el, t))
}

func TestIns(t *testing.T) {
	el := Ins(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "ins", el.Name())
	require.Equal(t, `<ins id="1">foo</ins>`, testRender(el, t))
}

func TestKbd(t *testing.T) {
	el := Kbd(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "kbd", el.Name())
	require.Equal(t, `<kbd id="1">foo</kbd>`, testRender(el, t))
}

func TestAddress(t *testing.T) {
	el := Address(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "address", el.Name())
	require.Equal(t, `<address id="1">foo</address>`, testRender(el, t))
}

func TestArea(t *testing.T) {
	el := Area(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "area", el.Name())
	require.Equal(t, `<area id="1">`, testRender(el, t))
}

func TestArticle(t *testing.T) {
	el := Article(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "article", el.Name())
	require.Equal(t, `<article id="1">foo</article>`, testRender(el, t))
}

func TestAside(t *testing.T) {
	el := Aside(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "aside", el.Name())
	require.Equal(t, `<aside id="1">foo</aside>`, testRender(el, t))
}

func TestAudio(t *testing.T) {
	el := Audio(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "audio", el.Name())
	require.Equal(t, `<audio id="1">foo</audio>`, testRender(el, t))
}

func TestBase(t *testing.T) {
	el := Base(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "base", el.Name())
	require.Equal(t, `<base id="1">`, testRender(el, t))
}

func TestBlockQuote(t *testing.T) {
	el := BlockQuote(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "blockquote", el.Name())
	require.Equal(t, `<blockquote id="1">foo</blockquote>`, testRender(el, t))
}

func TestBody(t *testing.T) {
	el := Body(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "body", el.Name())
	require.Equal(t, `<body id="1">foo</body>`, testRender(el, t))
}

func TestBr(t *testing.T) {
	el := Br(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "br", el.Name())
	require.Equal(t, `<br id="1">`, testRender(el, t))
}

func TestButton(t *testing.T) {
	el := Button(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "button", el.Name())
	require.Equal(t, `<button id="1">foo</button>`, testRender(el, t))
}

func TestCanvas(t *testing.T) {
	el := Canvas(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "canvas", el.Name())
	require.Equal(t, `<canvas id="1">foo</canvas>`, testRender(el, t))
}

func TestCite(t *testing.T) {
	el := Cite(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "cite", el.Name())
	require.Equal(t, `<cite id="1">foo</cite>`, testRender(el, t))
}

func TestCode(t *testing.T) {
	el := Code(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "code", el.Name())
	require.Equal(t, `<code id="1">foo</code>`, testRender(el, t))
}

func TestCol(t *testing.T) {
	el := Col(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "col", el.Name())
	require.Equal(t, `<col id="1">`, testRender(el, t))
}

func TestColGroup(t *testing.T) {
	el := ColGroup(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "colgroup", el.Name())
	require.Equal(t, `<colgroup id="1">foo</colgroup>`, testRender(el, t))
}

func TestDataElement(t *testing.T) {
	el := DataElement(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "data", el.Name())
	require.Equal(t, `<data id="1">foo</data>`, testRender(el, t))
}

func TestDataList(t *testing.T) {
	el := DataList(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "datalist", el.Name())
	require.Equal(t, `<datalist id="1">foo</datalist>`, testRender(el, t))
}

func TestDetails(t *testing.T) {
	el := Details(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "details", el.Name())
	require.Equal(t, `<details id="1">foo</details>`, testRender(el, t))
}

func TestDialog(t *testing.T) {
	el := Dialog(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "dialog", el.Name())
	require.Equal(t, `<dialog id="1">foo</dialog>`, testRender(el, t))
}

func TestDiv(t *testing.T) {
	el := Div(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "div", el.Name())
	require.Equal(t, `<div id="1">foo</div>`, testRender(el, t))
}

func TestDl(t *testing.T) {
	el := Dl(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "dl", el.Name())
	require.Equal(t, `<dl id="1">foo</dl>`, testRender(el, t))
}

func TestEmbed(t *testing.T) {
	el := Embed(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "embed", el.Name())
	require.Equal(t, `<embed id="1">`, testRender(el, t))
}

func TestForm(t *testing.T) {
	el := Form(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "form", el.Name())
	require.Equal(t, `<form id="1">foo</form>`, testRender(el, t))
}

func TestFieldSet(t *testing.T) {
	el := FieldSet(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "fieldset", el.Name())
	require.Equal(t, `<fieldset id="1">foo</fieldset>`, testRender(el, t))
}

func TestFigure(t *testing.T) {
	el := Figure(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "figure", el.Name())
	require.Equal(t, `<figure id="1">foo</figure>`, testRender(el, t))
}

func TestFooter(t *testing.T) {
	el := Footer(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "footer", el.Name())
	require.Equal(t, `<footer id="1">foo</footer>`, testRender(el, t))
}

func TestHead(t *testing.T) {
	el := Head(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "head", el.Name())
	require.Equal(t, `<head id="1">foo</head>`, testRender(el, t))
}

func TestHeader(t *testing.T) {
	el := Header(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "header", el.Name())
	require.Equal(t, `<header id="1">foo</header>`, testRender(el, t))
}

func TestHGroup(t *testing.T) {
	el := HGroup(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "hgroup", el.Name())
	require.Equal(t, `<hgroup id="1">foo</hgroup>`, testRender(el, t))
}

func TestHr(t *testing.T) {
	el := Hr(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "hr", el.Name())
	require.Equal(t, `<hr id="1">`, testRender(el, t))
}

func TestHtml(t *testing.T) {
	el := Html(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "html", el.Name())
	require.Equal(t, `<html id="1">foo</html>`, testRender(el, t))
}

func TestIFrame(t *testing.T) {
	el := IFrame(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "iframe", el.Name())
	require.Equal(t, `<iframe id="1">foo</iframe>`, testRender(el, t))
}

func TestImg(t *testing.T) {
	el := Img(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "img", el.Name())
	require.Equal(t, `<img id="1">`, testRender(el, t))
}

func TestInput(t *testing.T) {
	el := Input(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "input", el.Name())
	require.Equal(t, `<input id="1">`, testRender(el, t))
}

func TestLabel(t *testing.T) {
	el := Label(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "label", el.Name())
	require.Equal(t, `<label id="1">foo</label>`, testRender(el, t))
}

func TestLegend(t *testing.T) {
	el := Legend(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "legend", el.Name())
	require.Equal(t, `<legend id="1">foo</legend>`, testRender(el, t))
}

func TestLi(t *testing.T) {
	el := Li(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "li", el.Name())
	require.Equal(t, `<li id="1">foo</li>`, testRender(el, t))
}

func TestLink(t *testing.T) {
	el := Link(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "link", el.Name())
	require.Equal(t, `<link id="1">`, testRender(el, t))
}

func TestMainElement(t *testing.T) {
	el := Main(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "main", el.Name())
	require.Equal(t, `<main id="1">foo</main>`, testRender(el, t))
}

func TestMark(t *testing.T) {
	el := Mark(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "mark", el.Name())
	require.Equal(t, `<mark id="1">foo</mark>`, testRender(el, t))
}

func TestMenu(t *testing.T) {
	el := Menu(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "menu", el.Name())
	require.Equal(t, `<menu id="1">foo</menu>`, testRender(el, t))
}

func TestMeta(t *testing.T) {
	el := Meta(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "meta", el.Name())
	require.Equal(t, `<meta id="1">`, testRender(el, t))
}

func TestMeter(t *testing.T) {
	el := Meter(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "meter", el.Name())
	require.Equal(t, `<meter id="1">foo</meter>`, testRender(el, t))
}

func TestNav(t *testing.T) {
	el := Nav(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "nav", el.Name())
	require.Equal(t, `<nav id="1">foo</nav>`, testRender(el, t))
}

func TestNoScript(t *testing.T) {
	el := NoScript(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "noscript", el.Name())
	require.Equal(t, `<noscript id="1">foo</noscript>`, testRender(el, t))
}

func TestObject(t *testing.T) {
	el := Object(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "object", el.Name())
	require.Equal(t, `<object id="1">foo</object>`, testRender(el, t))
}

func TestOl(t *testing.T) {
	el := Ol(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "ol", el.Name())
	require.Equal(t, `<ol id="1">foo</ol>`, testRender(el, t))
}

func TestOptGroup(t *testing.T) {
	el := OptGroup(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "optgroup", el.Name())
	require.Equal(t, `<optgroup id="1">foo</optgroup>`, testRender(el, t))
}

func TestOption(t *testing.T) {
	el := Option(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "option", el.Name())
	require.Equal(t, `<option id="1">foo</option>`, testRender(el, t))
}

func TestP(t *testing.T) {
	el := P(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "p", el.Name())
	require.Equal(t, `<p id="1">foo</p>`, testRender(el, t))
}

func TestPara(t *testing.T) {
	el := Para(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "p", el.Name())
	require.Equal(t, `<p id="1">foo</p>`, testRender(el, t))
}

func TestParam(t *testing.T) {
	el := Param(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "param", el.Name())
	require.Equal(t, `<param id="1">`, testRender(el, t))
}

func TestPicture(t *testing.T) {
	el := Picture(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "picture", el.Name())
	require.Equal(t, `<picture id="1">foo</picture>`, testRender(el, t))
}

func TestPre(t *testing.T) {
	el := Pre(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "pre", el.Name())
	require.Equal(t, `<pre id="1">foo</pre>`, testRender(el, t))
}

func TestProgress(t *testing.T) {
	el := Progress(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "progress", el.Name())
	require.Equal(t, `<progress id="1">foo</progress>`, testRender(el, t))
}

func TestScript(t *testing.T) {
	el := Script(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "script", el.Name())
	require.Equal(t, `<script id="1">foo</script>`, testRender(el, t))
}

func TestSection(t *testing.T) {
	el := Section(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "section", el.Name())
	require.Equal(t, `<section id="1">foo</section>`, testRender(el, t))
}

func TestSelect(t *testing.T) {
	el := Select(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "select", el.Name())
	require.Equal(t, `<select id="1">foo</select>`, testRender(el, t))
}

func TestSlot(t *testing.T) {
	el := Slot(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "slot", el.Name())
	require.Equal(t, `<slot id="1">foo</slot>`, testRender(el, t))
}

func TestSource(t *testing.T) {
	el := Source(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "source", el.Name())
	require.Equal(t, `<source id="1">`, testRender(el, t))
}

func TestSpan(t *testing.T) {
	el := Span(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "span", el.Name())
	require.Equal(t, `<span id="1">foo</span>`, testRender(el, t))
}

func TestStyleElement(t *testing.T) {
	el := StyleElement(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "style", el.Name())
	require.Equal(t, `<style id="1">foo</style>`, testRender(el, t))
}

func TestSummary(t *testing.T) {
	el := Summary(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "summary", el.Name())
	require.Equal(t, `<summary id="1">foo</summary>`, testRender(el, t))
}

func TestTable(t *testing.T) {
	el := Table(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "table", el.Name())
	require.Equal(t, `<table id="1">foo</table>`, testRender(el, t))
}

func TestTBody(t *testing.T) {
	el := TBody(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "tbody", el.Name())
	require.Equal(t, `<tbody id="1">foo</tbody>`, testRender(el, t))
}

func TestTd(t *testing.T) {
	el := Td(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "td", el.Name())
	require.Equal(t, `<td id="1">foo</td>`, testRender(el, t))
}

func TestTemplate(t *testing.T) {
	el := Template(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "template", el.Name())
	require.Equal(t, `<template id="1">foo</template>`, testRender(el, t))
}

func TestTextarea(t *testing.T) {
	el := Textarea(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "textarea", el.Name())
	require.Equal(t, `<textarea id="1">foo</textarea>`, testRender(el, t))
}

func TestTFoot(t *testing.T) {
	el := TFoot(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "tfoot", el.Name())
	require.Equal(t, `<tfoot id="1">foo</tfoot>`, testRender(el, t))
}

func TestTh(t *testing.T) {
	el := Th(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "th", el.Name())
	require.Equal(t, `<th id="1">foo</th>`, testRender(el, t))
}

func TestTHead(t *testing.T) {
	el := THead(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "thead", el.Name())
	require.Equal(t, `<thead id="1">foo</thead>`, testRender(el, t))
}

func TestTr(t *testing.T) {
	el := Tr(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "tr", el.Name())
	require.Equal(t, `<tr id="1">foo</tr>`, testRender(el, t))
}

func TestUl(t *testing.T) {
	el := Ul(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "ul", el.Name())
	require.Equal(t, `<ul id="1">foo</ul>`, testRender(el, t))
}

func TestWbr(t *testing.T) {
	el := Wbr(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "wbr", el.Name())
	require.Equal(t, `<wbr id="1">`, testRender(el, t))
}

func TestQ(t *testing.T) {
	el := Q(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "q", el.Name())
	require.Equal(t, `<q id="1">foo</q>`, testRender(el, t))
}

func TestS(t *testing.T) {
	el := S(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "s", el.Name())
	require.Equal(t, `<s id="1">foo</s>`, testRender(el, t))
}

func TestSamp(t *testing.T) {
	el := Samp(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "samp", el.Name())
	require.Equal(t, `<samp id="1">foo</samp>`, testRender(el, t))
}

func TestSmall(t *testing.T) {
	el := Small(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "small", el.Name())
	require.Equal(t, `<small id="1">foo</small>`, testRender(el, t))
}

func TestStrong(t *testing.T) {
	el := Strong(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "strong", el.Name())
	require.Equal(t, `<strong id="1">foo</strong>`, testRender(el, t))
}

func TestSub(t *testing.T) {
	el := Sub(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "sub", el.Name())
	require.Equal(t, `<sub id="1">foo</sub>`, testRender(el, t))
}

func TestSup(t *testing.T) {
	el := Sup(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "sup", el.Name())
	require.Equal(t, `<sup id="1">foo</sup>`, testRender(el, t))
}

func TestTime(t *testing.T) {
	el := Time(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "time", el.Name())
	require.Equal(t, `<time id="1">foo</time>`, testRender(el, t))
}

func TestTitleElement(t *testing.T) {
	el := TitleElement(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "title", el.Name())
	require.Equal(t, `<title id="1">foo</title>`, testRender(el, t))
}

func TestU(t *testing.T) {
	el := U(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "u", el.Name())
	require.Equal(t, `<u id="1">foo</u>`, testRender(el, t))
}

func TestVar(t *testing.T) {
	el := Var(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "var", el.Name())
	require.Equal(t, `<var id="1">foo</var>`, testRender(el, t))
}

func TestVideo(t *testing.T) {
	el := Video(Id(1), Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "video", el.Name())
	require.Equal(t, `<video id="1">foo</video>`, testRender(el, t))
}
