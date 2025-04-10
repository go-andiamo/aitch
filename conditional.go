package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"strings"
)

type ConditionalFunc func(ctx *context.Context) bool

type conditional struct {
	nodes      []Node
	attributes []Node
	fn         ConditionalFunc
}

type conditionalAttribute struct {
	attribute  Node
	conditions []ConditionalFunc
}

func (a conditionalAttribute) evaluate(ctx *context.Context) bool {
	for _, condition := range a.conditions {
		if !condition(ctx) {
			return false
		}
	}
	return true
}

type conditionalAttributes []conditionalAttribute

func (c conditionalAttributes) evaluate(ctx *context.Context) map[string][]Node {
	result := make(map[string][]Node, len(c))
	for _, a := range c {
		if a.evaluate(ctx) {
			name := a.attribute.Name()
			result[name] = append(result[name], a.attribute)
		}
	}
	return result
}

func (c *conditional) Render(ctx *context.Context) error {
	if c.fn(ctx) {
		for _, n := range c.nodes {
			if ctx.Error = n.Render(ctx); ctx.Error != nil {
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
//
// Note: If the fn argument is nil, this will just return a Collection Node
func Conditional(fn ConditionalFunc, nodes ...Node) Node {
	if fn == nil {
		return Collection(nodes...)
	}
	result := &conditional{
		nodes:      make([]Node, 0, len(nodes)),
		attributes: make([]Node, 0, len(nodes)),
		fn:         fn,
	}
	for _, n := range nodes {
		if n.Type() == AttributeNode {
			result.attributes = append(result.attributes, n)
		} else {
			result.nodes = append(result.nodes, n)
		}
	}
	return result
}

// When creates a new conditional Node
//
// the condition is based on whether the specified key exists in the Context.Data
//
// Note: If the key exists and its Value is a boolean - that boolean Value is taken as the condition result
func When(key DynamicValueKey, nodes ...Node) Node {
	if strings.HasPrefix(string(key), "!") {
		key = key[1:]
		return Conditional(func(ctx *context.Context) bool {
			if v, ok := ctx.Data[string(key)]; ok {
				if b, ok := v.(bool); ok {
					return !b
				}
				return false
			}
			return true
		}, nodes...)
	}
	return Conditional(func(ctx *context.Context) bool {
		if v, ok := ctx.Data[string(key)]; ok {
			if b, ok := v.(bool); ok {
				return b
			}
			return true
		}
		return false
	}, nodes...)
}
