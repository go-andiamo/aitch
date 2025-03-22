package aitch

import (
	"bytes"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComment(t *testing.T) {
	c := Comment("foo")
	assert.Equal(t, CommentNode, c.Type())
	assert.Equal(t, "#comment", c.Name())
	rv := c.(*comment)
	assert.Equal(t, 1, len(rv.values))
	assert.Equal(t, 1, len(rv.getValues()))
}

func TestComment_Render(t *testing.T) {
	c := Comment("foo")
	var w bytes.Buffer
	err := c.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<!--foo-->", w.String())

	ew := &errorWriter{}
	err = c.Render(ew, &context.Context{})
	require.Error(t, err)
}
