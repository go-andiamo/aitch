package html

import (
	"errors"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestDocument(t *testing.T) {
	d := Document("en", []any{Meta(Charset("utf-8"))}, nil)
	assert.Equal(t, aitch.DocumentNode, d.Type())
	assert.Equal(t, "#document", d.Name())
	rv := d.(*document)
	assert.Equal(t, 1, len(rv.contents))
	outer := rv.contents[0]
	assert.Equal(t, "html", outer.Name())
	assert.Equal(t, aitch.ElementNode, outer.Type())
}

func TestDocument_Render(t *testing.T) {
	d := Document("en", []any{Meta(Charset("utf-8"))}, []any{P()})
	assert.Equal(t, `<!DOCTYPE html>
<html lang="en"><head><meta charset="utf-8"></head><body><p></p></body></html>`, testRender(d, t))

	err := d.Render(&context.Context{Writer: &errorWriter{}})
	require.Error(t, err)
}

type errorWriter struct{}

var _ io.Writer = (*errorWriter)(nil)

func (e *errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("error")
}
