package svg

import "github.com/go-andiamo/aitch"

func Svg(contents ...any) aitch.Node {
	return aitch.NewElement(tagSvg, contents...)
}

func AltGlyph(contents ...any) aitch.Node {
	return aitch.NewElement(tagAltGlyph, contents...)
}

func Animate(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimate, contents...)
}

func AnimateMotion(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimateMotion, contents...)
}

func AnimateTransform(contents ...any) aitch.Node {
	return aitch.NewElement(tagAnimateTransform, contents...)
}

func Circle(contents ...any) aitch.Node {
	return aitch.NewElement(tagCircle, contents...)
}

func ClipPath(contents ...any) aitch.Node {
	return aitch.NewElement(tagClipPath, contents...)
}

func Defs(contents ...any) aitch.Node {
	return aitch.NewElement(tagDefs, contents...)
}

func Desc(contents ...any) aitch.Node {
	return aitch.NewElement(tagDesc, contents...)
}

func Discard(contents ...any) aitch.Node {
	return aitch.NewElement(tagDiscard, contents...)
}

func Ellipse(contents ...any) aitch.Node {
	return aitch.NewElement(tagEllipse, contents...)
}

func FeBlend(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeBlend, contents...)
}

func FeColorMatrix(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeColorMatrix, contents...)
}

func FeComponentTransfer(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeComponentTransfer, contents...)
}

func FeComposite(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeComposite, contents...)
}

func FeConvolveMatrix(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeConvolveMatrix, contents...)
}

func FeDiffuseLighting(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDiffuseLighting, contents...)
}

func FeDisplacementMap(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDisplacementMap, contents...)
}

func FeDistantLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDistantLight, contents...)
}

func FeDropShadow(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeDropShadow, contents...)
}

func FeFlood(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFlood, contents...)
}

func FeFuncA(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncA, contents...)
}

func FeFuncB(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncB, contents...)
}

func FeFuncG(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncG, contents...)
}

func FeFuncR(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeFuncR, contents...)
}

func FeGaussianBlur(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeGaussianBlur, contents...)
}

func FeImage(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeImage, contents...)
}

func FeMerge(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMerge, contents...)
}

func FeMergeNode(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMergeNode, contents...)
}

func FeMorphology(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeMorphology, contents...)
}

func FeOffset(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeOffset, contents...)
}

func FePointLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFePointLight, contents...)
}

func FeSpecularLighting(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeSpecularLighting, contents...)
}

func FeSpotLight(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeSpotLight, contents...)
}

func FeTile(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeTile, contents...)
}

func FeTurbulence(contents ...any) aitch.Node {
	return aitch.NewElement(tagFeTurbulence, contents...)
}

func Filter(contents ...any) aitch.Node {
	return aitch.NewElement(tagFilter, contents...)
}

func ForeignObject(contents ...any) aitch.Node {
	return aitch.NewElement(tagForeignObject, contents...)
}

func G(contents ...any) aitch.Node {
	return aitch.NewElement(tagG, contents...)
}

func Glyph(contents ...any) aitch.Node {
	return aitch.NewElement(tagGlyph, contents...)
}

func Line(contents ...any) aitch.Node {
	return aitch.NewElement(tagLine, contents...)
}

func LinearGradient(contents ...any) aitch.Node {
	return aitch.NewElement(tagLinearGradient, contents...)
}

func Marker(contents ...any) aitch.Node {
	return aitch.NewElement(tagMarker, contents...)
}

func Mask(contents ...any) aitch.Node {
	return aitch.NewElement(tagMask, contents...)
}

func Metadata(contents ...any) aitch.Node {
	return aitch.NewElement(tagMetadata, contents...)
}

func Mpath(contents ...any) aitch.Node {
	return aitch.NewElement(tagMpath, contents...)
}

func Path(contents ...any) aitch.Node {
	return aitch.NewElement(tagPath, contents...)
}

func Pattern(contents ...any) aitch.Node {
	return aitch.NewElement(tagPattern, contents...)
}

func Polygon(contents ...any) aitch.Node {
	return aitch.NewElement(tagPolygon, contents...)
}

func Polyline(contents ...any) aitch.Node {
	return aitch.NewElement(tagPolyline, contents...)
}

func RadialGradient(contents ...any) aitch.Node {
	return aitch.NewElement(tagRadialGradient, contents...)
}

func Rect(contents ...any) aitch.Node {
	return aitch.NewElement(tagRect, contents...)
}

func Script(contents ...any) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

func Set(contents ...any) aitch.Node {
	return aitch.NewElement(tagSet, contents...)
}

func SolidColor(contents ...any) aitch.Node {
	return aitch.NewElement(tagSolidColor, contents...)
}

func Stop(contents ...any) aitch.Node {
	return aitch.NewElement(tagStop, contents...)
}

func StyleElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

func Symbol(contents ...any) aitch.Node {
	return aitch.NewElement(tagSymbol, contents...)
}

func TextElement(contents ...any) aitch.Node {
	return aitch.NewElement(tagText, contents...)
}

func TextPath(contents ...any) aitch.Node {
	return aitch.NewElement(tagTextPath, contents...)
}

func Title(contents ...any) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

func TRef(contents ...any) aitch.Node {
	return aitch.NewElement(tagTRef, contents...)
}

func TSpan(contents ...any) aitch.Node {
	return aitch.NewElement(tagTSpan, contents...)
}

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
