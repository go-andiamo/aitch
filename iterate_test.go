package aitch

import (
	"errors"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIterate(t *testing.T) {
	n := Iterate(func(ctx *context.Context) []any {
		return nil
	}, P(), Id())
	assert.Equal(t, iteratorNode, n.Type())
	assert.Equal(t, "!iterator", n.Name())
	rv := n.(*iterator)
	assert.Equal(t, 1, len(rv.nodes))
	assert.NotNil(t, rv.fn)
	require.Panics(t, func() {
		_ = Iterate(nil)
	})
}

func TestIterateKey(t *testing.T) {
	n := IterateKey("items", P(), Id())
	assert.Equal(t, iteratorNode, n.Type())
	assert.Equal(t, "!iterator", n.Name())
	rv := n.(*iterator)
	assert.Equal(t, 1, len(rv.nodes))
	assert.NotNil(t, rv.fn)
}

func TestIterate_Render(t *testing.T) {
	n := Iterate(func(ctx *context.Context) []any {
		return []any{1, 2, 3}
	}, P(), Id())
	assert.Equal(t, `<p></p><p></p><p></p>`, testRender(n, t))

	ctx := &context.Context{Error: errors.New("foo")}
	err := n.Render(ctx)
	require.Error(t, err)

	ctx = &context.Context{Writer: &errorWriter{}}
	err = n.Render(ctx)
	require.Error(t, err)
}

func TestIterateKey_Render(t *testing.T) {
	n := IterateKey("items", P(), Id())
	assert.Equal(t, `<p></p><p></p><p></p>`, testRender(n, t, map[string]any{
		"items": []any{1, 2, 3},
	}))
	assert.Equal(t, `<p></p><p></p><p></p>`, testRender(n, t, map[string]any{
		"items": []int{1, 2, 3},
	}))
	assert.Equal(t, ``, testRender(n, t))

	n = IterateKey("items", P(DynamicValueKey("text")))
	assert.Equal(t, `<p>foo</p><p>bar</p><p>baz</p>`, testRender(n, t, map[string]any{
		"items": []map[string]any{
			{"text": "foo"},
			{"text": "bar"},
			{"text": "baz"},
		},
	}))

	n = IterateKey("items", When("yes", P(DynamicValueKey("text"))))
	assert.Equal(t, `<p>foo</p><p>baz</p>`, testRender(n, t, map[string]any{
		"items": []map[string]any{
			{"text": "foo", "yes": true},
			{"text": "bar"},
			{"text": "baz", "yes": true},
		},
	}))
}

func TestIterateYield(t *testing.T) {
	n := IterateYield(func(ctx *context.Context, yield func(any)) {}, P(Id("foo"), Text(DynamicValueKey("."))))
	assert.Equal(t, iteratorNode, n.Type())
	assert.Equal(t, "!yield_iterator", n.Name())
	rv := n.(*yieldIterator)
	assert.Equal(t, 1, len(rv.nodes))
	assert.NotNil(t, rv.fn)
	require.Panics(t, func() {
		_ = IterateYield(nil)
	})
}

func TestIterateYield_Render(t *testing.T) {
	n := IterateYield(func(ctx *context.Context, yield func(any)) {
		for _, v := range []string{"foo", "bar", "baz"} {
			yield(v)
		}
	}, P(Text(DynamicValueKey("."))))
	assert.Equal(t, `<p>foo</p><p>bar</p><p>baz</p>`, testRender(n, t))

	n = IterateYield(func(ctx *context.Context, yield func(any)) {
		for _, v := range []string{"foo", "bar", "baz"} {
			yield(map[string]any{
				"label": v,
			})
		}
	}, P(Text(DynamicValueKey("label"))))
	assert.Equal(t, `<p>foo</p><p>bar</p><p>baz</p>`, testRender(n, t))

	ctx := &context.Context{Error: errors.New("foo")}
	err := n.Render(ctx)
	require.Error(t, err)

	ctx = &context.Context{Writer: &errorWriter{}}
	err = n.Render(ctx)
	require.Error(t, err)
}
