package aitch

import (
	"bytes"
	"errors"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func testRender(node Node, t *testing.T, data ...map[string]any) string {
	var w bytes.Buffer
	useData := map[string]any{}
	for _, d := range data {
		for k, v := range d {
			useData[k] = v
		}
	}
	err := node.Render(&context.Context{Writer: &w, Data: useData})
	require.NoError(t, err)
	return w.String()
}

type errorWriter struct{}

var _ io.Writer = (*errorWriter)(nil)

func (e *errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("error")
}
