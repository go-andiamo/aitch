package aitch

import "io"

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

func (f *fragment) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		for _, v := range f.values {
			if _, ctx.Error = v.render(w, ctx); ctx.Error != nil {
				return ctx.Error
			}
		}
	}
	return ctx.Error
}

func (f *fragment) Type() NodeType {
	return FragmentNode
}

func (f *fragment) Name() string {
	return "#fragment"
}

func (f *fragment) Append(nodes ...Node) Node {
	for _, n := range nodes {
		if n != nil {
			f.values = append(f.values, newValue(n)...)
		}
	}
	return f
}

func (f *fragment) getValues() []value {
	return f.values
}
