package aitch

import (
	"bytes"
	"fmt"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConditional(t *testing.T) {
	c := Conditional(func(ctx *context.Context) bool {
		return true
	}, P(), Class("foo"))
	assert.Equal(t, conditionalNode, c.Type())
	assert.Equal(t, "#conditional", c.Name())
	rv := c.(*conditional)
	assert.Equal(t, 1, len(rv.nodes))
	assert.NotNil(t, rv.fn)

	c = Conditional(nil)
	_, ok := c.(*collection)
	assert.True(t, ok)
}

func TestConditional_Render(t *testing.T) {
	condition := true
	c := Conditional(func(ctx *context.Context) bool {
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
	fnA := func(ctx *context.Context) bool { return conditionA }
	fnB := func(ctx *context.Context) bool { return conditionB }
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
	assert.NotNil(t, rv.fn)
}

func TestWhen_Render(t *testing.T) {
	c := When("foo", P())
	var w bytes.Buffer
	ctx := &context.Context{}
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

func TestConditional_Cases(t *testing.T) {
	trueFn := func(ctx *context.Context) bool { return true }
	falseFn := func(ctx *context.Context) bool { return false }
	testCases := []struct {
		node   Node
		expect string
	}{
		{
			node:   Conditional(trueFn, P()),
			expect: `<p></p>`,
		},
		{
			node:   Conditional(falseFn, P()),
			expect: ``,
		},
		{
			node:   P(Conditional(trueFn, A())),
			expect: `<p><a></a></p>`,
		},
		{
			node:   P(Conditional(falseFn, A())),
			expect: `<p></p>`,
		},
		{
			node:   P(Conditional(trueFn, Collection(A()))),
			expect: `<p><a></a></p>`,
		},
		{
			node:   P(Conditional(falseFn, Collection(A()))),
			expect: `<p></p>`,
		},
		{
			node:   P(Conditional(trueFn, Id("a"), Collection(A()))),
			expect: `<p id="a"><a></a></p>`,
		},
		{
			node:   P(Conditional(falseFn, Id("a"), Collection(A()))),
			expect: `<p></p>`,
		},
		{
			node:   P(Conditional(trueFn, Class("a"), Collection(A()))),
			expect: `<p class="a"><a></a></p>`,
		},
		{
			node:   P(Conditional(falseFn, Id("a"), Collection(A()))),
			expect: `<p></p>`,
		},
		{
			node:   P(Id("foo"), Conditional(trueFn, Id("bar"))),
			expect: `<p id="bar"></p>`,
		},
		{
			node:   P(Id("foo"), Conditional(falseFn, Id("bar"))),
			expect: `<p id="foo"></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(trueFn, A(), Class("bar"))),
			expect: `<p class="foo bar"><a></a></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(falseFn, A(), Class("bar"))),
			expect: `<p class="foo"></p>`,
		},
		{
			node:   P(Id("foo"), Conditional(trueFn, Id("bar"), A(), Conditional(trueFn, Id("baz")))),
			expect: `<p id="baz"><a></a></p>`,
		},
		{
			node:   P(Id("foo"), Conditional(trueFn, Id("bar"), A(), Conditional(falseFn, Id("baz")))),
			expect: `<p id="bar"><a></a></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(trueFn, Class("bar"), A(), Conditional(trueFn, Id("a")))),
			expect: `<p class="foo bar" id="a"><a></a></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(trueFn, Class("bar"), A(), Conditional(falseFn, Id("a")))),
			expect: `<p class="foo bar"><a></a></p>`,
		},
		{
			node:   P(Id("foo"), Conditional(trueFn, Collection(A(), Id("bar")))),
			expect: `<p id="bar"><a></a></p>`,
		},
		{
			node:   P(Id("foo"), Conditional(falseFn, Collection(A(), Id("bar")))),
			expect: `<p id="foo"></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(trueFn, Collection(A(), Class("bar")))),
			expect: `<p class="foo bar"><a></a></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(falseFn, Collection(A(), Class("bar")))),
			expect: `<p class="foo"></p>`,
		},
		{
			node:   P(Class("foo"), Id("a"), Collection(Conditional(trueFn, Collection(A(), Conditional(trueFn, Collection(B(), Id("b"), Class("bar"))))))),
			expect: `<p class="foo bar" id="b"><a></a><b></b></p>`,
		},
		{
			node:   P(Class("foo"), Conditional(trueFn, Id("a"), Collection(A(), Id("b")))),
			expect: `<p class="foo" id="b"><a></a></p>`,
		},
		{
			node:   P(Conditional(trueFn, Class("foo"), Conditional(trueFn, Class("bar")))),
			expect: `<p class="foo bar"></p>`,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			var w bytes.Buffer
			err := tc.node.Render(&w, nil)
			require.NoError(t, err)
			assert.Equal(t, tc.expect, w.String())
		})
	}
}
