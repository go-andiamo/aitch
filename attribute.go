package aitch

import "io"

type attribute struct {
	name   []byte
	values []value
}

var _ Node = (*attribute)(nil)
var _ valuesNode = (*attribute)(nil)

func (a *attribute) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(space); ctx.Error == nil {
			if _, ctx.Error = w.Write(a.name); ctx.Error == nil && len(a.values) > 0 {
				if _, ctx.Error = w.Write(attStart); ctx.Error == nil {
					for _, v := range a.values {
						if _, ctx.Error = v.render(w, ctx); ctx.Error != nil {
							return ctx.Error
						}
					}
					_, ctx.Error = w.Write(attEnd)
				}
			}
		}
	}
	return ctx.Error
}

func (a *attribute) Type() NodeType {
	return AttributeNode
}

func (a *attribute) Name() string {
	return string(a.name)
}

func (a *attribute) getValues() []value {
	return a.values
}

// Attribute creates a new attribute Node
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
func Attribute(name string, values ...any) Node {
	if !htmlTagRegex.MatchString(name) {
		if PanicOnInvalidName {
			panic("invalid html attribute name: " + name)
		}
		return nil
	}
	return newAttribute([]byte(name), values...)
}

func newAttribute(name []byte, values ...any) Node {
	return &attribute{
		name:   name,
		values: newValues(values...),
	}
}

type emptyAttribute struct {
	name []byte
}

var _ Node = (*emptyAttribute)(nil)

func (e *emptyAttribute) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(space); ctx.Error == nil {
			_, ctx.Error = w.Write(e.name)
		}
	}
	return ctx.Error
}

func (e *emptyAttribute) Type() NodeType {
	return AttributeNode
}

func (e *emptyAttribute) Name() string {
	return string(e.name)
}

func newEmptyAttribute(name []byte) Node {
	return &emptyAttribute{
		name: name,
	}
}

type delimitedAttribute struct {
	name      []byte
	delimiter []byte
	values    []value
}

var _ Node = (*delimitedAttribute)(nil)
var _ valuesNode = (*delimitedAttribute)(nil)

func (a *delimitedAttribute) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if _, ctx.Error = w.Write(space); ctx.Error == nil {
			if _, ctx.Error = w.Write(a.name); ctx.Error == nil {
				if _, ctx.Error = w.Write(attStart); ctx.Error == nil {
					something := false
					for _, v := range a.values {
						if something {
							if _, ctx.Error = w.Write(a.delimiter); ctx.Error != nil {
								return ctx.Error
							}
						}
						if something, ctx.Error = v.render(w, ctx); ctx.Error != nil {
							return ctx.Error
						}
					}
					_, ctx.Error = w.Write(attEnd)
				}
			}
		}
	}
	return ctx.Error
}

func (a *delimitedAttribute) Type() NodeType {
	return AttributeNode
}

func (a *delimitedAttribute) Name() string {
	return string(a.name)
}

func (a *delimitedAttribute) getValues() []value {
	return a.values
}

func DelimitedAttribute(name []byte, delimiter []byte, values ...any) Node {
	result := &delimitedAttribute{
		name:      name,
		delimiter: delimiter,
		values:    make([]value, 0, len(values)),
	}
	for _, v := range values {
		if v != nil {
			result.values = append(result.values, newValue(v)...)
		}
	}
	return result
}
