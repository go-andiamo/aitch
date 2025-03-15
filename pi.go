package aitch

import "io"

type pi struct {
	name   []byte
	values []value
}

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

func (p *pi) Append(nodes ...Node) Node {
	for _, n := range nodes {
		if n != nil {
			p.values = append(p.values, newValue(p)...)
		}
	}
	return p
}

func PI(name string, content ...any) Node {
	result := &pi{
		name:   []byte(name),
		values: make([]value, 0, len(content)),
	}
	for _, v := range content {
		if v != nil {
			result.values = append(result.values, newValue(v)...)
		}
	}
	return result
}

func DocType() Node {
	return PI("DOCTYPE", "html")
}
