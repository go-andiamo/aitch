package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"io"
)

type comment struct {
	values []value
}

var _ Node = (*comment)(nil)
var _ valuesNode = (*comment)(nil)

func (c *comment) Render(w io.Writer, ctx *context.Context) error {
	if ctx == nil {
		ctx = context.DefaultContext(w)
	} else {
		ctx.Writer = w
	}
	ctx.Write(openComment)
	for _, v := range c.values {
		_, _ = v.render(ctx)
	}
	ctx.Write(closeComment)
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
