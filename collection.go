package aitch

import (
	"github.com/go-andiamo/aitch/context"
)

type collection struct {
	nodes []Node
}

func (c *collection) Render(ctx *context.Context) error {
	for _, n := range c.nodes {
		if n.Type() != AttributeNode {
			if ctx.Error = n.Render(ctx); ctx.Error != nil {
				return ctx.Error
			}
		}
	}
	return ctx.Error
}

func (c *collection) Type() NodeType {
	return collectionNode
}

func (c *collection) Name() string {
	return "#collection"
}

// Collection creates a new Node collection
func Collection(nodes ...Node) Node {
	result := &collection{
		nodes: make([]Node, 0, len(nodes)),
	}
	for _, n := range nodes {
		if n != nil {
			result.nodes = append(result.nodes, n)
		}
	}
	return result
}

type contentCollector struct {
	collector func(ctx *context.Context) []Node
}

func (c *contentCollector) Render(ctx *context.Context) error {
	for _, n := range c.collector(ctx) {
		if n != nil && n.Type() != AttributeNode {
			_ = n.Render(ctx)
		}
	}
	return ctx.Error
}

func (c *contentCollector) Type() NodeType {
	return dynamicNode
}

func (c *contentCollector) Name() string {
	return "#collector"
}

// ContentCollect creates a new dynamic content Node
//
// Upon rendering, the supplied collector function is called to obtain
// the nodes for rendering
//
// The collector function should not return any attribute nodes - any
// returned attribute nodes will be ignored
func ContentCollect(collector func(ctx *context.Context) []Node) Node {
	if collector == nil {
		return &collection{}
	}
	return &contentCollector{collector: collector}
}
