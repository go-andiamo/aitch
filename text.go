package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"io"
)

// Text creates a new text Node
func Text(contents ...any) Node {
	return &text{
		values: newValues(contents...),
	}
}

type text struct {
	values []value
}

var _ Node = (*text)(nil)
var _ valuesNode = (*text)(nil)

func (t *text) Render(w io.Writer, ctx *context.Context) error {
	if ctx == nil {
		ctx = context.DefaultContext(w)
	} else {
		ctx.Writer = w
	}
	for _, v := range t.values {
		_, _ = v.render(ctx)
	}
	return ctx.Error
}

func (t *text) Type() NodeType {
	return TextNode
}

func (t *text) Name() string {
	return "#text"
}

func (t *text) getValues() []value {
	return t.values
}

var (
	nbsp  = []byte("&nbsp;")
	lnbsp = len(nbsp)
)

func Nbsp() []byte {
	return append([]byte(nil), nbsp...)
}

func Nbsps(n int) []byte {
	if n <= 0 {
		return nil
	}
	result := make([]byte, lnbsp*n)
	idx := 0
	for i := 0; i < n; i++ {
		copy(result[idx:], nbsp)
		idx += lnbsp
	}
	return result
}
