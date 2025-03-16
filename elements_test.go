package aitch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestA(t *testing.T) {
	el := A(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "a", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestAbbr(t *testing.T) {
	el := Abbr(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "abbr", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestB(t *testing.T) {
	el := B(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "b", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestCaption(t *testing.T) {
	el := Caption(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "caption", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDd(t *testing.T) {
	el := Dd(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "dd", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDel(t *testing.T) {
	el := Del(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "del", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDfn(t *testing.T) {
	el := Dfn(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "dfn", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDt(t *testing.T) {
	el := Dt(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "dt", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestEm(t *testing.T) {
	el := Em(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "em", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestFigCaption(t *testing.T) {
	el := FigCaption(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "figcaption", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestH1(t *testing.T) {
	el := H1(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "h1", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestH2(t *testing.T) {
	el := H2(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "h2", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestH3(t *testing.T) {
	el := H3(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "h3", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestH4(t *testing.T) {
	el := H4(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "h4", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestH5(t *testing.T) {
	el := H5(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "h5", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestH6(t *testing.T) {
	el := H6(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "h6", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestI(t *testing.T) {
	el := I(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "i", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestIns(t *testing.T) {
	el := Ins(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "ins", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestKbd(t *testing.T) {
	el := Kbd(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "kbd", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestAddress(t *testing.T) {
	el := Address(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "address", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestArea(t *testing.T) {
	el := Area(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "area", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestArticle(t *testing.T) {
	el := Article(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "article", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestAside(t *testing.T) {
	el := Aside(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "aside", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestAudio(t *testing.T) {
	el := Audio(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "audio", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestBase(t *testing.T) {
	el := Base(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "base", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestBlockQuote(t *testing.T) {
	el := BlockQuote(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "blockquote", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestBody(t *testing.T) {
	el := Body(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "body", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestBr(t *testing.T) {
	el := Br(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "br", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestButton(t *testing.T) {
	el := Button(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "button", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestCanvas(t *testing.T) {
	el := Canvas(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "canvas", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestCite(t *testing.T) {
	el := Cite(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "cite", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestCode(t *testing.T) {
	el := Code(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "code", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestCol(t *testing.T) {
	el := Col(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "col", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestColGroup(t *testing.T) {
	el := ColGroup(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "colgroup", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDataElement(t *testing.T) {
	el := DataElement(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "data", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDataList(t *testing.T) {
	el := DataList(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "datalist", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDetails(t *testing.T) {
	el := Details(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "details", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDialog(t *testing.T) {
	el := Dialog(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "dialog", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDiv(t *testing.T) {
	el := Div(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "div", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestDl(t *testing.T) {
	el := Dl(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "dl", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestEmbed(t *testing.T) {
	el := Embed(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "embed", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestForm(t *testing.T) {
	el := Form(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "form", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestFieldSet(t *testing.T) {
	el := FieldSet(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "fieldset", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestFigure(t *testing.T) {
	el := Figure(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "figure", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestFooter(t *testing.T) {
	el := Footer(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "footer", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestHead(t *testing.T) {
	el := Head(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "head", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestHeader(t *testing.T) {
	el := Header(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "header", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestHGroup(t *testing.T) {
	el := HGroup(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "hgroup", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestHr(t *testing.T) {
	el := Hr(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "hr", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestHtml(t *testing.T) {
	el := Html(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "html", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestIFrame(t *testing.T) {
	el := IFrame(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "iframe", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestImg(t *testing.T) {
	el := Img(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "img", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestInput(t *testing.T) {
	el := Input(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "input", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestLabel(t *testing.T) {
	el := Label(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "label", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestLegend(t *testing.T) {
	el := Legend(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "legend", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestLi(t *testing.T) {
	el := Li(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "li", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestLink(t *testing.T) {
	el := Link(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "link", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestMainElement(t *testing.T) {
	el := Main(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "main", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestMenu(t *testing.T) {
	el := Menu(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "menu", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestMeta(t *testing.T) {
	el := Meta(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "meta", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestMeter(t *testing.T) {
	el := Meter(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "meter", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestNav(t *testing.T) {
	el := Nav(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "nav", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestNoScript(t *testing.T) {
	el := NoScript(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "noscript", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestObject(t *testing.T) {
	el := Object(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "object", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestOl(t *testing.T) {
	el := Ol(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "ol", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestOptGroup(t *testing.T) {
	el := OptGroup(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "optgroup", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestOption(t *testing.T) {
	el := Option(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "option", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestP(t *testing.T) {
	el := P(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "p", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestPara(t *testing.T) {
	el := Para(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "p", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestParam(t *testing.T) {
	el := Param(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "param", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestPicture(t *testing.T) {
	el := Picture(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "picture", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestPre(t *testing.T) {
	el := Pre(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "pre", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestProgress(t *testing.T) {
	el := Progress(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "progress", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestScript(t *testing.T) {
	el := Script(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "script", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSection(t *testing.T) {
	el := Section(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "section", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSelect(t *testing.T) {
	el := Select(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "select", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSlot(t *testing.T) {
	el := Slot(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "slot", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSource(t *testing.T) {
	el := Source(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "source", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestSpan(t *testing.T) {
	el := Span(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "span", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestStyleElement(t *testing.T) {
	el := StyleElement(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "style", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSummary(t *testing.T) {
	el := Summary(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "summary", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSvg(t *testing.T) {
	el := Svg(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "svg", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTable(t *testing.T) {
	el := Table(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "table", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTBody(t *testing.T) {
	el := TBody(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "tbody", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTd(t *testing.T) {
	el := Td(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "td", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTemplate(t *testing.T) {
	el := Template(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "template", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTextarea(t *testing.T) {
	el := Textarea(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "textarea", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTFoot(t *testing.T) {
	el := TFoot(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "tfoot", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTh(t *testing.T) {
	el := Th(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "th", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTHead(t *testing.T) {
	el := THead(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "thead", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTr(t *testing.T) {
	el := Tr(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "tr", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestUl(t *testing.T) {
	el := Ul(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "ul", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestWbr(t *testing.T) {
	el := Wbr(Id(1), Text("foo"))
	require.IsType(t, &voidElement{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "wbr", el.Name())
	require.Equal(t, 1, len(el.(*voidElement).attributes))
}

func TestMark(t *testing.T) {
	el := Mark(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "mark", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestQ(t *testing.T) {
	el := Q(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "q", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestS(t *testing.T) {
	el := S(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "s", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSamp(t *testing.T) {
	el := Samp(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "samp", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSmall(t *testing.T) {
	el := Small(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "small", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestStrong(t *testing.T) {
	el := Strong(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "strong", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSub(t *testing.T) {
	el := Sub(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "sub", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestSup(t *testing.T) {
	el := Sup(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "sup", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTime(t *testing.T) {
	el := Time(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "time", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestTitleElement(t *testing.T) {
	el := TitleElement(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "title", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestU(t *testing.T) {
	el := U(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "u", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestVar(t *testing.T) {
	el := Var(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "var", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}

func TestVideo(t *testing.T) {
	el := Video(Id(1), Text("foo"))
	require.IsType(t, &element{}, el)
	require.Equal(t, ElementNode, el.Type())
	require.Equal(t, "video", el.Name())
	require.Equal(t, 1, len(el.(*element).attributes))
	require.Equal(t, 1, len(el.(*element).contents))
}
