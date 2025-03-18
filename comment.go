package aitch

import "io"

type comment struct {
	values []value
}

var _ Node = (*comment)(nil)
var _ valuesNode = (*comment)(nil)

func (c *comment) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(openComment)
	for _, v := range c.values {
		_, _ = v.render(ctx)
	}
	ctx.write(closeComment)
	return ctx.Error
}

func (c *comment) Type() NodeType {
	return CommentNode
}

func (c *comment) Name() string {
	return "#comment"
}

func (c *comment) getValues() []value {
	return c.values
}

// Comment creates a new comment Node
func Comment(contents ...any) Node {
	return &comment{
		values: newValues(contents...),
	}
}
