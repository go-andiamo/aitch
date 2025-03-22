package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"io"
)

// Fragment creates a new fragment node
//
// a fragment can be used to represent any data to be written to the writer - and
// does not have to be "well-formed" HTML (use with caution!)
func Fragment(content ...any) Node {
	return &fragment{
		values: newValues(content...),
	}
}

type fragment struct {
	values []value
}

var _ Node = (*fragment)(nil)
var _ valuesNode = (*fragment)(nil)

func (f *fragment) Render(w io.Writer, ctx *context.Context) error {
	if ctx == nil {
		ctx = context.DefaultContext(w)
	} else {
		ctx.Writer = w
	}
	for _, v := range f.values {
		_, _ = v.render(ctx)
	}
	return ctx.Error
}

func (f *fragment) Type() NodeType {
	return FragmentNode
}

func (f *fragment) Name() string {
	return "#fragment"
}

func (f *fragment) getValues() []value {
	return f.values
}
