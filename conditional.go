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
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	if c.fn(ctx) {
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

// Conditional creates a new conditional Node
func Conditional(fn ConditionalFunc, nodes ...Node) Node {
	if fn == nil {
		panic("conditional function is nil")
	}
	attrs, children := attributesAndContents(nodes)
	return &conditional{
		attributes: attrs,
		nodes:      children,
		fn:         fn,
	}
}

// When creates a new conditional Node
//
// the condition is based on whether the specified key exists in the Context.Data
//
// Note: If the key exists and its value is a boolean - that boolean value is taken as the condition result
func When(key DynamicKey, nodes ...Node) Node {
	if strings.HasPrefix(string(key), "!") {
		key = key[1:]
		return Conditional(func(ctx *Context) bool {
			if v, ok := ctx.Data[string(key)]; ok {
				if b, ok := v.(bool); ok {
					return !b
				}
				return false
			}
			return true
		}, nodes...)
	}
	return Conditional(func(ctx *Context) bool {
		if v, ok := ctx.Data[string(key)]; ok {
			if b, ok := v.(bool); ok {
				return b
			}
			return true
		}
		return false
	}, nodes...)
}
