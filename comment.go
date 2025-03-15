package aitch

import "io"

type comment struct {
	values []value
}

var _ Node = (*comment)(nil)
var _ valuesNode = (*comment)(nil)

func (c *comment) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(openComment); ctx.Error == nil {
			for _, v := range c.values {
				if _, ctx.Error = v.render(w, ctx); ctx.Error != nil {
					return ctx.Error
				}
			}
			_, ctx.Error = w.Write(closeComment)
		}
	}
	return ctx.Error
}

func (c *comment) Type() NodeType {
	return CommentNode
}

func (c *comment) Name() string {
	return "#comment"
}

func (c *comment) Append(nodes ...Node) Node {
	for _, n := range nodes {
		if n != nil {
			c.values = append(c.values, newValue(n)...)
		}
	}
	return c
}

func (c *comment) getValues() []value {
	return c.values
}

func Comment(contents ...any) Node {
	return &comment{
		values: newValues(contents...),
	}
}
