package svg

import (
	"bytes"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/require"
	"testing"
)

func testRender(node aitch.Node, t *testing.T) string {
	var w bytes.Buffer
	err := node.Render(&context.Context{Writer: &w})
	require.NoError(t, err)
	return w.String()
}

func TestText(t *testing.T) {
	n := Text("foo")
	require.Equal(t, "foo", testRender(n, t))
}

func TestComment(t *testing.T) {
	n := Comment("foo")
	require.Equal(t, "<!--foo-->", testRender(n, t))
}

func TestTextDynamic(t *testing.T) {
	var dn DynamicValueFunc = func(ctx *context.Context) []byte {
		return []byte{'b', 'a', 'r'}
	}
	dn2 := func(ctx *context.Context) []byte {
		return []byte{'b', 'a', 'z'}
	}
	n := Text(DynamicValueKey("foo"), " ", dn, " ", dn2)
	var w bytes.Buffer
	ctx := &context.Context{
		Data: map[string]any{
			"foo": "foo value",
		},
		Writer: &w,
	}
	err := n.Render(ctx)
	require.NoError(t, err)
	require.Equal(t, "foo value bar baz", w.String())
}
