package svg

import (
	"golang.org/x/net/html"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// GenerateCode reads SVG and writes fluent SVG code
func GenerateCode(r io.Reader, w io.Writer) error {
	if w == nil {
		w = os.Stdout
	}
	elemMap, attrMap := mappings()
	tokenizer := html.NewTokenizer(r)
	indent := []byte{}
	eof := false
	var err error
	nl := []byte{'\n'}
	tab := byte('\t')
	funcName := func(fn any) string {
		v := reflect.ValueOf(fn)
		result := "Unknown"
		if v.Kind() == reflect.Func {
			if fpp := runtime.FuncForPC(v.Pointer()); fpp != nil {
				name := fpp.Name()
				if ls := strings.LastIndex(name, "/"); ls > 0 {
					result = name[ls+1:]
				}
			}
		}
		return result
	}
	write := func(s string) {
		_, err = w.Write(indent)
		_, err = w.Write([]byte(s))
		_, err = w.Write(nl)
	}
	incIndent := func() {
		indent = append(indent, tab)
	}
	decIndent := func() {
		if len(indent) > 0 {
			indent = indent[:len(indent)-1]
		}
	}
	escape := func(s string) string {
		return strings.Replace(
			strings.Replace(
				strings.Replace(strings.Replace(s, "\\", `\\`, -1), "\t", `\t`, -1), "\n", `\n`, -1),
			`"`, `\"`, -1)
	}
	writeElementStart := func(name string, attrs []html.Attribute, selfClosing bool) {
		start := ""
		if elem, ok := elemMap[strings.ToLower(string(name))]; ok {
			start = funcName(elem) + "("
		} else {
			start = `aitch.Element("` + string(name) + `", `
		}
		_, err = w.Write(indent)
		_, err = w.Write([]byte(start))
		if selfClosing && len(attrs) == 0 {
			_, err = w.Write([]byte{')', ',', '\n'})
			return
		}
		for i, a := range attrs {
			if i > 0 {
				_, err = w.Write([]byte{',', ' '})
			}
			if attr, ok := attrMap[strings.ToLower(a.Key)]; ok {
				_, err = w.Write([]byte(funcName(attr)))
				_, err = w.Write([]byte{'(', '"'})
				_, err = w.Write([]byte(escape(a.Val)))
				_, err = w.Write([]byte{'"', ')'})
			} else {
				_, err = w.Write([]byte(`aitch.Attribute("`))
				_, err = w.Write([]byte(a.Key))
				_, err = w.Write([]byte{'"', ' ', ',', '"'})
				_, err = w.Write([]byte(escape(a.Val)))
				_, err = w.Write([]byte{'"', ')'})
			}
		}
		if selfClosing {
			_, err = w.Write([]byte{')', ','})
		} else if len(attrs) > 0 {
			_, err = w.Write([]byte{','})
		}
		_, err = w.Write([]byte{'\n'})
		if !selfClosing {
			incIndent()
		}
	}
	write("var template = aitch.Collection(")
	incIndent()
	for err == nil && !eof {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			if err = tokenizer.Err(); err == io.EOF {
				eof = true
				err = nil
			}
		case html.TextToken:
			token := tokenizer.Token()
			if strings.Trim(token.Data, " \n\t") != "" {
				write(`svg.Text("` + escape(token.Data) + `"),`)
			}
		case html.StartTagToken:
			token := tokenizer.Token()
			writeElementStart(token.Data, token.Attr, false)
		case html.SelfClosingTagToken:
			token := tokenizer.Token()
			writeElementStart(token.Data, token.Attr, true)
		case html.EndTagToken:
			decIndent()
			write("),")
		case html.CommentToken:
			token := tokenizer.Token()
			write(`svg.Comment("` + escape(token.Data) + `"),`)
		}
	}
	if err == nil {
		decIndent()
		write(")\n")
	}
	return err
}

func mappings() (elems map[string]any, attrs map[string]any) {
	elems = map[string]any{
		"altGlyph":            AltGlyph,
		"animate":             Animate,
		"animateMotion":       AnimateMotion,
		"animateTransform":    AnimateTransform,
		"circle":              Circle,
		"clipPath":            ClipPath,
		"defs":                Defs,
		"desc":                Desc,
		"discard":             Discard,
		"ellipse":             Ellipse,
		"feBlend":             FeBlend,
		"feColorMatrix":       FeColorMatrix,
		"feComponentTransfer": FeComponentTransfer,
		"feComposite":         FeComposite,
		"feConvolveMatrix":    FeConvolveMatrix,
		"feDiffuseLighting":   FeDiffuseLighting,
		"feDisplacementMap":   FeDisplacementMap,
		"feDistantLight":      FeDistantLight,
		"feDropShadow":        FeDropShadow,
		"feFlood":             FeFlood,
		"feFuncA":             FeFuncA,
		"feFuncB":             FeFuncB,
		"feFuncG":             FeFuncG,
		"feFuncR":             FeFuncR,
		"feGaussianBlur":      FeGaussianBlur,
		"feImage":             FeImage,
		"feMerge":             FeMerge,
		"feMergeNode":         FeMergeNode,
		"feMorphology":        FeMorphology,
		"feOffset":            FeOffset,
		"fePointLight":        FePointLight,
		"feSpecularLighting":  FeSpecularLighting,
		"feSpotLight":         FeSpotLight,
		"feTile":              FeTile,
		"feTurbulence":        FeTurbulence,
		"filter":              Filter,
		"foreignObject":       ForeignObject,
		"g":                   G,
		"glyph":               Glyph,
		"line":                Line,
		"linearGradient":      LinearGradient,
		"marker":              Marker,
		"mask":                Mask,
		"metadata":            Metadata,
		"mpath":               Mpath,
		"path":                Path,
		"pattern":             Pattern,
		"polygon":             Polygon,
		"polyline":            Polyline,
		"radialGradient":      RadialGradient,
		"rect":                Rect,
		"script":              Script,
		"set":                 Set,
		"solidcolor":          SolidColor,
		"stop":                Stop,
		"style":               StyleElement,
		"svg":                 Svg,
		"symbol":              Symbol,
		"text":                TextElement,
		"textPath":            TextPath,
		"title":               Title,
		"tref":                TRef,
		"tspan":               TSpan,
		"use":                 Use,
	}
	attrs = map[string]any{
		"accumulate":                   Accumulate,
		"additive":                     Additive,
		"alignment-baseline":           AlignmentBaseline,
		"attributeName":                AttributeName,
		"attributeType":                AttributeType,
		"baseline-shift":               BaselineShift,
		"begin":                        Begin,
		"by":                           By,
		"calcMode":                     CalcMode,
		"class":                        Class,
		"clip":                         Clip,
		"clip-path":                    ClipPathAttr,
		"clip-rule":                    ClipRule,
		"color":                        Color,
		"color-interpolation":          ColorInterpolation,
		"color-interpolation-filters":  ColorInterpolationFilters,
		"color-profile":                ColorProfile,
		"color-rendering":              ColorRendering,
		"cursor":                       Cursor,
		"cx":                           Cx,
		"cy":                           Cy,
		"d":                            D,
		"direction":                    Direction,
		"display":                      Display,
		"dominant-baseline":            DominantBaseline,
		"dur":                          Dur,
		"dx":                           Dx,
		"dy":                           Dy,
		"enable-background":            EnableBackground,
		"end":                          End,
		"externalResourcesRequired":    ExternalResourcesRequired,
		"fill":                         Fill,
		"fill-opacity":                 FillOpacity,
		"fill-rule":                    FillRule,
		"filter":                       FilterAttr,
		"filterUnits":                  FilterUnits,
		"flood-color":                  FloodColor,
		"flood-opacity":                FloodOpacity,
		"font-family":                  FontFamily,
		"font-size":                    FontSize,
		"font-size-adjust":             FontSizeAdjust,
		"font-stretch":                 FontStretch,
		"font-style":                   FontStyle,
		"font-variant":                 FontVariant,
		"font-weight":                  FontWeight,
		"from":                         From,
		"fx":                           Fx,
		"fy":                           Fy,
		"glyph-orientation-horizontal": GlyphOrientationHorizontal,
		"glyph-orientation-vertical":   GlyphOrientationVertical,
		"gradientTransform":            GradientTransform,
		"gradientUnits":                GradientUnits,
		"height":                       Height,
		"href":                         Href,
		"id":                           Id,
		"image-rendering":              ImageRendering,
		"kerning":                      Kerning,
		"keyPoints":                    KeyPoints,
		"keySplines":                   KeySplines,
		"keyTimes":                     KeyTimes,
		"lengthAdjust":                 LengthAdjust,
		"letter-spacing":               LetterSpacing,
		"lighting-color":               LightingColor,
		"marker-end":                   MarkerEnd,
		"marker-mid":                   MarkerMid,
		"marker-start":                 MarkerStart,
		"markerHeight":                 MarkerHeight,
		"markerUnits":                  MarkerUnits,
		"markerWidth":                  MarkerWidth,
		"mask":                         MaskAttr,
		"onblur":                       OnBlur,
		"onclick":                      OnClick,
		"onfocus":                      OnFocus,
		"onload":                       OnLoad,
		"onmousedown":                  OnMouseDown,
		"onmousemove":                  OnMouseMove,
		"onmouseout":                   OnMouseOut,
		"onmouseover":                  OnMouseOver,
		"onmouseup":                    OnMouseUp,
		"onunload":                     OnUnload,
		"opacity":                      Opacity,
		"orient":                       Orient,
		"path":                         PathAttr,
		"pathLength":                   PathLength,
		"patternContentUnits":          PatternContentUnits,
		"patternTransform":             PatternTransform,
		"patternUnits":                 PatternUnits,
		"pointer-events":               PointerEvents,
		"points":                       Points,
		"preserveAspectRatio":          PreserveAspectRatio,
		"primitiveUnits":               PrimitiveUnits,
		"r":                            R,
		"refX":                         RefX,
		"refY":                         RefY,
		"repeatCount":                  RepeatCount,
		"repeatDur":                    RepeatDur,
		"requiredExtensions":           RequiredExtensions,
		"requiredFeatures":             RequiredFeatures,
		"rotate":                       Rotate,
		"rx":                           Rx,
		"ry":                           Ry,
		"shape-rendering":              ShapeRendering,
		"spreadMethod":                 SpreadMethod,
		"stop-color":                   StopColor,
		"stop-opacity":                 StopOpacity,
		"stroke":                       Stroke,
		"stroke-dasharray":             StrokeDashArray,
		"stroke-dashoffset":            StrokeDashOffset,
		"stroke-linecap":               StrokeLineCap,
		"stroke-linejoin":              StrokeLineJoin,
		"stroke-miterlimit":            StrokeMiterLimit,
		"stroke-opacity":               StrokeOpacity,
		"stroke-width":                 StrokeWidth,
		"style":                        Style,
		"systemLanguage":               SystemLanguage,
		"tabindex":                     TabIndex,
		"text-anchor":                  TextAnchor,
		"text-decoration":              TextDecoration,
		"text-rendering":               TextRendering,
		"textLength":                   TextLength,
		"to":                           To,
		"type":                         Type,
		"unicode-bidi":                 UnicodeBidi,
		"values":                       Values,
		"visibility":                   Visibility,
		"width":                        Width,
		"word-spacing":                 WordSpacing,
		"writing-mode":                 WritingMode,
		"x":                            X,
		"x1":                           X1,
		"x2":                           X2,
		"xlink:actuate":                XlinkActuate,
		"xlink:arcrole":                XlinkArcrole,
		"xlink:href":                   XlinkHref,
		"xlink:role":                   XlinkRole,
		"xlink:show":                   XlinkShow,
		"xlink:title":                  XlinkTitle,
		"xml:base":                     XmlBase,
		"xml:lang":                     XmlLang,
		"xml:space":                    XmlSpace,
		"y":                            Y,
		"y1":                           Y1,
		"y2":                           Y2,
	}
	return
}
