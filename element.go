package aitch

import (
	"github.com/go-andiamo/aitch/context"
)

func renderAttributes(ctx *context.Context, attributes []Node, attIndices map[string]int, conditionals conditionalAttributes) {
	if len(conditionals) > 0 {
		if evaluated := conditionals.evaluate(ctx); len(evaluated) > 0 {
			use := &delimitedAttribute{}
			for name, index := range attIndices {
				attr := attributes[index]
				if eAttr, ok := evaluated[name]; ok {
					switch at := attr.(type) {
					case *delimitedAttribute:
						use.name = at.name
						use.delimiter = at.delimiter
						use.values = use.values[:0]
						use.values = append(use.values, at.values...)
						for _, e := range eAttr {
							if et, ok := e.(valuesNode); ok {
								use.values = append(use.values, et.getValues()...)
							}
						}
						_ = use.Render(ctx)
					default:
						_ = eAttr[len(eAttr)-1].Render(ctx)
					}
				} else {
					_ = attr.Render(ctx)
				}
			}
			for name, ca := range evaluated {
				if _, ok := attIndices[name]; !ok && len(ca) > 0 {
					first := ca[0]
					switch at := first.(type) {
					case *delimitedAttribute:
						use.name = at.name
						use.delimiter = at.delimiter
						use.values = use.values[:0]
						use.values = append(use.values, at.values...)
						for i := 1; i < len(ca); i++ {
							if et, ok := ca[i].(valuesNode); ok {
								use.values = append(use.values, et.getValues()...)
							}
						}
						_ = use.Render(ctx)
					default:
						_ = ca[len(ca)-1].Render(ctx)
					}
				}
			}
			return
		}
	}
	for _, attr := range attributes {
		_ = attr.Render(ctx)
	}
}

type voidElement struct {
	name         []byte
	attributes   []Node
	attIndices   map[string]int
	conditionals conditionalAttributes
}

func (e *voidElement) Render(ctx *context.Context) error {
	ctx.Write(openAngleBracket)
	ctx.Write(e.name)
	renderAttributes(ctx, e.attributes, e.attIndices, e.conditionals)
	ctx.Write(closeAngleBracket)
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
func VoidElement(name string, contents ...any) Node {
	if !htmlTagRegex.MatchString(name) {
		if PanicOnInvalidName {
			panic("invalid html tag name: " + name)
		}
		return nil
	}
	return NewVoidElement([]byte(name), contents...)
}

type element struct {
	name         []byte
	attributes   []Node
	attIndices   map[string]int
	conditionals conditionalAttributes
	contents     []Node
}

func (e *element) Render(ctx *context.Context) error {
	ctx.Write(openAngleBracket)
	ctx.Write(e.name)
	renderAttributes(ctx, e.attributes, e.attIndices, e.conditionals)
	ctx.Write(closeAngleBracket)
	for _, c := range e.contents {
		_ = c.Render(ctx)
	}
	ctx.Write(closeTagStart)
	ctx.Write(e.name)
	ctx.Write(closeAngleBracket)
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

// Element declares a new element Node
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
//
// The name is used to determine whether the element is an HTML void element - you want
// to specifically create a void element use VoidElement() instead
func Element(name string, contents ...any) Node {
	if !htmlTagRegex.MatchString(name) {
		if PanicOnInvalidName {
			panic("invalid html tag name: " + name)
		}
		return nil
	}
	if _, void := voidElementNames[name]; void {
		return NewVoidElement([]byte(name), contents...)
	}
	return NewElement([]byte(name), contents...)
}

// NewElement declares a new element Node
func NewElement(name []byte, contents ...any) Node {
	attrs, condAttrs, children := attributesAndContents(nil, contentsToNodes(contents))
	result := &element{
		name:         name,
		attributes:   make([]Node, 0, len(attrs)),
		attIndices:   make(map[string]int, len(attrs)),
		contents:     children,
		conditionals: condAttrs,
	}
	return result.addAttributes(attrs)
}

// NewVoidElement declares a new void element Node
//
// A void element has no content (and no closing tag) - therefore, any non-attribute
// contents are ignored
func NewVoidElement(name []byte, contents ...any) Node {
	attrs, condAttrs, _ := attributesAndContents(nil, contentsToNodes(contents))
	result := &voidElement{
		name:         name,
		attributes:   make([]Node, 0, len(attrs)),
		attIndices:   make(map[string]int, len(attrs)),
		conditionals: condAttrs,
	}
	return result.addAttributes(attrs)
}

func contentsToNodes(contents []any) []Node {
	result := make([]Node, len(contents))
	for i, c := range contents {
		switch v := c.(type) {
		case Node:
			result[i] = v
		default:
			result[i] = Text(c)
		}
	}
	return result
}

func attributesAndContents(conditions []ConditionalFunc, content []Node) (attrs []Node, condAttrs conditionalAttributes, contents []Node) {
	attrs = make([]Node, 0, len(content))
	condAttrs = make(conditionalAttributes, 0, len(content))
	contents = make([]Node, 0, len(content))
	for _, item := range content {
		if item != nil {
			if item.Type() == AttributeNode {
				if len(conditions) > 0 {
					condAttrs = append(condAttrs, conditionalAttribute{
						attribute:  item,
						conditions: conditions,
					})
				} else {
					attrs = append(attrs, item)
				}
			} else {
				contents = append(contents, item)
				switch cn := item.(type) {
				case *collection:
					addA, addCa, _ := attributesAndContents(conditions, cn.nodes)
					attrs = append(attrs, addA...)
					condAttrs = append(condAttrs, addCa...)
				case *conditional:
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
