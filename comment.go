package aitch

import "io"

type comment struct {
	values []value
}

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

func Comment(content ...any) Node {
	result := &comment{
		values: make([]value, 0, len(content)),
	}
	for _, v := range content {
		if v != nil {
			result.values = append(result.values, newValue(v)...)
		}
	}
	return result
}
