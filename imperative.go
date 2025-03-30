package aitch

import "github.com/go-andiamo/aitch/context"

// ImperativeContext is an interface that is passed to the func supplied to Imperative during rendering
type ImperativeContext interface {
	// Context obtains the outer rendering context
	Context() *context.Context
	// WriteRaw writes raw byte data to the output
	WriteRaw(data []byte)
	// WriteNodes writes nodes to the output
	//
	// Note: attribute type nodes will be ignored
	WriteNodes(nodes ...Node)
	// Start starts a new element
	//
	// If any of the attrs args are not AttributeNode types - they are ignored
	Start(elementName []byte, isVoid bool, attrs ...Node)
	// End ends the last started element (as started with Start)
	End()
	// EndAll ends all started elements
	EndAll()
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

func (ic *imperativeContext) WriteRaw(data []byte) {
	ic.ctx.Write(data)
}

func (ic *imperativeContext) WriteNodes(nodes ...Node) {
	for _, node := range nodes {
		if node != nil && node.Type() != AttributeNode {
			_ = node.Render(ic.ctx)
		}
	}
}

func (ic *imperativeContext) Start(elementName []byte, isVoid bool, attrs ...Node) {
	ic.ctx.Write(openAngleBracket)
	ic.ctx.Write(elementName)
	attValues := make([][]value, 0, len(attrs))
	attNames := make([][]byte, 0, len(attrs))
	attIndices := map[string]int{}
	for _, attr := range attrs {
		if attr != nil && attr.Type() == AttributeNode {
			if idx, ok := attIndices[attr.Name()]; ok {
				if da, ok := attr.(*delimitedAttribute); ok {
					avs := attValues[idx]
					for _, dv := range da.values {
						if len(avs) > 0 {
							avs = append(avs, value{value: da.delimiter})
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
					avs := make([]value, 0, len(da.values)*2)
					for i, v := range da.values {
						if i > 0 {
							avs = append(avs, value{value: da.delimiter})
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
}

func (ic *imperativeContext) End() {
	if l := len(ic.openedStack); l > 0 {
		l--
		ic.ctx.Write(closeTagStart)
		ic.ctx.Write(ic.openedStack[l])
		ic.ctx.Write(closeAngleBracket)
		ic.openedStack = ic.openedStack[:l]
	}
}

func (ic *imperativeContext) EndAll() {
	for i := len(ic.openedStack) - 1; i >= 0; i-- {
		ic.ctx.Write(closeTagStart)
		ic.ctx.Write(ic.openedStack[i])
		ic.ctx.Write(closeAngleBracket)
	}
	ic.openedStack = nil
}
