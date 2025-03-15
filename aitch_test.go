package aitch

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestDefaultContext(t *testing.T) {
	ctx := defaultContext()
	assert.NotNil(t, ctx)
	assert.NotNil(t, ctx.Data)
	assert.Nil(t, ctx.Error)
}

type errorWriter struct{}

var _ io.Writer = (*errorWriter)(nil)

func (e *errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("error")
}
