package html

import (
	"bytes"
	"errors"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestDocument(t *testing.T) {
	d := Document("en", []aitch.Node{Meta(Charset("utf-8"))}, nil)
	assert.Equal(t, aitch.DocumentNode, d.Type())
	assert.Equal(t, "#document", d.Name())
	rv := d.(*document)
	assert.Equal(t, 1, len(rv.contents))
	outer := rv.contents[0]
	assert.Equal(t, "html", outer.Name())
	assert.Equal(t, aitch.ElementNode, outer.Type())
}

func TestDocument_Render(t *testing.T) {
	d := Document("en", []aitch.Node{Meta(Charset("utf-8"))}, []aitch.Node{P()})
	var w bytes.Buffer
	err := d.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<!DOCTYPE html>\n<html lang=\"en\"><head><meta charset=\"utf-8\"></head><body><p></p></body></html>", w.String())

	ew := &errorWriter{}
	err = d.Render(ew, &context.Context{})
	require.Error(t, err)
}

type errorWriter struct{}

var _ io.Writer = (*errorWriter)(nil)

func (e *errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("error")
}
