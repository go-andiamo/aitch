package aitch

import (
	"github.com/go-andiamo/aitch/context"
)

type attribute struct {
	name   []byte
	values []Value
}

var _ Node = (*attribute)(nil)
var _ valuesNode = (*attribute)(nil)

func (a *attribute) Render(ctx *context.Context) error {
	ctx.Write(space)
	ctx.Write(a.name)
	ctx.Write(attStart)
	for _, v := range a.values {
		_, _ = v.render(ctx)
	}
	ctx.Write(attEnd)
	return ctx.Error
}

func (a *attribute) Type() NodeType {
	return AttributeNode
}

func (a *attribute) Name() string {
	return string(a.name)
}

func (a *attribute) getValues() []Value {
	return a.values
}

// Attribute creates a new attribute Node
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
func Attribute(name string, values ...any) Node {
	if htmlTagRegex.MatchString(name) {
		return NewAttribute([]byte(name), values...)
	} else if PanicOnInvalidName {
		panic("invalid html attribute name: " + name)
	}
	return nil
}

func NewAttribute(name []byte, values ...any) Node {
	return &attribute{
		name:   name,
		values: newValues(values...),
	}
}

type booleanAttribute struct {
	name []byte
}

var _ Node = (*booleanAttribute)(nil)

func (a *booleanAttribute) Render(ctx *context.Context) error {
	ctx.Write(space)
	ctx.Write(a.name)
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
// a boolean attribute has no Value
//
// the name is checked to ensure it's a valid name - returns nil if the name is invalid
// (or panics if PanicOnInvalidName is true)
func BooleanAttribute(name string) Node {
	if htmlTagRegex.MatchString(name) {
		return NewBooleanAttribute([]byte(name))
	} else if PanicOnInvalidName {
		panic("invalid html attribute name: " + name)
	}
	return nil
}

func NewBooleanAttribute(name []byte) Node {
	return &booleanAttribute{
		name: name,
	}
}

type delimitedAttribute struct {
	name      []byte
	delimiter []byte
	values    []Value
}

var _ Node = (*delimitedAttribute)(nil)
var _ valuesNode = (*delimitedAttribute)(nil)

func (a *delimitedAttribute) Render(ctx *context.Context) error {
	ctx.Write(space)
	ctx.Write(a.name)
	ctx.Write(attStart)
	something := false
	for _, v := range a.values {
		if something {
			ctx.Write(a.delimiter)
		}
		something, _ = v.render(ctx)
	}
	ctx.Write(attEnd)
	return ctx.Error
}

func (a *delimitedAttribute) Type() NodeType {
	return AttributeNode
}

func (a *delimitedAttribute) Name() string {
	return string(a.name)
}

func (a *delimitedAttribute) getValues() []Value {
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
		return NewDelimitedAttribute([]byte(name), []byte(delimiter), values...)
	} else if PanicOnInvalidName {
		panic("invalid html attribute name: " + name)
	}
	return nil
}

func NewDelimitedAttribute(name []byte, delimiter []byte, values ...any) Node {
	result := &delimitedAttribute{
		name:      name,
		delimiter: delimiter,
		values:    make([]Value, 0, len(values)),
	}
	for _, v := range values {
		if v != nil {
			result.values = append(result.values, newValue(v)...)
		}
	}
	return result
}
