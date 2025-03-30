package aitch

import (
	"github.com/go-andiamo/aitch/context"
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
	assert.Equal(t, `<foo class="a">`, testRender(e, t))

	err := e.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

func TestVoidElement_Render_MergedDelimitedAttributes(t *testing.T) {
	e := VoidElement("foo", Class("a"), Class("b"), Class("c"))
	assert.Equal(t, `<foo class="a b c">`, testRender(e, t))

	e = VoidElement("foo", Class("a"), Attribute("class", "overridden"))
	assert.Equal(t, `<foo class="overridden">`, testRender(e, t))
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
	assert.Equal(t, `<foo class="a">b</foo>`, testRender(e, t))

	err := e.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

func TestElement_Render_MergedDelimitedAttributes(t *testing.T) {
	e := Element("foo", Text("bar"), Class("a"), Class("b"), Class("c"))
	assert.Equal(t, `<foo class="a b c">bar</foo>`, testRender(e, t))

	e = Element("foo", Text("bar"), Class("a"), Attribute("class", "overridden"))
	assert.Equal(t, `<foo class="overridden">bar</foo>`, testRender(e, t))
}

func TestAttributesAndContents(t *testing.T) {
	attrs, _, contents := attributesAndContents(nil, []Node{Class(), P(), Text()})
	assert.Equal(t, 1, len(attrs))
	assert.Equal(t, 2, len(contents))

	attrs, _, contents = attributesAndContents(nil, []Node{Collection(Class(), P(), Text())})
	assert.Equal(t, 1, len(attrs))
	assert.Equal(t, 1, len(contents))

	attrs, _, contents = attributesAndContents(nil, []Node{Collection(Collection(Class(), P(), Text()))})
	assert.Equal(t, 1, len(attrs))
	assert.Equal(t, 1, len(contents))

	attrs, condAttrs, contents := attributesAndContents([]ConditionalFunc{func(ctx *context.Context) bool { return true }}, []Node{Class(), P(), Text()})
	assert.Equal(t, 0, len(attrs))
	assert.Equal(t, 1, len(condAttrs))
	assert.Equal(t, 2, len(contents))
}

func TestElement_ConditionalAttrs(t *testing.T) {
	condA := func(ctx *context.Context) bool { return true }
	condB := func(ctx *context.Context) bool { return false }
	p := P(Class("a"), Conditional(condA, Class("b"), Conditional(condB, Class("c"))))
	rv := p.(*element)
	assert.Equal(t, 1, len(rv.attributes))
	assert.Equal(t, 2, len(rv.conditionals))
	assert.Equal(t, 1, len(rv.conditionals[0].conditions))
	assert.True(t, rv.conditionals[0].conditions[0](nil))
	assert.Equal(t, 2, len(rv.conditionals[1].conditions))
	assert.True(t, rv.conditionals[1].conditions[0](nil))
	assert.False(t, rv.conditionals[1].conditions[1](nil))
}
