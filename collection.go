package aitch

import "io"

type collection struct {
	nodes []Node
}

func (c *collection) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		for _, n := range c.nodes {
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

func (c *collection) Append(nodes ...Node) Node {
	for _, n := range nodes {
		if n != nil {
			c.nodes = append(c.nodes, n)
		}
	}
	return c
}

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
