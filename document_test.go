package aitch

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDocument(t *testing.T) {
	d := Document("en", []Node{Meta(Charset("utf-8"))}, nil)
	assert.Equal(t, DocumentNode, d.Type())
	assert.Equal(t, "#document", d.Name())
	rv := d.(*document)
	assert.Equal(t, 1, len(rv.contents))
	outer := rv.contents[0]
	assert.Equal(t, "html", outer.Name())
	assert.Equal(t, ElementNode, outer.Type())
}

func TestDocument_Render(t *testing.T) {
	d := Document("en", []Node{Meta(Charset("utf-8"))}, []Node{P()})
	var w bytes.Buffer
	err := d.Render(&w, nil)
	require.NoError(t, err)
	assert.Equal(t, "<!DOCTYPE html>\n<html lang=\"en\"><head><meta charset=\"utf-8\"></head><body><p></p></body></html>", w.String())

	ew := &errorWriter{}
	err = d.Render(ew, &Context{})
	require.Error(t, err)
}
