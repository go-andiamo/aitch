package aitch

import "io"

var voidElements = map[string]struct{}{
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

type voidElement struct {
	name             []byte
	attributes       []Node
	attIndices       map[string]int
	conditionalAttrs []Node
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

func renderAttributes(ctx *Context, attrs []Node, conditionalAttrs []Node) {
	if len(conditionalAttrs) == 0 {
		for _, attr := range attrs {
			_ = attr.Render(ctx.w, ctx)
		}
	} else {
		//TODO
	}
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

type element struct {
	name             []byte
	attributes       []Node
	attIndices       map[string]int
	conditionalAttrs []Node
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
func Element(name string, contents ...Node) Node {
	if !htmlTagRegex.MatchString(name) {
		if PanicOnInvalidName {
			panic("invalid html tag name: " + name)
		}
		return nil
	}
	if _, void := voidElements[name]; void {
		return newVoidElement([]byte(name), contents...)
	}
	return newElement([]byte(name), contents...)
}

func newElement(name []byte, contents ...Node) Node {
	attrs, children := attributesAndContents(contents...)
	result := &element{
		name:       []byte(name),
		attributes: make([]Node, 0, len(attrs)),
		attIndices: make(map[string]int, len(attrs)),
		contents:   children,
	}
	return result.addAttributes(attrs)
}

func newVoidElement(name []byte, contents ...Node) Node {
	attrs, _ := attributesAndContents(contents...)
	result := &voidElement{
		name:       []byte(name),
		attributes: make([]Node, 0, len(attrs)),
		attIndices: make(map[string]int, len(attrs)),
	}
	return result.addAttributes(attrs)
}

func attributesAndContents(nodes ...Node) (attributes []Node, contents []Node) {
	attributes = make([]Node, 0, len(nodes))
	contents = make([]Node, 0, len(nodes))
	for _, item := range nodes {
		if item != nil {
			switch item.Type() {
			case AttributeNode:
				attributes = append(attributes, item)
			case collectionNode:
				contents = append(contents, item)
				if cn, ok := item.(*collection); ok {
					addA, _ := attributesAndContents(cn.nodes...)
					attributes = append(attributes, addA...)
				}
			case conditionalNode:
				contents = append(contents, item)
				//TODO more!!!
			default:
				contents = append(contents, item)
			}
		}
	}
	return
}
