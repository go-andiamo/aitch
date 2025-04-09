package svg

import "github.com/go-andiamo/aitch"

// Svg declares an <svg> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/svg
func Svg(contents ...any) aitch.Node {
	return aitch.NewElement(tagSvg, contents...)
}

// AltGlyph declares an <altGlyph> SVG element
func AltGlyph(contents ...any) aitch.Node {
	return aitch.NewElement(tagAltGlyph, contents...)
}

// Animate declares an <animate> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/animate
func Animate(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimate, contents...)
}

// AnimateMotion declares an <animateMotion> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/animateMotion
func AnimateMotion(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimateMotion, contents...)
}

// AnimateTransform declares an <animateTransform> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/animateTransform
func AnimateTransform(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimateTransform, contents...)
}

// Circle declares a <circle> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/circle
func Circle(contents ...any) aitch.Node {
	return aitch.NewElement(tagCircle, contents...)
}

// ClipPath declares a <clipPath> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/clipPath
func ClipPath(contents ...any) aitch.Node {
	return aitch.NewElement(tagClipPath, contents...)
}

// Defs declares a <defs> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/defs
func Defs(contents ...any) aitch.Node {
	return aitch.NewElement(tagDefs, contents...)
}

// Desc declares a <desc> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/desc
func Desc(contents ...any) aitch.Node {
	return aitch.NewElement(tagDesc, contents...)
}

// Discard declares a <discard> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/discard
func Discard(contents ...any) aitch.Node {
	return aitch.NewElement(tagDiscard, contents...)
}

// Ellipse declares a <ellipse> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/ellipse
func Ellipse(contents ...any) aitch.Node {
	return aitch.NewElement(tagEllipse, contents...)
}

// FeBlend declares a <feBlend> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feBlend
func FeBlend(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeBlend, contents...)
}

// FeColorMatrix declares a <feColorMatrix> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feColorMatrix
func FeColorMatrix(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeColorMatrix, contents...)
}

// FeComponentTransfer declares a <feComponentTransfer> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feComponentTransfer
func FeComponentTransfer(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeComponentTransfer, contents...)
}

// FeComposite declares a <feComposite> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feComposite
func FeComposite(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeComposite, contents...)
}

// FeConvolveMatrix declares a <feConvolveMatrix> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feConvolveMatrix
func FeConvolveMatrix(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeConvolveMatrix, contents...)
}

// FeDiffuseLighting declares a <feDiffuseLighting> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feDiffuseLighting
func FeDiffuseLighting(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDiffuseLighting, contents...)
}

// FeDisplacementMap declares a <feDisplacementMap> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feDisplacementMap
func FeDisplacementMap(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDisplacementMap, contents...)
}

// FeDistantLight declares a <feDistantLight> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feDistantLight
func FeDistantLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDistantLight, contents...)
}

// FeDropShadow declares a <feDropShadow> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feDropShadow
func FeDropShadow(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDropShadow, contents...)
}

// FeFlood declares a <feFlood> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feFlood
func FeFlood(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFlood, contents...)
}

// FeFuncA declares a <feFuncA> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feFuncA
func FeFuncA(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncA, contents...)
}

// FeFuncB declares a <feFuncB> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feFuncB
func FeFuncB(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncB, contents...)
}

// FeFuncG declares a <feFuncG> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feFuncG
func FeFuncG(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncG, contents...)
}

// FeFuncR declares a <feFuncR> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feFuncR
func FeFuncR(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncR, contents...)
}

// FeGaussianBlur declares a <feGaussianBlur> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feGaussianBlur
func FeGaussianBlur(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeGaussianBlur, contents...)
}

// FeImage declares a <feImage> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feImage
func FeImage(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeImage, contents...)
}

// FeMerge declares a <feMerge> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feMerge
func FeMerge(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMerge, contents...)
}

// FeMergeNode declares a <feMergeNode> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feMergeNode
func FeMergeNode(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMergeNode, contents...)
}

// FeMorphology declares a <feMorphology> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feMorphology
func FeMorphology(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMorphology, contents...)
}

// FeOffset declares a <feOffset> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feOffset
func FeOffset(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeOffset, contents...)
}

// FePointLight declares a <fePointLight> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/fePointLight
func FePointLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFePointLight, contents...)
}

// FeSpecularLighting declares a <feSpecularLighting> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feSpecularLighting
func FeSpecularLighting(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeSpecularLighting, contents...)
}

// FeSpotLight declares a <feSpotLight> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feSpotLight
func FeSpotLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeSpotLight, contents...)
}

// FeTile declares a <feTile> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feTile
func FeTile(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeTile, contents...)
}

// FeTurbulence declares a <feTurbulence> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/feTurbulence
func FeTurbulence(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeTurbulence, contents...)
}

// Filter declares a <filter> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/filter
func Filter(contents ...any) aitch.Node {
	return aitch.NewElement(tagFilter, contents...)
}

// ForeignObject declares a <foreignObject> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/foreignObject
func ForeignObject(contents ...any) aitch.Node {
	return aitch.NewElement(tagForeignObject, contents...)
}

// G declares a <g> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/g
func G(contents ...any) aitch.Node {
	return aitch.NewElement(tagG, contents...)
}

// Glyph declares a <glyph> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/glyph
func Glyph(contents ...any) aitch.Node {
	return aitch.NewElement(tagGlyph, contents...)
}

// Line declares a <line> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/line
func Line(contents ...any) aitch.Node {
	return aitch.NewElement(tagLine, contents...)
}

// LinearGradient declares a <linearGradient> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/linearGradient
func LinearGradient(contents ...any) aitch.Node {
	return aitch.NewElement(tagLinearGradient, contents...)
}

// Marker declares a <marker> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/marker
func Marker(contents ...any) aitch.Node {
	return aitch.NewElement(tagMarker, contents...)
}

// Mask declares a <mask> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/mask
func Mask(contents ...any) aitch.Node {
	return aitch.NewElement(tagMask, contents...)
}

// Metadata declares a <metadata> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/metadata
func Metadata(contents ...any) aitch.Node {
	return aitch.NewElement(tagMetadata, contents...)
}

// Mpath declares a <mpath> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/mpath
func Mpath(contents ...any) aitch.Node {
	return aitch.NewElement(tagMpath, contents...)
}

// Path declares a <path> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/path
func Path(contents ...any) aitch.Node {
	return aitch.NewElement(tagPath, contents...)
}

// Pattern declares a <pattern> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/pattern
func Pattern(contents ...any) aitch.Node {
	return aitch.NewElement(tagPattern, contents...)
}

// Polygon declares a <polygon> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/polygon
func Polygon(contents ...any) aitch.Node {
	return aitch.NewElement(tagPolygon, contents...)
}

// Polyline declares a <polyline> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/polyline
func Polyline(contents ...any) aitch.Node {
	return aitch.NewElement(tagPolyline, contents...)
}

// RadialGradient declares a <radialGradient> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/radialGradient
func RadialGradient(contents ...any) aitch.Node {
	return aitch.NewElement(tagRadialGradient, contents...)
}

// Rect declares a <rect> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/rect
func Rect(contents ...any) aitch.Node {
	return aitch.NewElement(tagRect, contents...)
}

// Script declares a <script> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/script
func Script(contents ...any) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

// Set declares a <set> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/set
func Set(contents ...any) aitch.Node {
	return aitch.NewElement(tagSet, contents...)
}

// SolidColor declares a <solidcolor> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/solidcolor
func SolidColor(contents ...any) aitch.Node {
	return aitch.NewElement(tagSolidColor, contents...)
}

// Stop declares a <stop> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/stop
func Stop(contents ...any) aitch.Node {
	return aitch.NewElement(tagStop, contents...)
}

// StyleElement declares a <style> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/style
func StyleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

// Symbol declares a <symbol> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/symbol
func Symbol(contents ...any) aitch.Node {
	return aitch.NewElement(tagSymbol, contents...)
}

// TextElement declares a <text> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/text
func TextElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagText, contents...)
}

// TextPath declares a <textPath> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/textPath
func TextPath(contents ...any) aitch.Node {
	return aitch.NewElement(tagTextPath, contents...)
}

// Title declares a <title> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/title
func Title(contents ...any) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

// TRef declares a <tref> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/tref
func TRef(contents ...any) aitch.Node {
	return aitch.NewElement(tagTRef, contents...)
}

// TSpan declares a <tspan> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/tspan
func TSpan(contents ...any) aitch.Node {
	return aitch.NewElement(tagTSpan, contents...)
}

// Use declares a <use> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/use
func Use(contents ...any) aitch.Node {
	return aitch.NewElement(tagUse, contents...)
}

// Switch declares a <switch> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/switch
func Switch(contents ...any) aitch.Node {
	return aitch.NewElement(tagSwitch, contents...)
}

// View declares a <view> SVG element
//
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Element/view
func View(contents ...any) aitch.Node {
	return aitch.NewElement(tagView, contents...)
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
	tagSwitch              = []byte("switch")
	tagView                = []byte("view")
)
