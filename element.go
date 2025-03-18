package aitch

import "io"

func renderAttributes(ctx *Context, attrs []Node, conditionals conditionalAttributes) {
	if len(conditionals) == 0 {
		for _, attr := range attrs {
			_ = attr.Render(ctx.w, ctx)
		}
	} else {
		panic(len(conditionals))
		/* TODO
		1. evaluate all conditional atts to determine the ones applicable
		2. merge applicable conditional atts with actual atts
		3. render final atts
		*/
	}
}

type voidElement struct {
	name             []byte
	attributes       []Node
	attIndices       map[string]int
	conditionalAttrs conditionalAttributes
}

func (e *voidElement) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(openAngleBracket)
	ctx.write(e.name)
	renderAttributes(ctx, e.attributes, e.conditionalAttrs)
	ctx.write(closeAngleBracket)
	return ctx.Error
}

func (e *voidElement) Type() NodeType {
	return ElementNode
}

func (e *voidElement) Name() string {
	return string(e.name)
}

func (e *voidElement) addAttributes(attrs []Node) Node {
	for _, att := range attrs {
		name := att.Name()
		if idx, ok := e.attIndices[name]; ok {
			existing := e.attributes[idx]
			if delimited, ok := existing.(*delimitedAttribute); ok {
				if addD, ok := att.(*delimitedAttribute); ok {
					delimited.values = append(delimited.values, addD.values...)
					continue
				}
			}
			e.attributes[idx] = att
		} else {
			e.attIndices[name] = len(e.attributes)
			e.attributes = append(e.attributes, att)
		}
	}
	return e
}

// VoidElement creates a new HTML void element Node
//
// a void element has no content (other than attributes) and has no closing tag
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
func VoidElement(name string, contents ...Node) Node {
	if !htmlTagRegex.MatchString(name) {
		if PanicOnInvalidName {
			panic("invalid html tag name: " + name)
		}
		return nil
	}
	return newVoidElement([]byte(name), contents...)
}

type element struct {
	name             []byte
	attributes       []Node
	attIndices       map[string]int
	conditionalAttrs conditionalAttributes
	contents         []Node
}

func (e *element) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(openAngleBracket)
	ctx.write(e.name)
	renderAttributes(ctx, e.attributes, e.conditionalAttrs)
	ctx.write(closeAngleBracket)
	for _, c := range e.contents {
		_ = c.Render(ctx.w, ctx)
	}
	ctx.write(closeTagStart)
	ctx.write(e.name)
	ctx.write(closeAngleBracket)
	return ctx.Error
}

func (e *element) Type() NodeType {
	return ElementNode
}

func (e *element) Name() string {
	return string(e.name)
}

func (e *element) addAttributes(attrs []Node) Node {
	for _, att := range attrs {
		name := att.Name()
		if idx, ok := e.attIndices[name]; ok {
			existing := e.attributes[idx]
			if delimited, ok := existing.(*delimitedAttribute); ok {
				if addD, ok := att.(*delimitedAttribute); ok {
					delimited.values = append(delimited.values, addD.values...)
					continue
				}
			}
			e.attributes[idx] = att
		} else {
			e.attIndices[name] = len(e.attributes)
			e.attributes = append(e.attributes, att)
		}
	}
	return e
}

// Element creates a new element Node
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
//
// The name is used to determine whether the element is an HTML void element - you want
// to specifically create a void element use VoidElement() instead
func Element(name string, contents ...Node) Node {
	if !htmlTagRegex.MatchString(name) {
		if PanicOnInvalidName {
			panic("invalid html tag name: " + name)
		}
		return nil
	}
	if _, void := voidElementNames[name]; void {
		return newVoidElement([]byte(name), contents...)
	}
	return newElement([]byte(name), contents...)
}

func newElement(name []byte, contents ...Node) Node {
	attrs, condAttrs, children := attributesAndContents(nil, contents)
	result := &element{
		name:             name,
		attributes:       make([]Node, 0, len(attrs)),
		attIndices:       make(map[string]int, len(attrs)),
		contents:         children,
		conditionalAttrs: condAttrs,
	}
	return result.addAttributes(attrs)
}

func newVoidElement(name []byte, contents ...Node) Node {
	attrs, condAttrs, _ := attributesAndContents(nil, contents)
	result := &voidElement{
		name:             name,
		attributes:       make([]Node, 0, len(attrs)),
		attIndices:       make(map[string]int, len(attrs)),
		conditionalAttrs: condAttrs,
	}
	return result.addAttributes(attrs)
}

func attributesAndContents(conditions []ConditionalFunc, nodes []Node) (attributes []Node, condAttrs conditionalAttributes, contents []Node) {
	attributes = make([]Node, 0, len(nodes))
	condAttrs = make(conditionalAttributes, 0, len(nodes))
	contents = make([]Node, 0, len(nodes))
	for _, item := range nodes {
		if item != nil {
			switch item.Type() {
			case AttributeNode:
				if len(conditions) > 0 {
					condAttrs = append(condAttrs, conditionalAttribute{
						attribute:  item,
						conditions: conditions,
					})
				} else {
					attributes = append(attributes, item)
				}
			case collectionNode:
				contents = append(contents, item)
				if cn, ok := item.(*collection); ok {
					addA, addCa, _ := attributesAndContents(conditions, cn.nodes)
					attributes = append(attributes, addA...)
					condAttrs = append(condAttrs, addCa...)
				}
			case conditionalNode:
				contents = append(contents, item)
				if cn, ok := item.(*conditional); ok {
					newConditions := append(conditions, cn.fn)
					for _, a := range cn.attributes {
						condAttrs = append(condAttrs, conditionalAttribute{
							attribute:  a,
							conditions: newConditions,
						})
					}
					_, addCa, _ := attributesAndContents(newConditions, cn.nodes)
					condAttrs = append(condAttrs, addCa...)
				}
			default:
				contents = append(contents, item)
			}
		}
	}
	return
}

var voidElementNames = map[string]struct{}{
	"area":   {},
	"base":   {},
	"br":     {},
	"col":    {},
	"embed":  {},
	"hr":     {},
	"img":    {},
	"input":  {},
	"link":   {},
	"meta":   {},
	"param":  {},
	"source": {},
	"track":  {},
	"wbr":    {},
}
