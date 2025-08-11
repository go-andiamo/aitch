package aitch

import "github.com/go-andiamo/aitch/context"

// ImperativeContext is an interface that is passed to the func supplied to Imperative during rendering
type ImperativeContext interface {
	// Context obtains the outer rendering context
	Context() *context.Context
	// WriteString writes the string (html escaped) to the output
	WriteString(s string) ImperativeContext
	// WriteRaw writes raw byte data to the output
	WriteRaw(data []byte) ImperativeContext
	// WriteNodes writes nodes to the output
	//
	// Note: attribute type nodes will be ignored
	WriteNodes(nodes ...Node) ImperativeContext
	// WriteElement starts and ends an element
	//
	// the element is assumed to be non-void (i.e. has a closing tag)
	//
	// the content arg can be a Node, string, []byte etc. (non-handled types are ignored)
	WriteElement(elementName []byte, content any, attrs ...Node) ImperativeContext
	// Start starts a new element
	//
	// If any of the attrs args are not AttributeNode types - they are ignored
	Start(elementName []byte, isVoid bool, attrs ...Node) ImperativeContext
	// End ends the last started element (as started with Start)
	End() ImperativeContext
	// EndAll ends all started elements
	EndAll() ImperativeContext
	// Marker returns the current length of started elements
	Marker() int
	// EndToMark ends all started elements to the supplied marker
	EndToMark(marker int) ImperativeContext
}

// Imperative creates a new imperative Node
//
// An imperative node allows you to write imperative markup
func Imperative(fn func(ctx ImperativeContext) error) Node {
	return &imperative{fn: fn}
}

type imperative struct {
	fn func(ctx ImperativeContext) error
}

func (i *imperative) Render(ctx *context.Context) (err error) {
	if i.fn != nil {
		err = i.fn(&imperativeContext{ctx: ctx})
	}
	return
}

func (i *imperative) Type() NodeType {
	return imperativeNode
}

func (i *imperative) Name() string {
	return "!imperative"
}

type imperativeContext struct {
	ctx         *context.Context
	openedStack [][]byte
}

func (ic *imperativeContext) Context() *context.Context {
	return ic.ctx
}

func (ic *imperativeContext) WriteString(s string) ImperativeContext {
	ic.ctx.Write(htmlEscapeString(s))
	return ic
}

func (ic *imperativeContext) WriteRaw(data []byte) ImperativeContext {
	ic.ctx.Write(data)
	return ic
}

func (ic *imperativeContext) WriteNodes(nodes ...Node) ImperativeContext {
	for _, node := range nodes {
		if node != nil && node.Type() != AttributeNode {
			_ = node.Render(ic.ctx)
		}
	}
	return ic
}

func (ic *imperativeContext) WriteElement(elementName []byte, content any, attrs ...Node) ImperativeContext {
	if cn, ok := content.(Node); ok {
		if cn.Type() == AttributeNode {
			ic.Start(elementName, false, append([]Node{cn}, attrs...)...)
			ic.End()
			return ic
		}
	}
	ic.Start(elementName, false, attrs...)
	if content != nil {
		switch v := content.(type) {
		case Node:
			_ = v.Render(ic.ctx)
		default:
			ic.ctx.Write(valueToBytes(content))
		}
	}
	ic.End()
	return ic
}

func (ic *imperativeContext) Start(elementName []byte, isVoid bool, attrs ...Node) ImperativeContext {
	ic.ctx.Write(openAngleBracket)
	ic.ctx.Write(elementName)
	attValues := make([][]Value, 0, len(attrs))
	attNames := make([][]byte, 0, len(attrs))
	attIndices := map[string]int{}
	for _, attr := range attrs {
		if attr != nil && attr.Type() == AttributeNode {
			if idx, ok := attIndices[attr.Name()]; ok {
				if da, ok := attr.(*delimitedAttribute); ok {
					avs := attValues[idx]
					for _, dv := range da.values {
						if len(avs) > 0 {
							avs = append(avs, Value{value: da.delimiter})
						}
						avs = append(avs, dv)
					}
					attValues[idx] = avs
				} else if va, ok := attr.(valuesNode); ok {
					attValues[idx] = va.getValues()
				} else {
					attValues[idx] = nil
				}
			} else {
				attIndices[attr.Name()] = len(attNames)
				attNames = append(attNames, []byte(attr.Name()))
				if da, ok := attr.(*delimitedAttribute); ok {
					avs := make([]Value, 0, len(da.values)*2)
					for i, v := range da.values {
						if i > 0 {
							avs = append(avs, Value{value: da.delimiter})
						}
						avs = append(avs, v)
					}
					attValues = append(attValues, avs)
				} else if vn, ok := attr.(valuesNode); ok {
					attValues = append(attValues, vn.getValues())
				} else {
					attValues = append(attValues, nil)
				}
			}
		}
	}
	for a, attName := range attNames {
		ic.ctx.Write(space)
		ic.ctx.Write(attName)
		if vs := attValues[a]; vs != nil {
			ic.ctx.Write(attStart)
			for _, vv := range vs {
				_, _ = vv.render(ic.ctx)
			}
			ic.ctx.Write(attEnd)
		}
	}
	ic.ctx.Write(closeAngleBracket)
	if !isVoid {
		ic.openedStack = append(ic.openedStack, elementName)
	}
	return ic
}

func (ic *imperativeContext) End() ImperativeContext {
	if l := len(ic.openedStack); l > 0 {
		l--
		ic.ctx.Write(closeTagStart)
		ic.ctx.Write(ic.openedStack[l])
		ic.ctx.Write(closeAngleBracket)
		ic.openedStack = ic.openedStack[:l]
	}
	return ic
}

func (ic *imperativeContext) EndAll() ImperativeContext {
	for i := len(ic.openedStack) - 1; i >= 0; i-- {
		ic.ctx.Write(closeTagStart)
		ic.ctx.Write(ic.openedStack[i])
		ic.ctx.Write(closeAngleBracket)
	}
	ic.openedStack = nil
	return ic
}

func (ic *imperativeContext) Marker() int {
	return len(ic.openedStack)
}

func (ic *imperativeContext) EndToMark(marker int) ImperativeContext {
	for len(ic.openedStack) > marker {
		ic.ctx.Write(closeTagStart)
		ic.ctx.Write(ic.openedStack[len(ic.openedStack)-1])
		ic.ctx.Write(closeAngleBracket)
		ic.openedStack = ic.openedStack[:len(ic.openedStack)-1]
	}
	return ic
}
