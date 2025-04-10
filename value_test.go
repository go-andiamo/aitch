package aitch

import (
	"bytes"
	"fmt"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValue_Render(t *testing.T) {
	v := Value{[]byte("foo"), nil, nil}
	var w bytes.Buffer
	ok, err := v.render(&context.Context{Writer: &w})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "foo", w.String())
}

func TestValue_Render_Dynamic(t *testing.T) {
	v := Value{nil, func(ctx *context.Context) []byte {
		return []byte("foo")
	}, nil}
	var w bytes.Buffer
	ok, err := v.render(&context.Context{Writer: &w})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "foo", w.String())
}

func TestValue_Render_Dynamic_Dot(t *testing.T) {
	vs := newValue(DynamicValueKey("."))
	v := vs[0]
	var w bytes.Buffer
	ok, err := v.render(&context.Context{Writer: &w, Cargo: "foo"})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "foo", w.String())
}

func TestValue_Render_Empty(t *testing.T) {
	v := Value{}
	var w bytes.Buffer
	ok, err := v.render(&context.Context{Writer: &w})
	require.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, 0, w.Len())
}

func TestNewValues(t *testing.T) {
	vs := newValues()
	assert.Equal(t, 0, len(vs))

	vs = newValues(nil)
	assert.Equal(t, 0, len(vs))

	vs = newValues(nil, "foo", nil)
	assert.Equal(t, 1, len(vs))
	assert.Equal(t, []byte("foo"), vs[0].value)
}

func TestNewValue(t *testing.T) {
	v := NewValue("foo")
	assert.Equal(t, []byte("foo"), v.value)
	assert.Nil(t, v.dynamicFunc)
	assert.Nil(t, v.concats)

	v = NewValue(Collection(Text("foo"), Text("bar")))
	assert.Nil(t, v.value)
	assert.Nil(t, v.dynamicFunc)
	assert.Len(t, v.concats, 2)
}

func TestNewConcatValue(t *testing.T) {
	v := NewConcatValue("foo", "bar")
	assert.Nil(t, v.value)
	assert.Nil(t, v.dynamicFunc)
	assert.Len(t, v.concats, 2)
}

func TestConcatValue_Render(t *testing.T) {
	v := NewConcatValue("foo", "bar")
	var w bytes.Buffer
	ok, err := v.render(&context.Context{Writer: &w})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "foobar", w.String())
}

func TestNewValue_Inner(t *testing.T) {
	vs := newValue(nil)
	assert.Equal(t, 0, len(vs))

	vs = newValue(Value{concats: []Value{{concats: []Value{}}}})
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.Nil(t, vs[0].dynamicFunc)
	assert.NotNil(t, vs[0].concats)

	vs = newValue(DynamicValueKey("foo"))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(func(*context.Context) []byte {
		return []byte("foo")
	})
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(DynamicValueFunc(func(*context.Context) []byte {
		return []byte("foo")
	}))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(DynamicValueKey("foo"))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(NewAttribute([]byte("att"), "foo", "bar"))
	assert.Equal(t, 2, len(vs))
	assert.Nil(t, vs[0].dynamicFunc)

	vs = newValue(Collection())
	assert.Equal(t, 0, len(vs))

	vs = newValue(Collection(NewAttribute([]byte("att"), "foo", "bar"), NewAttribute([]byte("att"), "baz")))
	assert.Equal(t, 3, len(vs))

	vs = newValue(P())
	assert.Equal(t, 0, len(vs))
}

func TestNewValue_DynamicKey(t *testing.T) {
	vs := newValue(DynamicValueKey("foo"))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	var w bytes.Buffer
	ok, err := vs[0].render(&context.Context{
		Data:   map[string]any{"foo": "bar"},
		Writer: &w,
	})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "bar", w.String())

	w = bytes.Buffer{}
	ok, err = vs[0].render(&context.Context{
		Data:   map[string]any{},
		Writer: &w,
	})
	require.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, 0, w.Len())
}

func TestValueToBytes(t *testing.T) {
	testCases := []struct {
		value  any
		expect []byte
	}{
		{},
		{
			value:  true,
			expect: []byte("true"),
		},
		{
			value:  false,
			expect: []byte("false"),
		},
		{
			value:  []byte("foo"),
			expect: []byte("foo"),
		},
		{
			value:  "foo",
			expect: []byte("foo"),
		},
		{
			value:  1,
			expect: []byte("1"),
		},
		{
			value:  int8(1),
			expect: []byte("1"),
		},
		{
			value:  int16(1),
			expect: []byte("1"),
		},
		{
			value:  int32(1),
			expect: []byte("1"),
		},
		{
			value:  int64(1),
			expect: []byte("1"),
		},
		{
			value:  uint(1),
			expect: []byte("1"),
		},
		{
			value:  uint8(1),
			expect: []byte("1"),
		},
		{
			value:  uint16(1),
			expect: []byte("1"),
		},
		{
			value:  uint32(1),
			expect: []byte("1"),
		},
		{
			value:  uint64(1),
			expect: []byte("1"),
		},
		{
			value:  float32(1.1),
			expect: []byte("1.1"),
		},
		{
			value:  float64(1.1),
			expect: []byte("1.1"),
		},
		{
			value:  stringy{},
			expect: []byte("foo"),
		},
		{
			value: func() any {
				return "foo"
			},
			expect: []byte("foo"),
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			result := valueToBytes(tc.value)
			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestHTMLEscapeString(t *testing.T) {
	data := htmlEscapeString("foo")
	assert.Equal(t, "foo", string(data))

	data = htmlEscapeString(`foo'&"<>` + "\u0000")
	assert.Equal(t, "foo&apos;&amp;&quot;&lt;&gt;", string(data))
}

type stringy struct{}

func (s stringy) String() string {
	return "foo"
}
