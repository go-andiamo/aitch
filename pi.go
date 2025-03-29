package aitch

import (
	"github.com/go-andiamo/aitch/context"
)

type pi struct {
	name   []byte
	values []value
}

var _ Node = (*pi)(nil)
var _ valuesNode = (*pi)(nil)

func (p *pi) Render(ctx *context.Context) error {
	ctx.Write(openPi)
	ctx.Write(p.name)
	ctx.Write(space)
	for _, v := range p.values {
		_, _ = v.render(ctx)
	}
	ctx.Write(closePi)
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
