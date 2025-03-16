package aitch

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditional(t *testing.T) {
	c := Conditional(func(ctx *Context) bool {
		return true
	}, P(), Class("foo"))
	assert.Equal(t, conditionalNode, c.Type())
	assert.Equal(t, "#conditional", c.Name())
	rv := c.(*conditional)
	assert.Equal(t, 1, len(rv.nodes))
	assert.Equal(t, 1, len(rv.attributes))
	assert.NotNil(t, rv.fn)

	assert.Panics(t, func() {
		Conditional(nil)
	})
}

func TestConditional_Render(t *testing.T) {
	condition := true
	c := Conditional(func(ctx *Context) bool {
		return condition
	}, P())

	var w bytes.Buffer
	err := c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<p></p>", w.String())

	w = bytes.Buffer{}
	condition = false
	err = c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, 0, w.Len())

	condition = true
	ew := &errorWriter{}
	err = c.Render(ew, nil)
	require.Error(t, err)
}

func TestConditional_Render_NestedElements(t *testing.T) {
	conditionA := true
	conditionB := true
	fnA := func(ctx *Context) bool { return conditionA }
	fnB := func(ctx *Context) bool { return conditionB }
	c := Conditional(fnA, P(Text("A")), Conditional(fnB, P(Text("B"))))

	var w bytes.Buffer
	err := c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<p>A</p><p>B</p>", w.String())

	conditionB = false
	w = bytes.Buffer{}
	err = c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<p>A</p>", w.String())

	conditionA = false
	conditionB = true
	w = bytes.Buffer{}
	err = c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, 0, w.Len())
}

func TestWhen(t *testing.T) {
	c := When("foo", P(), Class("foo"))
	assert.Equal(t, conditionalNode, c.Type())
	assert.Equal(t, "#conditional", c.Name())
	rv := c.(*conditional)
	assert.Equal(t, 1, len(rv.nodes))
	assert.Equal(t, 1, len(rv.attributes))
	assert.NotNil(t, rv.fn)
}

func TestWhen_Render(t *testing.T) {
	c := When("foo", P())
	var w bytes.Buffer
	ctx := &Context{}
	err := c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, w.Len())

	ctx.Data = map[string]any{"foo": nil}
	w = bytes.Buffer{}
	err = c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, "<p></p>", w.String())

	ctx.Data = map[string]any{"foo": true}
	w = bytes.Buffer{}
	err = c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, "<p></p>", w.String())

	ctx.Data = map[string]any{"foo": false}
	w = bytes.Buffer{}
	err = c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, w.Len())

	c = When("!foo", P())
	ctx.Data = nil
	w = bytes.Buffer{}
	err = c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, "<p></p>", w.String())

	ctx.Data = map[string]any{"foo": nil}
	w = bytes.Buffer{}
	err = c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, w.Len())

	ctx.Data = map[string]any{"foo": true}
	w = bytes.Buffer{}
	err = c.Render(&w, ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, w.Len())
}
