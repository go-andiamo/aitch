package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestText(t *testing.T) {
	txt := Text("foo", nil, "bar")
	assert.Equal(t, TextNode, txt.Type())
	assert.Equal(t, "#text", txt.Name())
	rv := txt.(*text)
	assert.Equal(t, 2, len(rv.values))
	assert.Equal(t, 2, len(rv.getValues()))
}

func TestText_Render(t *testing.T) {
	txt := Text("foo", nil, "bar")
	assert.Equal(t, "foobar", testRender(txt, t))

	err := txt.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

func TestNbsp_NoMutation(t *testing.T) {
	x := Nbsp()
	x[0] = '?'
	y := Nbsp()
	assert.Equal(t, "&nbsp;", string(y))
}

func TestNbsps(t *testing.T) {
	x := Nbsps(0)
	assert.Equal(t, 0, len(x))
	x = Nbsps(2)
	assert.Equal(t, "&nbsp;&nbsp;", string(x))
}
