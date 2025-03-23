package svg

import (
	"bytes"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTextDynamic(t *testing.T) {
	var dn DynamicValueFunc = func(ctx *context.Context) []byte {
		return []byte{'b', 'a', 'r'}
	}
	dn2 := func(ctx *context.Context) []byte {
		return []byte{'b', 'a', 'z'}
	}
	n := Text(DynamicValueKey("foo"), " ", dn, " ", dn2)
	ctx := &context.Context{
		Data: map[string]any{
			"foo": "foo value",
		},
	}
	var w bytes.Buffer
	err := n.Render(&w, ctx)
	require.NoError(t, err)
	require.Equal(t, "foo value bar baz", w.String())
}
