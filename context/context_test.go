package context

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestNew(t *testing.T) {
	var w bytes.Buffer
	ctx := New(&w, nil, nil)
	assert.NotNil(t, ctx)
	assert.Equal(t, &w, ctx.Writer)
}

func TestContext_Write(t *testing.T) {
	var w bytes.Buffer
	ctx := New(&w, nil, nil)
	ctx.Write([]byte("hello"))
	assert.Equal(t, "hello", w.String())
}

func TestContext_Write_Errors(t *testing.T) {
	ctx := New(&errorWriter{}, nil, nil)
	ctx.Write([]byte("hello"))
	ctx.Write([]byte("world"))
	assert.Error(t, ctx.Error)
	assert.Equal(t, "error 1", ctx.Error.Error())
}

type errorWriter struct {
	calls int
}

var _ io.Writer = (*errorWriter)(nil)

func (e *errorWriter) Write([]byte) (int, error) {
	e.calls++
	return 0, fmt.Errorf("error %d", e.calls)
}

func TestGet(t *testing.T) {
	_, ok := Get[string](nil, "key")
	assert.False(t, ok)

	ctx := New(nil, nil, nil)
	_, ok = Get[string](ctx, "key")
	assert.False(t, ok)

	ctx = New(nil, map[string]any{
		"key": 1,
	}, nil)
	_, ok = Get[string](ctx, "key")
	assert.False(t, ok)
	v, ok := Get[int](ctx, "key")
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	ctx = New(nil, nil, map[string]any{
		"key": 1,
	})
	_, ok = Get[string](ctx, "key")
	assert.False(t, ok)
	v, ok = Get[int](ctx, "key")
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	ctx = New(nil, nil, &struct {
		Key int
	}{Key: 1})
	_, ok = Get[string](ctx, "Key")
	assert.False(t, ok)
	v, ok = Get[int](ctx, "Key")
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestMustGet(t *testing.T) {
	ctx := New(nil, map[string]any{
		"key": 1,
	}, nil)
	v := MustGet[int](ctx, "key")
	assert.Equal(t, 1, v)
	assert.Panics(t, func() {
		_ = MustGet[string](ctx, "key")
	})
}
