package svg

import "github.com/go-andiamo/aitch"

func Svg(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagSvg, contents...)
}

func AltGlyph(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagAltGlyph, contents...)
}

func Animate(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagAnimate, contents...)
}

func AnimateMotion(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagAnimateMotion, contents...)
}

func AnimateTransform(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagAnimateTransform, contents...)
}

func Circle(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagCircle, contents...)
}

func ClipPath(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagClipPath, contents...)
}

func Defs(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagDefs, contents...)
}

func Desc(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagDesc, contents...)
}

func Discard(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagDiscard, contents...)
}

func Ellipse(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagEllipse, contents...)
}

func FeBlend(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeBlend, contents...)
}

func FeColorMatrix(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeColorMatrix, contents...)
}

func FeComponentTransfer(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeComponentTransfer, contents...)
}

func FeComposite(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeComposite, contents...)
}

func FeConvolveMatrix(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeConvolveMatrix, contents...)
}

func FeDiffuseLighting(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeDiffuseLighting, contents...)
}

func FeDisplacementMap(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeDisplacementMap, contents...)
}

func FeDistantLight(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeDistantLight, contents...)
}

func FeDropShadow(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeDropShadow, contents...)
}

func FeFlood(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeFlood, contents...)
}

func FeFuncA(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeFuncA, contents...)
}

func FeFuncB(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeFuncB, contents...)
}

func FeFuncG(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeFuncG, contents...)
}

func FeFuncR(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeFuncR, contents...)
}

func FeGaussianBlur(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeGaussianBlur, contents...)
}

func FeImage(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeImage, contents...)
}

func FeMerge(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeMerge, contents...)
}

func FeMergeNode(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeMergeNode, contents...)
}

func FeMorphology(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeMorphology, contents...)
}

func FeOffset(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeOffset, contents...)
}

func FePointLight(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFePointLight, contents...)
}

func FeSpecularLighting(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeSpecularLighting, contents...)
}

func FeSpotLight(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeSpotLight, contents...)
}

func FeTile(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeTile, contents...)
}

func FeTurbulence(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFeTurbulence, contents...)
}

func Filter(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagFilter, contents...)
}

func ForeignObject(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagForeignObject, contents...)
}

func G(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagG, contents...)
}

func Glyph(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagGlyph, contents...)
}

func Line(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagLine, contents...)
}

func LinearGradient(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagLinearGradient, contents...)
}

func Marker(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagMarker, contents...)
}

func Mask(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagMask, contents...)
}

func Metadata(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagMetadata, contents...)
}

func Mpath(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagMpath, contents...)
}

func Path(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagPath, contents...)
}

func Pattern(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagPattern, contents...)
}

func Polygon(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagPolygon, contents...)
}

func Polyline(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagPolyline, contents...)
}

func RadialGradient(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagRadialGradient, contents...)
}

func Rect(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagRect, contents...)
}

func Script(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagScript, contents...)
}

func Set(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagSet, contents...)
}

func SolidColor(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagSolidColor, contents...)
}

func Stop(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagStop, contents...)
}

func StyleElement(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagStyle, contents...)
}

func Symbol(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagSymbol, contents...)
}

func TextElement(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagText, contents...)
}

func TextPath(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagTextPath, contents...)
}

func Title(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagTitle, contents...)
}

func TRef(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagTRef, contents...)
}

func TSpan(contents ...aitch.Node) aitch.Node {
	return aitch.NewElement(tagTSpan, contents...)
}

func Use(contents ...aitch.Node) aitch.Node {
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
