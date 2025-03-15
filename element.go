package aitch

import "io"

func isVoidElement(name string) bool {
	_, ok := voidElements[name]
	return ok
}

var voidElements = map[string]struct{}{
	"area":    {},
	"base":    {},
	"br":      {},
	"col":     {},
	"command": {},
	"embed":   {},
	"hr":      {},
	"img":     {},
	"input":   {},
	"keygen":  {},
	"link":    {},
	"meta":    {},
	"param":   {},
	"source":  {},
	"track":   {},
	"wbr":     {},
}

type voidElement struct {
	name             []byte
	attributes       []Node
	attIndices       map[string]int
	conditionalAttrs []Node
}

func (e *voidElement) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(openAngleBracket); ctx.Error == nil {
			if _, ctx.Error = w.Write(e.name); ctx.Error == nil {
				if ctx.Error = renderAttributes(w, ctx, e.attributes, e.conditionalAttrs); ctx.Error == nil {
					_, ctx.Error = w.Write(closeAngleBracket)
				}
			}
		}
	}
	return ctx.Error
}

func renderAttributes(w io.Writer, ctx *Context, attrs []Node, conditionalAttrs []Node) error {
	if len(conditionalAttrs) == 0 {
		for _, attr := range attrs {
			if ctx.Error = attr.Render(w, ctx); ctx.Error != nil {
				return ctx.Error
			}
		}
	} else {
		//TODO
	}
	return ctx.Error
}

func (e *voidElement) Type() NodeType {
	return ElementNode
}

func (e *voidElement) Name() string {
	return string(e.name)
}

func (e *voidElement) Append(nodes ...Node) Node {
	attrs, _ := attributesAndContents(nodes...)
	e.addAttributes(attrs)
	return e
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
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(openAngleBracket); ctx.Error == nil {
			if _, ctx.Error = w.Write(e.name); ctx.Error == nil {
				if ctx.Error = renderAttributes(w, ctx, e.attributes, e.conditionalAttrs); ctx.Error != nil {
					return ctx.Error
				}
				if _, ctx.Error = w.Write(closeAngleBracket); ctx.Error == nil {
					for _, c := range e.contents {
						if ctx.Error = c.Render(w, ctx); ctx.Error != nil {
							return ctx.Error
						}
					}
				}
				if _, ctx.Error = w.Write(closeTagStart); ctx.Error == nil {
					if _, ctx.Error = w.Write(e.name); ctx.Error == nil {
						_, ctx.Error = w.Write(closeAngleBracket)
					}
				}
			}
		}
	}
	return ctx.Error
}

func (e *element) Type() NodeType {
	return ElementNode
}

func (e *element) Name() string {
	return string(e.name)
}

func (e *element) Append(nodes ...Node) Node {
	attrs, contents := attributesAndContents(nodes...)
	e.addAttributes(attrs)
	e.contents = append(e.contents, contents...)
	return e
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
	return newElement(name, contents...)
}

func newElement(name string, contents ...Node) Node {
	if isVoidElement(name) {
		attrs, _ := attributesAndContents(contents...)
		result := &voidElement{
			name:       []byte(name),
			attributes: make([]Node, 0, len(attrs)),
			attIndices: make(map[string]int, len(attrs)),
		}
		return result.addAttributes(attrs)

	}
	attrs, children := attributesAndContents(contents...)
	result := &element{
		name:       []byte(name),
		attributes: make([]Node, 0, len(attrs)),
		attIndices: make(map[string]int, len(attrs)),
		contents:   children,
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
				if cn, ok := item.(*collection); ok {
					addA, addC := attributesAndContents(cn.nodes...)
					attributes = append(attributes, addA...)
					contents = append(contents, addC...)
				} else {
					contents = append(contents, item)
				}
			case conditionalNode:
				panic("implement me!")
			default:
				contents = append(contents, item)
			}
		}
	}
	return
}
