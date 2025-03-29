package aitch

import (
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
	assert.Equal(t, "<!doctype foo>\n", testRender(d, t))

	err := d.Render(&context.Context{Writer: &errorWriter{}})
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
