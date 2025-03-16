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
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(openPi); ctx.Error == nil {
			if _, ctx.Error = w.Write(p.name); ctx.Error == nil {
				if _, ctx.Error = w.Write(space); ctx.Error == nil {
					for _, v := range p.values {
						if _, ctx.Error = v.render(w, ctx); ctx.Error != nil {
							return ctx.Error
						}
					}
					_, ctx.Error = w.Write(closePi)
				}
			}
		}
	}
	return ctx.Error
}

func (p *pi) Type() NodeType {
	return PINode
}

func (p *pi) Name() string {
	return "!" + string(p.name)
}

func PI(name string, contents ...any) Node {
	return &pi{
		name:   []byte(name),
		values: newValues(contents...),
	}
}

func (p *pi) getValues() []value {
	return p.values
}

func DocType() Node {
	return PI("DOCTYPE", "html")
}
