package aitch

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func A(contents ...Node) Node {
	return NewElement([]byte("a"), contents...)
}

func B(contents ...Node) Node {
	return NewElement([]byte("b"), contents...)
}

func P(contents ...Node) Node {
	return NewElement([]byte("p"), contents...)
}

func Div(contents ...Node) Node {
	return NewElement([]byte("div"), contents...)
}

func Class(v ...any) Node {
	return NewDelimitedAttribute([]byte("class"), space, v...)
}

func Id(v ...any) Node {
	return NewAttribute([]byte("id"), v...)
}

func TestCollection(t *testing.T) {
	c := Collection(P(), P())
	assert.Equal(t, collectionNode, c.Type())
	assert.Equal(t, "#collection", c.Name())
	rv := c.(*collection)
	assert.Equal(t, 2, len(rv.nodes))
}

func TestCollection_Render(t *testing.T) {
	c := Collection(P(Text("A")), P(Text("B")), Class("not seen"))
	var w bytes.Buffer
	err := c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<p>A</p><p>B</p>", w.String())

	ew := &errorWriter{}
	err = c.Render(ew, nil)
	require.Error(t, err)
}

func TestCollection_Render_Nested(t *testing.T) {
	collectionA := Collection(P(Text("A")), Class("a"))
	collectionB := Collection(collectionA, P(Text("B")), Class("b"))
	c := Collection(Div(Collection(collectionB)))
	var w bytes.Buffer
	err := c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, `<div class="a b"><p>A</p><p>B</p></div>`, w.String())
}
