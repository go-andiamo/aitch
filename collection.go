package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"io"
)

type collection struct {
	nodes []Node
}

func (c *collection) Render(w io.Writer, ctx *context.Context) error {
	if ctx == nil {
		ctx = context.DefaultContext(w)
	} else {
		ctx.Writer = w
	}
	for _, n := range c.nodes {
		if n.Type() != AttributeNode {
			if ctx.Error = n.Render(w, ctx); ctx.Error != nil {
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
