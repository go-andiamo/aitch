package svg

import "github.com/go-andiamo/aitch"

// Svg declares a <svg> SVG element
func Svg(contents ...any) aitch.Node {
	return aitch.NewElement(tagSvg, contents...)
}

// AltGlyph declares a <altGlyph> SVG element
func AltGlyph(contents ...any) aitch.Node {
	return aitch.NewElement(tagAltGlyph, contents...)
}

// Animate declares a <animate> SVG element
func Animate(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimate, contents...)
}

// AnimateMotion declares a <animateMotion> SVG element
func AnimateMotion(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimateMotion, contents...)
}

// AnimateTransform declares a <animateTransform> SVG element
func AnimateTransform(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimateTransform, contents...)
}

// Circle declares a <circle> SVG element
func Circle(contents ...any) aitch.Node {
	return aitch.NewElement(tagCircle, contents...)
}

// ClipPath declares a <clipPath> SVG element
func ClipPath(contents ...any) aitch.Node {
	return aitch.NewElement(tagClipPath, contents...)
}

// Defs declares a <defs> SVG element
func Defs(contents ...any) aitch.Node {
	return aitch.NewElement(tagDefs, contents...)
}

// Desc declares a <desc> SVG element
func Desc(contents ...any) aitch.Node {
	return aitch.NewElement(tagDesc, contents...)
}

// Discard declares a <discard> SVG element
func Discard(contents ...any) aitch.Node {
	return aitch.NewElement(tagDiscard, contents...)
}

// Ellipse declares a <ellipse> SVG element
func Ellipse(contents ...any) aitch.Node {
	return aitch.NewElement(tagEllipse, contents...)
}

// FeBlend declares a <feBlend> SVG element
func FeBlend(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeBlend, contents...)
}

// FeColorMatrix declares a <feColorMatrix> SVG element
func FeColorMatrix(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeColorMatrix, contents...)
}

// FeComponentTransfer declares a <feComponentTransfer> SVG element
func FeComponentTransfer(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeComponentTransfer, contents...)
}

// FeComposite declares a <feComposite> SVG element
func FeComposite(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeComposite, contents...)
}

// FeConvolveMatrix declares a <feConvolveMatrix> SVG element
func FeConvolveMatrix(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeConvolveMatrix, contents...)
}

// FeDiffuseLighting declares a <feDiffuseLighting> SVG element
func FeDiffuseLighting(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDiffuseLighting, contents...)
}

// FeDisplacementMap declares a <feDisplacementMap> SVG element
func FeDisplacementMap(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDisplacementMap, contents...)
}

// FeDistantLight declares a <feDistantLight> SVG element
func FeDistantLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDistantLight, contents...)
}

// FeDropShadow declares a <feDropShadow> SVG element
func FeDropShadow(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDropShadow, contents...)
}

// FeFlood declares a <feFlood> SVG element
func FeFlood(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFlood, contents...)
}

// FeFuncA declares a <feFuncA> SVG element
func FeFuncA(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncA, contents...)
}

// FeFuncB declares a <feFuncB> SVG element
func FeFuncB(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncB, contents...)
}

// FeFuncG declares a <feFuncG> SVG element
func FeFuncG(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncG, contents...)
}

// FeFuncR declares a <feFuncR> SVG element
func FeFuncR(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncR, contents...)
}

// FeGaussianBlur declares a <feGaussianBlur> SVG element
func FeGaussianBlur(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeGaussianBlur, contents...)
}

// FeImage declares a <feImage> SVG element
func FeImage(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeImage, contents...)
}

// FeMerge declares a <feMerge> SVG element
func FeMerge(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMerge, contents...)
}

// FeMergeNode declares a <feMergeNode> SVG element
func FeMergeNode(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMergeNode, contents...)
}

// FeMorphology declares a <feMorphology> SVG element
func FeMorphology(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMorphology, contents...)
}

// FeOffset declares a <feOffset> SVG element
func FeOffset(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeOffset, contents...)
}

// FePointLight declares a <fePointLight> SVG element
func FePointLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFePointLight, contents...)
}

// FeSpecularLighting declares a <feSpecularLighting> SVG element
func FeSpecularLighting(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeSpecularLighting, contents...)
}

// FeSpotLight declares a <feSpotLight> SVG element
func FeSpotLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeSpotLight, contents...)
}

// FeTile declares a <feTile> SVG element
func FeTile(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeTile, contents...)
}

// FeTurbulence declares a <feTurbulence> SVG element
func FeTurbulence(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeTurbulence, contents...)
}

// Filter declares a <filter> SVG element
func Filter(contents ...any) aitch.Node {
	return aitch.NewElement(tagFilter, contents...)
}

// ForeignObject declares a <foreignObject> SVG element
func ForeignObject(contents ...any) aitch.Node {
	return aitch.NewElement(tagForeignObject, contents...)
}

// G declares a <g> SVG element
func G(contents ...any) aitch.Node {
	return aitch.NewElement(tagG, contents...)
}

// Glyph declares a <glyph> SVG element
func Glyph(contents ...any) aitch.Node {
	return aitch.NewElement(tagGlyph, contents...)
}

// Line declares a <line> SVG element
func Line(contents ...any) aitch.Node {
	return aitch.NewElement(tagLine, contents...)
}

// LinearGradient declares a <linearGradient> SVG element
func LinearGradient(contents ...any) aitch.Node {
	return aitch.NewElement(tagLinearGradient, contents...)
}

// Marker declares a <marker> SVG element
func Marker(contents ...any) aitch.Node {
	return aitch.NewElement(tagMarker, contents...)
}

// Mask declares a <mask> SVG element
func Mask(contents ...any) aitch.Node {
	return aitch.NewElement(tagMask, contents...)
}

// Metadata declares a <metadata> SVG element
func Metadata(contents ...any) aitch.Node {
	return aitch.NewElement(tagMetadata, contents...)
}

// Mpath declares a <mpath> SVG element
func Mpath(contents ...any) aitch.Node {
	return aitch.NewElement(tagMpath, contents...)
}

// Path declares a <path> SVG element
func Path(contents ...any) aitch.Node {
	return aitch.NewElement(tagPath, contents...)
}

// Pattern declares a <pattern> SVG element
func Pattern(contents ...any) aitch.Node {
	return aitch.NewElement(tagPattern, contents...)
}

// Polygon declares a <polygon> SVG element
func Polygon(contents ...any) aitch.Node {
	return aitch.NewElement(tagPolygon, contents...)
}

// Polyline declares a <polyline> SVG element
func Polyline(contents ...any) aitch.Node {
	return aitch.NewElement(tagPolyline, contents...)
}

// RadialGradient declares a <radialGradient> SVG element
func RadialGradient(contents ...any) aitch.Node {
	return aitch.NewElement(tagRadialGradient, contents...)
}

// Rect declares a <rect> SVG element
func Rect(contents ...any) aitch.Node {
	return aitch.NewElement(tagRect, contents...)
}

// Script declares a <script> SVG element
func Script(contents ...any) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

// Set declares a <set> SVG element
func Set(contents ...any) aitch.Node {
	return aitch.NewElement(tagSet, contents...)
}

// SolidColor declares a <solidcolor> SVG element
func SolidColor(contents ...any) aitch.Node {
	return aitch.NewElement(tagSolidColor, contents...)
}

// Stop declares a <stop> SVG element
func Stop(contents ...any) aitch.Node {
	return aitch.NewElement(tagStop, contents...)
}

// StyleElement declares a <style> SVG element
func StyleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

// Symbol declares a <symbol> SVG element
func Symbol(contents ...any) aitch.Node {
	return aitch.NewElement(tagSymbol, contents...)
}

// TextElement declares a <text> SVG element
func TextElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagText, contents...)
}

// TextPath declares a <textPath> SVG element
func TextPath(contents ...any) aitch.Node {
	return aitch.NewElement(tagTextPath, contents...)
}

// Title declares a <title> SVG element
func Title(contents ...any) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

// TRef declares a <tref> SVG element
func TRef(contents ...any) aitch.Node {
	return aitch.NewElement(tagTRef, contents...)
}

// TSpan declares a <tspan> SVG element
func TSpan(contents ...any) aitch.Node {
	return aitch.NewElement(tagTSpan, contents...)
}

// Use declares a <use> SVG element
func Use(contents ...any) aitch.Node {
	return aitch.NewElement(tagUse, contents...)
}

var (
	tagAltGlyph            = []byte("altGlyph")
	tagAnimate             = []byte("animate")
	tagAnimateMotion       = []byte("animateMotion")
	tagAnimateTransform    = []byte("animateTransform")
	tagCircle              = []byte("circle")
	tagClipPath            = []byte("clipPath")
	tagDefs                = []byte("defs")
	tagDesc                = []byte("desc")
	tagDiscard             = []byte("discard")
	tagEllipse             = []byte("ellipse")
	tagFeBlend             = []byte("feBlend")
	tagFeColorMatrix       = []byte("feColorMatrix")
	tagFeComponentTransfer = []byte("feComponentTransfer")
	tagFeComposite         = []byte("feComposite")
	tagFeConvolveMatrix    = []byte("feConvolveMatrix")
	tagFeDiffuseLighting   = []byte("feDiffuseLighting")
	tagFeDisplacementMap   = []byte("feDisplacementMap")
	tagFeDistantLight      = []byte("feDistantLight")
	tagFeDropShadow        = []byte("feDropShadow")
	tagFeFlood             = []byte("feFlood")
	tagFeFuncA             = []byte("feFuncA")
	tagFeFuncB             = []byte("feFuncB")
	tagFeFuncG             = []byte("feFuncG")
	tagFeFuncR             = []byte("feFuncR")
	tagFeGaussianBlur      = []byte("feGaussianBlur")
	tagFeImage             = []byte("feImage")
	tagFeMerge             = []byte("feMerge")
	tagFeMergeNode         = []byte("feMergeNode")
	tagFeMorphology        = []byte("feMorphology")
	tagFeOffset            = []byte("feOffset")
	tagFePointLight        = []byte("fePointLight")
	tagFeSpecularLighting  = []byte("feSpecularLighting")
	tagFeSpotLight         = []byte("feSpotLight")
	tagFeTile              = []byte("feTile")
	tagFeTurbulence        = []byte("feTurbulence")
	tagFilter              = []byte("filter")
	tagForeignObject       = []byte("foreignObject")
	tagG                   = []byte("g")
	tagGlyph               = []byte("glyph")
	tagLine                = []byte("line")
	tagLinearGradient      = []byte("linearGradient")
	tagMarker              = []byte("marker")
	tagMask                = []byte("mask")
	tagMetadata            = []byte("metadata")
	tagMpath               = []byte("mpath")
	tagPath                = []byte("path")
	tagPattern             = []byte("pattern")
	tagPolygon             = []byte("polygon")
	tagPolyline            = []byte("polyline")
	tagRadialGradient      = []byte("radialGradient")
	tagRect                = []byte("rect")
	tagScript              = []byte("script")
	tagSet                 = []byte("set")
	tagSolidColor          = []byte("solidcolor")
	tagStop                = []byte("stop")
	tagStyle               = []byte("style")
	tagSvg                 = []byte("svg")
	tagSymbol              = []byte("symbol")
	tagText                = []byte("text")
	tagTextPath            = []byte("textPath")
	tagTitle               = []byte("title")
	tagTRef                = []byte("tref")
	tagTSpan               = []byte("tspan")
	tagUse                 = []byte("use")
)
