package aitch

import (
	"bytes"
	"fmt"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Async() Node {
	return NewBooleanAttribute([]byte("async"))
}

func TestImperativeContext(t *testing.T) {
	testCases := []struct {
		do             func(ic ImperativeContext, t *testing.T)
		expect         string
		expectStackLen int
	}{
		{
			do: func(ic ImperativeContext, t *testing.T) {
				//nothing
			},
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.WriteRaw([]byte("foo"))
			},
			expect: "foo",
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("div"), false)
			},
			expect:         `<div>`,
			expectStackLen: 1,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("br"), true)
			},
			expect:         `<br>`,
			expectStackLen: 0,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("div"), false, Id("1"), Class("foo"))
			},
			expect:         `<div id="1" class="foo">`,
			expectStackLen: 1,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("br"), true, Id("1"), Class("foo"))
			},
			expect:         `<br id="1" class="foo">`,
			expectStackLen: 0,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("div"), false)
				ic.End()
			},
			expect:         `<div></div>`,
			expectStackLen: 0,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("br"), true)
				ic.End()
			},
			expect:         `<br>`,
			expectStackLen: 0,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("div"), false)
				ic.Start([]byte("p"), false, Id("1"), Class("foo"))
				ic.Start([]byte("br"), true)
				ic.End()
			},
			expect:         `<div><p id="1" class="foo"><br></p>`,
			expectStackLen: 1,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("div"), false)
				ic.Start([]byte("p"), false, Id("1"), Class("foo"))
				ic.Start([]byte("br"), true)
				ic.EndAll()
			},
			expect:         `<div><p id="1" class="foo"><br></p></div>`,
			expectStackLen: 0,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.Start([]byte("div"), false, Async(), Id("1"), Class("foo", "bar"), Id("2"), Class("baz"), Async())
			},
			expect:         `<div async id="2" class="foo bar baz">`,
			expectStackLen: 1,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				ic.WriteNodes(Div(P(Text("bar"), Class("foo"))))
			},
			expect:         `<div><p class="foo">bar</p></div>`,
			expectStackLen: 0,
		},
		{
			do: func(ic ImperativeContext, t *testing.T) {
				assert.NotNil(t, ic.Context())
			},
			expect:         ``,
			expectStackLen: 0,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i+1), func(t *testing.T) {
			var w bytes.Buffer
			ctx := &context.Context{Writer: &w}
			ic := &imperativeContext{ctx: ctx}
			tc.do(ic, t)
			assert.Equal(t, tc.expect, w.String())
			assert.Equal(t, tc.expectStackLen, len(ic.openedStack))
		})
	}
}

func TestImperative(t *testing.T) {
	n := Imperative(nil)
	assert.NotNil(t, n)
	assert.Equal(t, imperativeNode, n.Type())
	assert.Equal(t, "!imperative", n.Name())
}

func TestImperative_Render(t *testing.T) {
	n := Imperative(func(ctx ImperativeContext) error {
		ctx.WriteRaw([]byte("foo"))
		return nil
	})
	assert.Equal(t, "foo", testRender(n, t))
}
