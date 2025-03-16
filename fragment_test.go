package aitch

import (
	"bytes"
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
}

func TestFragment_Render(t *testing.T) {
	txt := Fragment("foo", nil, "bar")
	var w bytes.Buffer
	err := txt.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "foobar", w.String())

	ew := &errorWriter{}
	err = txt.Render(ew, nil)
	require.Error(t, err)
}
