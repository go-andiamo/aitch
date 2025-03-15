package aitch

import (
	"io"
	"strings"
)

type ConditionalFunc func(ctx *Context) bool

type conditional struct {
	attributes []Node
	nodes      []Node
	fn         ConditionalFunc
}

func (c *conditional) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil && c.fn(ctx) {
		for _, n := range c.nodes {
			if ctx.Error = n.Render(w, ctx); ctx.Error != nil {
				break
			}
		}
	}
	return ctx.Error
}

func (c *conditional) Type() NodeType {
	return conditionalNode
}

func (c *conditional) Name() string {
	return "#conditional"
}

func (c *conditional) Append(nodes ...Node) Node {
	panic("conditionals cannot be appended")
}

func Conditional(fn ConditionalFunc, nodes ...Node) Node {
	attrs, children := attributesAndContents(nodes...)
	return &conditional{
		attributes: attrs,
		nodes:      children,
		fn:         fn,
	}
}

func When(key DynamicKey, nodes ...Node) Node {
	notted := strings.HasPrefix(string(key), "!")
	if notted {
		key = key[1:]
	}
	return Conditional(func(ctx *Context) bool {
		if v, ok := ctx.Data[string(key)]; ok {
			if b, ok := v.(bool); ok {
				if notted {
					return !b
				}
				return b
			}
			return !notted
		}
		return notted
	}, nodes...)
}
