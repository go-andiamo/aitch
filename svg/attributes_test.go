package svg

import (
	"github.com/go-andiamo/aitch"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccumulate(t *testing.T) {
	a := Accumulate("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "accumulate", a.Name())
	require.Equal(t, ` accumulate="foo"`, testRender(a, t))
}

func TestAdditive(t *testing.T) {
	a := Additive("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "additive", a.Name())
	require.Equal(t, ` additive="foo"`, testRender(a, t))
}

func TestAlignmentBaseline(t *testing.T) {
	a := AlignmentBaseline("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "alignment-baseline", a.Name())
	require.Equal(t, ` alignment-baseline="foo"`, testRender(a, t))
}

func TestAttributeName(t *testing.T) {
	a := AttributeName("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "attributeName", a.Name())
	require.Equal(t, ` attributeName="foo"`, testRender(a, t))
}

func TestAttributeType(t *testing.T) {
	a := AttributeType("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "attributeType", a.Name())
	require.Equal(t, ` attributeType="foo"`, testRender(a, t))
}

func TestBaselineShift(t *testing.T) {
	a := BaselineShift("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "baseline-shift", a.Name())
	require.Equal(t, ` baseline-shift="foo"`, testRender(a, t))
}

func TestBegin(t *testing.T) {
	a := Begin("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "begin", a.Name())
	require.Equal(t, ` begin="foo"`, testRender(a, t))
}

func TestBy(t *testing.T) {
	a := By("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "by", a.Name())
	require.Equal(t, ` by="foo"`, testRender(a, t))
}

func TestCalcMode(t *testing.T) {
	a := CalcMode("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "calcMode", a.Name())
	require.Equal(t, ` calcMode="foo"`, testRender(a, t))
}

func TestClass(t *testing.T) {
	a := Class("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "class", a.Name())
	require.Equal(t, ` class="foo"`, testRender(a, t))
}

func TestClip(t *testing.T) {
	a := Clip("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "clip", a.Name())
	require.Equal(t, ` clip="foo"`, testRender(a, t))
}

func TestClipPathAttr(t *testing.T) {
	a := ClipPathAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "clip-path", a.Name())
	require.Equal(t, ` clip-path="foo"`, testRender(a, t))
}

func TestClipRule(t *testing.T) {
	a := ClipRule("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "clip-rule", a.Name())
	require.Equal(t, ` clip-rule="foo"`, testRender(a, t))
}

func TestColor(t *testing.T) {
	a := Color("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "color", a.Name())
	require.Equal(t, ` color="foo"`, testRender(a, t))
}

func TestColorInterpolation(t *testing.T) {
	a := ColorInterpolation("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "color-interpolation", a.Name())
	require.Equal(t, ` color-interpolation="foo"`, testRender(a, t))
}

func TestColorInterpolationFilters(t *testing.T) {
	a := ColorInterpolationFilters("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "color-interpolation-filters", a.Name())
	require.Equal(t, ` color-interpolation-filters="foo"`, testRender(a, t))
}

func TestColorProfile(t *testing.T) {
	a := ColorProfile("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "color-profile", a.Name())
	require.Equal(t, ` color-profile="foo"`, testRender(a, t))
}

func TestColorRendering(t *testing.T) {
	a := ColorRendering("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "color-rendering", a.Name())
	require.Equal(t, ` color-rendering="foo"`, testRender(a, t))
}

func TestCursor(t *testing.T) {
	a := Cursor("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "cursor", a.Name())
	require.Equal(t, ` cursor="foo"`, testRender(a, t))
}

func TestCx(t *testing.T) {
	a := Cx("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "cx", a.Name())
	require.Equal(t, ` cx="foo"`, testRender(a, t))
}

func TestCy(t *testing.T) {
	a := Cy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "cy", a.Name())
	require.Equal(t, ` cy="foo"`, testRender(a, t))
}

func TestD(t *testing.T) {
	a := D("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "d", a.Name())
	require.Equal(t, ` d="foo"`, testRender(a, t))
}

func TestDirection(t *testing.T) {
	a := Direction("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "direction", a.Name())
	require.Equal(t, ` direction="foo"`, testRender(a, t))
}

func TestDisplay(t *testing.T) {
	a := Display("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "display", a.Name())
	require.Equal(t, ` display="foo"`, testRender(a, t))
}

func TestDominantBaseline(t *testing.T) {
	a := DominantBaseline("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "dominant-baseline", a.Name())
	require.Equal(t, ` dominant-baseline="foo"`, testRender(a, t))
}

func TestDur(t *testing.T) {
	a := Dur("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "dur", a.Name())
	require.Equal(t, ` dur="foo"`, testRender(a, t))
}

func TestDx(t *testing.T) {
	a := Dx("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "dx", a.Name())
	require.Equal(t, ` dx="foo"`, testRender(a, t))
}

func TestDy(t *testing.T) {
	a := Dy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "dy", a.Name())
	require.Equal(t, ` dy="foo"`, testRender(a, t))
}

func TestEnableBackground(t *testing.T) {
	a := EnableBackground("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "enable-background", a.Name())
	require.Equal(t, ` enable-background="foo"`, testRender(a, t))
}

func TestEnd(t *testing.T) {
	a := End("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "end", a.Name())
	require.Equal(t, ` end="foo"`, testRender(a, t))
}

func TestExternalResourcesRequired(t *testing.T) {
	a := ExternalResourcesRequired("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "externalResourcesRequired", a.Name())
	require.Equal(t, ` externalResourcesRequired="foo"`, testRender(a, t))
}

func TestFill(t *testing.T) {
	a := Fill("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "fill", a.Name())
	require.Equal(t, ` fill="foo"`, testRender(a, t))
}

func TestFillOpacity(t *testing.T) {
	a := FillOpacity("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "fill-opacity", a.Name())
	require.Equal(t, ` fill-opacity="foo"`, testRender(a, t))
}

func TestFillRule(t *testing.T) {
	a := FillRule("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "fill-rule", a.Name())
	require.Equal(t, ` fill-rule="foo"`, testRender(a, t))
}

func TestFilterAttr(t *testing.T) {
	a := FilterAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "filter", a.Name())
	require.Equal(t, ` filter="foo"`, testRender(a, t))
}

func TestFilterUnits(t *testing.T) {
	a := FilterUnits("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "filterUnits", a.Name())
	require.Equal(t, ` filterUnits="foo"`, testRender(a, t))
}

func TestFloodColor(t *testing.T) {
	a := FloodColor("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "flood-color", a.Name())
	require.Equal(t, ` flood-color="foo"`, testRender(a, t))
}

func TestFloodOpacity(t *testing.T) {
	a := FloodOpacity("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "flood-opacity", a.Name())
	require.Equal(t, ` flood-opacity="foo"`, testRender(a, t))
}

func TestFontFamily(t *testing.T) {
	a := FontFamily("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-family", a.Name())
	require.Equal(t, ` font-family="foo"`, testRender(a, t))
}

func TestFontSize(t *testing.T) {
	a := FontSize("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-size", a.Name())
	require.Equal(t, ` font-size="foo"`, testRender(a, t))
}

func TestFontSizeAdjust(t *testing.T) {
	a := FontSizeAdjust("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-size-adjust", a.Name())
	require.Equal(t, ` font-size-adjust="foo"`, testRender(a, t))
}

func TestFontStretch(t *testing.T) {
	a := FontStretch("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-stretch", a.Name())
	require.Equal(t, ` font-stretch="foo"`, testRender(a, t))
}

func TestFontStyle(t *testing.T) {
	a := FontStyle("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-style", a.Name())
	require.Equal(t, ` font-style="foo"`, testRender(a, t))
}

func TestFontVariant(t *testing.T) {
	a := FontVariant("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-variant", a.Name())
	require.Equal(t, ` font-variant="foo"`, testRender(a, t))
}

func TestFontWeight(t *testing.T) {
	a := FontWeight("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "font-weight", a.Name())
	require.Equal(t, ` font-weight="foo"`, testRender(a, t))
}

func TestFrom(t *testing.T) {
	a := From("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "from", a.Name())
	require.Equal(t, ` from="foo"`, testRender(a, t))
}

func TestFx(t *testing.T) {
	a := Fx("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "fx", a.Name())
	require.Equal(t, ` fx="foo"`, testRender(a, t))
}

func TestFy(t *testing.T) {
	a := Fy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "fy", a.Name())
	require.Equal(t, ` fy="foo"`, testRender(a, t))
}

func TestGlyphOrientationHorizontal(t *testing.T) {
	a := GlyphOrientationHorizontal("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "glyph-orientation-horizontal", a.Name())
	require.Equal(t, ` glyph-orientation-horizontal="foo"`, testRender(a, t))
}

func TestGlyphOrientationVertical(t *testing.T) {
	a := GlyphOrientationVertical("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "glyph-orientation-vertical", a.Name())
	require.Equal(t, ` glyph-orientation-vertical="foo"`, testRender(a, t))
}

func TestGradientTransform(t *testing.T) {
	a := GradientTransform("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "gradientTransform", a.Name())
	require.Equal(t, ` gradientTransform="foo"`, testRender(a, t))
}

func TestGradientUnits(t *testing.T) {
	a := GradientUnits("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "gradientUnits", a.Name())
	require.Equal(t, ` gradientUnits="foo"`, testRender(a, t))
}

func TestHeight(t *testing.T) {
	a := Height("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "height", a.Name())
	require.Equal(t, ` height="foo"`, testRender(a, t))
}

func TestHref(t *testing.T) {
	a := Href("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "href", a.Name())
	require.Equal(t, ` href="foo"`, testRender(a, t))
}

func TestId(t *testing.T) {
	a := Id("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "id", a.Name())
	require.Equal(t, ` id="foo"`, testRender(a, t))
}

func TestImageRendering(t *testing.T) {
	a := ImageRendering("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "image-rendering", a.Name())
	require.Equal(t, ` image-rendering="foo"`, testRender(a, t))
}

func TestKerning(t *testing.T) {
	a := Kerning("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "kerning", a.Name())
	require.Equal(t, ` kerning="foo"`, testRender(a, t))
}

func TestKeyPoints(t *testing.T) {
	a := KeyPoints("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "keyPoints", a.Name())
	require.Equal(t, ` keyPoints="foo"`, testRender(a, t))
}

func TestKeySplines(t *testing.T) {
	a := KeySplines("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "keySplines", a.Name())
	require.Equal(t, ` keySplines="foo"`, testRender(a, t))
}

func TestKeyTimes(t *testing.T) {
	a := KeyTimes("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "keyTimes", a.Name())
	require.Equal(t, ` keyTimes="foo"`, testRender(a, t))
}

func TestLengthAdjust(t *testing.T) {
	a := LengthAdjust("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "lengthAdjust", a.Name())
	require.Equal(t, ` lengthAdjust="foo"`, testRender(a, t))
}

func TestLetterSpacing(t *testing.T) {
	a := LetterSpacing("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "letter-spacing", a.Name())
	require.Equal(t, ` letter-spacing="foo"`, testRender(a, t))
}

func TestLightingColor(t *testing.T) {
	a := LightingColor("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "lighting-color", a.Name())
	require.Equal(t, ` lighting-color="foo"`, testRender(a, t))
}

func TestMarkerEnd(t *testing.T) {
	a := MarkerEnd("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "marker-end", a.Name())
	require.Equal(t, ` marker-end="foo"`, testRender(a, t))
}

func TestMarkerMid(t *testing.T) {
	a := MarkerMid("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "marker-mid", a.Name())
	require.Equal(t, ` marker-mid="foo"`, testRender(a, t))
}

func TestMarkerStart(t *testing.T) {
	a := MarkerStart("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "marker-start", a.Name())
	require.Equal(t, ` marker-start="foo"`, testRender(a, t))
}

func TestMarkerHeight(t *testing.T) {
	a := MarkerHeight("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "markerHeight", a.Name())
	require.Equal(t, ` markerHeight="foo"`, testRender(a, t))
}

func TestMarkerUnits(t *testing.T) {
	a := MarkerUnits("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "markerUnits", a.Name())
	require.Equal(t, ` markerUnits="foo"`, testRender(a, t))
}

func TestMarkerWidth(t *testing.T) {
	a := MarkerWidth("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "markerWidth", a.Name())
	require.Equal(t, ` markerWidth="foo"`, testRender(a, t))
}

func TestMaskAttr(t *testing.T) {
	a := MaskAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "mask", a.Name())
	require.Equal(t, ` mask="foo"`, testRender(a, t))
}

func TestOnBlur(t *testing.T) {
	a := OnBlur("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onblur", a.Name())
	require.Equal(t, ` onblur="foo"`, testRender(a, t))
}

func TestOnClick(t *testing.T) {
	a := OnClick("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onclick", a.Name())
	require.Equal(t, ` onclick="foo"`, testRender(a, t))
}

func TestOnFocus(t *testing.T) {
	a := OnFocus("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onfocus", a.Name())
	require.Equal(t, ` onfocus="foo"`, testRender(a, t))
}

func TestOnLoad(t *testing.T) {
	a := OnLoad("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onload", a.Name())
	require.Equal(t, ` onload="foo"`, testRender(a, t))
}

func TestOnMouseDown(t *testing.T) {
	a := OnMouseDown("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmousedown", a.Name())
	require.Equal(t, ` onmousedown="foo"`, testRender(a, t))
}

func TestOnMouseMove(t *testing.T) {
	a := OnMouseMove("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmousemove", a.Name())
	require.Equal(t, ` onmousemove="foo"`, testRender(a, t))
}

func TestOnMouseOut(t *testing.T) {
	a := OnMouseOut("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseout", a.Name())
	require.Equal(t, ` onmouseout="foo"`, testRender(a, t))
}

func TestOnMouseOver(t *testing.T) {
	a := OnMouseOver("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseover", a.Name())
	require.Equal(t, ` onmouseover="foo"`, testRender(a, t))
}

func TestOnMouseUp(t *testing.T) {
	a := OnMouseUp("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseup", a.Name())
	require.Equal(t, ` onmouseup="foo"`, testRender(a, t))
}

func TestOnUnload(t *testing.T) {
	a := OnUnload("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onunload", a.Name())
	require.Equal(t, ` onunload="foo"`, testRender(a, t))
}

func TestOpacity(t *testing.T) {
	a := Opacity("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "opacity", a.Name())
	require.Equal(t, ` opacity="foo"`, testRender(a, t))
}

func TestOrient(t *testing.T) {
	a := Orient("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "orient", a.Name())
	require.Equal(t, ` orient="foo"`, testRender(a, t))
}

func TestPathAttr(t *testing.T) {
	a := PathAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "path", a.Name())
	require.Equal(t, ` path="foo"`, testRender(a, t))
}

func TestPathLength(t *testing.T) {
	a := PathLength("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "pathLength", a.Name())
	require.Equal(t, ` pathLength="foo"`, testRender(a, t))
}

func TestPatternContentUnits(t *testing.T) {
	a := PatternContentUnits("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "patternContentUnits", a.Name())
	require.Equal(t, ` patternContentUnits="foo"`, testRender(a, t))
}

func TestPatternTransform(t *testing.T) {
	a := PatternTransform("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "patternTransform", a.Name())
	require.Equal(t, ` patternTransform="foo"`, testRender(a, t))
}

func TestPatternUnits(t *testing.T) {
	a := PatternUnits("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "patternUnits", a.Name())
	require.Equal(t, ` patternUnits="foo"`, testRender(a, t))
}

func TestPointerEvents(t *testing.T) {
	a := PointerEvents("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "pointer-events", a.Name())
	require.Equal(t, ` pointer-events="foo"`, testRender(a, t))
}

func TestPoints(t *testing.T) {
	a := Points("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "points", a.Name())
	require.Equal(t, ` points="foo"`, testRender(a, t))
}

func TestPreserveAspectRatio(t *testing.T) {
	a := PreserveAspectRatio("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "preserveAspectRatio", a.Name())
	require.Equal(t, ` preserveAspectRatio="foo"`, testRender(a, t))
}

func TestPrimitiveUnits(t *testing.T) {
	a := PrimitiveUnits("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "primitiveUnits", a.Name())
	require.Equal(t, ` primitiveUnits="foo"`, testRender(a, t))
}

func TestR(t *testing.T) {
	a := R("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "r", a.Name())
	require.Equal(t, ` r="foo"`, testRender(a, t))
}

func TestRefX(t *testing.T) {
	a := RefX("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "refX", a.Name())
	require.Equal(t, ` refX="foo"`, testRender(a, t))
}

func TestRefY(t *testing.T) {
	a := RefY("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "refY", a.Name())
	require.Equal(t, ` refY="foo"`, testRender(a, t))
}

func TestRepeatCount(t *testing.T) {
	a := RepeatCount("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "repeatCount", a.Name())
	require.Equal(t, ` repeatCount="foo"`, testRender(a, t))
}

func TestRepeatDur(t *testing.T) {
	a := RepeatDur("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "repeatDur", a.Name())
	require.Equal(t, ` repeatDur="foo"`, testRender(a, t))
}

func TestRequiredExtensions(t *testing.T) {
	a := RequiredExtensions("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "requiredExtensions", a.Name())
	require.Equal(t, ` requiredExtensions="foo"`, testRender(a, t))
}

func TestRequiredFeatures(t *testing.T) {
	a := RequiredFeatures("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "requiredFeatures", a.Name())
	require.Equal(t, ` requiredFeatures="foo"`, testRender(a, t))
}

func TestRotate(t *testing.T) {
	a := Rotate("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "rotate", a.Name())
	require.Equal(t, ` rotate="foo"`, testRender(a, t))
}

func TestRx(t *testing.T) {
	a := Rx("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "rx", a.Name())
	require.Equal(t, ` rx="foo"`, testRender(a, t))
}

func TestRy(t *testing.T) {
	a := Ry("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ry", a.Name())
	require.Equal(t, ` ry="foo"`, testRender(a, t))
}

func TestShapeRendering(t *testing.T) {
	a := ShapeRendering("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "shape-rendering", a.Name())
	require.Equal(t, ` shape-rendering="foo"`, testRender(a, t))
}

func TestSpreadMethod(t *testing.T) {
	a := SpreadMethod("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "spreadMethod", a.Name())
	require.Equal(t, ` spreadMethod="foo"`, testRender(a, t))
}

func TestStopColor(t *testing.T) {
	a := StopColor("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stop-color", a.Name())
	require.Equal(t, ` stop-color="foo"`, testRender(a, t))
}

func TestStopOpacity(t *testing.T) {
	a := StopOpacity("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stop-opacity", a.Name())
	require.Equal(t, ` stop-opacity="foo"`, testRender(a, t))
}

func TestStroke(t *testing.T) {
	a := Stroke("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke", a.Name())
	require.Equal(t, ` stroke="foo"`, testRender(a, t))
}

func TestStrokeDashArray(t *testing.T) {
	a := StrokeDashArray("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-dasharray", a.Name())
	require.Equal(t, ` stroke-dasharray="foo"`, testRender(a, t))
}

func TestStrokeDashOffset(t *testing.T) {
	a := StrokeDashOffset("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-dashoffset", a.Name())
	require.Equal(t, ` stroke-dashoffset="foo"`, testRender(a, t))
}

func TestStrokeLineCap(t *testing.T) {
	a := StrokeLineCap("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-linecap", a.Name())
	require.Equal(t, ` stroke-linecap="foo"`, testRender(a, t))
}

func TestStrokeLineJoin(t *testing.T) {
	a := StrokeLineJoin("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-linejoin", a.Name())
	require.Equal(t, ` stroke-linejoin="foo"`, testRender(a, t))
}

func TestStrokeMiterLimit(t *testing.T) {
	a := StrokeMiterLimit("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-miterlimit", a.Name())
	require.Equal(t, ` stroke-miterlimit="foo"`, testRender(a, t))
}

func TestStrokeOpacity(t *testing.T) {
	a := StrokeOpacity("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-opacity", a.Name())
	require.Equal(t, ` stroke-opacity="foo"`, testRender(a, t))
}

func TestStrokeWidth(t *testing.T) {
	a := StrokeWidth("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "stroke-width", a.Name())
	require.Equal(t, ` stroke-width="foo"`, testRender(a, t))
}

func TestStyle(t *testing.T) {
	a := Style("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "style", a.Name())
	require.Equal(t, ` style="foo"`, testRender(a, t))
}

func TestSystemLanguage(t *testing.T) {
	a := SystemLanguage("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "systemLanguage", a.Name())
	require.Equal(t, ` systemLanguage="foo"`, testRender(a, t))
}

func TestTabIndex(t *testing.T) {
	a := TabIndex("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "tabindex", a.Name())
	require.Equal(t, ` tabindex="foo"`, testRender(a, t))
}

func TestTextAnchor(t *testing.T) {
	a := TextAnchor("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "text-anchor", a.Name())
	require.Equal(t, ` text-anchor="foo"`, testRender(a, t))
}

func TestTextDecoration(t *testing.T) {
	a := TextDecoration("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "text-decoration", a.Name())
	require.Equal(t, ` text-decoration="foo"`, testRender(a, t))
}

func TestTextRendering(t *testing.T) {
	a := TextRendering("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "text-rendering", a.Name())
	require.Equal(t, ` text-rendering="foo"`, testRender(a, t))
}

func TestTextLength(t *testing.T) {
	a := TextLength("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "textLength", a.Name())
	require.Equal(t, ` textLength="foo"`, testRender(a, t))
}

func TestTo(t *testing.T) {
	a := To("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "to", a.Name())
	require.Equal(t, ` to="foo"`, testRender(a, t))
}

func TestType(t *testing.T) {
	a := Type("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "type", a.Name())
	require.Equal(t, ` type="foo"`, testRender(a, t))
}

func TestUnicodeBidi(t *testing.T) {
	a := UnicodeBidi("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "unicode-bidi", a.Name())
	require.Equal(t, ` unicode-bidi="foo"`, testRender(a, t))
}

func TestValues(t *testing.T) {
	a := Values("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "values", a.Name())
	require.Equal(t, ` values="foo"`, testRender(a, t))
}

func TestVisibility(t *testing.T) {
	a := Visibility("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "visibility", a.Name())
	require.Equal(t, ` visibility="foo"`, testRender(a, t))
}

func TestWidth(t *testing.T) {
	a := Width("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "width", a.Name())
	require.Equal(t, ` width="foo"`, testRender(a, t))
}

func TestWordSpacing(t *testing.T) {
	a := WordSpacing("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "word-spacing", a.Name())
	require.Equal(t, ` word-spacing="foo"`, testRender(a, t))
}

func TestWritingMode(t *testing.T) {
	a := WritingMode("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "writing-mode", a.Name())
	require.Equal(t, ` writing-mode="foo"`, testRender(a, t))
}

func TestX(t *testing.T) {
	a := X("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "x", a.Name())
	require.Equal(t, ` x="foo"`, testRender(a, t))
}

func TestX1(t *testing.T) {
	a := X1("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "x1", a.Name())
	require.Equal(t, ` x1="foo"`, testRender(a, t))
}

func TestX2(t *testing.T) {
	a := X2("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "x2", a.Name())
	require.Equal(t, ` x2="foo"`, testRender(a, t))
}

func TestXlinkActuate(t *testing.T) {
	a := XlinkActuate("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xlink:actuate", a.Name())
	require.Equal(t, ` xlink:actuate="foo"`, testRender(a, t))
}

func TestXlinkArcrole(t *testing.T) {
	a := XlinkArcrole("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xlink:arcrole", a.Name())
	require.Equal(t, ` xlink:arcrole="foo"`, testRender(a, t))
}

func TestXlinkHref(t *testing.T) {
	a := XlinkHref("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xlink:href", a.Name())
	require.Equal(t, ` xlink:href="foo"`, testRender(a, t))
}

func TestXlinkRole(t *testing.T) {
	a := XlinkRole("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xlink:role", a.Name())
	require.Equal(t, ` xlink:role="foo"`, testRender(a, t))
}

func TestXlinkShow(t *testing.T) {
	a := XlinkShow("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xlink:show", a.Name())
	require.Equal(t, ` xlink:show="foo"`, testRender(a, t))
}

func TestXlinkTitle(t *testing.T) {
	a := XlinkTitle("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xlink:title", a.Name())
	require.Equal(t, ` xlink:title="foo"`, testRender(a, t))
}

func TestXmlBase(t *testing.T) {
	a := XmlBase("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xml:base", a.Name())
	require.Equal(t, ` xml:base="foo"`, testRender(a, t))
}

func TestXmlLang(t *testing.T) {
	a := XmlLang("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xml:lang", a.Name())
	require.Equal(t, ` xml:lang="foo"`, testRender(a, t))
}

func TestXmlSpace(t *testing.T) {
	a := XmlSpace("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "xml:space", a.Name())
	require.Equal(t, ` xml:space="foo"`, testRender(a, t))
}

func TestY(t *testing.T) {
	a := Y("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "y", a.Name())
	require.Equal(t, ` y="foo"`, testRender(a, t))
}

func TestY1(t *testing.T) {
	a := Y1("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "y1", a.Name())
	require.Equal(t, ` y1="foo"`, testRender(a, t))
}

func TestY2(t *testing.T) {
	a := Y2("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "y2", a.Name())
	require.Equal(t, ` y2="foo"`, testRender(a, t))
}
