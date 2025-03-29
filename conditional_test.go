package aitch

import (
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

	assert.Equal(t, "<p></p>", testRender(c, t))

	condition = false
	assert.Equal(t, "", testRender(c, t))

	condition = true
	err := c.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

func TestConditional_Render_NestedElements(t *testing.T) {
	conditionA := true
	conditionB := true
	fnA := func(ctx *context.Context) bool { return conditionA }
	fnB := func(ctx *context.Context) bool { return conditionB }
	c := Conditional(fnA, P(Text("A")), Conditional(fnB, P(Text("B"))))

	assert.Equal(t, "<p>A</p><p>B</p>", testRender(c, t))

	conditionB = false
	assert.Equal(t, "<p>A</p>", testRender(c, t))

	conditionA = false
	conditionB = true
	assert.Equal(t, "", testRender(c, t))
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
	assert.Equal(t, "", testRender(c, t))

	assert.Equal(t, "<p></p>", testRender(c, t, map[string]any{"foo": nil}))
	assert.Equal(t, "<p></p>", testRender(c, t, map[string]any{"foo": true}))
	assert.Equal(t, "", testRender(c, t, map[string]any{"foo": false}))

	c = When("!foo", P())
	assert.Equal(t, "<p></p>", testRender(c, t))
	assert.Equal(t, "", testRender(c, t, map[string]any{"foo": nil}))
	assert.Equal(t, "", testRender(c, t, map[string]any{"foo": true}))
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
			assert.Equal(t, tc.expect, testRender(tc.node, t))
		})
	}
}
