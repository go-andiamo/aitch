package aitch

import (
	"bytes"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPI(t *testing.T) {
	d := PI("doctype", "foo")
	assert.Equal(t, PINode, d.Type())
	assert.Equal(t, "!doctype", d.Name())
	rv := d.(*pi)
	assert.Equal(t, 1, len(rv.values))
	assert.Equal(t, 1, len(rv.getValues()))
}

func TestPI_Render(t *testing.T) {
	d := PI("doctype", "foo")
	var w bytes.Buffer
	err := d.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<!doctype foo>\n", w.String())

	ew := &errorWriter{}
	err = d.Render(ew, &context.Context{})
	require.Error(t, err)
}

func TestDocType(t *testing.T) {
	d := DocType()
	assert.Equal(t, PINode, d.Type())
	assert.Equal(t, "!DOCTYPE", d.Name())
	rv := d.(*pi)
	assert.Equal(t, 1, len(rv.values))
	assert.Equal(t, "html", string(rv.values[0].value))
}
