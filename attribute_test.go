package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttribute(t *testing.T) {
	a := Attribute("foo", "bar")
	assert.Equal(t, AttributeNode, a.Type())
	assert.Equal(t, "foo", a.Name())
	rv := a.(*attribute)
	assert.Equal(t, 1, len(rv.values))
	assert.Equal(t, 1, len(rv.getValues()))
}

func TestAttribute_PanicsOnBadName(t *testing.T) {
	a := Attribute(" not a valid name")
	assert.Nil(t, a)

	defer func() {
		PanicOnInvalidName = false
	}()
	PanicOnInvalidName = true
	assert.Panics(t, func() {
		_ = Attribute(" not a valid name")
	})
}

func TestAttribute_Render(t *testing.T) {
	a := Attribute("foo", "bar")
	assert.Equal(t, ` foo="bar"`, testRender(a, t))

	err := a.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

func TestBooleanAttribute(t *testing.T) {
	a := BooleanAttribute("foo")
	assert.Equal(t, AttributeNode, a.Type())
	assert.Equal(t, "foo", a.Name())
}

func TestBooleanAttribute_PanicsOnBadName(t *testing.T) {
	a := BooleanAttribute(" not a valid name")
	assert.Nil(t, a)

	defer func() {
		PanicOnInvalidName = false
	}()
	PanicOnInvalidName = true
	assert.Panics(t, func() {
		_ = BooleanAttribute(" not a valid name")
	})
}

func TestBooleanAttribute_Render(t *testing.T) {
	a := BooleanAttribute("foo")
	assert.Equal(t, " foo", testRender(a, t))

	err := a.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

func TestDelimitedAttribute(t *testing.T) {
	a := DelimitedAttribute("foo", " ", "a", "b")
	assert.Equal(t, AttributeNode, a.Type())
	assert.Equal(t, "foo", a.Name())
	rv := a.(*delimitedAttribute)
	assert.Equal(t, 2, len(rv.values))
	assert.Equal(t, 2, len(rv.getValues()))
}

func TestDelimitedAttribute_PanicsOnBadName(t *testing.T) {
	a := DelimitedAttribute(" not a valid name", "")
	assert.Nil(t, a)

	defer func() {
		PanicOnInvalidName = false
	}()
	PanicOnInvalidName = true
	assert.Panics(t, func() {
		_ = DelimitedAttribute(" not a valid name", "")
	})
}

func TestDelimitedAttribute_Render(t *testing.T) {
	a := DelimitedAttribute("foo", " ", "a", "b")
	assert.Equal(t, ` foo="a b"`, testRender(a, t))

	err := a.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}
