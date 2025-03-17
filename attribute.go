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
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(space)
	ctx.write(a.name)
	ctx.write(attStart)
	for _, v := range a.values {
		_, _ = v.render(ctx)
	}
	ctx.write(attEnd)
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
	if htmlTagRegex.MatchString(name) {
		return newAttribute([]byte(name), values...)
	} else if PanicOnInvalidName {
		panic("invalid html attribute name: " + name)
	}
	return nil
}

func newAttribute(name []byte, values ...any) Node {
	return &attribute{
		name:   name,
		values: newValues(values...),
	}
}

type booleanAttribute struct {
	name []byte
}

var _ Node = (*booleanAttribute)(nil)

func (a *booleanAttribute) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(space)
	ctx.write(a.name)
	return ctx.Error
}

func (a *booleanAttribute) Type() NodeType {
	return AttributeNode
}

func (a *booleanAttribute) Name() string {
	return string(a.name)
}

// BooleanAttribute creates a new boolean attribute Node
//
// a boolean attribute has no value
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
func BooleanAttribute(name string) Node {
	if htmlTagRegex.MatchString(name) {
		return newBooleanAttribute([]byte(name))
	} else if PanicOnInvalidName {
		panic("invalid html attribute name: " + name)
	}
	return nil
}

func newBooleanAttribute(name []byte) Node {
	return &booleanAttribute{
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
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	ctx.write(space)
	ctx.write(a.name)
	ctx.write(attStart)
	something := false
	for _, v := range a.values {
		if something {
			ctx.write(a.delimiter)
		}
		something, _ = v.render(ctx)
	}
	ctx.write(attEnd)
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

// DelimitedAttribute creates a new delimited attribute Node
//
// examples of delimited attributes are Class() and Style()
//
// when delimited attributes are specified multiple times for an element - the final attribute
// appears only once - but with the values concatenated with the delimiter
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
func DelimitedAttribute(name string, delimiter string, values ...any) Node {
	if htmlTagRegex.MatchString(name) {
		return newDelimitedAttribute([]byte(name), []byte(delimiter), values...)
	} else if PanicOnInvalidName {
		panic("invalid html attribute name: " + name)
	}
	return nil
}

func newDelimitedAttribute(name []byte, delimiter []byte, values ...any) Node {
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
