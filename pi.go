package aitch

import "io"

type pi struct {
	name   []byte
	values []value
}

var _ Node = (*pi)(nil)
var _ valuesNode = (*pi)(nil)

func (p *pi) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(openPi)
	ctx.write(p.name)
	ctx.write(space)
	for _, v := range p.values {
		_, _ = v.render(ctx)
	}
	ctx.write(closePi)
	return ctx.Error
}

func (p *pi) Type() NodeType {
	return PINode
}

func (p *pi) Name() string {
	return "!" + string(p.name)
}

// PI creates a processing instruction Node
func PI(name string, contents ...any) Node {
	return &pi{
		name:   []byte(name),
		values: newValues(contents...),
	}
}

func (p *pi) getValues() []value {
	return p.values
}

// DocType creates a HTML5 doctype Node
func DocType() Node {
	return PI("DOCTYPE", "html")
}
