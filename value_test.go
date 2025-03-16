package aitch

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValue_Render(t *testing.T) {
	v := value{[]byte("foo"), nil}
	var w bytes.Buffer
	ok, err := v.render(&w, &Context{})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "foo", w.String())
}

func TestValue_Render_Dynamic(t *testing.T) {
	v := value{nil, func(ctx *Context) []byte {
		return []byte("foo")
	}}
	var w bytes.Buffer
	ok, err := v.render(&w, &Context{})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "foo", w.String())
}

func TestValue_Render_Empty(t *testing.T) {
	v := value{}
	var w bytes.Buffer
	ok, err := v.render(&w, &Context{})
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
	vs := newValue(nil)
	assert.Equal(t, 0, len(vs))

	vs = newValue(DynamicKey("foo"))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(func(*Context) []byte {
		return []byte("foo")
	})
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(DynamicFunc(func(*Context) []byte {
		return []byte("foo")
	}))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(DynamicKey("foo"))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	vs = newValue(newAttribute([]byte("att"), "foo", "bar"))
	assert.Equal(t, 2, len(vs))
	assert.Nil(t, vs[0].dynamicFunc)

	vs = newValue(Collection())
	assert.Equal(t, 0, len(vs))

	vs = newValue(Collection(newAttribute([]byte("att"), "foo", "bar"), newAttribute([]byte("att"), "baz")))
	assert.Equal(t, 3, len(vs))

	vs = newValue(P())
	assert.Equal(t, 0, len(vs))
}

func TestNewValue_DynamicKey(t *testing.T) {
	vs := newValue(DynamicKey("foo"))
	assert.Equal(t, 1, len(vs))
	assert.Nil(t, vs[0].value)
	assert.NotNil(t, vs[0].dynamicFunc)

	var w bytes.Buffer
	ok, err := vs[0].render(&w, &Context{
		Data: map[string]any{"foo": "bar"},
	})
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "bar", w.String())

	w = bytes.Buffer{}
	ok, err = vs[0].render(&w, &Context{
		Data: map[string]any{},
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
	data := HTMLEscapeString("foo")
	assert.Equal(t, "foo", string(data))

	data = HTMLEscapeString(`foo'&`)
	assert.Equal(t, "foo&#39;&amp;", string(data))
}

type stringy struct{}

func (s stringy) String() string {
	return "foo"
}
