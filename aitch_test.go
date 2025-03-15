package aitch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultContext(t *testing.T) {
	ctx := defaultContext()
	assert.NotNil(t, ctx)
	assert.NotNil(t, ctx.Data)
	assert.Nil(t, ctx.Error)
}
