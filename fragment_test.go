package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFragment(t *testing.T) {
	txt := Fragment("foo", nil, "bar")
	assert.Equal(t, FragmentNode, txt.Type())
	assert.Equal(t, "#fragment", txt.Name())
	rv := txt.(*fragment)
	assert.Equal(t, 2, len(rv.values))
	assert.Equal(t, 2, len(rv.getValues()))
}

func TestFragment_Render(t *testing.T) {
	txt := Fragment("foo", nil, "bar")
	assert.Equal(t, "foobar", testRender(txt, t))

	err := txt.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}
