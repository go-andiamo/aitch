package aitch

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVoidElement(t *testing.T) {
	e := VoidElement("foo", Class("a"))
	assert.Equal(t, ElementNode, e.Type())
	assert.Equal(t, "foo", e.Name())
	rv := e.(*voidElement)
	assert.Equal(t, 1, len(rv.attributes))
}

func TestVoidElement_PanicsOnBadName(t *testing.T) {
	e := VoidElement(" not a valid name")
	assert.Nil(t, e)

	defer func() {
		PanicOnInvalidName = false
	}()
	PanicOnInvalidName = true
	assert.Panics(t, func() {
		_ = VoidElement(" not a valid name")
	})
}

func TestVoidElement_Render(t *testing.T) {
	e := VoidElement("foo", Class("a"))
	var w bytes.Buffer
	err := e.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<foo class=\"a\">", w.String())

	ew := &errorWriter{}
	err = e.Render(ew, &Context{})
	require.Error(t, err)
}

func TestVoidElement_Render_MergedDelimitedAttributes(t *testing.T) {
	e := VoidElement("foo", Class("a"), Class("b"), Class("c"))
	var w bytes.Buffer
	err := e.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<foo class=\"a b c\">", w.String())

	e = VoidElement("foo", Class("a"), Attribute("class", "overridden"))
	w = bytes.Buffer{}
	err = e.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<foo class=\"overridden\">", w.String())
}

func TestElement(t *testing.T) {
	e := Element("foo", Class("a"), Text("b"))
	assert.Equal(t, ElementNode, e.Type())
	assert.Equal(t, "foo", e.Name())
	rv := e.(*element)
	assert.Equal(t, 1, len(rv.attributes))
	assert.Equal(t, 1, len(rv.contents))
}

func TestElement_VoidName(t *testing.T) {
	e := Element("meta")
	assert.Equal(t, ElementNode, e.Type())
	assert.Equal(t, "meta", e.Name())
	_, void := e.(*voidElement)
	require.True(t, void)
}

func TestElement_PanicsOnBadName(t *testing.T) {
	e := Element(" not a valid name")
	assert.Nil(t, e)

	defer func() {
		PanicOnInvalidName = false
	}()
	PanicOnInvalidName = true
	assert.Panics(t, func() {
		_ = Element(" not a valid name")
	})
}

func TestElement_Render(t *testing.T) {
	e := Element("foo", Text("b"), Class("a"))
	var w bytes.Buffer
	err := e.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<foo class=\"a\">b</foo>", w.String())

	ew := &errorWriter{}
	err = e.Render(ew, &Context{})
	require.Error(t, err)
}

func TestElement_Render_MergedDelimitedAttributes(t *testing.T) {
	e := Element("foo", Text("bar"), Class("a"), Class("b"), Class("c"))
	var w bytes.Buffer
	err := e.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<foo class=\"a b c\">bar</foo>", w.String())

	e = Element("foo", Text("bar"), Class("a"), Attribute("class", "overridden"))
	w = bytes.Buffer{}
	err = e.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<foo class=\"overridden\">bar</foo>", w.String())
}

func TestAttributesAndContents(t *testing.T) {
	attrs, contents := attributesAndContents([]Node{Class(), P(), Text()})
	assert.Equal(t, 1, len(attrs))
	assert.Equal(t, 2, len(contents))

	attrs, contents = attributesAndContents([]Node{Collection(Class(), P(), Text())})
	assert.Equal(t, 1, len(attrs))
	assert.Equal(t, 1, len(contents))

	attrs, contents = attributesAndContents([]Node{Collection(Collection(Class(), P(), Text()))})
	assert.Equal(t, 1, len(attrs))
	assert.Equal(t, 1, len(contents))

	attrs, contents = attributesAndContents([]Node{Conditional(func(ctx *Context) bool { return true }, Class(), P(), Text())})
	//TODO assert.Equal(t, 0, len(attrs))
	assert.Equal(t, 1, len(contents))
}
