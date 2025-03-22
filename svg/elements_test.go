package svg

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

func TestSvg(t *testing.T) {
	el := Svg(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "svg", el.Name())
	require.Equal(t, `<svg>foo</svg>`, testRender(el, t))
}

func TestAltGlyph(t *testing.T) {
	el := AltGlyph(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "altGlyph", el.Name())
	require.Equal(t, `<altGlyph>foo</altGlyph>`, testRender(el, t))
}

func TestAnimate(t *testing.T) {
	el := Animate(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "animate", el.Name())
	require.Equal(t, `<animate>foo</animate>`, testRender(el, t))
}

func TestAnimateMotion(t *testing.T) {
	el := AnimateMotion(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "animateMotion", el.Name())
	require.Equal(t, `<animateMotion>foo</animateMotion>`, testRender(el, t))
}

func TestAnimateTransform(t *testing.T) {
	el := AnimateTransform(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "animateTransform", el.Name())
	require.Equal(t, `<animateTransform>foo</animateTransform>`, testRender(el, t))
}

func TestCircle(t *testing.T) {
	el := Circle(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "circle", el.Name())
	require.Equal(t, `<circle>foo</circle>`, testRender(el, t))
}

func TestClipPath(t *testing.T) {
	el := ClipPath(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "clipPath", el.Name())
	require.Equal(t, `<clipPath>foo</clipPath>`, testRender(el, t))
}

func TestDefs(t *testing.T) {
	el := Defs(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "defs", el.Name())
	require.Equal(t, `<defs>foo</defs>`, testRender(el, t))
}

func TestDesc(t *testing.T) {
	el := Desc(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "desc", el.Name())
	require.Equal(t, `<desc>foo</desc>`, testRender(el, t))
}

func TestDiscard(t *testing.T) {
	el := Discard(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "discard", el.Name())
	require.Equal(t, `<discard>foo</discard>`, testRender(el, t))
}

func TestEllipse(t *testing.T) {
	el := Ellipse(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "ellipse", el.Name())
	require.Equal(t, `<ellipse>foo</ellipse>`, testRender(el, t))
}

func TestFeBlend(t *testing.T) {
	el := FeBlend(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feBlend", el.Name())
	require.Equal(t, `<feBlend>foo</feBlend>`, testRender(el, t))
}

func TestFeColorMatrix(t *testing.T) {
	el := FeColorMatrix(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feColorMatrix", el.Name())
	require.Equal(t, `<feColorMatrix>foo</feColorMatrix>`, testRender(el, t))
}

func TestFeComponentTransfer(t *testing.T) {
	el := FeComponentTransfer(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feComponentTransfer", el.Name())
	require.Equal(t, `<feComponentTransfer>foo</feComponentTransfer>`, testRender(el, t))
}

func TestFeComposite(t *testing.T) {
	el := FeComposite(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feComposite", el.Name())
	require.Equal(t, `<feComposite>foo</feComposite>`, testRender(el, t))
}

func TestFeConvolveMatrix(t *testing.T) {
	el := FeConvolveMatrix(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feConvolveMatrix", el.Name())
	require.Equal(t, `<feConvolveMatrix>foo</feConvolveMatrix>`, testRender(el, t))
}

func TestFeDiffuseLighting(t *testing.T) {
	el := FeDiffuseLighting(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feDiffuseLighting", el.Name())
	require.Equal(t, `<feDiffuseLighting>foo</feDiffuseLighting>`, testRender(el, t))
}

func TestFeDisplacementMap(t *testing.T) {
	el := FeDisplacementMap(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feDisplacementMap", el.Name())
	require.Equal(t, `<feDisplacementMap>foo</feDisplacementMap>`, testRender(el, t))
}

func TestFeDistantLight(t *testing.T) {
	el := FeDistantLight(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feDistantLight", el.Name())
	require.Equal(t, `<feDistantLight>foo</feDistantLight>`, testRender(el, t))
}

func TestFeDropShadow(t *testing.T) {
	el := FeDropShadow(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feDropShadow", el.Name())
	require.Equal(t, `<feDropShadow>foo</feDropShadow>`, testRender(el, t))
}

func TestFeFlood(t *testing.T) {
	el := FeFlood(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feFlood", el.Name())
	require.Equal(t, `<feFlood>foo</feFlood>`, testRender(el, t))
}

func TestFeFuncA(t *testing.T) {
	el := FeFuncA(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feFuncA", el.Name())
	require.Equal(t, `<feFuncA>foo</feFuncA>`, testRender(el, t))
}

func TestFeFuncB(t *testing.T) {
	el := FeFuncB(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feFuncB", el.Name())
	require.Equal(t, `<feFuncB>foo</feFuncB>`, testRender(el, t))
}

func TestFeFuncG(t *testing.T) {
	el := FeFuncG(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feFuncG", el.Name())
	require.Equal(t, `<feFuncG>foo</feFuncG>`, testRender(el, t))
}

func TestFeFuncR(t *testing.T) {
	el := FeFuncR(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feFuncR", el.Name())
	require.Equal(t, `<feFuncR>foo</feFuncR>`, testRender(el, t))
}

func TestFeGaussianBlur(t *testing.T) {
	el := FeGaussianBlur(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feGaussianBlur", el.Name())
	require.Equal(t, `<feGaussianBlur>foo</feGaussianBlur>`, testRender(el, t))
}

func TestFeImage(t *testing.T) {
	el := FeImage(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feImage", el.Name())
	require.Equal(t, `<feImage>foo</feImage>`, testRender(el, t))
}

func TestFeMerge(t *testing.T) {
	el := FeMerge(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feMerge", el.Name())
	require.Equal(t, `<feMerge>foo</feMerge>`, testRender(el, t))
}

func TestFeMergeNode(t *testing.T) {
	el := FeMergeNode(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feMergeNode", el.Name())
	require.Equal(t, `<feMergeNode>foo</feMergeNode>`, testRender(el, t))
}

func TestFeMorphology(t *testing.T) {
	el := FeMorphology(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feMorphology", el.Name())
	require.Equal(t, `<feMorphology>foo</feMorphology>`, testRender(el, t))
}

func TestFeOffset(t *testing.T) {
	el := FeOffset(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feOffset", el.Name())
	require.Equal(t, `<feOffset>foo</feOffset>`, testRender(el, t))
}

func TestFePointLight(t *testing.T) {
	el := FePointLight(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "fePointLight", el.Name())
	require.Equal(t, `<fePointLight>foo</fePointLight>`, testRender(el, t))
}

func TestFeSpecularLighting(t *testing.T) {
	el := FeSpecularLighting(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feSpecularLighting", el.Name())
	require.Equal(t, `<feSpecularLighting>foo</feSpecularLighting>`, testRender(el, t))
}

func TestFeSpotLight(t *testing.T) {
	el := FeSpotLight(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feSpotLight", el.Name())
	require.Equal(t, `<feSpotLight>foo</feSpotLight>`, testRender(el, t))
}

func TestFeTile(t *testing.T) {
	el := FeTile(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feTile", el.Name())
	require.Equal(t, `<feTile>foo</feTile>`, testRender(el, t))
}

func TestFeTurbulence(t *testing.T) {
	el := FeTurbulence(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "feTurbulence", el.Name())
	require.Equal(t, `<feTurbulence>foo</feTurbulence>`, testRender(el, t))
}

func TestFilter(t *testing.T) {
	el := Filter(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "filter", el.Name())
	require.Equal(t, `<filter>foo</filter>`, testRender(el, t))
}

func TestForeignObject(t *testing.T) {
	el := ForeignObject(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "foreignObject", el.Name())
	require.Equal(t, `<foreignObject>foo</foreignObject>`, testRender(el, t))
}

func TestG(t *testing.T) {
	el := G(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "g", el.Name())
	require.Equal(t, `<g>foo</g>`, testRender(el, t))
}

func TestGlyph(t *testing.T) {
	el := Glyph(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "glyph", el.Name())
	require.Equal(t, `<glyph>foo</glyph>`, testRender(el, t))
}

func TestLine(t *testing.T) {
	el := Line(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "line", el.Name())
	require.Equal(t, `<line>foo</line>`, testRender(el, t))
}

func TestLinearGradient(t *testing.T) {
	el := LinearGradient(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "linearGradient", el.Name())
	require.Equal(t, `<linearGradient>foo</linearGradient>`, testRender(el, t))
}

func TestMarker(t *testing.T) {
	el := Marker(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "marker", el.Name())
	require.Equal(t, `<marker>foo</marker>`, testRender(el, t))
}

func TestMask(t *testing.T) {
	el := Mask(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "mask", el.Name())
	require.Equal(t, `<mask>foo</mask>`, testRender(el, t))
}

func TestMetadata(t *testing.T) {
	el := Metadata(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "metadata", el.Name())
	require.Equal(t, `<metadata>foo</metadata>`, testRender(el, t))
}

func TestMpath(t *testing.T) {
	el := Mpath(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "mpath", el.Name())
	require.Equal(t, `<mpath>foo</mpath>`, testRender(el, t))
}

func TestPath(t *testing.T) {
	el := Path(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "path", el.Name())
	require.Equal(t, `<path>foo</path>`, testRender(el, t))
}

func TestPattern(t *testing.T) {
	el := Pattern(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "pattern", el.Name())
	require.Equal(t, `<pattern>foo</pattern>`, testRender(el, t))
}

func TestPolygon(t *testing.T) {
	el := Polygon(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "polygon", el.Name())
	require.Equal(t, `<polygon>foo</polygon>`, testRender(el, t))
}

func TestPolyline(t *testing.T) {
	el := Polyline(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "polyline", el.Name())
	require.Equal(t, `<polyline>foo</polyline>`, testRender(el, t))
}

func TestRadialGradient(t *testing.T) {
	el := RadialGradient(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "radialGradient", el.Name())
	require.Equal(t, `<radialGradient>foo</radialGradient>`, testRender(el, t))
}

func TestRect(t *testing.T) {
	el := Rect(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "rect", el.Name())
	require.Equal(t, `<rect>foo</rect>`, testRender(el, t))
}

func TestScript(t *testing.T) {
	el := Script(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "script", el.Name())
	require.Equal(t, `<script>foo</script>`, testRender(el, t))
}

func TestSet(t *testing.T) {
	el := Set(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "set", el.Name())
	require.Equal(t, `<set>foo</set>`, testRender(el, t))
}

func TestSolidColor(t *testing.T) {
	el := SolidColor(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "solidcolor", el.Name())
	require.Equal(t, `<solidcolor>foo</solidcolor>`, testRender(el, t))
}

func TestStop(t *testing.T) {
	el := Stop(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "stop", el.Name())
	require.Equal(t, `<stop>foo</stop>`, testRender(el, t))
}

func TestStyleElement(t *testing.T) {
	el := StyleElement(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "style", el.Name())
	require.Equal(t, `<style>foo</style>`, testRender(el, t))
}

func TestSymbol(t *testing.T) {
	el := Symbol(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "symbol", el.Name())
	require.Equal(t, `<symbol>foo</symbol>`, testRender(el, t))
}

func TestText(t *testing.T) {
	el := Text(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "text", el.Name())
	require.Equal(t, `<text>foo</text>`, testRender(el, t))
}

func TestTextPath(t *testing.T) {
	el := TextPath(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "textPath", el.Name())
	require.Equal(t, `<textPath>foo</textPath>`, testRender(el, t))
}

func TestTitle(t *testing.T) {
	el := Title(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "title", el.Name())
	require.Equal(t, `<title>foo</title>`, testRender(el, t))
}

func TestTRef(t *testing.T) {
	el := TRef(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "tref", el.Name())
	require.Equal(t, `<tref>foo</tref>`, testRender(el, t))
}

func TestTSpan(t *testing.T) {
	el := TSpan(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "tspan", el.Name())
	require.Equal(t, `<tspan>foo</tspan>`, testRender(el, t))
}

func TestUse(t *testing.T) {
	el := Use(aitch.Text("foo"))
	require.Equal(t, aitch.ElementNode, el.Type())
	require.Equal(t, "use", el.Name())
	require.Equal(t, `<use>foo</use>`, testRender(el, t))
}
